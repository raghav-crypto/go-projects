package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "\n In memory database \n\n"+
		"Use the following commands: \n"+
		"SET key value \n"+
		"GET key \n"+
		"DEL key \n\n"+
		"Example: \n"+
		"SET fav chocolate \n"+
		"GET fav \n\n")

	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		switch fs[0] {
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintf(conn, "EXPECTED VALUE \n")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "GET":
			k := fs[1]
			v := data[k]
			fmt.Fprintf(conn, "%s \n", v)

		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintf(conn, "INVALID COMMAND %s \n", fs[0])
		}
	}

}
