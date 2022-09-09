# sharepoint-scanner
Multithreaded Microsoft SharePoint version / vulnerability scanner

## Installation
### Binaries
Compiled 64-bit executable files for Windows, Mac and Linux are available [here](https://github.com/robhax/sharepoint-scanner/releases/)

### Go get
If you would prefer to build yourself (and Go is setup [correctly](https://golang.org/doc/install)):
```
go get -u github.com/robhax/sharepoint-scanner
```

# Usage 
```
Usage: ./sharepoint-scanner [options]

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
```
