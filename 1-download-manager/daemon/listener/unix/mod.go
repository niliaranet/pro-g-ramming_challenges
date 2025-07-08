package unix

import (
	"log"
	"net"
	"os"
)

const socketUrl string = "/tmp/carrega.sock"

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)

	response := buf[:n]
	conn.Write(response)
	log.Println(string(response))
}

func StartServer() {
	os.Remove(socketUrl)
	listener, _ := net.Listen("unix", socketUrl)
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
