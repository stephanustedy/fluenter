package fluenter

import (
	"log"
	"net"
)

type fluentdUDPConn struct {
	conn *net.UDPConn
}

func initUDP(host string) *fluentdUDPConn {
	if host == "" {
		host = "ip-10-0-11-83.ap-southeast-1.compute.internal:7777"
	}

	//default AWS IP
	fluentd, err := net.ResolveUDPAddr("udp", host)
	if err != nil {
		log.Fatal(err)
	}

	tdConn, err := net.DialUDP("udp", nil, fluentd)
	if err != nil {
		log.Fatal(err)
	}

	return &fluentdUDPConn{
		conn : tdConn,
	}
}

func (tdClient *fluentdUDPConn)UdpToS3(tag string, jsonPacket string) {
	buf := []byte(jsonPacket)
	tdClient.conn.Write(buf)
}
