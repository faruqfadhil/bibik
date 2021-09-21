/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

// execCmd represents the exec command
var execCmd = &cobra.Command{
	Use:   "exec",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cliHandler := initContainer()
		out, err := cliHandler.Exec(context.Background(), keyToBeExec, useDirr)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		fmt.Println(out)
		return nil
	},
}

var keyToBeExec string
var useDirr bool

func init() {
	rootCmd.AddCommand(execCmd)
	execCmd.Flags().StringVarP(&keyToBeExec, "key", "k", "", "key to be get command to execution")
	execCmd.MarkFlagRequired("key")
	execCmd.Flags().BoolVar(&useDirr, "use-dirr", false, "exec command in specified dir, specified dir is added when you run bibik save")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// execCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// execCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
