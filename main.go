package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/mitchellh/cli"
	"github.com/stayradiated/gpx/compress"
	"github.com/stayradiated/gpx/coords"
	"github.com/stayradiated/gpx/truncate"
)

func coordsAction() (cli.Command, error) {
	return &coordsCommand{}, nil
}

type coordsCommand struct{}

func (c *coordsCommand) Help() string {
  return "Usage: gpx json-coords [FILENAME]"
}

func (c *coordsCommand) Synopsis() string {
	return "Extract coordinates from a GPX file"
}

func (c *coordsCommand) Run(args []string) int {
	data, err := readInput(args)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	result, err := coords.Coords(data)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println(result)
	return 0
}

func reducePointsAction() (cli.Command, error) {
	return &reducePointsCommand{}, nil
}

type reducePointsCommand struct {
	count int
}

func (c *reducePointsCommand) Help() string {
	return "Usage: gpx reduce-points [OPTIONS] FILENAME\n\n" +
		"Options:\n" +
		"  --count=N     Use N coordinates (default 2)\n"
}

func (c *reducePointsCommand) Synopsis() string {
	return "Compress a GPX file by reducing the number of points"
}

func (c *reducePointsCommand) Run(args []string) int {
	// Parse command line flags
	flags := flag.NewFlagSet("reduce-points", flag.ContinueOnError)
	flags.IntVar(&c.count, "count", 2, "Use N coordinates")
	flags.Parse(args)

	data, err := readInput(flag.Args())
	if err != nil {
		fmt.Println(err)
		return 1
	}

	result, err := compress.Compress(data, c.count)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println(result)
	return 0
}

func reducePrecisionAction() (cli.Command, error) {
	return &reducePrecisionCommand{}, nil
}

type reducePrecisionCommand struct {
	precision int
}

func (c *reducePrecisionCommand) Help() string {
	return "Usage: gpx reduce-precisionAME]\n\n" +
		"Options:\n" +
		"  --precision=N Set precision to N decimal places (default 2)"
}

func (c *reducePrecisionCommand) Synopsis() string {
	return "Reduce the precision of a GPX file's lat/lon values"
}

func (c *reducePrecisionCommand) Run(args []string) int {
	// Parse command line flags
	flags := flag.NewFlagSet("reduce-precision", flag.ContinueOnError)
	flags.IntVar(&c.precision, "precision", 2, "Set precision to N decimal places")
	flags.Parse(args)

	data, err := readInput(flags.Args())
	if err != nil {
		fmt.Println(err)
		return 1
	}

	result, err := truncate.Truncate(data, c.precision)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println(result)
	return 0
}

func readInput(args []string) ([]byte, error) {
	var data []byte
	var err error

	useStdin, err := hasStdinAvailable()
	if err != nil {
		return data, err
	}

	if useStdin {
		return readInputFromStdin()
	}

	if len(args) <= 0 {
		return data, errors.New("No filename specified and no data received on stdin.")
	}
	filename := args[len(args)-1]
	return readInputFromFilename(filename)
}

func hasStdinAvailable() (bool, error) {
	stat, err := os.Stdin.Stat()
	if err != nil {
		return false, err
	}
	return (stat.Mode() & os.ModeCharDevice) == 0, nil
}

func readInputFromStdin() ([]byte, error) {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func readInputFromFilename(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return data, err
}

func main() {
	c := cli.NewCLI("gpx", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"json-coords":   coordsAction,
		"reduce-points": reducePointsAction,
		"reduce-precision": reducePrecisionAction,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(exitStatus)
}
