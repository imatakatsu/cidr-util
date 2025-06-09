## cidr-util
small lib made to generate ip address using cidr.
work with IPv4 and IPv6 and can be used for proxy servers and etc
return net.IP type

developed for NORA proxy server

### example code
```go
package main

import (
	"fmt"
	"log"

	util "github.com/imatakatsu/cidr-util"
)

func main() {
	ipv4_cidr, err := util.ParseCIDR("0.0.0.0/16")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4_cidr.RandomIP().String())
	ipv6_cidr, err := util.ParseCIDR("fe80::f816:3eff:fec3:754c/64")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv6_cidr.RandomIP().String())
}

```

### example output
```terminal
# go run test.go
0.0.172.39
fe80::b831:166d:febb:487a
```
