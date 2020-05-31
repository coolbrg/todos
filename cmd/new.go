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
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Add a new task",
	Long:  `Add a new task to the existing list of tasks`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("task") != "" {
			taskInfo := strings.Split(viper.GetString("task"), ",")
			tasks := allTasks()
			newTaskNum := lastTaskNum() + 1
			newTasks := []Task{
				{Num: newTaskNum, Done: false, Date: taskInfo[0], Priority: taskInfo[1], Name: taskInfo[2]},
			}

			for _, task := range tasks {
				newTasks = append(newTasks, task)
			}

			if err := updateTasks(newTasks); err != nil {
				fmt.Println(err)
			}
			fmt.Printf("Task #%d added successfully.\n", newTaskNum)
		} else {
			fmt.Println("Task entry is empty.")
		}
	},
}

func init() {
	newCmd.Flags().StringP("task", "", "", "Add new task (format: 'DD/MM/YYYY,high|medium|low,Task name)'")
	viper.BindPFlag("task", newCmd.Flags().Lookup("task"))
	rootCmd.AddCommand(newCmd)
}
