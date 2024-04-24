package cpu

import (
	"log"

	"github.com/pkg/errors"
)

type (
	operation func() error
	opDecoder func(Opcode) operation
)

// NewCPU constructs and returns a pointer to a CPU instance with the
// stack pointer and program counter set to their initial values.
func NewCPU() *CPU {
	c := CPU{
		sp: 0xfa0,
		pc: 0x200,
	}
	c.registerOpDecoder()

	return &c
}

// CPU represents a CHIP-8 CPU and its hardware components for use in running a CHIP-8 program.
type CPU struct {
	V [16]byte
	I uint16

	pc uint16

	memory [4096]byte
	screen [2048]byte

	delay byte
	sound byte

	stack [16]uint16
	sp    uint16

	keyboard [16]byte

	opDecoder
}

// Cycle performs one CPU cycle by fetching, decoding, and executing an opcode.
func (c *CPU) Cycle() {
	// fetch the opcode corresponding to the current pc address
	b := c.memory[c.pc : c.pc+2]
	opcode := OpcodeFromBytes(b)

	// decode the opcode operation
	op := c.opDecoder(opcode)

	// execute the operation on the CPU
	if err := op(); err != nil {
		log.Println(errors.WithStack(err))
	}
}

func (c CPU) registerOpDecoder() {
	var _0x0map = map[byte]operation{
		0x00: c._0x0000,
		0xe0: c._0x00E0,
		0xee: c._0x00EE,
	}
	var _0x8map = map[byte]operation{
		0x0: c._0x8xy0,
		0x1: c._0x8xy1,
		0x2: c._0x8xy2,
		0x3: c._0x8xy3,
		0x4: c._0x8xy4,
		0x5: c._0x8xy5,
		0x6: c._0x8xy6,
		0x7: c._0x8xy7,
		0xe: c._0x8xyE,
	}
	var _0xEmap = map[byte]operation{
		0x9e: c._0xEx9E,
		0xa1: c._0xExA1,
	}
	var _0xFmap = map[byte]operation{
		0x07: c._0xFx07,
		0x0a: c._0xFx0A,
		0x15: c._0xFx15,
		0x18: c._0xFx18,
		0x1e: c._0xFx1E,
		0x29: c._0xFx29,
		0x33: c._0xFx33,
		0x55: c._0xFx55,
		0x65: c._0xFx65,
	}

	var opcodeMap = map[byte]func(byte) operation{
		0x0: func(b byte) operation { return _0x0map[b] },
		0x1: func(b byte) operation { return c._0x1nnn },
		0x2: func(b byte) operation { return c._0x2nnn },
		0x3: func(b byte) operation { return c._0x3xkk },
		0x4: func(b byte) operation { return c._0x4xkk },
		0x5: func(b byte) operation { return c._0x5xy0 },
		0x6: func(b byte) operation { return c._0x6xkk },
		0x7: func(b byte) operation { return c._0x7xkk },
		0x8: func(b byte) operation { return _0x8map[b&0xf] },
		0x9: func(b byte) operation { return c._0x9xy0 },
		0xa: func(b byte) operation { return c._0xAnnn },
		0xb: func(b byte) operation { return c._0xBnnn },
		0xc: func(b byte) operation { return c._0xCxkk },
		0xd: func(b byte) operation { return c._0xDxyn },
		0xe: func(b byte) operation { return _0xEmap[b] },
		0xf: func(b byte) operation { return _0xFmap[b] },
	}

	c.opDecoder = func(opcode Opcode) operation {
		firstByte, secondByte := opcode.Bytes()

		op := opcodeMap[firstByte](secondByte)
		if op == nil {
			return c.unknownOp
		}

		return op
	}
}

func (c *CPU) unknownOp() error {
	return ErrUnknownOpcode
}

func (c *CPU) _0x0000() error {
	return nil
}

func (c *CPU) _0x00E0() error {
	return nil
}

func (c *CPU) _0x00EE() error {
	return nil
}

func (c *CPU) _0x1nnn() error {
	return nil
}

func (c *CPU) _0x2nnn() error {
	return nil
}

func (c *CPU) _0x3xkk() error {
	return nil
}

func (c *CPU) _0x4xkk() error {
	return nil
}

func (c *CPU) _0x5xy0() error {
	return nil
}

func (c *CPU) _0x6xkk() error {
	return nil
}

func (c *CPU) _0x7xkk() error {
	return nil
}

func (c *CPU) _0x8xy0() error {
	return nil
}

func (c *CPU) _0x8xy1() error {
	return nil
}

func (c *CPU) _0x8xy2() error {
	return nil
}

func (c *CPU) _0x8xy3() error {
	return nil
}

func (c *CPU) _0x8xy4() error {
	return nil
}

func (c *CPU) _0x8xy5() error {
	return nil
}

func (c *CPU) _0x8xy6() error {
	return nil
}

func (c *CPU) _0x8xy7() error {
	return nil
}

func (c *CPU) _0x8xyE() error {
	return nil
}

func (c *CPU) _0x9xy0() error {
	return nil
}

func (c *CPU) _0xAnnn() error {
	return nil
}

func (c *CPU) _0xBnnn() error {
	return nil
}

func (c *CPU) _0xCxkk() error {
	return nil
}

func (c *CPU) _0xDxyn() error {
	return nil
}

func (c *CPU) _0xEx9E() error {
	return nil
}

func (c *CPU) _0xExA1() error {
	return nil
}

func (c *CPU) _0xFx07() error {
	return nil
}

func (c *CPU) _0xFx0A() error {
	return nil
}

func (c *CPU) _0xFx15() error {
	return nil
}

func (c *CPU) _0xFx18() error {
	return nil
}

func (c *CPU) _0xFx1E() error {
	return nil
}

func (c *CPU) _0xFx29() error {
	return nil
}

func (c *CPU) _0xFx33() error {
	return nil
}

func (c *CPU) _0xFx55() error {
	return nil
}

func (c *CPU) _0xFx65() error {
	return nil
}
