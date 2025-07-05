package signal

import "net"

const socketUrl = "/tmp/carrega.sock"

func Send(message string) {
	addr, _ := net.ResolveUnixAddr("unixgram", socketUrl)
	conn, _ := net.DialUnix("unixgram", nil, addr)
	conn.Write([]byte(message))
	conn.Close()
}
