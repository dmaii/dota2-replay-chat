package lib

import (
	"strings"
)

func GetLocalHero(unlocalized string) string {
	prefix := "npc_dota_hero_"
	suffix := unlocalized[len(prefix):len(unlocalized)]

	split := strings.Split(suffix, "_")
	for i := range split {
		split[i] = strings.ToUpper(split[i][0:1]) +
			split[i][1:len(split[i])]
	}

	return strings.Join(split, " ")
}
