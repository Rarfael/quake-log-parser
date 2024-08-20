package parser

type Game struct {
	TotalKills   int
	Players      map[string]*Player
	KillsByMeans map[string]int
}

type Player struct {
	Name  string
	Kills int
}
