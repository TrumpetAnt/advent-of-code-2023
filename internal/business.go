package internal

import (
	"advent/internal/facade"
	"advent/internal/utils"
	"io"
	"os"
)

func DownloadInputAndInstructions(day string) {

	filePath := "/home/henrik/source/repos/advent-of-code-2023/inputs/" + day + ".txt"
	_, err := os.Open(filePath)
	if err != nil && err.Error() == "open "+filePath+": no such file or directory" {
		newFile, err := os.Create(filePath)
		utils.IfErrorPrintAndExit("Unexpected error attempting to create file for input", err)
		response := facade.RequestInputData(day)

		data, err := io.ReadAll(response.Body)
		utils.IfErrorPrintAndExit("Unexpected error attempting to read response body", err)
		newFile.Write(data)
		newFile.Close()
		DownloadInputAndInstructions(day)
	} else {
		utils.IfErrorPrintAndExit("Unexpected error attempting to open file for input", err)
	}

}
