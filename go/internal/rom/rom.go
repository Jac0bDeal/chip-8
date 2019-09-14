package rom

import (
	"bytes"
	"chip-8/internal/cpu"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/pkg/errors"
)

// CHIP-8 programs are set at 0x2 in memory by convention
// so we skip ahead to that address.
const romMemStartOffset = 0x200

// Load opens a file and returns the rom bytes.
func Load(filepath string) (io.Reader, error) {
	romBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return bytes.NewReader(romBytes), nil
}

// Disassemble parses the passed ROM bytes into a human readable assembly format.
// It parses 2 bytes at a time, maps instruction, then appends it to the returned io.Reader.
func Disassemble(rom io.Reader) (io.Reader, error) {
	romBytes, err := ioutil.ReadAll(rom)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	instructions := bytes.NewBuffer([]byte{})

	for pc := 0; pc < len(romBytes); pc += 2 {
		b := romBytes[pc : pc+2]
		op := cpu.OpcodeFromBytes(b)
		lineOut := fmt.Sprintf("%04x %02x %02x %s", pc+romMemStartOffset, b[0], b[1], op.Instruction())
		_, err := instructions.WriteString(lineOut + "\n")
		if err != nil {
			return nil, errors.Wrapf(err, "failed to write instruction at byte %d", pc)
		}
	}

	return instructions, nil
}
