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

	listDepsCmd.Flags().StringP("name", "n", "", "list dependencies for the given distro name")
}

var listDepsCmd = &cobra.Command{
	Use:   "listDeps",
	Short: "List dependencies for a given distro name.",
	Long:  `Lists all runtime required dependencies for a give CPAN perl distribution name.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		//log.Println("stub: list dependencies here....")

		//log.Printf("MDM.BHooksEndOfScope:  %v", dlist.MDM.BHooksEndOfScope)
		//log.Printf("CoreModulesMap[CPAN::Complete]: %v", dlist.CoreModulesMap["CPAN::Complete"])

		//fmt.Println(result["users"])

		//log.Printf("Result[author]: %v", dlist.Result["author"])

		/*
			log.Printf("metadata file count: %v", len(dlist.DistroMetaDataFilepathMap))

			for k, v := range dlist.DistroMetaDataFilepathMap {
				log.Printf("k: %v v: %v", k, v)
			}
		*/

		for k, v := range dlist.DistroMetaDataMap {

			switch k {
			case "DateTime":
				//var r = make(map[string]interface{})
				r := v["prereqs"].(map[string]interface{})
				s := r["runtime"].(map[string]interface{})
				t := s["requires"].(map[string]interface{})
				//log.Printf("DateTime author: %v", v["prereqs"])
				//log.Printf("DateTime prereqs configure requires: %v", s["requires"])
				//log.Printf("DateTime prereqs configure requires: %v", t)

				for k, _ := range t {
					log.Printf("required module: %v", k)
				}

			default:
			}

		}

		return nil
	},
}
