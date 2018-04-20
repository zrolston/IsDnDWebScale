package main

func makeInt(blah string) Skill {
	mySkill := Skill{
		Name:        blah,
		PrimaryStat: "Intelligence",
	}
	return mySkill
}

func makeStr(blah string) Skill {
	mySkill := Skill{
		Name:        blah,
		PrimaryStat: "Strength",
	}
	return mySkill
}

func makeDex(blah string) Skill {
	mySkill := Skill{
		Name:        blah,
		PrimaryStat: "Dexterity",
	}
	return mySkill
}

func makeWis(blah string) Skill {
	mySkill := Skill{
		Name:        blah,
		PrimaryStat: "Wisdom",
	}
	return mySkill
}

func makeCha(blah string) Skill {
	mySkill := Skill{
		Name:        blah,
		PrimaryStat: "Charisma",
	}
	return mySkill
}

func makeAllSkills() []Skill {
	skillz := []Skill{
		makeStr("Athletics"),
		makeDex("Acrobatics"),
		makeDex("Sleight of Hand"),
		makeDex("Stealth"),
		makeInt("Arcana"),
		makeInt("History"),
		makeInt("Investigation"),
		makeInt("Nature"),
		makeInt("Religion"),
		makeWis("Animal Handling"),
		makeWis("Insight"),
		makeWis("Medicine"),
		makeWis("Perception"),
		makeWis("Survival"),
		makeCha("Deception"),
		makeCha("Intimidation"),
		makeCha("Performance"),
		makeCha("Persuasion"),
	}
	return skillz
}
