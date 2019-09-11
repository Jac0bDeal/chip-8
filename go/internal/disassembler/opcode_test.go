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
		expectedName        string
		expectedInstruction string
	}
	cases := []testCase{
		{
			label:               "6xkk",
			opcode:              disassembler.Opcode{MostSignificantByte: 0x62, LeastSignificantByte: 0x08},
			expectedName:        "MVI",
			expectedInstruction: "V2,#$08",
		},
		{
			label:               "Annn",
			opcode:              disassembler.Opcode{MostSignificantByte: 0xA2, LeastSignificantByte: 0x20},
			expectedName:        "MVI",
			expectedInstruction: "I,#$220",
		},
	}
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			name, instruction, err := c.opcode.Disassemble()
			require.NoError(t, err)

			assert.Equal(t, c.expectedName, name)
			assert.Equal(t, c.expectedInstruction, instruction)
		})
	}
}

func TestOpcode_DisassembleUnknownOpcode(t *testing.T) {
	unknownOpcode := disassembler.Opcode{MostSignificantByte: 0x00, LeastSignificantByte: 0x00}
	_, _, err := unknownOpcode.Disassemble()
	require.Error(t, err)
	require.Equal(t, err, disassembler.UnknownOpError)
}
