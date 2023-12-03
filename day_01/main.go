package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var regex = regexp.MustCompile(`[^0-9 ]+`)

func clearString(str string) string {
	return regex.ReplaceAllString(str, "")
}

func main() {
	var c int = 0
	content, err := os.ReadFile("input.txt")
	if err != nil {
		println(err)
		return
	}

	lines := strings.Split(string(content), "\n")

	for _, v := range lines {
		a := clearString(v)
		if len(a) != 0 {
			if len(a) > 2 {
				first := a[0:1]
				last := a[len(a)-1:]
				combined := first + last
				t, err := strconv.Atoi(combined)
				if err != nil {
					fmt.Println(err)
					return
				}
				c += t
			} else if len(a) == 2 {
				number, err := strconv.Atoi(a)
				if err != nil {
					fmt.Println(err)
					return
				}
				c += number
			} else {
				number := a + a
				n, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println(err)
					return
				}
				c += n
			}
		}
	}

	fmt.Println(c)
}
