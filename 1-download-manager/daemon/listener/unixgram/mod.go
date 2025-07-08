package unixgram

import (
	"log"
	"net"
	"os"
	"strings"

	"carrega/daemon/models"
	"carrega/daemon/process"
)

const socketUrl string = "/tmp/carregagram.sock"

func StartListener() {
	socket := startControlSocket(socketUrl)
	listenForMessages(socket)
}

func startControlSocket(path string) *net.UnixConn {
	os.Remove(path)
	addr, _ := net.ResolveUnixAddr("unixgram", path)
	conn, _ := net.ListenUnixgram("unixgram", addr)
	return conn
}

func listenForMessages(conn *net.UnixConn) {
	buf := make([]byte, 1024)
	for {
		n, _ := conn.Read(buf)
		call := string(buf[:n])
		args := strings.Split(call, " ")
		switch args[0] {
		case "download":
			downloadCall(args)
		}
	}
}

func downloadCall(args []string) {
	if len(args) < 2 {
		return
	}
	url := args[1]

	var ops models.DownloadOptions

	err := process.Download(ops.From(url))
	if err != nil {
		log.Fatal(err)
	}
}
