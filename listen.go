package main

import (
	"fmt"
	"net"
    goopt "github.com/droundy/goopt"
)

var amVerbose = goopt.Flag([]string{"-v", "--verbose"}, []string{"--quiet"}, "output verbosely", "be quiet, instead")
var listenAddr = goopt.String([]string{"-a", "--addr"}, "0.0.0.0", "listen address")
var listenPort = goopt.Int([]string{"-p", "--port"}, 6666, "port to listen tos")
var isUdp = goopt.Flag([]string{"-u", "--udp"}, []string{}, "UDP instead of TCP", "")


func main() {
    goopt.Summary = "listen"
	goopt.Parse(nil)
	if (*isUdp) {
		udp(*listenAddr, *listenPort)
	} else {
		tcp(*listenAddr, *listenPort)
	}
}

func udp(addr string, port int) {
	udpAddr := &net.UDPAddr{IP: net.ParseIP(addr), Port: port}
	sock, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Listening UDP on %s:%d\n--\n", addr, port)
	defer sock.Close()

	// Create buffer for receiving data.
	buf := make([]byte, 1024)

	// Enter infinite loop to handle clients.
	for {
		n, addr, err := sock.ReadFromUDP(buf)
		if err != nil {
			panic(err)
		}
		go func() {
			fmt.Printf("[%s]: %s", addr.IP.String(), string(buf[:n]))
		}()
	}
}

func tcp(addr string, port int) {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Printf("Listening TCP on %s:%d\n--\n", addr, port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		go func() {
			fmt.Printf("[%s]: %s", addr, string(buf[:n]))
		}()
	}
}
