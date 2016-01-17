package lib

import (
	"dota2parser/lib/structs"
	"fmt"
	"github.com/dotabuff/manta/dota"
)

func GetPlayers(fileInfo *dota.CDemoFileInfo) []structs.Player {
	playerInfos := fileInfo.GetGameInfo().GetDota().GetPlayerInfo()

	var players []structs.Player

	for i := range playerInfos {
		pi := playerInfos[i]
		p := structs.Player{
			Name:    pi.GetPlayerName(),
			Hero:    pi.GetHeroName(),
			SteamId: fmt.Sprint(pi.GetSteamid()),
		}

		players = append(players, p)
	}

	return players
}
