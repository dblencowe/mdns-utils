package main

import (
	"flag"

	"github.com/dblencowe/mdns-utils/pkg/command"
	"github.com/dblencowe/mdns-utils/pkg/mdns"
)

var mdnsAddress string
var mdnsPort string

func init() {
	flag.StringVar(&mdnsAddress, "mdns-address", mdns.MDNSAddress, "Override default mdns address")
	flag.StringVar(&mdnsPort, "mdns-port", mdns.MDNSPort, "Override default mdns port")
}

func main() {
	command.DoListen(mdnsAddress, mdnsPort)
}