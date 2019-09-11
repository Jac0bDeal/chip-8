package disassembler

import (
	"errors"
	"fmt"
)

// UnknownOpError represents an unknown Opcode
var UnknownOpError = errors.New("unknown opcode")

// Opcode represents a single CHIP-8 operation.
type Opcode uint16

func (o Opcode) Bytes() (high, low byte) {
	return byte(o >> 8), byte(o)
}

// Disassemble returns the Opcode's name and instruction. If the Opcode is
// unknown then UnknownOpError will be returned.
func (o Opcode) Disassemble() (name, instruction string, err error) {
	high, low := o.Bytes()

	switch high >> 4 {
	case 0x06:
		name = "MVI"
		instruction = fmt.Sprintf(
			"V%01X,#$%02x",
			high&0xF,
			low,
		)
	case 0x0A:
		name = "MVI"
		instruction = fmt.Sprintf(
			"I,#$%01x%02x",
			high&0x0F,
			low,
		)
	default:
		err = UnknownOpError
	}
	return
}
