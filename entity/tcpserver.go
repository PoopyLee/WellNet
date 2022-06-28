package entity

import (
	"fmt"
	_ "github.com/lvwei25/WellNet/interface"
)

//TcpServer Entity
type TcpServer struct {
	Server *Server //服务基本信息
	Addr   string
}

func (this *TcpServer) Run() {
	fmt.Println(this.Addr)
}
