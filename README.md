# Parker

Not peter parker, just a parker in Golang.

## Getting Started

These instructions will guide you to know about the project and running this on your machine.

### Prerequisites

Before started, please assure that your environment follows these specs:

- Go version 1.12+
- Linux/Unix Machine

The folder structure in this project is based on [aws-sdk-go/service](https://github.com/aws/aws-sdk-go/tree/master/service) directory structure. Each folder represents a package.

- `cmd` folder contains handlers to handle commands given from interactive shell.
- `park` folder contains the core algorithms to implement parker.
- `utils` folder contains some helper functions.

Each of the folders must contain a file named `errors.go` which declares the error types, and another file named `service.go` which is the core logic of the package.

### Running

To run the app, you need to build the binary first, just run the following command

```
make build / make b
```

Now, you can run the app

```
make run
```

Or simply run this command to build and run with a single tap

```
make
```

After you run this, it will display an interactive shell. These are list of available commands

<pre>
$ init   [cap]           initialize parker with given capacity
$ park   [reg] [color]   park a car with reg number and color
$ leave  [id]            remove a car from given slot id  	
$ get    [id]            get a car by slot id
$ search [term]          search cars by reg number / color 
$ exit                   exit 
</pre>

## Running the tests

To run the unit tests, run this

```
make test
```
