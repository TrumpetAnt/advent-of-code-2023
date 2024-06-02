package main

import (
	"advent/internal"
	"advent/internal/day4"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func main() {
	if len(os.Args) == 3 && os.Args[1] == "download" {
		internal.DownloadInputAndInstructions(validateDay(os.Args[2]))
		os.Exit(0)
	}

	if len(os.Args) == 3 && os.Args[1] == "solve" {
		switch validateDay(os.Args[2]) {
		case "1":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "2":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "3":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "4":
			day4.Day4()
			return
		case "5":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "6":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "7":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "8":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "9":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "10":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "11":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "12":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "13":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "14":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "15":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "16":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "17":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "18":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "19":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "20":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "21":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "22":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "23":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "24":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		case "25":
			fmt.Println("Day not yet implemented")
			os.Exit(1)
		default:
			fmt.Println("days only between 1-25")
			os.Exit(1)
		}
	}

	fmt.Println("Usage:")
	prefix := "    " + filepath.Base(os.Args[0]) + " "
	fmt.Println(prefix + "download [day]")
	fmt.Println(prefix + "solve [day]")
	os.Exit(1)
}

func validateDay(day string) string {
	i, err := strconv.ParseInt(day, 10, 8)
	if err != nil || i < 1 || i > 25 {
		fmt.Println("Error: day must be number between 1-25")
		os.Exit(1)
	}
	return day
}
