package main

import (
	"encoding/json"
	"fmt"
)

type ManageAuthConfig struct {
	EnableSSH bool `json:"enableSSH"`
	EnableSMP bool `json:"enableSMP"`
}

func main() {
	var manageAuth ManageAuthConfig

	str := "{\"enableSSH\":true,\"enableSMP\":false}"
	//authJson, err := json.Marshal(str)
	json.Unmarshal([]byte(str), &manageAuth)
	fmt.Println(manageAuth)

}
