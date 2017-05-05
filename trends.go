package trends

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
)

const (
	MethodInterestOverTime = "TIMESERIES"
	MethodInterestByRegion = "GEO_MAP"
	MethodRelatedTopics    = "RELATED_TOPICS"
	MethodRelatedQueries   = "RELATED_QUERIES"
)

var (
	base = "https://trends.google.com/trends/"
)

func Get(method string, params url.Values) ([]byte, error) {
	body := []byte{}
	req, err := http.NewRequest(http.MethodGet, base+method, nil)
	if err != nil {
		return body, err
	}
	req.URL.RawQuery = params.Encode()
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/57.0.2987.133 Safari/537.36")

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return body, err
	}
	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return body, err
	}
	return trim(body), nil
}

func trim(body []byte) []byte {
	// for some reason the string starts with garbage
	re := regexp.MustCompile("^([^{]+)")
	return re.ReplaceAll(body, []byte(""))
}
