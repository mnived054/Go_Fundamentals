package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

var HasMX, HasSPF, HasDMARC bool
var Domain, SPFRecord, DMARCRecord string

func main() {
	fmt.Printf("Enter Domain Name: ")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	Domain = scanner.Text()
	CheckDomain()

	if err := scanner.Err(); err != nil {
		log.Printf("Error: could not read from input: %v", err)
	}
}

func CheckDomain() {
	mxRecords, err := net.LookupMX(Domain)
	if err != nil {
		log.Printf("Error: %v\n", err)
	}

	if len(mxRecords) > 0 {
		HasMX = true
	}

	txtRecords, err := net.LookupTXT(Domain)
	if err != nil {
		log.Println("Error: ", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			HasSPF = true
			SPFRecord = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + Domain)
	if err != nil {
		log.Println("Error:", err)
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			HasDMARC = true
			DMARCRecord = record
			break
		}
	}

	Result()
}

func Result() {
	fmt.Println("Domain: ", Domain)
	fmt.Println("HasMX: ", HasMX)
	fmt.Println("HasSPF: ", HasSPF)
	fmt.Println("SPFRecord: ", SPFRecord)
	fmt.Println("HasDMAC: ", HasDMARC)
	fmt.Println("DMARCRecord: ", DMARCRecord)
}
