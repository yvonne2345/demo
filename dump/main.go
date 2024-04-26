package main

import (
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func main() {
	// 远程设备的SSH连接信息
	remoteAddress := "root@10.25.30.126:22"

	// 建立SSH连接
	sshConfig := &ssh.ClientConfig{
		User: "root", // 远程设备的用户名
		Auth: []ssh.AuthMethod{
			ssh.Password("netvine123"),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 这是一个示例，实际中请使用更安全的方式验证主机
	}

	// 连接到远程设备
	client, err := ssh.Dial("tcp", remoteAddress, sshConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// 在远端服务器上执行抓包命令
	cmd := exec.Command("tcpdump", "-i", "eth0", "-w", "/path/to/output/file.pcap")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to run tcpdump: %v", err)
	}

	// 将抓包文件从远端服务器传回本地电脑
	remoteFilePath := "/path/to/output/file.pcap"
	localFilePath := "/path/to/local/output/file.pcap"

	session, err = client.NewSession()
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	defer session.Close()

	// 执行命令并捕获标准输出
	cmdStr := "cat " + remoteFilePath
	output, err := session.CombinedOutput(cmdStr)
	if err != nil {
		log.Fatalf("Failed to run cat command: %v", err)
	}

	// 将标准输出写入本地文件
	err = ioutil.WriteFile(localFilePath, output, 0644)
	if err != nil {
		log.Fatalf("Failed to write local file: %v", err)
	}

	log.Printf("File copied successfully from remote to local.")

}
