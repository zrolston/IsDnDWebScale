package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

func testBody(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)

	decoder := json.NewDecoder(r.Body)
	var char Character
	err := decoder.Decode(&char)
	if err != nil {
		w.WriteHeader(420)
		writeMessage(w, "Fug")
	}

	w.WriteHeader(200)
	writeMessage(w, "Successful decode of "+char.Name)
	/*
		decoder := json.NewDecoder(bod)

		var char Character
		err = decoder.Decode(&char)
		if err != nil {
			w.WriteHeader(400)
			writeMessage(w, err.Error())
			return
		}
		/*
			response, _ := json.Marshal(char)
			w.WriteHeader(200)
			w.Write(response)
	*/
}

func putGerald(w http.ResponseWriter, r *http.Request) {
	char := makeGerald()
	setupResponse(&w, r)

	ctx := appengine.NewContext(r)
	gKey := datastore.NewKey(ctx, "character", "Gerald", 0, nil)

	_, err := datastore.Put(ctx, gKey, &char)
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, (err))
		return
	}

	w.WriteHeader(200)
	fmt.Fprintln(w, "Correctly Stored Gerald")
}

func getGeraldDB(w http.ResponseWriter, r *http.Request) {

	setupResponse(&w, r)

	var char Character

	ctx := appengine.NewContext(r)
	gKey := datastore.NewKey(ctx, "character", "Gerald", 0, nil)

	if err := datastore.Get(ctx, gKey, &char); err != nil {
		w.WriteHeader(418)
		fmt.Fprintln(w, "Not found")
		return
	}

	response, _ := json.Marshal(char)
	w.WriteHeader(200)
	w.Write(response)
}

func getGeraldWep(w http.ResponseWriter, r *http.Request) {
	char := makeGerald()
	wep := char.EquipedWeapon[0]
	response, _ := json.Marshal(wep)
	setupResponse(&w, r)
	w.WriteHeader(200)
	w.Write(response)
}

func addSkill(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var skill Skill
	err := decoder.Decode(&skill)
	setupResponse(&w, r)
	if err == nil {
		w.WriteHeader(500)
	}
	log.Print(skill.Name)
	w.WriteHeader(200)
}

func welcomeTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the fucking show")
}

func makeGerald() Character {

	bonusMap := Modifiers{
		IntelligenceMod: 1,
	}

	human := Race{
		Name:        "Human",
		Image:       "benis.jpg",
		Description: "Pretty boring for now.",
		Traits:      []string{"Meme master: Mastery of memes"},
		Bonuses:     bonusMap,
	}

	warrior := Class{
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
	}

	stat := Stats{
		Strength:     15,
		Dexterity:    16,
		Constitution: 17,
		Intelligence: 18,
		Wisdom:       19,
		Charisma:     20,
	}

	mods := NewMod(stat)

	gerald := Character{
		Name:        "Gerald",
		Race:        human,
		Class:       warrior,
		Portrait:    "Fug.jpg",
		Description: desc,
		Exp:         15,
		Stats:       stat,
		Modifiers:   mods,
		Inventory: []Item{
			Item{
				Name:        "Sandwich",
				Worth:       3,
				Description: "Sub-Par",
			},
		},
		EquipedClothes: []Wearable{
			Wearable{
				Item: Item{
					Name:        "Chestpiece",
					Worth:       20,
					Description: "Leather",
				},
				AC: 10,
			},
		},
		EquipedWeapon: []Weapon{Weapon{
			Item: Item{
				Name:        "memeSword",
				Worth:       69,
				Description: "maymay powered"},
			WeaponType: "MayMays",
			Damage:     DiceRoll{NumRolls: 2, Sides: 10},
		}},
		Spells: []Spell{
			Spell{
				Name:        "Dab",
				Description: "Dabs on the haters",
				Type:        "Per Encounter",
				MaxUses:     5,
			},
		},
		Skills: makeAllSkills(),
	}

	return gerald
}
