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

	"github.com/faruqfadhil/bibik/internal/entity"
	"github.com/spf13/cobra"
)

// saveCmd represents the save command
var saveCmd = &cobra.Command{
	Use:   "save",
	Short: "Save data",
	Long:  `Tell bibik to remember the data`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cliHandler := initContainer()
		payload := &entity.Command{
			Key:   key,
			Value: value,
			Options: &entity.Options{
				Dir: dir,
			},
		}
		err := cliHandler.UpsertCommand(context.Background(), payload)
		if err != nil {
			fmt.Println(err.Error())
			return err
		}
		fmt.Println(" done")
		return nil
	},
}

var key string
var value string
var dir string

func init() {
	rootCmd.AddCommand(saveCmd)
	saveCmd.Flags().StringVarP(&key, "key", "k", "", "set key data")
	saveCmd.MarkFlagRequired("key")
	saveCmd.Flags().StringVarP(&value, "value", "v", "", "set value")
	saveCmd.MarkFlagRequired("value")
	saveCmd.Flags().StringVarP(&dir, "dir", "d", "", "set the directory where the saved command will be exec (using bibik exec --usedir command)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// saveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// saveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
