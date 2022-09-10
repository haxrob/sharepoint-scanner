/*
Microsoft SharePoint build/version scanner
Author: Twitter @haxrob

https://github.com/robhax/sharepoint-scanner
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {

	var (
		hostsList        []string
		workerCount      int
		networkRange     string
		hostListFile     string
		timeout          int
		verbose          bool
		userAgent        string
		singleHost       string
		hostsListScanner *bufio.Scanner
	)
	flag.IntVar(&workerCount, "w", 20, "Number of concurrent workers")
	flag.StringVar(&networkRange, "n", "", "Network in CIDR format (e.g. 192.168.0.0/24)")
	flag.StringVar(&hostListFile, "f", "", "File containing list of hosts")
	flag.IntVar(&timeout, "t", 2, "HTTP timeout (seconds)")
	flag.BoolVar(&verbose, "v", false, "Verbose")
	flag.StringVar(&userAgent, "u", "", "Custom user agent string")
	flag.StringVar(&singleHost, "h", "", "Single host")
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Sharepoint Scanner v0.01 - https://github.com/robhax/sharepoint-scanner\n\n")
		fmt.Fprintf(os.Stderr, "Usage: %s [options]\n\n", os.Args[0])
		flag.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\nHosts can be IP addresses or URIs.\n")
		fmt.Fprint(os.Stderr, "Both port 80 and 443 (https) will be scanned.\n\n")
	}

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Error: must specify -h <host> or -n <subnet> or -f <hosts file>\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if userAgent == "" {
		userAgent = "Mozilla/5.0 (Windows NT 6.4; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2225.0 Safari/537.36"
	}

	if networkRange != "" {
		addNetwork(networkRange, &hostsList)
	}
	if singleHost != "" {
		hostsList = append(hostsList, singleHost)
		workerCount = 1
	}
	if hostListFile != "" {
		file, err := os.Open(hostListFile)
		defer file.Close()
		if err != nil {
			log.Fatal(err)
		}

		hostsListScanner = bufio.NewScanner(file)
		for hostsListScanner.Scan() {
			host := hostsListScanner.Text()
			if len(host) < 8 {
				continue
			}
			// network has been specified
			if host[len(host)-3] == '/' {
				addNetwork(host, &hostsList)
			} else {
				hostsList = append(hostsList, host)
			}
		}
	}

	var finalList []string
	for _, host := range hostsList {
		if addr := net.ParseIP(host); addr != nil {
			finalList = append(finalList, "http://"+host)
			finalList = append(finalList, "https://"+host)
		} else {
			finalList = append(finalList, host)
		}

	}

	s := SharePointScanner{
		UserAgent:   userAgent,
		WorkerCount: workerCount,
		Verbose:     verbose,
		HttpTimeout: timeout,
	}
	s.PopulateBuilds()

	fmt.Printf("[\033[92m+\033[0m] Testing %d host(s) with %d concurrent worker(s) ..\n\n", len(hostsList), workerCount)
	s.Run(finalList)

}

func addNetwork(network string, list *[]string) {
	for _, ip := range netExpand(network) {
		*list = append(*list, ip)
	}
}

func netExpand(network string) []string {
	var ips []string
	ip, ipnet, err := net.ParseCIDR(network)
	if err != nil {
		log.Fatal(err)
	}
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	return ips[1 : len(ips)-1]
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
