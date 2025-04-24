# Logger Package (Example project)

This repository contains a simple logging utility written in Go. The package provides a lightweight and customizable way to log messages with different levels of severity.

## Features

- Log messages with different levels: `INFO`, `WARN`, `ERROR`, etc.
- Easy-to-use API for logging.
- Customizable output format and destination.

## Files

### `logger.go`

This file contains the core implementation of the logging utility. It defines the logger structure, log levels, and methods for logging messages.

Key functionalities:
- Define log levels (e.g., `INFO`, `DEBUG`, `ERROR`).
- Format log messages with timestamps and severity levels.
- Write logs to the console, file or any destination which supports `io.Writer`.

### `main.go`

This file demonstrates how to use the logger package. It includes examples of logging messages with different severity levels and configuring the logger.

Key functionalities:
- Initialize the logger.
- Log messages using various log levels.
- Showcase customization options.

## Usage

1. Import the logger package in your Go project.
2. Initialize the logger.
3. Use the provided methods to log messages.

Example:
```go
package main

import "github.com/vorozhko/logger"

func main() {
    log := logger.New()
    log.Info("This is an info message")
    log.Warn("This is a warning message")
    log.Error("This is an error message")
}
```

## Installation

Clone the repository:
```bash
git clone https://github.com/vorozhko/logger.git
```

Build and run the example:
```bash
go run main.go
```

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.
