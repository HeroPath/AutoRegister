package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"net/http"
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

func PostRequest(url string, data interface{}, token string) string {
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir los datos a JSON:", err)
		return ""
	}
	if data == "" {
		payload = []byte(`""`)
	}

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if token != "" {
		req.Header.Set("Authorization", "Bearer "+token)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	return resp.Status

}
