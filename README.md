# Microsoft SharePoint Build / Vulnerability scanner

A Multithreaded Microsoft SharePoint version / vulnerability scanner

## Installation
### Binaries
Compiled 64-bit executable files for Windows, Mac and Linux are available [here](https://github.com/robhax/sharepoint-scanner/releases/)

### Building
If you would prefer to build yourself (and Go is setup [correctly](https://golang.org/doc/install)):
```
git clone https://github.com/robhax/sharepoint-scanner
cd sharepoint-scanner
go build .
```

# Usage 
```
Usage: sharepoint-scanner [options]

  -f string
        File containing list of hosts
  -h string
        Single host
  -n string
        Network in CIDR format (e.g. 192.168.0.0/24)
  -t int
        HTTP timeout (seconds) (default 2)
  -u string
        Custom user agent string
  -v    Verbose
  -w int
        Number of concurrent workers (default 20)

Hosts can be either IP addresses or domain names or URIs.
Both port 80 and 443 (https) will be scanned.
```
## Example
```
❯ ./sharepoint-scanner -n 192.168.0.1/24
[+] Testing 254 host(s) with 20 concurrent worker(s) ..

[-] https://192.168.0.215 is running 16.0.0.10354
        Build number: 16.0.10354
        Build Name: January 2020 CU
        Release Date: 2020 January 14
        Component: SharePoint Server 2019 (language independent)
        Description: KB4484224

[+] Done!
```
Hosts in the host file (`-f`) or single host (`-h`) can be IP addresses, domains or URIs, e.g.
```
192.168.0.1
https://example.server.com/
http://10.10.0.1:8443
```
