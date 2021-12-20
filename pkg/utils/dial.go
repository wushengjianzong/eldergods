package utils

import (
	"net"
	"time"

	"google.golang.org/grpc"
)

func DialGrpcSocketWithTimeout(sockPath string, timeout time.Duration) (*grpc.ClientConn, error) {
	return grpc.Dial(
		sockPath,
		grpc.WithBlock(),
		grpc.WithInsecure(),
		grpc.WithTimeout(timeout),
		grpc.WithDialer(func(s string, d time.Duration) (net.Conn, error) {
			return net.DialTimeout("unix", s, d)
		}),
	)
}
