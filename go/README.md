# CHIP-8 Emulator: Go Implementation

[![Go Report Card](https://goreportcard.com/badge/github.com/Jac0bDeal/chip-8)](https://goreportcard.com/report/github.com/Jac0bDeal/chip-8)
[![CircleCI](https://circleci.com/gh/Jac0bDeal/chip-8.svg?style=shield)](https://circleci.com/gh/Jac0bDeal/chip-8)

A Go implementation of the CHIP-8 emulator. 
This project is mostly for academic purposes as 
a first foray into emulators and is not intended
to serve as a polished product.

## Build
Building the binary is super simple (given you have `make` installed, which you should):
```shell
make
```

Binary is built and found inside the `bin/` directory when using `make`.

Cleaning up the project can be done with:
```shell
make clean
```

To run just the tests run:
```shell
make test
```

## Usage
The main purpose of chip8 is to load and run a CHIP-8 ROM in the emulator:

```shell
chip8 <filepath>
```

### Disassembler
The disassembler subcommand reads in a ROM file and dumps the diassembled instructions
to either stdout or a file for inspection.

Disassemble ROM and print to stdout:
```shell
chip8 disassemble <filepath>
```

Disassemble ROM and dump to file:
```shell
chip8 disassemble <filepath> -output <filepath>
chip8 disassemble <filepath> -o <filepath>
```
