package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/BakeRolls/trends"
)

func main() {
	d := flag.String("d", ",", "Delimiter")
	f := flag.String("f", "2006-01-02", "Time format")
	flag.Parse()

	iot, err := trends.InterestOverTime(flag.Args()...)
	if err != nil {
		fmt.Println("could not get trends: %v", err)
		os.Exit(2)
	}
	for _, r := range iot {
		fmt.Print(r.Time.Format(*f))
		for _, v := range r.Values {
			fmt.Printf("%s%d", *d, v)
		}
		fmt.Println()
	}
}
