# Gorem

Gorem is a command-line tool that converts pixel values to rem in CSS files.

## Installation

To install `gorem`, you can use the following `go get` command:

```bash
go get -u github.com/rahuld109/gorem
```

Make sure your Go bin directory is included in your system's `PATH`.

## Usage

```bash
gorem <directory>
```

Converts pixel values to rem in CSS files within the specified directory.

## Options

- `-h`, `--help`: Show help message.
- `-v`, `--version`: Show version information.

## Examples

Convert CSS files in the current directory:

```bash
gorem .
```

## License

This project is licensed under the [MIT License](LICENSE).
