package main

import (
	"ddbms/lib/lb"
	"ddbms/lib/manager"
	"ddbms/lib/serv"
	"fmt"
	"log"
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

	// send config info to instances
	op := serv.OpArgs{N: 1, Status: "Ready"}
	err = lb.SendCfgToServer((*data)["server1"].Address, "Ops.SetConfig", op)
	if err != nil {
		log.Println(err)
	}

	// test system availability

	// ========== Backgroud activity ===========
	// start heartbeat system

	// getting metrics from instances

	// ========== Work with Requests ===========
	// init listener

	// requests redirecting

}
