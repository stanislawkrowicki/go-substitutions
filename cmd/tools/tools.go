package tools

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	DayThresholdHour      = 15
	CantReadResponseBody  = "failed to read the response body"
	CantUnmarshalResponse = "failed to unmarshal response body"
	CantParseHTML         = "failed to parse HTML from the response"
)

type Response struct {
	R string `json:"r"`
}

func GetRequestDate() string {
	date := time.Now()
	if date.Hour() >= DayThresholdHour {
		date.AddDate(0, 0, 1)
	}

	return date.Format("2006-01-02")
}

func UnmarshalResponse(resp *http.Response) (*Response, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.New(CantReadResponseBody)
	}

	parsedResponse := Response{}

	err = json.Unmarshal(body, &parsedResponse)
	if err != nil {
		return nil, errors.New(CantUnmarshalResponse)
	}

	return &parsedResponse, nil
}

func ExtractChanges(apiResponse *http.Response) (*[]*html.Node, error) {
	parsedResp, err := UnmarshalResponse(apiResponse)
	if err != nil {
		return nil, err
	}

	doc, err := htmlquery.Parse(strings.NewReader(parsedResp.R))
	if err != nil {
		return nil, errors.New(CantParseHTML)
	}

	class := os.Getenv("CLASS")
	xpath := fmt.Sprintf("/html/body/div/div/div/div/span[text() = '%s']/../../div[2]/*", class)
	list := htmlquery.Find(doc, xpath)

	if len(list) == 0 {
		return nil, nil
	}
	return &list, nil
}
