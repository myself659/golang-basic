package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.20.181:22")
	if err != nil {
		fmt.Println("net dial error:", err)
		return
	}
	fmt.Println(conn)
	fmt.Println(conn.RemoteAddr())

	fmt.Println("net dial ok")
	var send [128]byte
	send[0] = 1
	conn.Write(send[0:1])
	ip := net.ParseIP("192.168.70.141")
	raddr := net.TCPAddr{IP: ip, Port: 8000}
	laddr := net.TCPAddr{}
	tconn, terr := net.DialTCP("tcp", &laddr, &raddr)
	if terr != nil {
		fmt.Println("dail tcp failed ", raddr, terr)
		return
	}
	tconn.SetKeepAlive(true)
	fmt.Println(ip)

	tconn.Write(send[0:1])
	time.Sleep(600 * time.Second)
	ifname, e := net.Interfaces()
	if e != nil {
		fmt.Println("errno")
	}

	fmt.Println(ifname)
}
