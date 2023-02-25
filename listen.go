package main

import (
	"fmt"
	goopt "github.com/droundy/goopt"
	"net"
	"time"
)

var amVerbose = goopt.Flag([]string{"-v", "--verbose"}, []string{"--quiet"}, "output verbosely", "be quiet, instead")
var listenAddr = goopt.String([]string{"-a", "--addr"}, "0.0.0.0", "listen address")
var listenPort = goopt.Int([]string{"-p", "--port"}, 6666, "port to listen tos")
var isUdp = goopt.Flag([]string{"-u", "--udp"}, []string{}, "listen for UDP", "")
var isTcp = goopt.Flag([]string{"-t", "--tcp"}, []string{}, "listen for TCP", "")

func main() {
	goopt.Summary = "listen"
	goopt.Parse(nil)
	if *isUdp {
		go udp(*listenAddr, *listenPort)
	}
	if *isTcp {
		go tcp(*listenAddr, *listenPort)
	}
    if *isUdp || *isTcp {
        fmt.Printf("Listen (tcp:%t, udp:%t) on %s:%d\n", *isTcp, *isUdp, *listenAddr, *listenPort)
        for {
            time.Sleep(10 * time.Second)
        }
    } else {
        panic("Provide -t or -u option")
    }
}

func udp(addr string, port int) {
	udpAddr := &net.UDPAddr{IP: net.ParseIP(addr), Port: port}
	sock, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic(err)
	}
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
