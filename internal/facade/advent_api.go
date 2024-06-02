package facade

import (
	"advent/internal/utils"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const sessionCookie = "53616c7465645f5f525edceba92dae50eccbe3c178aecc19de67f947dfb539a01a55e87afe4cc80cc8bae91d4343e044e1760ef0681d038bd57ffc51d7f1ad7f"

func RequestInputData(day string) *http.Response {
	client := http.Client{}
	urlString := fmt.Sprintf("https://adventofcode.com/2023/day/%s/input", day)
	req, err := http.NewRequest("GET", urlString, nil)
	utils.IfErrorPrintAndExit("Error creating http request", err)

	addAuth(req)

	fmt.Println("Performing http request to " + urlString)
	resp, err := client.Do(req)
	utils.IfErrorPrintAndExit("Error while performing http request", err)
	if resp.StatusCode != 200 {
		content, _ := io.ReadAll(resp.Body)
		fmt.Printf(string(content) + "\n")
		utils.IfErrorPrintAndExit("Error while performing http request", errors.New("unexpected status code, wanted 200, got "+strconv.Itoa(resp.StatusCode)))
	}
	return resp
}

func addAuth(req *http.Request) {
	c := http.Cookie{
		Name:  "session",
		Value: sessionCookie,
	}
	req.AddCookie(&c)
}
