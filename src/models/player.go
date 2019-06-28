package models

import "fmt"

type player struct {
	name     string
	clan     string
	country  int
	score    int
	isIngame bool
}

//NewPlayer returns a new instance of the struct player
func NewPlayer(name string, clan string, country int, score int, isIngame bool) player {
	if country == 0 {
		country = -1
	}
	p := player{name, clan, country, score, isIngame}
	return p
}

func (p player) String() string {
	return fmt.Sprintf("name: %s, clan: %s, country: %d, score: %d, isIngame: %t", p.name, p.clan, p.country, p.score, p.isIngame)
}
