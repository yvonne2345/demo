package main

//
//import (
//	"fmt"
//	"github.com/go-redis/redis"
//	"gorm.io/gorm"
//	_ "net/http/pprof"
//	"strconv"
//	"strings"
//	"time"
//)
//
//func initRedis() (err error) {
//	NETVINE_REDIS = redis.NewClient(&redis.Options{
//		Addr:     "10.25.10.125:6379", // 指定
//		Password: "Netvine123#@!",
//		DB:       1, // redis一共16个库，指定其中一个库即可
//	})
//	_, err = NETVINE_REDIS.Ping().Result()
//	return
//}
//
//func main() {
//	initRedis()
//	t1 := time.Now()
//	for i := 0; i < 10; i++ {
//		s := NETVINE_REDIS.Get("base_line_content").Val()
//		fmt.Println("", s)
//	}
//	t2 := time.Now()
//	d := t2.Sub(t1)
//	fmt.Println("", d)
//	//go func() {
//	//dsn := "root:root@tcp(10.25.10.134:3306)/audit?charset=utf8mb4&parseTime=True&loc=Local"
//	//db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	//fmt.Println("", db)
//	//message := "'2000-01-01 20:54:43','2000-01-01 20:54:43','c8:5b:76:3e:a5:5d','3e:c0:18:af:9d:d9', 4,'192.168.99.45',3812,'192.168.99.31',502,6,'TCP',64,'','modbus'----func----''----"
//	//InsertContent(message)
//
//	//for {
//	//	for i := 0; i < 100000; i++ {
//	//	}
//	//	time.Sleep(time.Second)
//	//}
//	//}()
//
//	//http.ListenAndServe("0.0.0.0:8081", nil)
//}
//
//func InsertContent(message string) {
//	split := strings.Split(message, "----")
//	//"'2000-01-01 20:54:43','2000-01-01 20:54:43','c8:5b:76:3e:a5:5d','3e:c0:18:af:9d:d9', 4,'192.168.99.45',3812,'192.168.99.31',502,6,'TCP',64,'','modbus'----func----''----"
//	flowHead := split[0]
//	//protSeg := split[1] //"length,windows,test"
//	protVal := split[2] //"'^','^','^'"
//	flowHeadSplit := strings.Split(flowHead, ",")
//	protoName := flowHeadSplit[13]
//	if protoName == "eniptcp" || protoName == "enip tcp" || protoName == "ENIP" {
//		protoName = "ENIP TCP"
//	}
//	var flowDataHead FlowDataHeads
//	flowDataHead.ID = 0
//	flowDataHead.FlowTimestamp = strings.Trim(flowHeadSplit[0], "'")
//	flowDataHead.PacketTimestamp = strings.Trim(flowHeadSplit[1], "'")
//	flowDataHead.SourceMac = strings.Trim(flowHeadSplit[2], "'")
//	flowDataHead.DestinationMac = strings.Trim(flowHeadSplit[3], "'")
//	flowDataHead.IpVersion, _ = strconv.Atoi(strings.Trim(flowHeadSplit[4], " "))
//	flowDataHead.SourceIp = strings.Trim(flowHeadSplit[5], "'")
//	flowDataHead.SourcePort, _ = strconv.Atoi(flowHeadSplit[6])
//	flowDataHead.DestinationIp = strings.Trim(flowHeadSplit[7], "'")
//	flowDataHead.DestinationPort, _ = strconv.Atoi(flowHeadSplit[8])
//	flowDataHead.ProtocolType, _ = strconv.Atoi(flowHeadSplit[9])
//	flowDataHead.ProtocolTypeName = strings.Trim(flowHeadSplit[10], "'")
//	flowDataHead.PacketLenth, _ = strconv.Atoi(flowHeadSplit[11])
//	flowDataHead.PacketContent = strings.Trim(flowHeadSplit[12], "'")
//	flowDataHead.ProtocolSourceName = strings.Trim(protoName, "'")
//	//camel := utils.ConvertCamel(protSeg)
//	flowDataHead.ProtVal = protVal
//	flowDataHead.FlowHeadSplit = flowHeadSplit[1]
//	//flowDataHead.Camel = camel
//	// TODO 修改成批量处理，现在每个数据都要去匹配，效率很低
//	handlerBaseLineStudy(flowDataHead)
//}
//
//func handlerBaseLineStudy(flowDataHead FlowDataHeads) {
//	//查询当前是否正在基线学习中，并确定学习基线的信息
//	s := NETVINE_REDIS.Get("base_line_study").Val()
//	if s != "" {
//		val := NETVINE_REDIS.Get("base_line_content").Val()
//		baseLineGroupId, _ := strconv.Atoi(s)
//		var tempArray []BaseLineStudyTemp
//		//如果正在学习中，将协议审计中能获取到的基线信息写入临时表中
//		if strings.Contains(val, "1") {
//			temp1 := BaseLineStudyTemp{
//				BaseLineGroupId: baseLineGroupId,
//				SrcIp:           flowDataHead.SourceIp,
//				About:           BaseLineTypeIp,
//			}
//			tempArray = append(tempArray, temp1)
//			temp2 := BaseLineStudyTemp{
//				BaseLineGroupId: baseLineGroupId,
//				SrcIp:           flowDataHead.DestinationIp,
//				About:           BaseLineTypeIp,
//			}
//			tempArray = append(tempArray, temp2)
//		}
//		if strings.Contains(val, "2") {
//			temp := BaseLineStudyTemp{
//				BaseLineGroupId: baseLineGroupId,
//				SrcIp:           flowDataHead.DestinationIp,
//				Port:            strconv.Itoa(flowDataHead.DestinationPort),
//				About:           BaseLineTypePort,
//			}
//			tempArray = append(tempArray, temp)
//			temp2 := BaseLineStudyTemp{
//				BaseLineGroupId: baseLineGroupId,
//				SrcIp:           flowDataHead.SourceIp,
//				Port:            strconv.Itoa(flowDataHead.SourcePort),
//				About:           BaseLineTypePort,
//			}
//			tempArray = append(tempArray, temp2)
//		}
//		if strings.Contains(val, "3") {
//			temp := BaseLineStudyTemp{
//				BaseLineGroupId: baseLineGroupId,
//				SrcIp:           flowDataHead.SourceIp,
//				DestIp:          flowDataHead.DestinationIp,
//				Port:            strconv.Itoa(flowDataHead.DestinationPort),
//				Protocol:        flowDataHead.ProtocolTypeName,
//				About:           BaseLineTypeNetworkConnect,
//			}
//			tempArray = append(tempArray, temp)
//		}
//		if strings.Contains(val, "4") {
//			temp := BaseLineStudyTemp{
//				BaseLineGroupId: baseLineGroupId,
//				SrcIp:           flowDataHead.SourceIp,
//				DestIp:          flowDataHead.DestinationIp,
//				Port:            strconv.Itoa(flowDataHead.DestinationPort),
//				Protocol:        flowDataHead.ProtocolTypeName,
//				About:           BaseLineTypeFlow,
//				PacketLength:    flowDataHead.PacketLenth, // 用于统计上下行流量
//			}
//			tempArray = append(tempArray, temp)
//		}
//
//		// 批量存储
//		NETVINE_DB.Create(&tempArray)
//		if len(tempArray) > 0 {
//		}
//	}
//}
//
//type FlowDataHeads struct {
//	NETVINE_MODEL
//	FlowTimestamp      string `json:"flowTimestamp"`
//	PacketTimestamp    string `json:"packetTimestamp"`
//	SourceMac          string `json:"sourceMac"`
//	DestinationMac     string `json:"destinationMac"`
//	IpVersion          int    `json:"ipVersion"`
//	SourcePort         int    `json:"sourcePort"`
//	SourceIp           string `json:"sourceIp"`
//	DestinationIp      string `json:"destinationIp"`
//	DestinationPort    int    `json:"destinationPort"`
//	ProtocolType       int    `json:"protocolType"`
//	ProtocolTypeName   string `json:"protocolTypeName"`
//	PacketLenth        int    `json:"packetLenth"`
//	PacketContent      string `json:"packetContent"`
//	ProtocolSourceName string `json:"protocolSourceName"`
//	Directions         int    `json:"directions"`
//	Camel              string `json:"camel" sql:"-" gorm:"-"`
//	ProtVal            string `json:"protVal" sql:"-" gorm:"-"`
//	FlowHeadSplit      string `json:"flowHeadSplit" sql:"-" gorm:"-"`
//}
//
//type NETVINE_MODEL struct {
//	ID        uint           `gorm:"primarykey" json:"id"` // 主键ID
//	CreatedAt FormatTime     `json:"createdAt"`            // 创建时间
//	UpdatedAt FormatTime     `json:"updatedAt"`            // 更新时间
//	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`       // 删除时间
//}
//
//type FormatTime struct {
//	time.Time
//}
//
//var (
//	NETVINE_DB    *gorm.DB
//	NETVINE_REDIS *redis.Client
//)
//
//const (
//	BaseLineTypeIp             = 1
//	BaseLineTypePort           = 2
//	BaseLineTypeNetworkConnect = 3
//	BaseLineTypeFlow           = 4
//)
//
//type BaseLineStudyTemp struct {
//	NETVINE_MODEL
//	SrcIp           string `json:"srcIp"`
//	DestIp          string `json:"destIp"`
//	Port            string `json:"port"`
//	About           int    `json:"about"` // 1 IP基线  2 活动端口基线
//	BaseLineGroupId int    `json:"baseLineGroupId"`
//	Protocol        string `json:"protocol"`
//	PacketLength    int    `json:"packetLength"` // 后增加，用于统计上下行流量
//}
//
//func (p *BaseLineStudyTemp) TableName() string {
//	return "base_line_study_temp"
//}
