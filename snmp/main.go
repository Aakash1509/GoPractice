package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gosnmp/gosnmp"
)

func main() {
	// Check if enough arguments are provided
	if len(os.Args) < 4 {
		log.Fatal("Usage: snmpHostnameFetcher <community> <version> <ip>")
	}

	// Read SNMP connection details from command-line arguments
	community := os.Args[1]
	version := os.Args[2]
	ip := os.Args[3]

	// Determine SNMP version
	var snmpVersion gosnmp.SnmpVersion
	switch version {
	case "1":
		snmpVersion = gosnmp.Version1
	case "2c":
		snmpVersion = gosnmp.Version2c
	default:
		log.Fatal("Unsupported SNMP version. Use '1' or '2c'.")
	}

	// Set up the SNMP connection parameters
	params := &gosnmp.GoSNMP{
		Target:    ip,
		Port:      161,
		Community: community,
		Version:   snmpVersion,
		Timeout:   gosnmp.Default.Timeout,
		Retries:   gosnmp.Default.Retries,
	}

	// Connect to the SNMP device
	err := params.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to SNMP device: %v", err)
	}
	defer params.Conn.Close() // Close the connection when done

	// The OID for hostname (sysName) in SNMP
	hostnameOID := ".1.3.6.1.2.1.1.5.0"

	// Fetch the hostname using SNMP Get
	result, err := params.Get([]string{hostnameOID})
	if err != nil {
		log.Fatalf("Failed to fetch hostname: %v", err)
	}

	// Check if the result contains the expected value
	if len(result.Variables) == 0 {
		log.Fatal("No SNMP response received for hostname.")
	}

	// Extract and print the hostname
	hostname := result.Variables[0].Value
	fmt.Printf("Connected to host: %v\n", hostname)
}
