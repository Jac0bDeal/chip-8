package main

import (
	"chip-8/internal/rom"
	"io"
	"log"
	"os"

	"github.com/pkg/errors"
)

func main() {
	rawRom, err := rom.Load("./test/roms/test_opcode.ch8")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to load rom file"))
	}

	disassembledRom, err := rom.Disassemble(rawRom)
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to disassemble rom file"))
	}

	file, err := os.Create("output.asm")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to open output file for writing"))
	}

	if _, err = io.Copy(file, disassembledRom); err != nil {
		log.Fatal(errors.Wrap(err, "failed to write out file body"))
	}
}
