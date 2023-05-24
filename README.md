# Pastefile Command

`Powered by Chat-GPT`

The `pastefile` command is a simple utility that reads input from standard input and writes it to a specified file.

## Usage

```
pastefile [OPTIONS] FILENAME
```

## Options

- `-f, --force`: Overwrite the file if it already exists.

## Examples

1. Basic usage:

   ```
   $ echo "Hello, world!" | pastefile output.txt
   ```

   This will create a file named `output.txt` with the content "Hello, world!".

2. Overwrite existing file:

   ```
   $ echo "New content" | pastefile --force existing.txt
   ```

   This will overwrite the existing file `existing.txt` with the new content.

## Installation

To install the `pastefile` command, you can download the binary for your platform from the [Releases](https://github.com/sat8bit/releases) page, or you can build it from source using Go:

```
$ go get -u github.com/sat8bit/pastefile
```

Make sure that your Go environment is properly set up.

## License

This project is licensed under the [MIT License](LICENSE).

