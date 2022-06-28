package wellnet

import (
	"github.com/lvwei25/WellNet/entity"
	face "github.com/lvwei25/WellNet/interface"
)

func WNTcpServer() face.ServerInterface {
	return &entity.TcpServer{
		Server: entity.NewServer(),
		Addr:   "127.0.0.1:9999",
	}
}
