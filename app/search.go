package app

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"stardust.db/schema"
)

type Search struct {
}

func NewSearch() *Search {
	return &Search{}
}

func (s *Search) Stardust(name string) string {
	return fmt.Sprintf("Hello %s, Welcome to the Star Realm", name)
}

func (s *Search) GetResourceParams(resource string) []string {
	order := []string{"code", "client_name", "shipping_address.name", "status", "order_type"}
	arn := []string{"code", "client_name", "customer.name", "status"}

	switch resource {
	case "arn":
		return arn
	case "order":
		return order
	default:
		return []string{}
	}
}

func (s *Search) LocalSearch(query, index, clientID, params string) (*schema.LocalSearchResponse, error) {
	query = strings.ReplaceAll(query, " ", "%20")
	url := fmt.Sprintf("https://stageapi.leanafywms.com/core/search/local?query=%s&index=%s&clientID=%s&params=%s", query, index, clientID, params)
	fmt.Println(url)
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "")
	req.Header.Add("platform_id", "wms_central")
	// req.Header.Add("User-Agent", "Apidog/1.0.0 (https://apidog.com)")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var results schema.LocalSearchResponse

	err = json.Unmarshal(body, &results)
	if err != nil {
		return nil, err
	}
	return &results, nil
}
