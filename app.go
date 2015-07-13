package fluenter

import (
	"log"
	"net"
)

var tdConn *net.UDPConn

func initUDP(host string) {
	if host == "" {
		host = "ip-10-0-11-83.ap-southeast-1.compute.internal:7777"
	}

	//default AWS IP
	fluentd, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		log.Fatal(err)
	}

	tdConn, err = net.DialUDP("udp", nil, fluentd)
	if err != nil {
		log.Fatal(err)
	}
}

func UdpToS3(tag string, jsonPacket string) {
	buf := []byte(packet)
	tdConn.Write(buf)
}
