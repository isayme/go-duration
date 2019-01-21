## Duration
[![Build Status](https://travis-ci.org/isayme/go-duration.svg?branch=master)](https://travis-ci.org/isayme/go-duration)

A wraper for time.Duration, for prettier print

## Example
```
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

	fmt.Printf("%s\n", d.String()) // 1h2m3s
	fmt.Printf("%v\n", d)          // 1h2m3s
	fmt.Printf("%+v\n", d)         // 1h2m3s
	fmt.Printf("%#v\n", d)         // "1h2m3s"
}
```