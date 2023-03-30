README for GPX CLI program
===========================

The GPX CLI program is a tool for processing GPS data in GPX format. It provides four subcommands:

* `gpx from-toml [FILENAME]`: converts a TOML file into a GPX file
* `gpx json-coords [FILENAME]`: extracts coordinates from a GPX file
* `gpx reduce-points [OPTIONS] FILENAME`: compresses a GPX file by reducing the number of points
* `gpx reduce-precision [OPTIONS] FILENAME`: reduces the precision of a GPX file's lat/lon values

## Installation

To install the GPX CLI program, run the following command:

```
go get github.com/stayradiated/gpx-cli
```

## Usage

To use the GPX CLI program, run the `gpx` command followed by the subcommand you want to use. For example, to extract coordinates from a GPX file, run:

```
gpx json-coords path/to/gpx/file.gpx
```

Each subcommand has its own set of options and arguments. To see the available options and arguments for a subcommand, run:

```
gpx [SUBCOMMAND] --help
```

## Examples

### Converting a TOML file into a GPX file

To convert a TOML file into a GPX file, run:

```
gpx from-toml path/to/toml/file.toml > path/to/gpx/file.gpx
```

This will write the GPX file to the specified file path.

### Extracting coordinates from a GPX file

To extract coordinates from a GPX file, run:

```
gpx json-coords path/to/gpx/file.gpx
```

This will print the coordinates to standard output in JSON format.

### Compressing a GPX file by reducing the number of points

To compress a GPX file by reducing the number of points, run:

```
gpx reduce-points --count=100 path/to/gpx/file.gpx > path/to/compressed/gpx/file.gpx
```

This will compress the GPX file by reducing the number of points to 100 and write the compressed file to the specified file path.

### Reducing the precision of a GPX file's lat/lon values

To reduce the precision of a GPX file's lat/lon values, run:

```
gpx reduce-precision --precision=4 path/to/gpx/file.gpx > path/to/reduced-precision/gpx/file.gpx
```

This will reduce the precision of the GPX file's lat/lon values to 4 decimal places and write the result to the specified file path.

## License

This program is licensed under the MIT license. See the `LICENSE` file for more information.
