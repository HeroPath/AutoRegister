package src

import (
	"fmt"
	"strconv"
)

type Npc struct {
	Name        string `json:"name"`
	Level       int    `json:"level"`
	GiveMaxExp  int    `json:"giveMaxExp"`
	GiveMinExp  int    `json:"giveMinExp"`
	GiveMaxGold int    `json:"giveMaxGold"`
	GiveMinGold int    `json:"giveMinGold"`
	Hp          int    `json:"hp"`
	MaxHp       int    `json:"maxHp"`
	MinDmg      int    `json:"minDmg"`
	MaxDmg      int    `json:"maxDmg"`
	Defense     int    `json:"defense"`
	Zone        string `json:"zone"`
}

func GetNpcs() []Npc {
	npcs := make([]Npc, 0)
	sheet := readFile("Npcs")

	for _, row := range sheet {
		if _, err := strconv.Atoi(row.Cells[0].String()); err == nil {
			npc := Npc{}
			npc.Name = row.Cells[1].String()
			npc.Level, _ = strconv.Atoi(row.Cells[2].String())
			npc.GiveMaxExp, _ = strconv.Atoi(row.Cells[3].String())
			npc.GiveMinExp, _ = strconv.Atoi(row.Cells[4].String())
			npc.GiveMaxGold, _ = strconv.Atoi(row.Cells[5].String())
			npc.GiveMinGold, _ = strconv.Atoi(row.Cells[6].String())
			npc.Hp, _ = strconv.Atoi(row.Cells[7].String())
			npc.MaxHp, _ = strconv.Atoi(row.Cells[8].String())
			npc.MinDmg, _ = strconv.Atoi(row.Cells[9].String())
			npc.MaxDmg, _ = strconv.Atoi(row.Cells[10].String())
			npc.Defense, _ = strconv.Atoi(row.Cells[11].String())
			npc.Zone = row.Cells[12].String()
			npcs = append(npcs, npc)
		}
	}
	return npcs
}

func RegisterNpcs(url string, token string) {
	npcs := GetNpcs()
	for _, npc := range npcs {
		_, status := PostRequest(
			url+"npcs",
			npc,
			token,
		)
		fmt.Println(status)
	}
}
