package main

import (
	"fmt"
)

func main() {
	str := "a,b,c,"
	//reStr := strings.Replace(str, ",", "", 4)
	reStr := str[:len(str)-1]
	fmt.Println("", reStr)
	//fmt.Println("", reStr)
}
