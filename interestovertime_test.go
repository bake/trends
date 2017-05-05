package trends

import (
	"fmt"
	"log"

	"github.com/BakeRolls/trends"
)

func ExampleInterestOverTime() {
	iot, err := trends.InterestOverTime("gorillaz")
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range iot.Default.Timeline {
		fmt.Println(item.Time, item.Value[0])
	}
	// Output: 2012-05-27 02:00:00 +0200 CEST 28
	// 2012-06-03 02:00:00 +0200 CEST 27
	// …
}
