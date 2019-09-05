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

	rootCmd.AddCommand(listDependenciesCmd)

	listDependenciesCmd.Flags().StringP("name", "n", "", "list all dependencies for the given distro name")
}

var listDependenciesCmd = &cobra.Command{
	Use:   "listDependencies",
	Short: "List all dependencies for a given distro name.",
	Long:  `Lists all runtime required dependencies for a give CPAN perl distribution name.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		log.Println("stub: list dependencies here....")

		/*
			chkall, _ := cmd.Flags().GetBool("all")
			if chkall {
				checkAll()
			} else {
				svcname, _ := cmd.Flags().GetString("name")
				if svcname == "" {
					checkAll()
				} else {
					resp, err := checkByName(svcname)
					if err != nil {
						log.Printf("%v service health check call failed with error: %v", svcname, err)
					} else {
						log.Printf("%v service health check response code   : %v", svcname, resp.Details.Code)
						log.Printf("%v service health check response message: %s", svcname, resp.Details.Message)
					}
				}
			}
		*/

		return nil
	},
}
