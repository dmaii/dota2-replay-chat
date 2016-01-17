package main

import (
	"fmt"
	"os"

	"github.com/suisha/dota2-replay-chat/lib"
	"github.com/suisha/dota2-replay-chat/lib/structs"

	"github.com/dotabuff/manta"
	"github.com/dotabuff/manta/dota"
)

func main() {
	p, _ := manta.NewParserFromFile(os.Args[1])
	//m := make(map[string]bool)
	var gameTime float64 = 0
	var startTime float64 = 0

	game := &structs.Game{}

	p.Callbacks.OnCUserMessageSayText2(func(m *dota.CUserMessageSayText2) error {
		clock := lib.GetGameClock(gameTime, startTime)
		msg := structs.Message{
			Message:    m.GetParam2(),
			PlayerName: m.GetParam1(),
			Clock:      clock,
		}
		game.Messages = append(game.Messages, msg)

		return nil
	})

	p.Callbacks.OnCDemoFileInfo(func(m *dota.CDemoFileInfo) error {
		game.Players = lib.GetPlayers(m)
		game.MatchId = fmt.Sprint(m.GetGameInfo().GetDota().GetMatchId())

		return nil
	})

	p.OnPacketEntity(func(pe *manta.PacketEntity, pet manta.EntityEventType) error {
		if pe.ClassName == "CDOTAGamerulesProxy" {
			gameTime32, _ := pe.FetchFloat32("CDOTAGamerules.m_fGameTime")
			startTime32, _ := pe.FetchFloat32("CDOTAGamerules.m_flGameStartTime")

			gameTime = float64(gameTime32)
			startTime = float64(startTime32)
		}

		return nil
	})

	p.Start()
	fmt.Println(lib.StructToJson(game))
}
