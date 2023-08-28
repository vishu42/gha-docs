/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/vishu42/gha-docs/cmd/impl"
)

type CommandGroup struct{}

func (c *CommandGroup) RootCmd() *cobra.Command {
	o := &impl.RootOpts{}

	rootCmd := &cobra.Command{
		Use:   "gha-docs",
		Short: "gha-docs is a cli tool to generate documentation for github actions",
		Run: func(cmd *cobra.Command, args []string) {
			impl.RunRoot(cmd, args, o)
		},
		PreRun: func(cmd *cobra.Command, args []string) {
			// get debug flag
			debug, err := cmd.Flags().GetBool("debug")
			if err != nil {
				fmt.Println("error getting debug flag")
			}
			o.Debug = debug

			// get action file flag
			actionFile, err := cmd.Flags().GetString("action-yaml")
			if err != nil {
				fmt.Println("error getting action file flag")
			}

			o.ActionFile = actionFile
		},
	}
	// set debug flag
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "enable debug mode")

	// set action file flag
	rootCmd.PersistentFlags().StringP("action-yaml", "f", "action.yml", "action file name")
	return rootCmd
}

func (c *CommandGroup) All() *cobra.Command {
	root := c.RootCmd()
	return root
}
