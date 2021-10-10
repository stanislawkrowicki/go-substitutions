package tools

import (
	"github.com/antchfx/htmlquery"
	httpapi "go-substitutions/pkg/http-api"
)

type Substitutions struct {
	Date    string   `json:"date"`
	Changes []string `json:"changes"`
}

func GetSubstitutions(date string) (*Substitutions, error) {
	apiResponse, err := httpapi.GetResponse(date)
	if err != nil {
		return nil, err
	}

	nodes, err := ExtractChanges(apiResponse)
	if err != nil {
		return nil, err
	}
	if nodes == nil {
		return nil, nil
	}

	var substitutions Substitutions
	substitutions.Date = date

	for _, node := range *nodes {
		substitutions.Changes = append(substitutions.Changes, htmlquery.InnerText(node))
	}

	return &substitutions, nil
}
