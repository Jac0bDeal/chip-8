package disassembler_test

import (
	"testing"

	"chip-8/internal/disassembler"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestOpcode_Disassemble(t *testing.T) {
	type testCase struct {
		label               string
		opcode              disassembler.Opcode
		expectedInstruction string
	}
	cases := []testCase{
		{
			label:               "00E0 clear the display",
			opcode:              0x00E0,
			expectedInstruction: "CLS       ",
		},
		{
			label:               "00EE return from a subroutine",
			opcode:              0x00EE,
			expectedInstruction: "RTS       ",
		},
		{
			label:               "1nnn jump to location nnn",
			opcode:              0x128A,
			expectedInstruction: "JUMP       $28A",
		},
		{
			label:               "2nnn call subroutine at nnn",
			opcode:              0x228A,
			expectedInstruction: "CALL       $28A",
		},
		{
			label:               "3xkk skip next instruction if Vx == kk",
			opcode:              0x3A00,
			expectedInstruction: "SKIP.EQ    VA,#$00",
		},
		{
			label:               "4xkk skip next instruction if Vx != kk",
			opcode:              0x4800,
			expectedInstruction: "SKIP.NE    V8,#$00",
		},
		{
			label:               "5xy0 skip next instruction if Vx == Vy",
			opcode:              0x5A70,
			expectedInstruction: "SKIP.EQ    VA,V7",
		},
		{
			label:               "6xkk set Vx = kk",
			opcode:              0x6208,
			expectedInstruction: "MVI        V2,#$08",
		},
		{
			label:               "7xkk set Vx = Vx + kk",
			opcode:              0x7B93,
			expectedInstruction: "ADI        VB,#$93",
		},
		{
			label:               "8xy0 set Vx = Vy",
			opcode:              0x83A0,
			expectedInstruction: "MOV        V3,VA",
		},
		{
			label:               "8xy1 set Vx = Vx OR Vy",
			opcode:              0x83A1,
			expectedInstruction: "OR         V3,VA",
		},
		{
			label:               "8xy2 set Vx = Vx AND Vy",
			opcode:              0x83A2,
			expectedInstruction: "AND        V3,VA",
		},
		{
			label:               "8xy3 set Vx = Vx XOR Vy",
			opcode:              0x83A3,
			expectedInstruction: "XOR        V3,VA",
		},
		{
			label:               "8xy4 set Vx = Vx + Vy, set VF = carry",
			opcode:              0x83B4,
			expectedInstruction: "ADD.       V3,VB",
		},
		{
			label:               "8xy5 set Vx = Vx - Vy, set VF = NOT borrow",
			opcode:              0x83B5,
			expectedInstruction: "SUB.       V3,VB",
		},
		{
			label:               "8xy6 set Vx = Vx SHR 1",
			opcode:              0x83B6,
			expectedInstruction: "SHR.       V3",
		},
		{
			label:               "8xy7 set Vx = Vy - Vx, set VF = NOT borrow",
			opcode:              0x83B7,
			expectedInstruction: "SUBB.      V3,VB",
		},
		{
			label:               "8xyE set Vx = Vx SHL 1",
			opcode:              0x83BE,
			expectedInstruction: "SHL.       V3",
		},
		{
			label:               "9xy0 skip next instruction if Vx != Vy",
			opcode:              0x93B0,
			expectedInstruction: "SKIP.NE    V3,VB",
		},
		{
			label:               "Annn set I = nnn",
			opcode:              0xA220,
			expectedInstruction: "MVI        I,#$220",
		},
		{
			label:               "Bnnn jump to location nnn + V0",
			opcode:              0xB290,
			expectedInstruction: "JUMP       $290(V0)",
		},
		{
			label:               "Cxkk set Vx = random byte AND kk",
			opcode:              0xC9F2,
			expectedInstruction: "RND        V9,#$F2",
		},
		{
			label:               "Dxyn display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision",
			opcode:              0xD229,
			expectedInstruction: "SPRITE.    V2,V2,#$9",
		},
		{
			label:               "Ex9E skip next instruction if key with the value of Vx is pressed",
			opcode:              0xE79E,
			expectedInstruction: "SKIP.KEY   V7",
		},
		{
			label:               "ExA1 skip next instruction if key with the value of Vx is not pressed",
			opcode:              0xE7A1,
			expectedInstruction: "SKIP.NOKEY V7",
		},
		{
			label:               "Fx07 set Vx = delay timer value",
			opcode:              0xF907,
			expectedInstruction: "MOV        V9,DELAY",
		},
		{
			label:               "Fx0A wait for a key press, store the value of the key in Vx",
			opcode:              0xFC0A,
			expectedInstruction: "WAITKEY    VC",
		},
		{
			label:               "Fx15 set delay timer = Vx",
			opcode:              0xF215,
			expectedInstruction: "MOV        DELAY,V2",
		},
		{
			label:               "Fx18 set sound timer = Vx",
			opcode:              0xF318,
			expectedInstruction: "MOV        SOUND,V3",
		},
		{
			label:               "Fx1E set I = I + Vx",
			opcode:              0xF81E,
			expectedInstruction: "ADD        I,V8",
		},
		{
			label:               "Fx29 set I = location of sprite for digit Vx",
			opcode:              0xFD29,
			expectedInstruction: "SPRITECHAR VD",
		},
		{
			label:               "Fx33 store BCD representation of Vx in memory locations I, I+1, and I+2",
			opcode:              0xF533,
			expectedInstruction: "MOVBCD     V5",
		},
		{
			label:               "Fx55 store registers V0 through Vx in memory starting at location I",
			opcode:              0xF755,
			expectedInstruction: "MOVM       (I),V0-V7",
		},
		{
			label:               "Fx65 read registers V0 through Vx from memory starting at location I",
			opcode:              0xFD65,
			expectedInstruction: "MOVM       V0-VD,(I)",
		},
	}
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			instruction, err := c.opcode.Instruction()
			require.NoError(t, err, "received 0x%X", c.opcode, "opcode not recognized")

			assert.Equal(t, c.expectedInstruction, instruction, "incorrect instruction returned for opcode")
		})
	}
}

func TestOpcode_DisassembleUnknownOpcodes(t *testing.T) {
	type testCase struct {
		label  string
		opcode disassembler.Opcode
	}
	cases := []testCase{
		{
			label:  "unknown 0 code",
			opcode: 0x0000,
		},
		{
			label:  "unknown 8 code",
			opcode: 0x800f,
		},
		{
			label:  "unknown e code",
			opcode: 0xe000,
		},
		{
			label:  "unknown f code",
			opcode: 0xf000,
		},
	}
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			_, err := c.opcode.Instruction()
			require.Error(t, err, "should have received an error")
			require.Equal(t, err, disassembler.ErrUnknownOp, "error is not the correct type")
		})
	}
}
