package cpu

import (
	"encoding/binary"
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

func OpcodeFrom(b []byte) Opcode {
	return Opcode(binary.BigEndian.Uint16(b))
}

// Instruction returns the Opcode's name and instruction. If the Opcode is
// unknown then ErrUnknownOp will be returned.
func (o Opcode) Instruction() string {
	firstByte, secondByte := o.Bytes()

	firstNib := firstByte >> 4
	secondNib := firstByte & 0xf
	thirdNib := secondByte >> 4
	fourthNib := secondByte & 0xf

	switch firstNib {
	case 0x0:
		switch secondByte {
		case 0xe0:
			return fmt.Sprintf("%-10s", "CLS")
		case 0xee:
			return fmt.Sprintf("%-10s", "RTS")
		default:
			return "unknown 0x0XXX opcode"
		}
	case 0x1:
		return fmt.Sprintf("%-10s $%01x%02X", "JUMP", secondNib, secondByte)
	case 0x2:
		return fmt.Sprintf("%-10s $%01x%02X", "CALL", secondNib, secondByte)
	case 0x3:
		return fmt.Sprintf("%-10s V%01X,#$%02X", "SKIP.EQ", secondNib, secondByte)
	case 0x4:
		return fmt.Sprintf("%-10s V%01X,#$%02X", "SKIP.NE", secondNib, secondByte)
	case 0x5:
		return fmt.Sprintf("%-10s V%01X,V%01X", "SKIP.EQ", secondNib, thirdNib)
	case 0x6:
		return fmt.Sprintf("%-10s V%01X,#$%02X", "MVI", secondNib, secondByte)
	case 0x7:
		return fmt.Sprintf("%-10s V%01X,#$%02X", "ADI", secondNib, secondByte)
	case 0x8:
		switch secondByte & 0xf {
		case 0x0:
			return fmt.Sprintf("%-10s V%01X,V%01X", "MOV", secondNib, thirdNib)
		case 0x1:
			return fmt.Sprintf("%-10s V%01X,V%01X", "OR", secondNib, thirdNib)
		case 0x2:
			return fmt.Sprintf("%-10s V%01X,V%01X", "AND", secondNib, thirdNib)
		case 0x3:
			return fmt.Sprintf("%-10s V%01X,V%01X", "XOR", secondNib, thirdNib)
		case 0x4:
			return fmt.Sprintf("%-10s V%01X,V%01X", "ADD.", secondNib, thirdNib)
		case 0x5:
			return fmt.Sprintf("%-10s V%01X,V%01X", "SUB.", secondNib, thirdNib)
		case 0x6:
			return fmt.Sprintf("%-10s V%01X", "SHR.", secondNib)
		case 0x7:
			return fmt.Sprintf("%-10s V%01X,V%01X", "SUBB.", secondNib, thirdNib)
		case 0xe:
			return fmt.Sprintf("%-10s V%01X", "SHL.", secondNib)
		default:
			return "unknown 0x8XXX opcode"
		}
	case 0x9:
		return fmt.Sprintf("%-10s V%01X,V%01X", "SKIP.NE", secondNib, thirdNib)
	case 0xa:
		return fmt.Sprintf("%-10s I,#$%01x%02X", "MVI", secondNib, secondByte)
	case 0xb:
		return fmt.Sprintf("%-10s $%01x%02X(V0)", "JUMP", secondNib, secondByte)
	case 0xc:
		return fmt.Sprintf("%-10s V%01X,#$%02X", "RND", secondNib, secondByte)
	case 0xd:
		return fmt.Sprintf("%-10s V%01X,V%01X,#$%01X", "SPRITE.", secondNib, thirdNib, fourthNib)
	case 0xe:
		switch secondByte {
		case 0x9E:
			return fmt.Sprintf("%-10s V%01X", "SKIP.KEY", secondNib)
		case 0xA1:
			return fmt.Sprintf("%-10s V%01X", "SKIP.NOKEY", secondNib)
		default:
			return "unknown 0xeXXX opcode"
		}
	case 0xf:
		switch secondByte {
		case 0x07:
			return fmt.Sprintf("%-10s V%01X,DELAY", "MOV", secondNib)
		case 0x0a:
			return fmt.Sprintf("%-10s V%01X", "WAITKEY", secondNib)
		case 0x15:
			return fmt.Sprintf("%-10s DELAY,V%01X", "MOV", secondNib)
		case 0x18:
			return fmt.Sprintf("%-10s SOUND,V%01X", "MOV", secondNib)
		case 0x1e:
			return fmt.Sprintf("%-10s I,V%01X", "ADD", secondNib)
		case 0x29:
			return fmt.Sprintf("%-10s V%01X", "SPRITECHAR", secondNib)
		case 0x33:
			return fmt.Sprintf("%-10s V%01X", "MOVBCD", secondNib)
		case 0x55:
			return fmt.Sprintf("%-10s (I),V0-V%01X", "MOVM", secondNib)
		case 0x65:
			return fmt.Sprintf("%-10s V0-V%01X,(I)", "MOVM", secondNib)
		default:
			return "unknown 0xfXXX opcode"
		}
	default:
		return "unknown opcode"
	}
}
