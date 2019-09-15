package cli

import (
	"chip-8/internal/rom"
	"io"
	"os"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var disassembleOut string

var cmdDisassemble = &cobra.Command{
	Use:   "disassemble <rom file>",
	Short: "Disassemble a CHIP-8 ROM file",
	Long: "disassemble reads the specified ROM file, disassembles it into opcodes,\n" +
		"then maps those opcodes into their respective instructions. Can return\n" +
		"a file containing the decompiled instructions, but writes to stdout by\n" +
		"default.",
	Args: cobra.ExactArgs(1),
	Run:  disassembleROM,
}

func init() {
	cmdDisassemble.Flags().StringVarP(&disassembleOut, "output", "o", "stdout", "Output file to write to.")
	rootCmd.AddCommand(cmdDisassemble)
}

func disassembleROM(_ *cobra.Command, args []string) {
	fileIn := args[0]
	rawRom, err := rom.Load(fileIn)
	if err != nil {
		logErrorAndExit(errors.Wrapf(err, "failed to load %s", fileIn))
	}

	disassembledRom, err := rom.Disassemble(rawRom)
	if err != nil {
		logErrorAndExit(errors.Wrapf(err, "failed to disassemble %s", fileIn))
	}

	var output io.Writer
	if disassembleOut == "stdout" || disassembleOut == "" {
		output = os.Stdout
	} else {
		output, err = os.Create(disassembleOut)
		if err != nil {
			logErrorAndExit(errors.Wrap(err, "failed to open output file for writing"))
		}
	}

	n, err := io.Copy(output, disassembledRom)
	if err != nil {
		logErrorAndExit(errors.Wrap(err, "failed to write out disassemble instructions"))
	}

	logAndExit(0, "wrote out %d bytes to %s", n, disassembleOut)
}
