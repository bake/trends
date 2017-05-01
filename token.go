package trends

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type TokenWidget struct {
	ID    string `json:"id"`
	Type  string `json:"widgetType"`
	Title string `json:"title"`
	Token string `json:"token"`
}

type TokenComparisonItem struct {
	Geo     string `json:"geo"`
	Time    string `json:"time"`
	Keyword string `json:"keyword"`
}

type TokenComparisonRequest struct {
	Category       int                   `json:"category"`
	Property       string                `json:"property"`
	ComparisonItem []TokenComparisonItem `json:"comparisonItem"`
}

type TokenComparisonResponse struct {
	Widgets  []TokenWidget `json:"widgets"`
	Keywords []struct {
		Keyword string `json:"keyword"`
		Name    string `json:"name"`
		Type    string `json:"type"`
	} `json:"keywords"`
}

func Token(id string, keywords ...string) (string, error) {
	js := TokenComparisonRequest{}
	for _, keyword := range keywords {
		js.ComparisonItem = append(js.ComparisonItem, TokenComparisonItem{
			Keyword: keyword,
			Time:    "today 5-y",
		})
	}
	req, err := json.Marshal(js)
	if err != nil {
		return "", err
	}

	body, err := Get("api/explore", url.Values{
		"hl":  {"de"},
		"tz":  {"-120"},
		"req": {string(req)},
	})
	if err != nil {
		return "", err
	}

	res := TokenComparisonResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return "", err
	}

	for _, w := range res.Widgets {
		if w.ID == id {
			return w.Token, nil
		}
	}

	return "", fmt.Errorf("cound not get token with id %s", id)
}
