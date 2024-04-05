package main

import (
	"ddbms/lib/manager"
	"fmt"
)

const (
	initCfgFileName = "InitConfig.json"
	filePath        = "../../cfg/"
)

func main() {
	// ============ Starting system ============
	// init logger
	manager.InitLogger()

	// get config
	data, err := manager.ReadInitCfgFile(filePath + initCfgFileName)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data)

	// test system availability

	// ========== Backgroud activity ===========
	// start heartbeat system

	// getting metrics from instances

	// ========== Work with Requests ===========
	// init listener

	// requests redirecting

}
