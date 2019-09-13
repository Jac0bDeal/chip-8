package rom

import (
	"chip-8/internal/cpu"
	"fmt"
	"io/ioutil"
)

// CHIP-8 programs are set at 0x2 in memory by convention
// so we skip ahead to that address.
const romStartOffset = 0x200

// Load opens a file and returns the bytes in as a byte array.
func Load(filepath string) ([]byte, error) {
	return ioutil.ReadFile(filepath)
}

// Disassemble parses the passed ROM bytes into a human readable assembly format.
func Disassemble(romBytes []byte) {
	for pc := 0; pc < len(romBytes); pc += 2 {
		bytes := romBytes[pc : pc+2]
		op := cpu.OpcodeFromBytes(bytes)
		lineOut := fmt.Sprintf("%04x %02x %02x %s", pc+romStartOffset, bytes[0], bytes[1], op.Instruction())
		fmt.Println(lineOut)
	}
}
