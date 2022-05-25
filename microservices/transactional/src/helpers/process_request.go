package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func ProcessRequest(method, url, token, id string, payload, output interface{}) error {
	var res *http.Response
	var err error
	var buff []byte
	client := &http.Client{}

	buff, err = json.Marshal(payload)
	if err != nil {
		fmt.Println("Error in Marshal payload", err.Error())
		return err
	}

	req, _ := http.NewRequest(method, url, bytes.NewReader(buff))

	if method == "GET" {
		req.Header.Set("session_token", token)
		req.Header.Set("bank_account_id", id)
	}

	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	res, err = client.Do(req)

	if err != nil {
		fmt.Println("Error on request", err.Error())
		return err
	}

	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(output); err != nil {
		fmt.Println("Error in Decode response", err.Error())
	}
	return err
}
