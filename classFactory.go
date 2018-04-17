package main

func makeBarbarian() Class {
	myClass := Class{
		Name:          "Barbarian",
		Subclass:      "None",
		Description:   "A fierce warrior of primitive background who can enter a battle rage.",
		Proficiencies: []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeBard() Class {
	myClass := Class{
		Name:          "Bard",
		Subclass:      "None",
		Description:   "An inspiring magician whose power echoes the music of creation.",
		Proficiencies: []string{"Light Armor", "Simple Weapons", "Hand Crossbows", "Longswords", "Rapiers", "Shortswords"},
	}

	return myClass
}

func makeCleric() Class {
	myClass := Class{
		Name:          "Cleric",
		Subclass:      "None",
		Description:   "A preistly companion who weilds divine magic in service of a higher power.",
		Proficiencies: []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons"},
	}

	return myClass
}

func makeDruid() Class {
	myClass := Class{
		Name:          "Druid",
		Subclass:      "None",
		Description:   "A priest of the Old Faith, weilding the powers of nature - moonlight and plant growth, fire and lightning - and adopting animal forms.",
		Proficiencies: []string{"Light Nonmetal Armor", "Medium Nonmetal Armor", "Nonmetal Shields", "Clubs", "Daggers", "Darts", "Javelins", "Maces", "Quarterstaffs", "Scimitars", "Sickles", "Slings", "Spears"},
	}

	return myClass
}

func makeFighter() Class {
	myClass := Class{
		Name:          "Fighter",
		Subclass:      "None",
		Description:   "A master of martial combat, skilled with a variety of weapons of armor.",
		Proficiencies: []string{"All Armor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeMonk() Class {
	myClass := Class{
		Name:          "Monk",
		Subclass:      "None",
		Description:   "A master of martial arts, harnessing the power of the body in pursuit of physical and spiritual perfection.",
		Proficiencies: []string{"Simple Weapons", "Shortswords"},
	}

	return myClass
}

func makePaladin() Class {
	myClass := Class{
		Name:          "Paladin",
		Subclass:      "None",
		Description:   "A holy warrior bound to a sacred oath.",
		Proficiencies: []string{"All Armor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeRanger() Class {
	myClass := Class{
		Name:          "Ranger",
		Subclass:      "None",
		Description:   "A warior who uses martial prowess and nature magic to combat threats on the edges of civilization.",
		Proficiencies: []string{"Light Armor", "Medium Aromor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeRouge() Class {
	myClass := Class{
		Name:          "Rouge",
		Subclass:      "None",
		Description:   "A scoundrel who uses stealth and trickery to overcome obstacles and enemies.",
		Proficiencies: []string{"Light Armor", "Simple Weapons", "Hand Crossbows", "Longswords", "Rapiers", "Shortswords"},
	}

	return myClass
}

func makeSorcerer() Class {
	myClass := Class{
		Name:          "Sorcerer",
		Subclass:      "None",
		Description:   "A spellcaster who draws on inherent magic from a gift or bloodline",
		Proficiencies: []string{"Daggers", "Darts", "Slings", "Quarterstaffs", "Light Crossbows"},
	}

	return myClass
}

func makeWarlock() Class {
	myClass := Class{
		Name:          "Warlock",
		Subclass:      "None",
		Description:   "A wielder of magic that is derived from a bargain with an extraplanar entity",
		Proficiencies: []string{"Light Armor", "Simple Weapons"},
	}

	return myClass
}

func makeWizard() Class {
	myClass := Class{
		Name:          "Wizard",
		Subclass:      "None",
		Description:   "A scholarly magic-user capable of manipulating the sturctures of reality",
		Proficiencies: []string{"Daggers", "Darts", "Slings", "Quarterstaffs", "Light Crossbows"},
	}

	return myClass
}

func getAllClasses() []Class {
	classes := make([]Class, 12)

	classes[0] = makeBarbarian()
	classes[1] = makeBard()
	classes[2] = makeCleric()
	classes[3] = makeDruid()
	classes[4] = makeFighter()
	classes[5] = makeMonk()
	classes[6] = makePaladin()
	classes[7] = makeRanger()
	classes[8] = makeRouge()
	classes[9] = makeSorcerer()
	classes[10] = makeWarlock()
	classes[11] = makeWizard()

	return classes
}
