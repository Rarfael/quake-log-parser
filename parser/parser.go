package parser

import (
	"bufio"
	"io"
	"regexp"
)

type Player struct {
	Name  string
	Kills int
}

type Game struct {
	TotalKills   int
	Players      map[string]*Player
	KillsByMeans map[string]int
}

type GameParser struct {
	games       []Game
	currentGame *Game
}

func NewGameParser() *GameParser {
	return &GameParser{
		games: []Game{},
	}
}

func (gp *GameParser) ParseLogFile(reader io.Reader) []Game {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		gp.processLine(line)
	}

	gp.finalizeGame()
	return gp.games
}

func (gp *GameParser) processLine(line string) {
	if gp.isInitGame(line) {
		gp.startNewGame()
	} else if gp.isKillEvent(line) {
		gp.processKillEvent(line)
	} else if gp.isPlayerInfoChanged(line) {
		gp.processPlayerInfoChange(line)
	}
}

func (gp *GameParser) isInitGame(line string) bool {
	reInitGame := regexp.MustCompile(`InitGame:`)
	return reInitGame.MatchString(line)
}

func (gp *GameParser) isKillEvent(line string) bool {
	reKill := regexp.MustCompile(`Kill: (\d+) (\d+) (\d+): (.+) killed (.+) by (.+)`)
	return reKill.MatchString(line)
}

func (gp *GameParser) isPlayerInfoChanged(line string) bool {
	reClientUserinfoChanged := regexp.MustCompile(`ClientUserinfoChanged: \d+ n\\([^\\]+)`)
	return reClientUserinfoChanged.MatchString(line)
}

func (gp *GameParser) startNewGame() {
	if gp.currentGame != nil {
		gp.finalizeGame()
	}
	gp.currentGame = &Game{
		TotalKills:   0,
		Players:      make(map[string]*Player),
		KillsByMeans: make(map[string]int),
	}
}

func (gp *GameParser) processKillEvent(line string) {
	reKill := regexp.MustCompile(`Kill: (\d+) (\d+) (\d+): (.+) killed (.+) by (.+)`)
	match := reKill.FindStringSubmatch(line)

	killerID := match[1]
	killerName := match[4]
	victim := match[5]
	meansOfDeath := match[6]

	gp.currentGame.TotalKills++
	gp.currentGame.KillsByMeans[meansOfDeath]++

	if killerID != "1022" {
		gp.addOrUpdatePlayer(killerName, 1)
	}

	gp.addOrUpdatePlayer(victim, -1)
}

func (gp *GameParser) processPlayerInfoChange(line string) {
	reClientUserinfoChanged := regexp.MustCompile(`ClientUserinfoChanged: \d+ n\\([^\\]+)`)
	match := reClientUserinfoChanged.FindStringSubmatch(line)
	playerName := match[1]

	if _, exists := gp.currentGame.Players[playerName]; !exists {
		gp.currentGame.Players[playerName] = &Player{Name: playerName}
	}
}

func (gp *GameParser) addOrUpdatePlayer(playerName string, killChange int) {
	if _, exists := gp.currentGame.Players[playerName]; !exists {
		gp.currentGame.Players[playerName] = &Player{Name: playerName}
	}
	gp.currentGame.Players[playerName].Kills += killChange
}

func (gp *GameParser) finalizeGame() {
	if gp.currentGame != nil {
		gp.games = append(gp.games, *gp.currentGame)
		gp.currentGame = nil
	}
}
