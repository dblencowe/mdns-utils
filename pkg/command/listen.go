package command

import (
	"fmt"
	"net"
	"os"
	"encoding/json"
	"github.com/dblencowe/mdns-utils/pkg/mdns"
)

var OutputStream = os.Stdout

const bufferSize = 8 * 1024

func DoListen(ip string, port string) error {
	c := &Listen{}
	return c.Execute(ip, port)
}

type Listen struct{}

func (c *Listen) Execute(ip string, port string) error {
	address := fmt.Sprintf("%s:%s", ip, port)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		return err
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return err
	}

	buffer := make([]byte, bufferSize)
	enc := json.NewEncoder(OutputStream)
	for {
		n, from, err := conn.ReadFrom(buffer)
		if err != nil {
			return err
		}

		p, err := mdns.ParsePacket(buffer[:n], from)
		if err != nil {
			return err
		}

		err = enc.Encode(p)
		if err != nil {
			return err
		}
	}
}
