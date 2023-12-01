package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


var GPT_ADDRESS string = "http://127.0.0.1:8080/message"
var GPT_MODEL   string = "gpt-3.5-turbo"
var MAX_TOKEN   string = "100"

func Send_gpt_message (messages *[]map[string]string) error {
	request := map[string][]map[string]string{}
	request["info"] = []map[string]string{{"model" : GPT_MODEL, "max_token": MAX_TOKEN}}
	request["messages"] = *messages

	json_data, err := json.Marshal(request)

    if err != nil {
        return err
    }

	response, err := http.Post(GPT_ADDRESS, "application/json", bytes.NewBuffer(json_data))

	if err != nil {
        return err
    }
	defer response.Body.Close()

	var data map[string]string

    body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	st := []byte(body)
	json.Unmarshal(st, &data);
	// fmt.Println(st)
	// fmt.Println(data)
	*messages = append(*messages, data)
	return nil
}

func Get_gpt_message(messages *[]map[string]string) (string, error) {
	if len(*messages) == 0 {
		return "", fmt.Errorf("Нет сообщений")
	}

	if (*messages)[len(*messages) - 1]["role"] != "assistant" {
		return "", fmt.Errorf("Последнее сообщение не бота.")
	}

	return (*messages)[len(*messages) - 1]["content"], nil	
}

func Add_user_message(messages *[]map[string]string, mes string) {
	*messages = append(*messages, map[string]string{"role" : "user", "content" : mes})
}