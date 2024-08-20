package reports

import (
	"fmt"
	"io"
	"os"
	"sort"

	"github.dom/Rarfael/quake-log-parser/parser"
)

func GenerateReports(games []parser.Game) {
	var stdOut io.Writer = os.Stdout
	for i, game := range games {
		gameNumber := i + 1
		fmt.Fprintf(stdOut, "\"game_%d\": {\n", gameNumber)
		fmt.Fprintf(stdOut, "  \"total_kills\": %d,\n", game.TotalKills)

		playerNames := make([]string, 0, len(game.Players))
		for _, player := range game.Players {
			playerNames = append(playerNames, player.Name)
		}
		sort.Strings(playerNames)
		fmt.Fprintf(stdOut, "  \"players\": [\"%s\"", playerNames[0])
		for _, player := range playerNames[1:] {
			fmt.Fprintf(stdOut, ", \"%s\"", player)
		}
		fmt.Fprintf(stdOut, "],\n")

		fmt.Fprintf(stdOut, "  \"kills\": {\n")
		first := true
		for _, playerName := range playerNames {
			player := game.Players[playerName]
			if !first {
				fmt.Fprintf(stdOut, ",\n")
			}
			fmt.Fprintf(stdOut, "    \"%s\": %d", player.Name, player.Kills)
			first = false
		}
		fmt.Fprintf(stdOut, "\n  },\n")

		fmt.Fprintf(stdOut, "  \"kills_by_means\": {\n")
		first = true
		for means, count := range game.KillsByMeans {
			if !first {
				fmt.Fprintf(stdOut, ",\n")
			}
			fmt.Fprintf(stdOut, "    \"%s\": %d", means, count)
			first = false
		}
		fmt.Fprintf(stdOut, "\n  }\n")
		fmt.Fprintf(stdOut, "}\n\n")
	}
}
