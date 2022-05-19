package main

import (
	"browser-sync/internal/common"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	config  string
	cmdRoot = &cobra.Command{
		Use:     "browser-sync",
		Long:    "browser-sync",
		Version: common.Version,
	}
	cmdInit = &cobra.Command{
		Use:   "init",
		Short: "init configuration file",
		Long:  "init",
		Run:   initialization,
	}
	cmdRun = &cobra.Command{
		Use:   "run",
		Short: "run listen",
		Long:  "run",
		Run:   run,
	}
)

func main() {
	cmdRun.Flags().StringVarP(&config, "config", "c", "./config.yaml", "configuration file")

	cmdRoot.AddCommand(cmdInit)
	cmdRoot.AddCommand(cmdRun)

	if err := cmdRoot.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}
