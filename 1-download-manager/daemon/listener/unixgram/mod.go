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
	socket, err := startControlSocket(socketUrl)
	if err != nil {
		log.Panicln(err)
	}

	listenForMessages(socket)
}

func startControlSocket(path string) (*net.UnixConn, error) {
	os.Remove(path)
	addr, err := net.ResolveUnixAddr("unixgram", path)
	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUnixgram("unixgram", addr)
	if err != nil {
		return nil, err
	}

	return conn, nil
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

	var ops models.DownloadProcess

	err := process.Download(ops.From(url))
	if err != nil {
		log.Fatal(err)
	}
}
