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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// removeCmd represents the remove command
var removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a task",
	Long:  `Remove the task from the list`,
	Run: func(cmd *cobra.Command, args []string) {
		var newTasks []Task

		if viper.GetString("task-num") != "" {
			tasks := allTasks()
			taskNum, _ := strconv.ParseInt(viper.GetString("task-num"), 10, 64)
			for _, task := range tasks {
				if int64(task.Num) != taskNum {
					newTasks = append(newTasks, task)
				}
			}

			if err := updateTasks(newTasks); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Task #%d added successfully.\n", taskNum)
		} else {
			fmt.Println("Need to specify task number to remove any task.")
		}
	},
}

func init() {
	removeCmd.Flags().StringP("task-num", "r", "", "Remove a task by specifying its number")
	viper.BindPFlag("task-num", removeCmd.Flags().Lookup("task-num"))
	rootCmd.AddCommand(removeCmd)
}
