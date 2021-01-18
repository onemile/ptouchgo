package ptouchgo

import (
	"io"
	"net"
)

func OpenTCP(address string) (io.ReadWriteCloser, error) {
	return net.Dial("tcp", address)
}
