package day4

import (
	"advent/internal/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day4() int {
	filePath := "/home/henrik/source/repos/advent-of-code-2023/inputs/4.txt"

	file, err := os.Open(filePath)
	utils.IfErrorPrintAndExit("Error reading input data", err)
	scanner := bufio.NewScanner(file)
	totalPoints := 0
	for scanner.Scan() {
		s := scanner.Text()[10:]
		ticketsSubstring := s[:29]
		winningSubstring := s[32:]
		ticketNumbers := extractTicketNumbers(ticketsSubstring)
		winningNumbers := extractTicketNumbers(winningSubstring)

		wins := 0
		for _, ticket := range ticketNumbers {
			for _, winning := range winningNumbers {
				if ticket == winning {
					wins += 1
					break
				}
			}
		}
		points := 0
		if wins > 0 {
			points = 1 << (wins - 1)
		}
		// fmt.Printf("%s %v | %v | %d\n", scanner.Text()[:10], ticketNumbers, winningNumbers, points)
		totalPoints += points
	}
	fmt.Printf("%d\n", totalPoints)
	return totalPoints
}

func extractTicketNumbers(input string) []int {
	s := input
	res := make([]int, 0)
	for len(s) > 1 {
		i, err := strconv.Atoi(strings.Trim(s[:2], " "))
		utils.IfErrorPrintAndExit("Unable to read ticker number from string:"+input, err)
		res = append(res, i)

		if len(s) < 3 {
			s = ""
		} else {
			s = s[3:]
		}
	}
	return res
}
