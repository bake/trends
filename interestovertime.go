package trends

import (
	"encoding/json"
	"fmt"
	"net/url"
	"sort"
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
		Timeline iotTimelineItems `json:"timelineData"`
	} `json:"default"`
}

type iotTimelineItem struct {
	Time  timestamp `json:"time"`
	Value []int     `json:"value"`
}

type iotTimelineItems []iotTimelineItem

func (is iotTimelineItems) Len() int           { return len(is) }
func (is iotTimelineItems) Swap(i, j int)      { is[i], is[j] = is[j], is[i] }
func (is iotTimelineItems) Less(i, j int) bool { return is[i].Time.Unix() < is[j].Time.Unix() }

type IotResult struct {
	Time   time.Time
	Values []int
}

func InterestOverTime(keywords ...string) ([]IotResult, error) {
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

	resp := iotResponse{}
	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, fmt.Errorf("could not unmarshal response: %v", err)
	}
	sort.Sort(resp.Default.Timeline)

	res := []IotResult{}
	for _, i := range resp.Default.Timeline {
		res = append(res, IotResult{i.Time.Time, i.Value})
	}

	return res, nil
}
