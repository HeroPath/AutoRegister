package src

import (
	"fmt"
	"strconv"
)

type Item struct {
	Name          string
	Type          string
	LvlMin        int
	ClassRequired string
	Price         int
	Strength      int
	Dexterity     int
	Intelligence  int
	Vitality      int
	Luck          int
}

func GetItems() {

	items := make([]Item, 0)
	sheet := readFile("Items")

	for _, row := range sheet {
		if _, err := strconv.Atoi(row.Cells[0].String()); err == nil {
			item := Item{}
			item.Name = row.Cells[1].String()
			item.Type = row.Cells[2].String()
			item.LvlMin, _ = strconv.Atoi(row.Cells[3].String())
			item.ClassRequired = row.Cells[4].String()
			item.Price, _ = strconv.Atoi(row.Cells[5].String())
			item.Strength, _ = strconv.Atoi(row.Cells[6].String())
			item.Dexterity, _ = strconv.Atoi(row.Cells[7].String())
			item.Intelligence, _ = strconv.Atoi(row.Cells[8].String())
			item.Vitality, _ = strconv.Atoi(row.Cells[9].String())
			item.Luck, _ = strconv.Atoi(row.Cells[10].String())
			items = append(items, item)
		}
	}
	fmt.Println(items)
}
