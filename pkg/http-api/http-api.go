package http_api

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

const (
	CantMarshalBody = "failed to marshal request body"
	CantGetFromApi  = "failed to get data from edupage api"
)

func GetResponse(date string) (*http.Response, error) {
	args := []map[string]interface{}{
		nil,
		{
			"date": date,
			"mode": "classes",
		},
	}
	gsh := "00000000"

	requestBody, err := json.Marshal(map[string]interface{}{
		"__args": args,
		"__gsh":  gsh,
	})
	if err != nil {
		return nil, errors.New(CantMarshalBody)
	}

	resp, err := http.Post(os.Getenv("SUBSTITUTIONS_API"), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New(CantGetFromApi)
	}

	return resp, nil
}
