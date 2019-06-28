package models

type gameServer struct {
	ip         string
	port       int
	players    []player
	serverType string
	version    string
	name       string
	mapName    string
	mapCrc     int
	mapSize    int
	gameType   string
	flags      int
	numPlayers int
	maxPlayers int
	numClients int
	maxClients int
}

//TODO: implement String function

//NewGameServer returns a new instance of the struct gameServer
func NewGameServer(ip string, port int) gameServer {
	gs := gameServer{ip: ip, port: port}
	return gs
}

func (gs gameServer) Players() []player {
	return gs.players
}

func (gs gameServer) addPlayer(player player) {
	gs.players = append(gs.players, player)
}
