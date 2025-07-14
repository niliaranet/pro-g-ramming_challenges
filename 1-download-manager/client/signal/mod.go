package signal

import "net"

const unixgramUrl = "/tmp/carregagram.sock"
const unixUrl = "/tmp/carrega.sock"

func Send(message string) (string, error) {
	conn, err := net.Dial("unix", unixUrl)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	conn.Write([]byte(message))

	buf := make([]byte, 1024)
	n, _ := conn.Read(buf)
	return string(buf[:n]), nil
}

func SendDownload(message string) error {
	addr, err := net.ResolveUnixAddr("unixgram", unixgramUrl)
	if err != nil {
		return err
	}

	conn, err := net.DialUnix("unixgram", nil, addr)
	if err != nil {
		return err
	}

	conn.Write([]byte(message))
	conn.Close()

	return nil
}
