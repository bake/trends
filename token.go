package trends

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type tokenWidget struct {
	ID    string `json:"id"`
	Type  string `json:"widgetType"`
	Title string `json:"title"`
	Token string `json:"token"`
}

type tokenComparisonItem struct {
	Geo     string `json:"geo"`
	Time    string `json:"time"`
	Keyword string `json:"keyword"`
}

type tokenComparisonRequest struct {
	Category       int                   `json:"category"`
	Property       string                `json:"property"`
	ComparisonItem []tokenComparisonItem `json:"comparisonItem"`
}

type tokenComparisonResponse struct {
	Widgets  []tokenWidget `json:"widgets"`
	Keywords []struct {
		Keyword string `json:"keyword"`
		Name    string `json:"name"`
		Type    string `json:"type"`
	} `json:"keywords"`
}

func token(id string, keywords ...string) (string, error) {
	js := tokenComparisonRequest{}
	for _, keyword := range keywords {
		js.ComparisonItem = append(js.ComparisonItem, tokenComparisonItem{
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

	res := tokenComparisonResponse{}
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
