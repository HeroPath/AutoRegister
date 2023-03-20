package src

import (
	"strconv"
)

type Quest struct {
	Name           string `json:"name"`
	Description    string `json:"description"`
	LevelRequired  int    `json:"levelRequired"`
	NameNpcKill    string `json:"nameNpcKill"`
	NpcAmountNeed  int    `json:"npcAmountNeed"`
	UserAmountNeed int    `json:"userAmountNeed"`
	GiveExp        int    `json:"giveExp"`
	GiveGold       int    `json:"giveGold"`
	GiveDiamonds   int    `json:"giveDiamonds"`
}

func GetQuests() []Quest {
	quests := make([]Quest, 0)
	sheet := readFile("Quests")

	for _, row := range sheet {
		if _, err := strconv.Atoi(row.Cells[0].String()); err == nil {
			quest := Quest{}
			quest.Name = row.Cells[1].String()
			quest.Description = row.Cells[2].String()
			quest.LevelRequired, _ = strconv.Atoi(row.Cells[3].String())
			quest.NameNpcKill = row.Cells[4].String()
			quest.NpcAmountNeed, _ = strconv.Atoi(row.Cells[5].String())
			quest.UserAmountNeed, _ = strconv.Atoi(row.Cells[6].String())
			quest.GiveExp, _ = strconv.Atoi(row.Cells[7].String())
			quest.GiveGold, _ = strconv.Atoi(row.Cells[8].String())
			quest.GiveDiamonds, _ = strconv.Atoi(row.Cells[9].String())
			quests = append(quests, quest)
		}
	}
	return quests
}

func RegisterQuests(url string, token string) string {
	quests := GetQuests()
	for _, quest := range quests {
		PostRequest(url+"quests", quest, token)
	}
	return ""
}
