package trends

import (
	"fmt"
	"log"
)

func ExampleInterestOverTime() {
	iot, err := InterestOverTime("gorillaz")
	if err != nil {
		log.Fatal(err)
	}
	for time, value := range iot {
		fmt.Println(time, value)
	}
	// Output: 2012-05-27 02:00:00 +0200 CEST 28
	// 2012-06-03 02:00:00 +0200 CEST 27
	// …
}
