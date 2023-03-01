package src

import (
	"fmt"
	"github.com/tealeg/xlsx"
)

func readFile(fileName string) []*xlsx.Row {
	file, err := xlsx.OpenFile("./data/" + fileName + ".xlsx")
	if err != nil {
		fmt.Println("Error in open file", err)
		return nil
	}

	sheet := file.Sheet["Hoja1"]
	return sheet.Rows
}

func PostRequest() {

}
