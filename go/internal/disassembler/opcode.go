package disassembler

import (
	"errors"
	"fmt"
)

// ErrUnknownOp represents an unknown Opcode.
var ErrUnknownOp = errors.New("unknown opcode")

// Opcode represents a single CHIP-8 operation.
type Opcode uint16

// Bytes splits the opcode into its respective bytes and returns them.
func (o Opcode) Bytes() (firstByte, secondByte byte) {
	return byte(o >> 8), byte(o)
}

// Disassemble returns the Opcode's name and instruction. If the Opcode is
// unknown then ErrUnknownOp will be returned.
func (o Opcode) Disassemble() (name, instruction string, err error) {
	firstByte, secondByte := o.Bytes()

	switch firstByte >> 4 {
	case 0x06:
		name = "MVI"
		instruction = fmt.Sprintf(
			"V%01X,#$%02x",
			firstByte&0xF,
			secondByte,
		)
	case 0x0A:
		name = "MVI"
		instruction = fmt.Sprintf(
			"I,#$%01x%02x",
			firstByte&0x0F,
			secondByte,
		)
	default:
		err = ErrUnknownOp
	}
	return
}
