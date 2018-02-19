package main

import (
	"fmt"
	"math/rand"
)

type DiceRoll struct {
	Sides    int `json:"sides"`
	NumRolls int `json:"rolls"`
}

func (d *DiceRoll) Roll() int {
	sum := 0
	for i := 0; i < d.NumRolls; i += 1 {
		roll := (rand.Intn(d.Sides) + 1)
		fmt.Println("Roll " + string(i) + " is " + string(roll))
		sum += roll
	}

	return sum
}

type Character struct {
	Name           string
	Race           *Race       `json:"race"`     //Making it a pointer is more effiecient because there are a
	Class          *Class      `json:"class"`    //set number of instatiated races and classes that we can just reference
	Portrait       string      `json:"portrait"` //Maybe an image name???
	Description    Description `json:"desc"`
	Exp            int         `json:"exp"`
	Stats          Stats       `json:"stats"`
	Inventory      []Item      `json:"inventory"`
	EquipedClothes []Wearable  `json:"clothes"`
	EquipedWeapon  []Weapon    `json:"weapons"`
	Skills         []Skill     `json:"skills"`
}

type Race struct {
	Name        string         `json:"name"`
	Image       string         `json:"image"`
	Description string         `json:"desc"` //Possible front end use???
	Bonuses     map[string]int `json:"bonuses"`
}

type Class struct {
	Name          string   `json:"name"`
	Subclass      string   `json:"subclass"`
	Description   string   `json:"desc"` //Possible front end use???
	Proficiencies []string `json:"proficiencies"`
}

type Description struct {
	Alignment  string   `json:"alignment"`
	Ideals     string   `json:"ideals"`
	Bonds      string   `json:"bonds"`
	Flaws      string   `json:"flaws"`
	Background string   `json:"background"`
	Languages  []string `json:"languages"`
}

type Stats struct {
	RawStats  map[string]int `json:"rawStats"`
	Modifiers map[string]int `json:"modifiers"`
}

type Item struct {
	Name        string `json:"name"`
	Worth       int    `json:"worth"` //Worth in GP
	Description string `json:"desc"`
}

type Wearable struct {
	Item
	AC int `json:"ac"`
}

type Weapon struct {
	Item
	WeaponType string   `json:"type"`
	Damage     DiceRoll `json:"damage"` //Possibly create a struct of some kind to represent dicerolls???
}

type Skill struct {
	Name        string `json:"name"`
	Description string `json:"desc"`
	Type        string `json:"type"` //Defines what type of skill i.e. usable per day/hour/etc
	MaxUses     int    `json:"numUses"`
}

func makeGerald() *Character {
	bonusMap := map[string]int{"Strength": 1, "Intelligence": 1}

	human := &Race{
		Name:        "Human",
		Image:       "benis.jpg",
		Description: "Pretty boring for now.",
		Bonuses:     bonusMap,
	}

	warrior := &Class{
		Name:          "Warrior",
		Subclass:      "Meme Crusader",
		Description:   "Nothing for now",
		Proficiencies: []string{"Maymays", "Swords"},
	}

	desc := Description{
		Alignment:  "n/a",
		Ideals:     "Nihilism",
		Bonds:      "n/a",
		Flaws:      "most things",
		Background: "idk",
		Languages:  []string{"memes", "Latvian"},
	}

	stat := Stats{
		RawStats:  map[string]int{"Strength": 15},
		Modifiers: map[string]int{"Strength": 1},
	}

	gerald := Character{
		Name:           "Gerald",
		Race:           human,
		Class:          warrior,
		Portrait:       "Fug.jpg",
		Description:    desc,
		Exp:            15,
		Stats:          stat,
		Inventory:      []Item{},
		EquipedClothes: []Wearable{},
		EquipedWeapon: []Weapon{Weapon{
			Item: Item{
				Name:        "memeSword",
				Worth:       69,
				Description: "maymay powered"},
			WeaponType: "MayMays",
			Damage:     DiceRoll{NumRolls: 2, Sides: 10},
		}},
		Skills: []Skill{},
	}

	return &gerald
}
