package reports

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.dom/Rarfael/quake-log-parser/parser"
)

func TestGenerateReportsWithKillsByMeans(t *testing.T) {
	var stdOut io.Writer = os.Stdout

	games := []parser.Game{
		{
			TotalKills: 45,
			Players: map[string]*parser.Player{
				"Dono da bola": {Name: "Dono da bola", Kills: 5},
				"Isgalamido":   {Name: "Isgalamido", Kills: 18},
				"Zeh":          {Name: "Zeh", Kills: 20},
			},
			KillsByMeans: map[string]int{
				"MOD_SHOTGUN":  10,
				"MOD_RAILGUN":  2,
				"MOD_GAUNTLET": 1,
			},
		},
	}

	expectedOutput := "\"game_1\": {\n" +
		"  \"total_kills\": 45,\n" +
		"  \"players\": [\"Dono da bola\", \"Isgalamido\", \"Zeh\"],\n" +
		"  \"kills\": {\n" +
		"    \"Dono da bola\": 5,\n" +
		"    \"Isgalamido\": 18,\n" +
		"    \"Zeh\": 20\n" +
		"  },\n" +
		"  \"kills_by_means\": {\n" +
		"    \"MOD_SHOTGUN\": 10,\n" +
		"    \"MOD_RAILGUN\": 2,\n" +
		"    \"MOD_GAUNTLET\": 1\n" +
		"  }\n" +
		"}\n\n"

	var buf bytes.Buffer
	oldStdout := stdOut
	defer func() { stdOut = oldStdout }()
	stdOut = &buf

	GenerateReports(games)

	actualOutput := buf.String()
	if actualOutput != expectedOutput {
		t.Errorf("GenerateReports() output did not match expected output.\nActual Output:\n%s\nExpected:\n%s", actualOutput, expectedOutput)
	}
}
