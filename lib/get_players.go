package lib

import (
	"fmt"
	"github.com/dotabuff/manta/dota"
	"github.com/suisha/dota2-replay-chat/lib/structs"
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
			Team:    pi.GetGameTeam(),
		}

		players = append(players, p)
	}

	return players
}
