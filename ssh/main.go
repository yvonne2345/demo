package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/ssh"
	"net/http"
	"os"
	"time"
)

type ReqInfo struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type CmdReq struct {
	Cmd      string `json:"cmd"`
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

var (
	Client  *ssh.Client
	Session *ssh.Session
)

func sshTest(w http.ResponseWriter, r *http.Request) {
	var cmdReq CmdReq
	err := json.NewDecoder(r.Body).Decode(&cmdReq)
	//处理
	var reqParam ReqInfo
	reqParam.Host = cmdReq.Host + ":22"
	reqParam.User = cmdReq.User
	reqParam.Password = cmdReq.Password

	err = InitConnect(reqParam)
	// 创建一个新的会话
	Session, err = Client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: ", err)
		os.Exit(1)
	}
	defer Session.Close()
	// 设置标准输出和标准错误输出
	Session.Stdout = os.Stdout
	Session.Stderr = os.Stderr

	// 执行远程命令
	err = Session.Run(cmdReq.Cmd)
	if err != nil {
		fmt.Println("Failed to run command: ", err)
		os.Exit(1)
	}

}

// sshHandler HTTP 请求的处理程序
func sshConnect(w http.ResponseWriter, r *http.Request) {
	var reqParam ReqInfo
	var resp Response
	err := json.NewDecoder(r.Body).Decode(&reqParam)
	if err != nil {
		// 处理解码错误
		fmt.Println("获取ip,用户.........")
	}
	if reqParam.Host == "" || reqParam.User == "" || reqParam.Password == "" {
		resp.Msg = "Please provide host and User and Password."
		return
	}

	//处理
	reqParam.Host = reqParam.Host + ":22"

	err = InitConnect(reqParam)

	//// 创建一个新的会话
	//Session, err = Client.NewSession()
	//if err != nil {
	//	fmt.Println("Failed to create session: ", err)
	//	os.Exit(1)
	//}
	//defer Session.Close()
	//// 设置标准输出和标准错误输出
	//Session.Stdout = os.Stdout
	//Session.Stderr = os.Stderr
	//
	//// 执行远程命令
	//err = Session.Run("ls -l")
	//if err != nil {
	//	fmt.Println("Failed to run command: ", err)
	//	os.Exit(1)
	//}

	data := "ssh success"
	response, err := json.Marshal(data)
	// 设置响应头
	w.Header().Set("Content-Type", "application/json")

	// 发送 JSON 响应
	w.Write(response)
}

func InitConnect(reqParam ReqInfo) error {
	// 配置SSH客户端
	config := &ssh.ClientConfig{
		User: reqParam.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(reqParam.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 不检查主机密钥，实际应用中请慎用
	}

	// 连接到远程服务器
	var err error
	Client, err = ssh.Dial("tcp", reqParam.Host, config)
	if err != nil {
		fmt.Println("Failed to dial: ", err)
		os.Exit(1)
	}
	//defer Client.Close()

	return nil
}

func main() {
	// 设置路由和处理程序
	http.HandleFunc("/ssh", sshConnect)
	http.HandleFunc("/query", sshTest)

	// 启动 HTTP 服务器，监听端口 8088
	http.ListenAndServe(":8088", nil)
	time.Sleep(10 * time.Minute)
}
