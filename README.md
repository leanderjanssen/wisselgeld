# Wisselgeld
Calculate the minimum amount of coins to match the change

## Usage
`wisselgeld` can be run both with and without arguments.
When no argument is provided, it will ask you for a number. 

You can also provide one or more arguments on the commandline.  

Please note, only valid float numbers (using a `.` as the seperator) will be processed.

## Compilation
This program is written in [`golang`](https://golang.org/).
To compile the sourcecode you'll need to have the go compiler [installed](https://golang.org/doc/install).
Compiling the code can be done using any of the following commands:

- `go run wisselgeld.go`
- `go install`
- `go build`

For your convenience the binaries (for Windows, Linux and Mac OS) have also been added to this repository in the `bin` directory.
