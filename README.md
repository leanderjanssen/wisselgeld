# Wisselgeld
Calculate the minimum amount of coins to match the change

## Usage
`wisselgeld` can be run both with and without arguments.
When no argument is provided, it will ask you for a number. 

You can also provide one or more arguments on the commandline.  

```
# ./wisselgeld 1.23

Ingevoerd wisselgeld bedrag:   € 1.23
Afgerond wisselgeld bedrag:    € 1.25

Het minimum aantal munten is:

   1x  1 euro
   1x 20 cent
   1x  5 cent
```

Please note, only valid float numbers (using a `.` as the seperator) will be processed.

## Compilation
This program is written in [`golang`](https://golang.org/).
To compile the sourcecode you'll need to have the go compiler [installed](https://golang.org/doc/install).
Compiling the code can be done using any of the following commands:

- `go run wisselgeld.go`
- `go install`
- `go build`

For your convenience the precompiled binaries (for Windows, Linux and Mac OS) have also been added to this repository in the `bin` directory.
