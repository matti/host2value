package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func suffle(elements []string) []string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(elements), func(i, j int) { elements[i], elements[j] = elements[j], elements[i] })

	return elements
}

func main() {
	var recordType string
	flag.StringVar(&recordType, "recordType", "A", "DNS record type")
	var results int
	flag.IntVar(&results, "results", 1, "DNS record type")

	flag.Parse()

	host := flag.Arg(0)

	var records []string
	switch recordType {
	case "A":
		ips, err := net.LookupIP(host)
		if err != nil {
			log.Fatalln(err)
		}
		for _, ip := range ips {
			if ip.To4() != nil {
				records = append(records, ip.To4().String())
			}
		}
	case "TXT":
		txts, err := net.LookupTXT(host)
		if err != nil {
			log.Fatalln(err)
		}

		records = txts
	case "NS":
		nss, err := net.LookupNS(host)
		if err != nil {
			log.Fatalln(err)
		}
		for _, ns := range nss {
			records = append(records, ns.Host)
		}
	}
	for i, record := range suffle(records) {
		if i == results {
			os.Exit(0)
		}
		fmt.Println(record)
	}
}
