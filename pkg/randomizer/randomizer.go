package randomizer

import (
	"encoding/json"
	"math/rand"
	"time"

	"github.com/KazuhiroIto0127/gachagacha-cli/resources"
)

type Rarity struct {
	Name string
	Rate int
	Start int
	End int
}

type Item struct {
	Name string `json:"name"`
	Aa string `json:"aa"`
}

type Items struct {
	SSR []Item `json:"SSR"`
	SR []Item `json:"SR"`
	R []Item `json:"R"`
}

var rarities = []Rarity{
	{"SSR", 3, 0, 3},
	{"SR", 18, 3, 21},
	{"R", 79, 21, 100},
}

func GetRarity() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	result := r.Intn(100)

	for _, rarity := range rarities {
		if result >= rarity.Start && result < rarity.End {
			return rarity.Name
		}
	}
	return "Unknown"
}

func GetRandomItem(rarity string) (Item, error) {
	data, err := resources.Items.ReadFile("items.json")

	if err != nil{
		return Item{}, err
	}

	var items Items
	err = json.Unmarshal(data, &items)
	if err != nil {
		return Item{}, err
	}

	var itemList []Item
	switch rarity {
	case "SSR":
		itemList = items.SSR
	case "SR":
		itemList = items.SR
	case "R":
		itemList = items.R
	default:
		return Item{}, nil
	}

	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	return itemList[r.Intn(len(itemList))], nil
}
