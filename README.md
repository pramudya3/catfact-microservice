# catfact-microservice

This project implements a microservice written in Go that provides cat facts.

## Getting Started

### Prerequisites

* Golang (version 1.18 or higher) installed on your system. You can download it from the official website [https://go.dev/doc/install](https://go.dev/doc/install).
* A code editor or IDE of your choice (e.g., Visual Studio Code, GoLand).

### Building and Running

1. Clone this repository to your local machine.
2. Open a terminal window and navigate to the project directory.
3. Run the following command to build the microservice executable:

```bash
make build
```

This will create a file named `bin/fact` in the project directory.

4. Start the microservice using the following command:

```bash
make run
```

This will run the `fact` binary, which is the microservice responsible for providing cat facts.

## Testing

To run the unit tests for the project, execute the following command in your terminal:

```bash
make test
```

This will run all the tests defined within the project and report the results.

## Makefile Explained

The provided `Makefile` defines three targets:

* `build`: This target builds the microservice executable named `bin/fact` using the `go build` command.
* `run`: This target depends on the `build` target and subsequently runs the built executable `./bin/fact`.
* `test`: This target executes all the unit tests defined within the project using the `go test` command with the `-v` flag for verbose output.
