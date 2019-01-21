package main

import (
	"encoding/json"
	"fmt"

	duration "github.com/isayme/go-duration"
)

func main() {
	var d duration.Duration
	err := json.Unmarshal([]byte(`"1h2m3s"`), &d)
	if err != nil {
		panic(err)
	}

	fmt.Println("duration.Duration")
	fmt.Printf("%%s: %s\n", d.String())
	fmt.Printf("%%v: %v\n", d)
	fmt.Printf("%%+v: %+v\n", d)
	fmt.Printf("%%#v: %#v\n", d)

	fmt.Println("\nstring")
	s := "str"
	fmt.Printf("%%s: %s\n", s)
	fmt.Printf("%%v: %v\n", s)
	fmt.Printf("%%+v: %+v\n", s)
	fmt.Printf("%%#v: %#v\n", s)
}
