package network

import (
	"fmt"
	"net"
)

func handleClient(conn net.Conn) {
	fmt.Printf("New connection: %s", conn.LocalAddr().String())
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		_, err2 := conn.Write(buf[0:n])
		fmt.Print(string(buf[0:]))
		if err2 != nil {
			return
		}

	}
}
func SingleThreadedEcho() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
		conn.Close()
	}
}

func MultiThreadedEcho() {
	service := ":1201"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	fmt.Println("Tcp Addr", tcpAddr)
	checkError(err)

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
		// conn.Close()
	}
}
