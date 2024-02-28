package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

//type BacnetSocketField struct {
//	BacnetField
//	Message string
//}

//type BacnetField struct {
//	SourceIp           string `json:"sourceIp"`
//	DestinationIp      string `json:"destinationIp"`
//	NpduType           string `json:"npduType"`
//	ApduType           string `json:"apduType"`
//	ServiceChoice      string `json:"serviceChoice"`
//	ObjectType         string `json:"objectType"`
//	InstanceNumber     string `json:"instanceNumber"`
//	PropertyIdentifier string `json:"propertyIdentifier"`
//}

//type Manifest struct {
//	Key             string `json:"key"`
//	Name            string `json:"name"`
//	IsForward       bool   `json:"isForward"`
//	Url             string `json:"url"`
//	Method          string `json:"method"`
//	ContentType     string `json:"contentType"`
//	AllowBatch      bool   `json:"allowBatch"`
//	ResponseType    string `json:"responseType"`
//	ShowProgressBar bool   `json:"showProgressBar"`
//	RequireLog      bool   `json:"requireLog"`
//}

func main() {
	sourcePath := "C:\\Users\\liyang\\Desktop\\7\\7.5\\1111111222" // 替换为你的文件夹路径
	outputFilePath := "C:\\Users\\liyang\\Desktop\\7\\7.5\\11.sql" // 替换为你的输出文件路径
	content, err := ioutil.ReadFile(sourcePath)
	if err != nil {
		fmt.Printf("无法读取文件：%s，错误：%s\n", sourcePath, err)
		return
	}
	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		fmt.Println("无法创建输出文件：", err)
		return
	}
	defer outputFile.Close()
	// 按行切分内容
	lines := strings.Split(string(content), "\n")
	id := 15213
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		id++
		sid01 := strings.Index(line, "sid:")
		sid02 := strings.Index(line, "; rev")
		sid := line[sid01:sid02]
		sid1, _ := strconv.Atoi(strings.Split(sid, "sid:")[1])

		split01 := strings.Split(line, "msg:")
		split02 := strings.Split(split01[1], "\"")
		name := split02[1]
		//strconv.Atoi(strings.Split(sid, "sid:")[1])
		//sql := fmt.Sprintf("INSERT INTO `security_blacklist_vulnerability`(`id` ,`risk_level`, `name`, `status`, `action`, `severity`, `sid`, `message`, `sign_name`) VALUES (%d, 2, \"%s\", 1, 1, 1, %s, '%s', '');\n", id, name, sid, line)
		sql := fmt.Sprintf("INSERT INTO `policy_intrusion_dictionary` (`id`, `sid`, `rule_name`, `rule_source`, `occur_date`, `class_type`, `rule_message`, `risk_level`, `vulnerability_source`, `cve`, `attack_requirement`, `description`, `sign_name`, `type`) VALUES (%d, %d, \"%s\", '', '', '', '%s',1, '', '','','','','worm');\n", id, sid1, name, line)
		_, err = outputFile.WriteString(sql)
		if err != nil {
			fmt.Printf("无法写入文件：%s，错误：%s\n", outputFilePath, err)
			return
		}
	}
}

