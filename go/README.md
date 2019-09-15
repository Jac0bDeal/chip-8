# CHIP-8 Emulator: Go Implementation

[![Go Report Card](https://goreportcard.com/badge/github.com/Jac0bDeal/chip-8)](https://goreportcard.com/report/github.com/Jac0bDeal/chip-8)
[![CircleCI](https://circleci.com/gh/Jac0bDeal/chip-8.svg?style=shield)](https://circleci.com/gh/Jac0bDeal/chip-8)

A Go implementation of the CHIP-8 emulator. 
This project is mostly for academic purposes as 
a first foray into emulators and is not intended
to serve as a polished product.

## Build
Building the binaries is super simple (given you have `make` installed, which you should):
```shell
make
```

Binaries are built and found inside the `bin/` directory when using `make`.

Cleaning up the project can be done with:
```shell
make clean
```

To run just the tests run:
```shell
make test
```

## Usage
There are two binaries built by the project:
- [diassemble](#disassembler)
- [chip8](#emulator)

### Disassembler
The disassembler is a CLI that reads in a ROM file and dumps the diassembled instructions
to either stdout or a file for inspection.

Disassemble ROM and print to stdout:
```shell
disassemble <filepath>
```

Disassemble ROM and dump to file:
```shell
disassemble <filepath> -o <filepath>
```

### Emulator
Not implemented yet, coming soon!
