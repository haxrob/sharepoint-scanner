/*
Microsoft SharePoint build/version scanner
Author: Twitter @haxrob

https://github.com/robhax/sharepoint-scanner
*/
package main

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type SharePointScanner struct {
	Builds      map[string]Build
	UserAgent   string
	WorkerCount int
	Verbose     bool
	HttpTimeout int
	wg          sync.WaitGroup
	mu          sync.Mutex
}

func (s *SharePointScanner) Run(hostList []string) {
	hostsChan := make(chan string)
	for i := 0; i < s.WorkerCount; i++ {
		s.wg.Add(1)
		go func() {
			for host := range hostsChan {
				s.CheckHost(host)
			}
			s.wg.Done()
		}()

	}

	for _, host := range hostList {
		hostsChan <- host
	}
	close(hostsChan)
	s.wg.Wait()
	fmt.Printf("\n[\033[92m+\033[0m] Done!\n")
}
func (s *SharePointScanner) CheckHost(url string) {

	header := s.FetchHeader(url)
	if header != "" {

		if s.Verbose == true {
			fmt.Println("Got build number: " + header)
		}
		p := strings.Split(header, ".")

		if len(p) == 4 {
			fixedVersion := fmt.Sprintf("%s.0.%s", p[0], p[3])
			if s.Verbose == true {
				fmt.Println("Fixed build number: " + fixedVersion)
			}

			if _, ok := s.Builds[fixedVersion]; ok {
				fmt.Printf("[\033[92m-\033[0m] %s is running %s\n", url, header)
				b := s.Builds[fixedVersion]
				fmt.Println("\tBuild number: " + fixedVersion)
				b.Print()
				return
			}
		}
		fmt.Printf("[\033[92m-\033[0m] %s is returns known version %s. Contact author.\n", url, header)
	}

}
func (s *SharePointScanner) FetchHeader(url string) string {
	if s.Verbose == true {
		fmt.Println("Checking: " + url)
	}
	tr := &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		IdleConnTimeout:   time.Second,
		DisableKeepAlives: true,
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(s.HttpTimeout) * time.Second,
			KeepAlive: time.Second,
		}).DialContext,
	}
	//client := &http.Client{Transport: tr}
	//
	//http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("HEAD", url, nil)

	req.Close = true
	req.Header.Add("User-Agent", s.UserAgent)
	resp, err := client.Do(req)

	if err == nil {
		header := "MicrosoftSharePointTeamServices"
		return resp.Header.Get(header)
	} else {
		//fmt.Println(err)
	}
	return ""
}
