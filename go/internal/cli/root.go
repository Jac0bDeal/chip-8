package cli

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chip8",
	Short: "chip8 is an emulator to run, debug, and (dis)assemble CHIP-8 ROM's.",
	Long:  "",
	Run:   runROM,
}

// Execute loads and executes the cli app.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logErrorAndExit(errors.WithStack(err))
	}
}

func runROM(cmd *cobra.Command, args []string) {
	_, _ = cmd, args
}
