package trends

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type iotRequest struct {
	Time            string              `json:"time"`
	Resolution      string              `json:"resolution"`
	Locale          string              `json:"locale"`
	ComparisonItems []iotComparisonItem `json:"comparisonItem"`
	Options         iotOptions          `json:"requestOptions"`
}

type iotComparisonItem struct {
	Geo         struct{} `json:"geo"`
	Restriction struct {
		Keywords []iotKeyword `json:"keyword"`
	} `json:"complexKeywordsRestriction"`
}

type iotKeyword struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type iotOptions struct {
	Property string `json:"property"`
	Backend  string `json:"backend"`
	Category int    `json:"category"`
}

type iotResponse struct {
	Default struct {
		Timeline []iotTimelineItem `json:"timelineData"`
	} `json:"default"`
}

type iotTimelineItem struct {
	Time  jsonTime `json:"time"`
	Value []int    `json:"value"`
}

func InterestOverTime(keywords ...string) (map[time.Time]int, error) {
	token, err := token(MethodInterestOverTime, keywords...)
	if err != nil {
		return nil, err
	}

	items := make([]iotComparisonItem, len(keywords))
	for i, keyword := range keywords {
		items[i] = iotComparisonItem{}
		items[i].Restriction.Keywords = make([]iotKeyword, 1)
		items[i].Restriction.Keywords[0] = iotKeyword{
			Type:  "BROAD",
			Value: keyword,
		}
	}

	options := iotOptions{
		Backend: "IZG",
	}

	end := time.Now()
	start := end.AddDate(-5, 0, 0)
	req := iotRequest{
		Time:            fmt.Sprintf("%s %s", start.Format("2006-01-02"), end.Format("2006-01-02")),
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

	res := iotResponse{}
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}

	list := map[time.Time]int{}
	for _, item := range res.Default.Timeline {
		list[item.Time.Time] = item.Value[0]
	}

	return list, nil
}
