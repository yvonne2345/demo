package main

import (
	"fmt"
	"strings"
)

func main() {
	//获取handle
	//str := "table ip netvine-count-table {\n\tchain flow-detection-chain { # handle 1\n\t\ttype filter hook forward priority filter; policy accept;\n\t\tip saddr 192.168.1.11 ip daddr 192.168.1.222 tcp dport 502 meta hour \"00:00\"-\"02:00\" meta day \"Thursday\" counter packets 0 bytes 0 ct mark set 0x05f5e10c # handle 3\n\t\tip saddr 192.168.247.131 ip daddr 192.168.3.10 udp dport 20000 meta hour \"00:00\"-\"02:00\" meta day \"Thursday\" counter packets 0 bytes 0 ct mark set 0x05f5e10d # handle 4\n\t\tip saddr 192.168.3.251 ip daddr 192.168.8"
	//split := strings.Split(str, "mark")
	//var handleList []string
	//for _, v := range split {
	//	if strings.Contains(v, "set") {
	//		handle := strings.Split(v, "\n\t")
	//
	//		handleList = append(handleList, handle[0])
	//	}
	//}

	//replace bytes重新统计
	//str := "udp dport 502 counter packets 0 bytes 87777 ct mark set 0x05f5e10f # handle 3"
	//
	//re, err := regexp.Compile(`bytes \d*`)
	//if err != nil {
	//
	//}
	//
	//newStr := re.ReplaceAllString(str, "bytes 0")
	//
	//fmt.Println(newStr)

	//获取bytes
	//str := "table ip netvine-count-table {\n\tchain flow-detection-chain { # handle 1\n\t\ttype filter hook forward priority filter; policy accept;\n\t\tip saddr 192.168.1.11 ip daddr 192.168.1.222 tcp dport 502 meta hour \"00:00\"-\"02:00\" meta day \"Thursday\" counter packets 0 bytes 0 ct mark set 0x05f5e10c # handle 3\n\t\tip saddr 192.168.247.131 ip daddr 192.168.3.10 udp dport 20000 meta hour \"00:00\"-\"02:00\" meta day \"Thursday\" counter packets 0 bytes 0 ct mark set 0x05f5e10d # handle 4\n\t\tip saddr 192.168.3.251 ip daddr 192.168.8"
	//split := strings.Split(str, "counter")
	//var handleList []string
	//for _, s := range split {
	//	if strings.Contains(s, "packets") {
	//		handle := strings.Split(s, "\n\t")
	//		handleList = append(handleList, handle[0])
	//	}
	//}
	//
	//fmt.Println("", handleList)

	//int转float
	//s := "1000"
	//num, _ := strconv.Atoi(s)         //int
	//f, _ := strconv.ParseFloat(s, 64) //float64
	//
	//var str string
	//if num < 1024 {
	//	str = s + "B"
	//} else if num > 1024 {
	//	str = fmt.Sprintf("%.2f", f/1024)
	//	str = str + "M"
	//}
	//fmt.Println(str)

	//time计算时间

	//timeDuration := time.Now()
	//time2 := time.Now().Add(20 * time.Minute)
	//fmt.Println(time2)
	//
	//if timeDuration.Sub(time2) >= 20*time.Minute {
	//	// log表新增
	//	fmt.Println("1")
	//} else {
	//	fmt.Println("2")
	//
	//}
	//var FlowDetectionWarningMap = make(map[string]time.Time)
	//for i, _ := range split {
	//	if FlowDetectionWarningMap[string(i)].IsZero() {
	//		fmt.Println(1111)
	//	}
	//}
	//var t time.Time
	//fmt.Printf("%v %v\n\n", t, t.IsZero())
	//t = time.Now()
	//fmt.Printf("%v %v", t, t.IsZero())

	//str := "[{\"value\":1218,\"name\":\"Modbus\"},{\"value\":319,\"name\":\"CIP\"},{\"value\":142273,\"name\":\"TRDP\"},{\"value\":83593,\"name\":\"OMRON FINS\"},{\"value\":56821,\"name\":\"BACnet\"}]"
	//str = strings.ReplaceAll(str, "\"value\"", "value")
	//str = strings.ReplaceAll(str, "\"name\"", "name")
	//fmt.Println(str)

	str := "{\"time\":null,\"arrays\":null}"
	str = getResultEmpty(str)
	fmt.Println(str)

}

func getResultEmpty(str string) string {
	if strings.Contains(str, "null") {
		str = strings.ReplaceAll(str, "null", "[]")
		return str
	} else {
		return str
	}
}
