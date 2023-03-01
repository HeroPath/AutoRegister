package src

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tealeg/xlsx"
	"io/ioutil"
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

func PostRequest(url string, data interface{}, token string) (interface{}, int) {
	payload, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error al convertir los datos a JSON:", err)
		return nil, -1
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
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al hacer la petición:", err)
		return nil, -1
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer el cuerpo de la respuesta:", err)
		return nil, -1
	}

	if len(body) == 0 {
		return resp.StatusCode, -1
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		fmt.Println("Error al decodificar la respuesta JSON:", err)
		return nil, -1
	}

	return result, resp.StatusCode
}
