/*
Copyright © 2021 Berk Safran

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
	"go-redis-example/normal"

	"github.com/spf13/cobra"
)

// normalCmd represents the normal command
var normalCmd = &cobra.Command{
	Use:   "normal",
	Short: "Run with Normal Mode including GET, SET, EXISTS examples",
	Run: func(cmd *cobra.Command, args []string) {
		normal.RunNormalMode()
	},
}

func init() {
	rootCmd.AddCommand(normalCmd)
}
