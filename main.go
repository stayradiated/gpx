package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/mitchellh/cli"
	"github.com/stayradiated/gpx/compress"
	"github.com/stayradiated/gpx/coords"
	"github.com/stayradiated/gpx/truncate"
)

type coordsCommand struct{}

func (c *coordsCommand) Help() string {
	return "Extract coordinates from a GPX file"
}

func (c *coordsCommand) Synopsis() string {
	return "Extract coordinates from a GPX file"
}

func (c *coordsCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("Missing filename argument")
		return 1
	}

	data, err := readInput(args[len(args)-1])
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

type compressCommand struct {
	Count int
}

func (c *compressCommand) Help() string {
	return "Compress a GPX file by reducing the number of points"
}

func (c *compressCommand) Synopsis() string {
	return "Compress a GPX file by reducing the number of points"
}

func (c *compressCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("Missing filename argument")
		return 1
	}

	data, err := readInput(args[len(args)-1])
	if err != nil {
		fmt.Println(err)
		return 1
	}

	result, err := compress.Compress(data, c.Count)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println(result)
	return 0
}

func (c *compressCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("compress", flag.ContinueOnError)
	fs.IntVar(&c.Count, "count", 2, "Number of points to keep in the compressed file")
	return fs
}

type truncateCommand struct {
	Precision int
}

func (c *truncateCommand) Help() string {
	return "Truncate the precision of a GPX file's lat/lon values"
}

func (c *truncateCommand) Synopsis() string {
	return "Truncate the precision of a GPX file's lat/lon values"
}

func (c *truncateCommand) Run(args []string) int {
	if len(args) == 0 {
		fmt.Println("Missing filename argument")
		return 1
	}

	data, err := readInput(args[len(args)-1])
	if err != nil {
		fmt.Println(err)
		return 1
	}

	result, err := truncate.Truncate(data, c.Precision)
	if err != nil {
		fmt.Println(err)
		return 1
	}

	fmt.Println(result)
	return 0
}

func (c *truncateCommand) Flags() *flag.FlagSet {
	fs := flag.NewFlagSet("truncate", flag.ContinueOnError)
	fs.IntVar(&c.Precision, "precision", 6, "Number of decimal places to truncate the lat/lon values to")
	return fs
}

func readInput(filename string) ([]byte, error) {
	var data []byte
	var err error

	// check if input is being piped in via stdin
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		data, err = ioutil.ReadAll(os.Stdin)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	} else {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer file.Close()

		data, err = ioutil.ReadAll(file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	return data, err
}
