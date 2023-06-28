package main

import (
	"sync"
)

var (
	WaitGroup sync.WaitGroup
	//sql       = "'2022-07-20 00:00:00','2022-07-20 00:00:00','08:00:27:67:05:3c','08:00:27:ff:ee:23', 4,'192.168.0.222',502,'192.168.0.211',41758,6,'TCP',77, 080027FFEE2308002767053C08004500003F820E4000400635A9C0A800DEC0A800D301F6A31E59F6431C017F6264801800E3B0B100000101080A080AC1B3080AC82D000C000000050003020000,'modbus'----func----'Read_Holding_Registers'----"
	sql = "'2007-09-12 18:18:56','2007-09-12 18:18:56','02:14:4f:23:98:cf','ff:ff:ff:ff:ff:ff', 0,'NULL',0,'NULL',0,255,'NULL',113, FFFFFFFFFFFF02144F2398CF88A461100BF000200100018004000F010007F105003001028004000800050001F205003001028004000800010001F304003001028004000800010001F403003001028004000800010001F502003001028004000800010001F6010030010200040008000100,'ethercat'----cmd,Logaddr,offsetaddr----'{LWR,BRD,APRD,APRD,APRD,APRD,APRD}','{12000}','{130,130,130,130,130,130}'"
)

func main() {
	//声明unixsocket
	//us := utils.NewUnixSocket("/data/socket/obtain_proto_sock")
	////发送数据unixsocket并返回服务端处理结果
	//for z := 0; z < 10; z++ {
	//	//循环中启动10个携程
	//	WaitGroup.Add(1)
	//	go func() {
	//		//单个携程循环1000次向unix socket发送执行数据
	//		for i := 0; i < 1000; i++ {
	//			us.ClientSendContext(sql)
	//		}
	//		defer WaitGroup.Done()
	//	}()
	//}
	WaitGroup.Add(1)
	go func() {
		//单个携程循环1000次向unix socket发送执行数据
		for i := 0; i < 1; i++ {
			//us.ClientSendContext(sql)

		}
		defer WaitGroup.Done()
	}()
	//挂起等待携程执行完成后结束主线程
	WaitGroup.Wait()
}
