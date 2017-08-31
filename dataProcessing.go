package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type ProcessedData struct {
	DeviceID       int
	DeviceOID      string
	OIDdescription string
	Value          string
	Time           string
}

var ProcessedDataObject ProcessedData

// DataProcessing get information about device and write it to log file, returns array to send it to server
func DataProcessing(oidsID []int, oidsValues []string, oidsDescript []string) (processedInfo []string) {
	var processedData []string
	for i := 0; i < len(oidsID); i++ {
		getSnmpInfo(oidsValues[i], SNMPip, Community, Port)

		var stringlogData []string
		t := time.Now().Format("2006-01-02 15:04:05")
		tLog := time.Now().Format("2006-01-02")
		stringlogData = append(stringlogData, t)
		stringlogData = append(stringlogData, oidsValues[i])
		stringlogData = append(stringlogData, oidsDescript[i])
		stringlogData = append(stringlogData, ReternsData)

		a := strings.Join(stringlogData, ", ")
		f, err := os.OpenFile("log/log"+tLog+".txt", os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			println(err)
		}
		defer f.Close()

		if _, err = f.WriteString(a + "\r\n"); err != nil {
			println(err)
		}
		ProcessedDataObject.DeviceID = oidsID[i]
		ProcessedDataObject.DeviceOID = oidsValues[i]
		ProcessedDataObject.OIDdescription = oidsDescript[i]
		ProcessedDataObject.Value = ReternsData
		ProcessedDataObject.Time = t
		ProcessedInformation, err1 := json.Marshal(ProcessedDataObject)
		if err1 != nil {
			fmt.Println(`Marshal json problem in dataProcessing`)
			fmt.Println(err1)
		}
		processedData = append(processedData, string(ProcessedInformation))
	}
	return processedData
}
