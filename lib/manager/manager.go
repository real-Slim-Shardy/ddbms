package manager

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Structure that implements unit's data from InitConfig file from user
type InitCfgUnit struct {
	Number  int    `json:"number"`
	Address string `json:"address"`
}

// Structure that implements All json file data format
type CfgData struct {
	Lb      []InitCfgUnit `json:"load-balancers"`
	Servers []InitCfgUnit `json:"servers"`
	DBC     []InitCfgUnit `json:"db-clusters"`
}

// Structure that implements Instance information
type Instance struct {
	Type    string
	Number  int
	Address string
	Status  string
}

// Read InitConfig.json file and transfers raw data to special struct
// Function read json file with configuration
// Return a map ["InstanceName"]:"Instance{Type:string; Number: int; Address: string; Status: string}"
func ReadInitCfgFile(fileName string) (*map[string]Instance, error) {
	// Configuration file from user stored in cfg/*InitCfgFileName*
	buf, e := os.ReadFile(fileName)
	if e != nil {
		return nil, errors.New("Can't open Init Configuration File or it doesn't exists")
	}

	// Unmarshal data from json file to variable 'data'
	var data CfgData
	json.Unmarshal(buf, &data)

	// init output variable cfg
	cfg := make(map[string]Instance)

	// Parse data to result format
	// Parse LB information
	for i := 0; i < len(data.Lb); i++ {
		cfg["LB"+strconv.Itoa(i)] = Instance{
			Type:    "LB",
			Number:  data.Lb[i].Number,
			Address: data.Lb[i].Address,
			Status:  "Ready"}
	}

	// Parse Servers information
	for i := 0; i < len(data.Servers); i++ {
		cfg["Server"+strconv.Itoa(i)] = Instance{
			Type:    "Server",
			Number:  data.Servers[i].Number,
			Address: data.Servers[i].Address,
			Status:  "Ready"}
	}

	// Parse DB clusters information
	for i := 0; i < len(data.DBC); i++ {
		cfg["DB"+strconv.Itoa(i)] = Instance{
			Type:    "DB",
			Number:  data.DBC[i].Number,
			Address: data.DBC[i].Address,
			Status:  "Ready"}
	}

	// ОТЛАДКА
	fmt.Println(data)
	return &cfg, nil
}

// Create or Open file logs for storing all logs in it
func InitLogger() {
	file, err := os.OpenFile("logs", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		log.SetOutput(file)
	} else {
		log.Fatal("Failed to log to file, using default stderr")
	}
}
