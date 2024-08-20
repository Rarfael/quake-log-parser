package parser

import (
	"strings"
	"testing"
)

func TestParseLogFile(t *testing.T) {
	logData := `
20:37 InitGame:
20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
21:07 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
21:15 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\xian/default
`
	file := strings.NewReader(logData)

	parser := NewGameParser()
	games := parser.ParseLogFile(file)

	if len(games) != 1 {
		t.Fatalf("Expected 1 game, got %d", len(games))
	}

	game := games[0]
	if game.TotalKills != 2 {
		t.Errorf("Expected 2 total kills, got %d", game.TotalKills)
	}

	if player, exists := game.Players["Isgalamido"]; !exists {
		t.Errorf("Expected player 'Isgalamido' to be in the game")
	} else if player.Kills != -2 {
		t.Errorf("Expected -2 kills for Isgalamido, got %d", player.Kills)
	}

	if count, exists := game.KillsByMeans["MOD_TRIGGER_HURT"]; !exists {
		t.Errorf("Expected 2 kills by MOD_TRIGGER_HURT, got none")
	} else if count != 2 {
		t.Errorf("Expected 2 kills by MOD_TRIGGER_HURT, got %d", count)
	}
}
