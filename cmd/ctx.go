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

	"github.com/spf13/cobra"
)

// ctxCmd represents the ctx command
var ctxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "context lab",
	Long:  `context lab`,
	Run: func(cmd *cobra.Command, args []string) {

		list := []string{
			"8.8.8.8",
			"5.5.5.5",
			"4.4.4.4",
			"3.3.3.3",
			"1.1.1.1",
			"7.7.7.7",
		}
		amount := 3
		ch_data := make(chan []string, 1)
		data := []string{}
		ch_data <- data

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		for i := 0; i < len(list); i++ {
			go func(c context.Context, n string) {
				select {
				case <-c.Done():
					return
				default:
					d := <-ch_data
					if len(d) == amount {
						ch_data <- d
						return
					}
					d = append(d, n)
					ch_data <- d
				}
			}(ctx, list[i])
		}

		for {
			r := <-ch_data
			if len(r) == amount {
				cancel()
				ch_data <- r
				break
			}
			ch_data <- r
		}

		result := <-ch_data
		if len(result) == amount {
			fmt.Println(result)
		} else {
			fmt.Println("error")
		}

	},
}

func init() {
	rootCmd.AddCommand(ctxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ctxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ctxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
