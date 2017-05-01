package trends

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type IOTRequest struct {
	Time            string              `json:"time"`
	Resolution      string              `json:"resolution"`
	Locale          string              `json:"locale"`
	ComparisonItems []IOTComparisonItem `json:"comparisonItem"`
	Options         IOTOptions          `json:"requestOptions"`
}

type IOTComparisonItem struct {
	Geo         struct{} `json:"geo"`
	Restriction struct {
		Keywords []IOTKeyword `json:"keyword"`
	} `json:"complexKeywordsRestriction"`
}

type IOTKeyword struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type IOTOptions struct {
	Property string `json:"property"`
	Backend  string `json:"backend"`
	Category int    `json:"category"`
}

type IOTResponse struct {
	Default struct {
		Timeline []IOTTimelineItem `json:"timelineData"`
	} `json:"default"`
}

type IOTTimelineItem struct {
	Time  JSONTime `json:"time"`
	Value []int    `json:"value"`
}

func InterestOverTime(keywords ...string) (*IOTResponse, error) {
	token, err := Token(MethodInterestOverTime, keywords...)
	if err != nil {
		return nil, err
	}

	items := make([]IOTComparisonItem, len(keywords))
	for i, keyword := range keywords {
		items[i] = IOTComparisonItem{}
		items[i].Restriction.Keywords = make([]IOTKeyword, 1)
		items[i].Restriction.Keywords[0] = IOTKeyword{
			Type:  "BROAD",
			Value: keyword,
		}
	}

	options := IOTOptions{
		Backend: "IZG",
	}

	req := IOTRequest{
		Time:            "2012-05-01 2017-05-01",
		Resolution:      "WEEK",
		Locale:          "de",
		ComparisonItems: items,
		Options:         options,
	}

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	body, err := Get("api/widgetdata/multiline", url.Values{
		"hl":    {"de"},
		"tz":    {"-120"},
		"token": {token},
		"req":   {string(payload)},
	})
	if err != nil {
		return nil, err
	}

	res := IOTResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}
	return &res, nil
}
