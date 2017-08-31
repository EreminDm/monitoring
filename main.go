package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"

	DB "./pkg/netPingGetOidbyObject"
	SNMP "./pkg/snmpWrapper"
)

// SNMPconfig structure
type SNMPconfig struct {
	SNMPip    string
	Community string
	Port      uint16
}

//ReternsData processed data
var ReternsData string

//SNMPconfigObject of SNMPconfig structure
var SNMPconfigObject SNMPconfig

// SNMPip ip connect adress
var SNMPip string

//Community secret word
var Community string

//Port connect
var Port uint16

// return information about all devices in time interval
func main() {

	var processedInfo []string
	os.Mkdir("log", 0777)
	println(`start timer`)
	ticker := time.NewTicker(time.Millisecond * 1500)
	go func() {
		println(`start go routine function device monitoring`)
		for t := range ticker.C {
			// i=1 Netping; i=2 Teltonica ; i=3 PC
			for i := 1; i <= 3; i++ {
				query := ConnectConfig(i)
				oidsID, oidsValues, oidsDescript := DB.GetOIDSFromDB(query)
				//write to local file information
				processedInfo = DataProcessing(oidsID, oidsValues, oidsDescript)
				//postRequestToServer(processedInfo)

			}
			fmt.Println(processedInfo, t)
		}
	}()

	println(`start web server`)
	http.HandleFunc("/snmp", handler)
	http.ListenAndServe(":8080", nil)
}

// requst function with device param which return respoce with processed data
func handler(w http.ResponseWriter, r *http.Request) {
	coonType := r.URL.Query().Get("type")
	var query string
	if coonType == "" {
		deviceType := 1
		query = ConnectConfig(deviceType)
	} else {
		deviceType, err := strconv.Atoi(coonType)
		if err != nil {
			println(err)
		}
		query = ConnectConfig(deviceType)
	}
	oidsID, oidsValues, oidsDescript := DB.GetOIDSFromDB(query)
	processedInfo := DataProcessing(oidsID, oidsValues, oidsDescript)
	for i := 0; i < len(processedInfo); i++ {
		io.WriteString(w, processedInfo[i])
	}
}

// Request to get snmp information
func getSnmpInfo(value string, ip string, community string, port uint16) {
	//snmp.GetOIDInfo(`1.3.6.1.2.1.1.1`, `127.0.0.1`, `public`, 161)
	ReternsData = SNMP.GetOIDInfo(value, ip, community, port)
	fmt.Println(ReternsData)
}

// func postRequestToServer(processedInfo []string) {
// 	configFile := `./config/serverURLconfig.txt`
// 	// read txt config
// 	b, err := ioutil.ReadFile(configFile)
// 	if err != nil {
// 		fmt.Print(err)
// 	}
// 	url := string(b)

// 	if url == "" {
// 		println("There is no server url, check config file")
// 	} else {
// 		fmt.Println([]byte(processedInfo))
// 		var jsonStr = []byte(processedInfo)
// 		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
// 		req.Header.Set("X-Custom-Header", "myvalue")
// 		req.Header.Set("Content-Type", "application/json")

// 		client := &http.Client{}
// 		resp, err := client.Do(req)
// 		if err != nil {
// 			panic(err)
// 		}
// 		defer resp.Body.Close()
// 	}
// }
