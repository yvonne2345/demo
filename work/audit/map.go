package main

import "fmt"

func main() {

	//var RelationAnalysisMap = make(map[int64]int)
	//RelationAnalysisMap[1] = 0
	//RelationAnalysisMap[2] = 0
	//RelationAnalysisMap[3] = 0
	//RelationAnalysisMap[4] = 0
	//
	//for i, _ := range RelationAnalysisMap {
	//	RelationAnalysisMap[i] = int(+i)
	//}
	//fmt.Println(RelationAnalysisMap)
	//fmt.Println("----")
	//
	//for k, v := range RelationAnalysisMap {
	//	fmt.Println(RelationAnalysisMap[k])
	//	fmt.Println(v)
	//	fmt.Println("======")
	//}

	var tempList = make(map[int64]map[int64]int64)
	var temp = make(map[int64]int64)
	temp[1] = 1
	tempList[1] = temp
	temp[2] = 2
	tempList[1] = temp
	m := tempList[1]
	fmt.Println(len(tempList[1]))
	fmt.Println(len(m))
}
