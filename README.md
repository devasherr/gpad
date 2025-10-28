# gpad

`gpad` is a Go CLI Tool that automatically **reorders struct fileds across your codebase** to optimize memory layout and reduce padding. It helps improve memory efficiency in Go programs by analyzing structs and rearranging fields for better alignment. 

## Installation

Install `gpad` using `go install`:

```bash
go install github.com/devasherr/gpad@v0.1.2
```

## Usage

After installation, run gpad from the root of your Go project:
```bash
gpad
```
Thats it, all structs in the codebase have been sorted!!


You can also specify a custom directory to start from using the -path flag:
```bash
gpad -path "/internal/models"
```
This will only process structs within the specified directory and its subdirectories
