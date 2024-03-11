package idgenclient

import "github.com/prakash-p-3121/idgenclient/impl"

func NewIDGenClient(host string, port uint) IDGenClient {
	return &impl.IDGenClientImpl{Host: host, Port: port}
}
