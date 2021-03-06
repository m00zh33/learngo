// For more tutorials: https://blog.learngoprogramming.com
//
// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {
		var c string

		switch rand.Intn(3) {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		}
		fmt.Printf("\r%s", c)
		time.Sleep(time.Millisecond * 150)
	}
}
