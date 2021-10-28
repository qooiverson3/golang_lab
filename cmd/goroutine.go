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
	"fmt"
	goroutinewg "golang_lab/pkg/goroutine_wg"
	"sync"

	"github.com/spf13/cobra"
)

var amount int

var goroutineCmd = &cobra.Command{
	Use:   "goroutine",
	Short: "以協程方式，快速取得多個使用者回應",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		if amount != 0 && amount > 0 {
			user := goroutinewg.User{}
			wg := &sync.WaitGroup{}
			wg.Add(amount)

			for i := 0; i < amount; i++ {
				user.ID = i
				go user.Says(wg)
			}
			wg.Wait()
			fmt.Println("Hello everyone!!")
			return
		}
		fmt.Println("illegal")
	},
}

func init() {
	rootCmd.AddCommand(goroutineCmd)
	goroutineCmd.Flags().IntVarP(&amount, "amount", "m", amount, "指定的數量")
}
