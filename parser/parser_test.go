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

func TestParseLogFile_MultipleGames(t *testing.T) {
	logData := `
0:00 InitGame:
0:20 Kill: 3 2 7: Zeh killed Isgalamido by MOD_ROCKET
1:00 InitGame:
1:20 Kill: 2 3 10: Isgalamido killed Zeh by MOD_RAILGUN
1:50 Kill: 2 3 10: Isgalamido killed Zeh by MOD_RAILGUN
`
	file := strings.NewReader(logData)

	parser := NewGameParser()
	games := parser.ParseLogFile(file)

	if len(games) != 2 {
		t.Fatalf("Expected 2 games, got %d", len(games))
	}

	game1 := games[0]
	if game1.TotalKills != 1 {
		t.Errorf("Expected 1 total kill in game 1, got %d", game1.TotalKills)
	}
	if player, exists := game1.Players["Zeh"]; !exists {
		t.Errorf("Expected player 'Zeh' to be in game 1")
	} else if player.Kills != 1 {
		t.Errorf("Expected 1 kill for Zeh in game 1, got %d", player.Kills)
	}
	if count, exists := game1.KillsByMeans["MOD_ROCKET"]; !exists || count != 1 {
		t.Errorf("Expected 1 kill by MOD_ROCKET in game 1, got %d", count)
	}

	game2 := games[1]
	if game2.TotalKills != 2 {
		t.Errorf("Expected 2 total kills in game 2, got %d", game2.TotalKills)
	}
	if player, exists := game2.Players["Isgalamido"]; !exists {
		t.Errorf("Expected player 'Isgalamido' to be in game 2")
	} else if player.Kills != 2 {
		t.Errorf("Expected 2 kills for Isgalamido in game 2, got %d", player.Kills)
	}
	if count, exists := game2.KillsByMeans["MOD_RAILGUN"]; !exists || count != 2 {
		t.Errorf("Expected 2 kills by MOD_RAILGUN in game 2, got %d", count)
	}
}

func TestParseLogFile_NoKills(t *testing.T) {
	logData := `
0:00 InitGame:
1:00 ClientUserinfoChanged: 2 n\Isgalamido\t\0\model\xian/default
2:00 ClientUserinfoChanged: 3 n\Zeh\t\0\model\uriel/default
`
	file := strings.NewReader(logData)

	// Create an instance of GameParser
	parser := NewGameParser()
	games := parser.ParseLogFile(file)

	if len(games) != 1 {
		t.Fatalf("Expected 1 game, got %d", len(games))
	}

	game := games[0]
	if game.TotalKills != 0 {
		t.Errorf("Expected 0 total kills, got %d", game.TotalKills)
	}

	if len(game.Players) != 2 {
		t.Errorf("Expected 2 players, got %d", len(game.Players))
	}

	if _, exists := game.KillsByMeans["MOD_ROCKET"]; exists {
		t.Errorf("Expected no kills by MOD_ROCKET, but found one")
	}
}

func TestParseLogFile_WorldKillsPlayer(t *testing.T) {
	logData := `
0:00 InitGame:
0:30 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT
`
	file := strings.NewReader(logData)

	// Create an instance of GameParser
	parser := NewGameParser()
	games := parser.ParseLogFile(file)

	if len(games) != 1 {
		t.Fatalf("Expected 1 game, got %d", len(games))
	}

	game := games[0]
	if game.TotalKills != 1 {
		t.Errorf("Expected 1 total kill, got %d", game.TotalKills)
	}

	if player, exists := game.Players["Isgalamido"]; !exists {
		t.Errorf("Expected player 'Isgalamido' to be in the game")
	} else if player.Kills != -1 {
		t.Errorf("Expected -1 kills for Isgalamido, got %d", player.Kills)
	}

	if count, exists := game.KillsByMeans["MOD_TRIGGER_HURT"]; !exists || count != 1 {
		t.Errorf("Expected 1 kill by MOD_TRIGGER_HURT, got %d", count)
	}
}
