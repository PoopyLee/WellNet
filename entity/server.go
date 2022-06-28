package entity

//通用服务的实体
type Server struct {
	ServerName    string `json:"server_name"`    //服务名称
	ServerVersion string `json:"server_version"` //服务版本
}

func NewServer() *Server {
	return new(Server)
}
