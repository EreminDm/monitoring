package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// ConnectConfig Configurate snmp connection
func ConnectConfig(device int) (query string) {
	var configFile string
	var dbquery string
	if device == 1 {
		configFile = `./config/snmpNPconfig.json`
		dbquery = `SELECT id, value, description FROM oids WHERE type = 'Netping'`
	}
	if device == 2 {
		configFile = `./config/snmpTLconfig.json`
		dbquery = `SELECT id, value, description FROM oids WHERE type = 'Teltonika'`
	}
	if device == 3 {
		configFile = `./config/snmpPCconfig.json`
		dbquery = `SELECT id, value, description FROM oids WHERE type = 'PC'`
	}
	plan, _ := ioutil.ReadFile(configFile)
	err := json.Unmarshal(plan, &SNMPconfigObject)
	if err != nil {
		log.Fatal("Cannot unmarshal the json in read connection config ", err)
	}
	SNMPip = SNMPconfigObject.SNMPip
	Community = SNMPconfigObject.Community
	Port = SNMPconfigObject.Port

	return dbquery
}
