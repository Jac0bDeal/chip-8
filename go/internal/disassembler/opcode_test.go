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
			opcode:              0x6208,
			expectedName:        "MVI",
			expectedInstruction: "V2,#$08",
		},
		{
			label:               "Annn",
			opcode:              0xA220,
			expectedName:        "MVI",
			expectedInstruction: "I,#$220",
		},
	}
	for _, c := range cases {
		t.Run(c.label, func(t *testing.T) {
			name, instruction, err := c.opcode.Disassemble()
			require.NoError(t, err, "received 0x%X", c.opcode)

			assert.Equal(t, c.expectedName, name)
			assert.Equal(t, c.expectedInstruction, instruction)
		})
	}
}

func TestOpcode_DisassembleUnknownOpcode(t *testing.T) {
	unknownOpcode := disassembler.Opcode(0x0000)
	_, _, err := unknownOpcode.Disassemble()
	require.Error(t, err)
	require.Equal(t, err, disassembler.UnknownOpError)
}
