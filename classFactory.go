package main

func makeBarbarian() Class {
	myClass := Class{
		Name:          "Barbarian",
		Image:         "https://i.imgur.com/srZIxJQ.jpg",
		Subclass:      "None",
		Description:   "A fierce warrior of primitive background who can enter a battle rage.",
		Proficiencies: []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeBard() Class {
	myClass := Class{
		Name:          "Bard",
		Image:         "https://i.imgur.com/SHJVAms.jpg",
		Subclass:      "None",
		Description:   "An inspiring magician whose power echoes the music of creation.",
		Proficiencies: []string{"Light Armor", "Simple Weapons", "Hand Crossbows", "Longswords", "Rapiers", "Shortswords"},
	}

	return myClass
}

func makeCleric() Class {
	myClass := Class{
		Name:          "Cleric",
		Image:         "https://i.imgur.com/1d4nT1e.jpg",
		Subclass:      "None",
		Description:   "A preistly companion who weilds divine magic in service of a higher power.",
		Proficiencies: []string{"Light Armor", "Medium Armor", "Shields", "Simple Weapons"},
	}

	return myClass
}

func makeDruid() Class {
	myClass := Class{
		Name:          "Druid",
		Image:         "https://i.imgur.com/Jd5XsBX.jpg",
		Subclass:      "None",
		Description:   "A priest of the Old Faith, weilding the powers of nature - moonlight and plant growth, fire and lightning - and adopting animal forms.",
		Proficiencies: []string{"Light Nonmetal Armor", "Medium Nonmetal Armor", "Nonmetal Shields", "Clubs", "Daggers", "Darts", "Javelins", "Maces", "Quarterstaffs", "Scimitars", "Sickles", "Slings", "Spears"},
	}

	return myClass
}

func makeFighter() Class {
	myClass := Class{
		Name:          "Fighter",
		Image:         "https://i.imgur.com/dQjxoND.jpg",
		Subclass:      "None",
		Description:   "A master of martial combat, skilled with a variety of weapons of armor.",
		Proficiencies: []string{"All Armor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeMonk() Class {
	myClass := Class{
		Name:          "Monk",
		Image:         "https://i.imgur.com/yhofka7.jpg",
		Subclass:      "None",
		Description:   "A master of martial arts, harnessing the power of the body in pursuit of physical and spiritual perfection.",
		Proficiencies: []string{"Simple Weapons", "Shortswords"},
	}

	return myClass
}

func makePaladin() Class {
	myClass := Class{
		Name:          "Paladin",
		Image:         "https://i.imgur.com/Mit70Hq.jpg",
		Subclass:      "None",
		Description:   "A holy warrior bound to a sacred oath.",
		Proficiencies: []string{"All Armor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeRanger() Class {
	myClass := Class{
		Name:          "Ranger",
		Image:         "https://i.imgur.com/uvl7tEd.jpg",
		Subclass:      "None",
		Description:   "A warior who uses martial prowess and nature magic to combat threats on the edges of civilization.",
		Proficiencies: []string{"Light Armor", "Medium Aromor", "Shields", "Simple Weapons", "Martial Weapons"},
	}

	return myClass
}

func makeRogue() Class {
	myClass := Class{
		Name:          "Rogue",
		Image:         "https://i.imgur.com/F7LcsMG.jpg",
		Subclass:      "None",
		Description:   "A scoundrel who uses stealth and trickery to overcome obstacles and enemies.",
		Proficiencies: []string{"Light Armor", "Simple Weapons", "Hand Crossbows", "Longswords", "Rapiers", "Shortswords"},
	}

	return myClass
}

func makeSorcerer() Class {
	myClass := Class{
		Name:          "Sorcerer",
		Image:         "https://i.imgur.com/Qi4A1bl.jpg",
		Subclass:      "None",
		Description:   "A spellcaster who draws on inherent magic from a gift or bloodline",
		Proficiencies: []string{"Daggers", "Darts", "Slings", "Quarterstaffs", "Light Crossbows"},
	}

	return myClass
}

func makeWarlock() Class {
	myClass := Class{
		Name:          "Warlock",
		Image:         "https://i.imgur.com/9CWHkxo.jpg",
		Subclass:      "None",
		Description:   "A wielder of magic that is derived from a bargain with an extraplanar entity",
		Proficiencies: []string{"Light Armor", "Simple Weapons"},
	}

	return myClass
}

func makeWizard() Class {
	myClass := Class{
		Name:          "Wizard",
		Image:         "https://i.imgur.com/fMoZ1Ax.jpg",
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
	classes[8] = makeRogue()
	classes[9] = makeSorcerer()
	classes[10] = makeWarlock()
	classes[11] = makeWizard()

	return classes
}
