package cmd

/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

import (
	flag "github.com/spf13/pflag"

	"github.com/spf13/cobra"
)

var (
	tasks    []Task
	skipDone bool
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks",
	Long:  `List all the tasks (completed as well)`,
	Run: func(cmd *cobra.Command, args []string) {
		if !skipDone {
			showTasks(allTasks())
		} else {
			// fetch incomplete tasks
			var incompleteTasks []Task
			for _, task := range allTasks() {
				if !task.Done {
					incompleteTasks = append(incompleteTasks, task)
				}
			}

			showTasks(incompleteTasks)
		}
	},
}

var (
	skipDoneFlag = &flag.Flag{
		Name:      "skip-done",
		Shorthand: "s",
		Usage:     "Skip the completed tasks",
	}
)

func init() {
	// listCmd.Flags().AddFlag(skipDoneFlag)
	// viper.BindPFlags(listCmd.Flags())

	listCmd.Flags().BoolVarP(&skipDone, "skip-done", "s", false, "Skip the completed tasks")
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
