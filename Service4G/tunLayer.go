package Service4G

import "net"

type Tun2Pipe interface {
	GetTarget(conn net.Conn) string
	Finish()
	InputPacket(data []byte)
	Proxying(chan error)
	RemoveFromSession(keyPort int)
}

