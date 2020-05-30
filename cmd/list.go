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
package cmd

import (
  "text/template"
	"github.com/spf13/cobra"
  "os"
)

const listTmpl = `
#   Priority          Date             Name
------------------------------------------------------------
{{ range $val := . }}
{{$val.Num}}   {{$val.Priority}}          {{$val.Date}}             {{$val.Name}}
{{end}}
`

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the tasks",
	Long: `List all the tasks (completed as well)`,
	Run: func(cmd *cobra.Command, args []string) {
    tasks := []Task {
      { Num: 1, Name: "This is first task", Priority: "high", Date: "01/05/2020" },
      { Num: 2, Name: "This is second task", Priority: "medium", Date: "10/05/2020" },
      { Num: 3, Name: "This is third task", Priority: "low", Date: "21/05/2020" },
    }

    // display list of tasks
    t := template.Must(template.New("listTmpl").Parse(listTmpl))
    t.Execute(os.Stdout, tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
