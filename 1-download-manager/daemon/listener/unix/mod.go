package unix

import (
	"log"
	"net"
	"os"

	"carrega/daemon/memory"
)

const socketUrl string = "/tmp/carrega.sock"

func StartServer() {
	os.Remove(socketUrl)
	listener, _ := net.Listen("unix", socketUrl)
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	response := buf[:n]
	response = append(response, memory.Finished.ToBytes()...)
	log.Println(memory.Finished)

	conn.Write(response)
}
