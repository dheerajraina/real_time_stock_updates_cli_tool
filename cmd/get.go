/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		for i := 0; i < int(math.Pow(10, 10)); i++ {

			resultChan := make(chan string)
			go func() {
				result := getStockPrice(args[0])
				resultChan <- *result
				close(resultChan)
			}()

			// Wait for the update call to complete
			result := <-resultChan

			clearTerminal()
			dataArray := strings.Split(result, ",")
			basePrice := dataArray[0]
			priceChange := dataArray[1]
			percentPriceChange := dataArray[2]
			lastPrice := dataArray[3]
			previousClose := dataArray[4]

			fmt.Printf("Symbol(%s)\tBase_Price(%s)\tPrice_Change(%s)\tChange_Percent(%s)\tLast_Price(%s)\tPrevious_Close(%s)\n", args[0], basePrice, priceChange, percentPriceChange, lastPrice, previousClose)
			time.Sleep(time.Second * 1)
		}

	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func getStockPrice(stockSymbol string) *string {

	arg1 := stockSymbol

	cmd := exec.Command("python3", "web_scraper.py", arg1)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing Python script", err)
		return nil
	}
	result := string(output)
	return &result

}

func clearTerminal() {
	cmd := exec.Command("clear") // Use "clear" for Unix-like systems, "cls" for Windows
	cmd.Stdout = os.Stdout
	cmd.Run()
}
