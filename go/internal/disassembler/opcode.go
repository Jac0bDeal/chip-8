package disassembler

import (
	"errors"
	"fmt"
)

// UnknownOpError represents an unknown Opcode
var UnknownOpError = errors.New("unknown opcode")

// Opcode represents a single CHIP-8 operation.
type Opcode struct {
	MostSignificantByte  byte
	LeastSignificantByte byte
}

// Disassemble returns the Opcode's name and instruction. If the Opcode is
// unknown then UnknownOpError will be returned.
func (o Opcode) Disassemble() (name, instruction string, err error) {
	switch o.MostSignificantByte >> 4 {
	case 0x06:
		name = "MVI"
		instruction = fmt.Sprintf(
			"V%01X,#$%02x",
			o.MostSignificantByte&0xF,
			o.LeastSignificantByte,
		)
	case 0x0A:
		name = "MVI"
		instruction = fmt.Sprintf(
			"I,#$%01x%02x",
			o.MostSignificantByte&0x0F,
			o.LeastSignificantByte,
		)
	default:
		err = UnknownOpError
	}
	return
}
