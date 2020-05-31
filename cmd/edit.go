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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit a task",
	Long:  `Edit a task with any changes`,
	Run: func(cmd *cobra.Command, args []string) {
		var newTasks []Task

		fmt.Println("Task-num: ", viper.GetString("task-num"))

		if viper.GetString("task-num") != "" {
			if viper.GetString("task") != "" {
				taskInfo := strings.Split(viper.GetString("task"), ",")
				tasks := allTasks()
				taskNumToEdit, _ := strconv.ParseInt(viper.GetString("task-num"), 10, 64)

				for _, task := range tasks {
					if int64(task.Num) == taskNumToEdit {
						task.Date = taskInfo[0]
						task.Priority = taskInfo[1]
						task.Name = taskInfo[2]
					}
					newTasks = append(newTasks, task)
				}

				if err := updateTasks(newTasks); err != nil {
					fmt.Println(err)
				}
				fmt.Printf("Task #%d edited successfully.\n", taskNumToEdit)
			} else {
				fmt.Println("Specified empty change.")
			}
		} else {
			fmt.Println("Need to specify task number to edit task.")
		}
	},
}

func init() {
	editCmd.Flags().StringP("task-num", "n", "", "Task number to be edited")
	viper.BindPFlag("task-num", editCmd.Flags().Lookup("task-num"))
	editCmd.Flags().StringP("task", "t", "", "Updated task (format: 'DD/MM/YYYY,high|medium|low,Task name)'")
	viper.BindPFlag("task", editCmd.Flags().Lookup("task"))
	rootCmd.AddCommand(editCmd)
}
