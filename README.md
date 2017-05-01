# go-trends

Google Trends library for Go

```go
package main

import (
	"fmt"
	"log"

	"git.192k.pw/bake/go-trends"
)

func main() {
	iot, err := trends.InterestOverTime("gorillaz")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range iot.Default.Timeline {
		fmt.Println(item.Time, item.Value[0])
	}
}
```

```
2012-05-06 02:00:00 +0200 CEST 28
2012-05-13 02:00:00 +0200 CEST 28
2012-05-20 02:00:00 +0200 CEST 28
2012-05-27 02:00:00 +0200 CEST 28
2012-06-03 02:00:00 +0200 CEST 27
…
```
