package SnmpWrapper

import (
	"fmt"
	"log"

	g "./packages/gosnmp-master"
)

func snmpConnect(ip string, community string, port uint16) {
	g.Default.Target = ip
	g.Default.Community = community
	g.Default.Port = port //1611
	err := g.Default.Connect()
	if err != nil {
		log.Fatalf("Connect() err: %v", err)
	}
}

//GetArrayInfo return oids values information from device
func GetArrayInfo(OIDS []string, ip string, community string, port uint16) {
	//snmpConnect("172.10.10.182", "SWITCH", 1611)
	snmpConnect(ip, community, port)
	//[]string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
	result, err2 := g.Default.Get(OIDS) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}

	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.
			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
		}
	}
	defer g.Default.Conn.Close()
}

func GetOIDInfo(OID string, ip string, community string, port uint16) (Data string) {
	snmpConnect(ip, community, port)
	var oids []string
	var returnsData string
	oids = append(oids, OID)
	result, err2 := g.Default.Get(oids) // Get() accepts up to g.MAX_OIDS
	if err2 != nil {
		log.Fatalf("Get() err: %v", err2)
	}
	for i, variable := range result.Variables {
		fmt.Printf("%d: oid: %s ", i, variable.Name)

		// the Value of each variable returned by Get() implements
		// interface{}. You could do a type switch...
		switch variable.Type {
		case g.OctetString:
			fmt.Printf("string: %s\n", string(variable.Value.([]byte)))
			abc := string(variable.Value.([]byte))
			returnsData = abc
		default:
			// ... or often you're just interested in numeric values.
			// ToBigInt() will return the Value as a BigInt, for plugging
			// into your calculations.

			fmt.Printf("number: %d\n", g.ToBigInt(variable.Value))
			val := g.ToBigInt(variable.Value)
			returnsData = val.String()

		}
	}
	return returnsData
}
