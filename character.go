package character

import (
	"fmt"
	"math/rand"
	"encoding/json"
)

type DiceRoll struct {
	Sides    int 'json:"sides"'
	NumRolls int 'json:"rolls"'
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
	Name           string 'json:"name"'
	Race           Race	'json:"race"'
	Class          Class 'json:"class"'
	Portrait       string 'json:"portrait"'//Maybe an image name???
	Description    Description 'json:"desc"'
	Exp            int 'json:"exp"'
	Stats          Stats 'json:"stats"'
	Inventory      []Item 'json:"inventory"'
	EquipedClothes []Wearable 'json:"clothes"'
	EquipedWeapon  []Weapon 'json:"weapons"'
	Skills         []Skill 'json:"skills"'
}

type Race struct {
	Name        string 'json:"name"'
	Image       string 'json:"image"'
	Description string 'json:"desc"'//Possible front end use???
	Bonuses     map[string]int 'json:"bonuses"'
}

type Class struct {
	Name          string 'json:"name"'
	Subclass      string 'json:"subclass"'
	Description   string 'json:"desc"'//Possible front end use???
	Proficiencies []string 'json:"proficiencies"'
}

type Description struct {
	Alignment  string 'json:"alignment"'
	Ideals     string 'json:"ideals"'
	Bonds      string 'json:"bonds"'
	Flaws      string 'json:"flaws"'
	Background string 'json:"background"'
	Languages  []string 'json:"languages"'
}

type Stats struct {
	RawStats  map[string]int 'json:"rawStats"'
	Modifiers map[string]int 'json:"modifiers"'
}

type Item struct {
	Name        string 'json:"name"'
	Worth       int 'json:"worth"'//Worth in GP
	Description string 'json:"desc"'
}

type Wearable struct {
	*Item
	AC int 'json:"ac"'
}

type Weapon struct {
	*Item
	WeaponType string 'json:"type"'
	Damage     DiceRoll 'json:"damage"'//Possibly create a struct of some kind to represent dicerolls???
}

type Skill struct {
	Name        string 'json:"name"'
	Description string 'json:"desc"'
	Type        string 'json:"type"'//Defines what type of skill i.e. usable per day/hour/etc
	MaxUses     int 'json:"numUses"'
}

func main() {
	d := DiceRoll{6, 2}
	fmt.Println(d.Roll())
}
