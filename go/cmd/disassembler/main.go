package main

import (
	"chip-8/internal/rom"
	"log"

	"github.com/pkg/errors"
)

func main() {
	romBytes, err := rom.Load("./test/roms/test_opcode.ch8")
	if err != nil {
		log.Fatal(errors.Wrap(err, "failed to load rom file"))
	}
	rom.Disassemble(romBytes)
}
