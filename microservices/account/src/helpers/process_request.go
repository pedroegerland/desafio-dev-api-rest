package helpers

import (
	"banktest_account/src/entity"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func ProcessRequest(path string, payload, output interface{}) error {
	buff, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error in Marshal payload", err.Error())
		return err
	}

	res, err := http.Post(entity.URL+path, "application/json", bytes.NewBuffer(buff))
	if err != nil {
		fmt.Println("Error on request", err.Error())
		return err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(output); err != nil {
		fmt.Println("Error in Decode response", err.Error())
		return err
	}
	return err
}
