package main

import (
	"fmt"
	"math/rand"
)

type DiceRoll struct {
	Sides    int
	NumRolls int
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
	Race           Race
	Class          Class
	Portrait       string //Maybe an image name???
	Description    Description
	Exp            int
	Stats          Stats
	Inventory      []Item
	EquipedClothes []Wearable
	EquipedWeapon  []Weapon
	Skills         []Skill
}

type Race struct {
	Name        string
	Image       string
	Description string //Possible front end use???
	Bonuses     map[string]int
}

type Class struct {
	Name          string
	Subclass      string
	Description   string //Possible front end use???
	Proficiencies []string
}

type Description struct {
	Alignment  string
	Ideals     string
	Bonds      string
	Flaws      string
	Background string
	Languages  []string
}

type Stats struct {
	RawStats  map[string]int
	Modifiers map[string]int
}

type Item struct {
	Name        string
	Worth       int //Worth in GP
	Description string
}

type Wearable struct {
	*Item
	AC int
}

type Weapon struct {
	*Item
	WeaponType string
	Damage     DiceRoll //Possibly create a struct of some kind to represent dicerolls???
}

type Skill struct {
	Name        string
	Description string
	Type        string //Defines what type of skill i.e. usable per day/hour/etc
	MaxUses     int
}

func main() {
	d := DiceRoll{6, 2}
	fmt.Println(d.Roll())
}
