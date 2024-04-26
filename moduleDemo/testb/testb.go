package testb

import "fmt"

func Sum(i1 int, i2 int) int {
	return i1 + i2
}

func Hello(s string) string {
	res := fmt.Sprintf("name:%s", s)
	return res
}
