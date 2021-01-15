package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	ips, err := net.LookupIP(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(ips), func(i, j int) { ips[i], ips[j] = ips[j], ips[i] })

	for _, ip := range ips {
		if ip.To4() != nil {
			fmt.Println(ip.To4().String())
			os.Exit(0)
		}
	}

	os.Exit(1)
}
