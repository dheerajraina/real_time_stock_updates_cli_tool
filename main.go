package main

import (
	"fmt"
	"os/exec"
)

func main() {

	arg1:="INFY"

	cmd := exec.Command("python3", "web_scraper.py",arg1)

	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing Python script", err)
		return
	}
	result := string(output)
	fmt.Println("Result from Python script", result)
}
