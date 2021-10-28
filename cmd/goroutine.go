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
	Short: "以多線程方式，快速取得多個使用者回應",
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
