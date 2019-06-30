package models

type masterServer struct {
	ip       string
	port     int
	hostname string
	server   []gameServer
}

//TODO: implement String function

// NewMasterServer returns a new instance of the struct masterServer
func NewMasterServer(ip string, port int, hostname string) masterServer {
	ms := masterServer{ip: ip, port: port, hostname: hostname}
	return ms
}

func (ms masterServer) addServer	