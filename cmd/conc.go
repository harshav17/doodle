package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(concCommand)
}

var concCommand = &cobra.Command{
	Use:   "conc",
	Short: "simple program to test computer's cores",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Get ready for some heat")
		testCores(10)
	},
}

func testCores(n int) {
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for i := 0; i < n; i++ {
		go worker(jobs, results)
	}

	for i := 0; i < 100; i++ {
		jobs <- i
	}
	close(jobs)

	for j := 0; j < 100; j++ {
		fmt.Println(<-results)
	}
}

func worker(jobs <-chan int, results chan<- int) {
	for n := range jobs {
		results <- fib(n)
	}
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
