package main

import (
	"encoding/binary"
	"fmt"
	"net"
)

type netflowPacketHeader struct {
	Version   int16
	Count     int16
	Uptime    int32
	Sequence  int32
	Id        int32
	FlowSetId netflowPacketFlowsetId
	Length    netflowPacketTemplate
}

type netflowPacketFlowsetId struct {
	FlowSetID int16
}

type netflowPacketTemplate struct {
	Length int16
}

func main() {

	addr := net.UDPAddr{
		Port: 9999,
		IP:   net.ParseIP("127.0.0.1"),
	}
	conn, err := net.ListenUDP("udp", &addr)

	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	p := netflowPacketHeader{}
	err = binary.Read(conn, binary.BigEndian, &p)
	if err != nil {
		fmt.Printf("Some error %v\n", err)
		return
	}

	fmt.Printf("(%v) Int: %v %v %v\n", conn.RemoteAddr(), p.Version, p.Count, p.Length)
	if p.FlowSetId.FlowSetID == 0 {
		t := netflowPacketTemplate{}
		err = binary.Read(conn, binary.BigEndian, &t)
		if err != nil {
			fmt.Printf("Some error %v\n", err)
			return
		}
	}
	// Buffer creates an array of bytes
	//buffer := make([]byte, 1024)

	// Read the number of bytes (1024) into a variable length slice of bytes, 'Buffer'
	//count, _ := conn.Read(buffer)

	// 'Count' refers to the number of bytes received in the slice
	// Below we decode that amount (buffer_slice[:number_of_bytes]) as a string
	// fmt.Println(string(buffer[:count]))

}
