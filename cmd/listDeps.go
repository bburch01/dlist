// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/bburch01/dlist/internal/app/dlist"
	"github.com/spf13/cobra"
)

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkServiceAlivenessCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkServiceAlivenessCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.AddCommand(listDepsCmd)

	listDepsCmd.Flags().StringSliceP("name", "n", nil, "list dependencies for the given distro name")
}

var listDepsCmd = &cobra.Command{
	Use:   "listDeps",
	Short: "List dependencies for a given distro name.",
	Long:  `Lists all runtime required dependencies for a give CPAN perl distribution name.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		distroList, _ := cmd.Flags().GetStringSlice("name")

		if json, err := dlist.GetDepList(distroList); err == nil {
			log.Print(json)
		} else {
			log.Printf("err: %v", err)
		}

		return nil
	},
}