//func ManifestImport() {
//	//row 0旧的url 1新的url 2描述 3请求方式
//	outputFilePath := "C:\\Users\\ZqWrold\\Desktop\\1.txt" // 替换为你的输出文件路径
//	var oldMap = make(map[string]bool)
//	var list []Manifest
//	// 打开Excel文件
//	f, err := excelize.OpenFile("C:\\Users\\ZqWrold\\Desktop\\1.xlsx")
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 读取数据
//	rows, err := f.GetRows("Sheet1") // "Sheet1" 是您要读取的工作表名称
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 遍历行和列来获取数据
//	for _, row := range rows {
//		oldMap[row[0]] = true
//	}
//
//	for _, row := range rows {
//		if _, ok := oldMap[row[1]]; !ok {
//			if row[1][:1] == "/" {
//				row[1] = row[1][1:]
//			}
//			key := strings.ReplaceAll(row[1], "/", "_") + "_" + strings.ToLower(row[3])
//			log := true
//			if strings.Contains(row[1], "Query") {
//				log = false
//			}
//			var manifest = Manifest{
//				Key:             key,
//				Name:            row[2],
//				IsForward:       true,
//				Url:             "/" + row[1],
//				Method:          strings.ToLower(row[3]),
//				ContentType:     "application/json",
//				AllowBatch:      true,
//				ResponseType:    "realtime",
//				ShowProgressBar: false,
//				RequireLog:      log,
//			}
//			list = append(list, manifest)
//		}
//	}
//	fmt.Println(len(list))
//	outputFile, err := os.Create(outputFilePath)
//	if err != nil {
//		fmt.Println("无法创建输出文件：", err)
//		return
//	}
//	defer outputFile.Close()
//	listJson, err := json.Marshal(list)
//	if err != nil {
//		fmt.Println("111：", err)
//		return
//	}
//	outputFile.WriteString(string(listJson))
//}
//
//func scan() {
//	folderPath := "C:\\Users\\ZqWrold\\Desktop\\emerging.rules\\rules"       // 替换为你的文件夹路径
//	outputFilePath := "C:\\Users\\ZqWrold\\Desktop\\emerging.rules\\1.rules" // 替换为你的输出文件路径
//
//	filePaths, err := filepath.Glob(filepath.Join(folderPath, "*.rules"))
//	if err != nil {
//		fmt.Println("无法读取文件路径：", err)
//		return
//	}
//
//	outputFile, err := os.Create(outputFilePath)
//	if err != nil {
//		fmt.Println("无法创建输出文件：", err)
//		return
//	}
//	defer outputFile.Close()
//
//	for _, filePath := range filePaths {
//		content, err := ioutil.ReadFile(filePath)
//		if err != nil {
//			fmt.Printf("无法读取文件：%s，错误：%s\n", filePath, err)
//			continue
//		}
//		if err != nil {
//			log.Fatal(err)
//		}
//
//		// 按行切分内容
//		lines := strings.Split(string(content), "\n")
//
//		// 遍历并打印每行内容
//		for _, line := range lines {
//			if len(line) == 0 || line[:1] == "#" {
//				continue
//			}
//			_, err = outputFile.WriteString(string(line) + "\n")
//			if err != nil {
//				fmt.Printf("无法写入文件：%s，错误：%s\n", outputFilePath, err)
//				return
//			}
//		}
//	}
//
//	// 读取整个文件内容
//	content, err := ioutil.ReadFile("C:\\Users\\ZqWrold\\Desktop\\新建文本文档 (2).txt")
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	// 按行切分内容
//	lines := strings.Split(string(content), "\n")
//
//	// 遍历并打印每行内容
//	for _, line := range lines {
//		if strings.Contains(line, "报文") {
//			//fmt.Println(line)
//			ss := strings.Split(line, "bacnet报文：")
//			Bacnet(ss[1])
//		}
//	}
//}
//
//var m = make(map[uint64]bool)
//
//func Bacnet(message string) {
//	splitNftRule := strings.Split(message, " ")
//	var bacnetSocketField BacnetSocketField
//	bacnetSocketField.SourceIp = splitNftRule[2]
//	bacnetSocketField.DestinationIp = splitNftRule[5]
//	//对pd1处理，获取功能码
//	a := strings.Index(splitNftRule[7], ":")
//	b := strings.Index(splitNftRule[7], ",")
//	c := strings.Index(splitNftRule[7], ";")
//	pd1 := splitNftRule[7][a+1 : b]
//	pd1Uint, _ := strconv.ParseUint(pd1, 10, 64)
//	pd1Bin := strconv.FormatUint(pd1Uint, 2)
//	build := strings.Builder{}
//	for i := 0; i < 64-len(pd1Bin); i++ {
//		build.WriteString("0")
//	}
//	pd1Bin = build.String() + pd1Bin
//	//对二进制的pd1处理
//	//注意大小端的问题
//	serviceChoice, _ := strconv.ParseUint(pd1Bin[40:48], 2, 64)
//	apduType, _ := strconv.ParseUint(pd1Bin[48:56], 2, 64)
//	npduType, _ := strconv.ParseUint(pd1Bin[56:], 2, 64)
//	if npduType == 128 {
//		//代表npdu mesgtype不存在
//		bacnetSocketField.ServiceChoice = strconv.FormatUint(serviceChoice, 10)
//		bacnetSocketField.ApduType = strconv.FormatUint(apduType, 10)
//		pd2 := splitNftRule[7][b+1 : c]
//		pd2Uint, _ := strconv.ParseUint(pd2, 10, 64)
//		pd2Bin := strconv.FormatUint(pd2Uint, 2)
//		build2 := strings.Builder{}
//		for i := 0; i < 64-len(pd2Bin); i++ {
//			build2.WriteString("0")
//		}
//		pd2Bin = build2.String() + pd2Bin
//		propertyId, _ := strconv.ParseUint(pd2Bin[:8], 2, 64)
//		objectType, _ := strconv.ParseUint(pd2Bin[8:18], 2, 64)
//		instanceNumber, _ := strconv.ParseUint(pd2Bin[18:40], 2, 64)
//		bacnetSocketField.PropertyIdentifier = strconv.FormatUint(propertyId, 10)
//		bacnetSocketField.ObjectType = strconv.FormatUint(objectType, 10)
//		bacnetSocketField.InstanceNumber = strconv.FormatUint(instanceNumber, 10)
//	} else {
//		bacnetSocketField.NpduType = strconv.FormatUint(npduType, 10)
//	}
//
//	//计算特征值
//	eigenvalue, _ := hashstructure.Hash(&bacnetSocketField.BacnetField, hashstructure.FormatV2, nil)
//	if m[eigenvalue] {
//		fmt.Println(eigenvalue, message)
//		return
//	}
//	m[eigenvalue] = true
//}
//
//func BcryptHash(password string) string {
//	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
//	return string(bytes)
//}
