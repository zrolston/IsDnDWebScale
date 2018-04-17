package main

func makeDwarf() Race {
  mod := Modifiers{
    StrengthMod:     0,
    DexterityMod:    0,
    ConstitutionMod: 2,
    IntelligenceMod: 0,
    WisdomMod:       0,
    CharismaMod:     0,
  }
  myRace := Race{
    Name : "Dwarf",
    Image : "dwarf.jpg",
    Description : "Kingdoms rich in ancient grandeur, ahlls carved into the roots of mountains, deep mines and blazing forges, a commitment to clans and tradition, and a burning hatred of goblins and orcs - these common threads unite all dwarves.",
    Traits : []string{"Age: Dwarves mature at the same rate as humans, but they’re considered young until they reach the age of 50. On average, they live about 350 years.",
    "Alignment: Most dwarves are lawful, believing firmly in the benefits of a well-ordered society. They tend toward good as well, with a strong sense of Fair Play and a belief that everyone deserves to share in the benefits of a just order.",
    "Size: Dwarves stand between 4 and 5 feet tall and average about 150 pounds. Your size is Medium.",
    "Speed: Your base walking speed is 25 feet. Your speed is not reduced by wearing Heavy Armor.",
    "Darkvision: Accustomed to life underground, you have superior vision in dark and dim Conditions. You can see in dim light within 60 feet of you as if it were bright light, and in Darkness as if it were dim light. You can’t discern color in Darkness, only shades of gray.",
    "Dwarven Resilience: You have advantage on saving throws against poison, and you have Resistance against poison damage.",
    "Dwarven Combat Training: You have proficiency with the Battleaxe, Handaxe, Light Hammer, and Warhammer.",
    "Tool Proficiency: You gain proficiency with the artisan’s tools of your choice: smith’s tools, brewer’s supplies, or mason’s tools.",
    "Stonecunning: Whenever you make an Intelligence (History) check related to the origin of stonework, you are considered proficient in the History skill and add double your proficiency bonus to the check, instead of your normal proficiency bonus.",
    "Languages: You can speak, read, and write Common and Dwarvish. Dwarvish is full of hard consonants and guttural sounds, and those characteristics spill over into whatever other Language a dwarf might speak.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeElf() Race{
  mod := Modifiers{
    StrengthMod:     0,
    DexterityMod:    2,
    ConstitutionMod: 0,
    IntelligenceMod: 0,
    WisdomMod:       0,
    CharismaMod:     0,
  }
  myRace := Race{
    Name : "Elf",
    Image : "elf.jpg",
    Description : "A magical people of otherworldly grace, living in the world but not entirely part of it.",
    Traits : []string{"Age: Although elves reach physical maturity at about the same age as humans, the elven understanding of adulthood goes beyond physical growth to encompass worldly experience. An elf typically claims adulthood and an adult name around the age of 100 and can live to be 750 years old.",
    "Alignment: Elves love freedom, variety, and self- expression, so they lean strongly toward the gentler aspects of chaos. They value and protect others’ freedom as well as their own, and they are more often good than not.",
    "Size: Elves range from under 5 to over 6 feet tall and have slender builds. Your size is Medium.",
    "Speed: Your base walking speed is 30 feet.",
    "Darkvision: Accustomed to twilit forests and the night sky, you have superior vision in dark and dim Conditions. You can see in dim light within 60 feet of you as if it were bright light, and in Darkness as if it were dim light. You can’t discern color in Darkness, only shades of gray.",
    "Keen Senses: You have proficiency in the Perception skill.",
    "Fey Ancestry: You have advantage on saving throws against being Charmed, and magic can’t put you to sleep.",
    "Trance: Elves don’t need to sleep. Instead, they meditate deeply, remaining semiconscious, for 4 hours a day. (The Common word for such meditation is “trance.”) While meditating, you can dream after a fashion; such dreams are actually mental exercises that have become reflexive through years of practice. After Resting in this way, you gain the same benefit that a human does from 8 hours of sleep.",
    "Languages: You can speak, read, and write Common and Elvish. Elvish is fluid, with subtle intonations and intricate grammar. Elven literature is rich and varied, and their songs and poems are famous among other races. Many bards learn their Language so they can add Elvish ballads to their repertoires.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeHalfling() Race{
  mod := Modifiers{
    StrengthMod:     0,
    DexterityMod:    2,
    ConstitutionMod: 0,
    IntelligenceMod: 0,
    WisdomMod:       0,
    CharismaMod:     0,
  }
  myRace := Race{
    Name : "Halfling",
    Image : "halfling.jpg",
    Description : "A diminutive race who aim to live a comfortable and peaceful life.",
    Traits : []string{"Age: A halfling reaches adulthood at the age of 20 and generally lives into the middle of his or her second century.",
    "Alignment: Most halflings are lawful good. As a rule, they are good-hearted and kind, hate to see others in pain, and have no tolerance for oppression. They are also very orderly and traditional, leaning heavily on the support of their community and the comfort of their old ways.",
    "Size: Halflings average about 3 feet tall and weigh about 40 pounds. Your size is Small.",
    "Speed: Your base walking speed is 25 feet.",
    "Lucky: When you roll a 1 on the d20 for an Attack roll, ability check, or saving throw, you can reroll the die and must use the new roll.",
    "Brave: You have advantage on saving throws against being Frightened.",
    "Halfling Nimbleness: You can move through the space of any creature that is of a size larger than yours.",
    "Languages: You can speak, read, and write Common and Halfling. The Halfling Language isn’t secret, but halflings are loath to share it with others. They write very little, so they don’t have a rich body of literature. Their oral tradition, however, is very strong. Almost all halflings speak Common to converse with the people in whose lands they dwell or through which they are traveling.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeHuman() Race{
  mod := Modifiers{
    StrengthMod:     1,
    DexterityMod:    1,
    ConstitutionMod: 1,
    IntelligenceMod: 1,
    WisdomMod:       1,
    CharismaMod:     1,
  }
  myRace := Race{
    Name : "Human",
    Image : "human.jpg",
    Description : "Humans are the youngest of the known races.  Whatever drives them , humans are innovators, achievers, and pioneers of the world.",
    Traits : []string{"Age: Humans reach adulthood in their late teens and live less than a century.",
    "Alignment: Humans tend toward no particular alignment. The best and the worst are found among them.",
    "Size: Humans vary widely in height and build, from barely 5 feet to well over 6 feet tall. Regardless of your position in that range, your size is Medium.",
    "Speed: Your base walking speed is 30 feet.",
    "Languages: You can speak, read, and write Common and one extra Language of your choice. Humans typically learn the languages of other peoples they deal with, including obscure dialects. They are fond of sprinkling their Speech with words borrowed from other tongues: Orc curses, Elvish musical expressions, Dwarvish military phrases, and so on.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeDragonborn() Race{
  mod := Modifiers{
    StrengthMod:     2,
    DexterityMod:    0,
    ConstitutionMod: 0,
    IntelligenceMod: 0,
    WisdomMod:       0,
    CharismaMod:     1,
  }
  myRace := Race{
    Name : "Dragonborn",
    Image : "dragonborn.jpg",
    Description : "Born of dragons, as their name proclaims, Dragonborn walk proudly through a world that greets them with a fearful incomprehension.",
    Traits : []string{"Age: Young dragonborn grow quickly. They walk hours after hatching, attain the size and development of a 10-year-old human child by the age of 3, and reach adulthood by 15. They live to be around 80.",
    "Alignment: Dragonborn tend to extremes, making a conscious choice for one side or the other in the cosmic war between good and evil. Most dragonborn are good, but those who side with evil can be terrible villains.",
    "Size: Dragonborn are taller and heavier than humans, standing well over 6 feet tall and averaging almost 250 pounds. Your size is Medium.",
    "Speed: Your base walking speed is 30 feet.",
    "Draconic Ancestry: You have draconic ancestry. Choose one type of dragon from the Draconic Ancestry table. Your breath weapon and damage Resistance are determined by the dragon type, as shown in the table.",
    "Breath Weapon: You can use your action to exhale destructive energy. Your draconic ancestry determines the size, shape, and damage type of the exhalation.",
    "Damage Resistance: You have Resistance to the damage type associated with your draconic ancestry.",
    "Languages: You can speak, read, and write Common and Draconic. Draconic is thought to be one of the oldest languages and is often used in the study of magic. The Language sounds harsh to most other creatures and includes numerous hard consonants and sibilants.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeGnome() Race{
  mod := Modifiers{
    StrengthMod:     0,
    DexterityMod:    0,
    ConstitutionMod: 0,
    IntelligenceMod: 2,
    WisdomMod:       0,
    CharismaMod:     0,
  }
  myRace := Race{
    Name : "Gnome",
    Image : "gnome.jpg",
    Description : "A constant hum of busy activity pervades the warrens and neighborhoods where gnomes form close knit communities.  Gnomes take a delight in life, enjoying every moment of invention, exploration, investigation, creation, and play.",
    Traits : []string{"Age: Gnomes mature at the same rate humans do, and most are expected to settle down into an adult life by around age 40. They can live 350 to almost 500 years.", "Alignment: Gnomes are most often good. Those who tend toward law are sages, engineers, researchers, scholars, investigators, or inventors. Those who tend toward chaos are minstrels, tricksters, wanderers, or fanciful jewelers. Gnomes are good-hearted, and even the tricksters among them are more playful than vicious.",
    "Size: Gnomes are between 3 and 4 feet tall and average about 40 pounds. Your size is Small.", "Speed: Your base walking speed is 25 feet.", "Darkvision: Accustomed to life underground, you have superior vision in dark and dim Conditions. You can see in dim light within 60 feet of you as if it were bright light, and in Darkness as if it were dim light. You can’t discern color in Darkness, only shades of gray.", "Gnome Cunning: You have advantage on all Intelligence, Wisdom, and Charisma saving throws against magic.", "Languages: You can speak, read, and write Common and Gnomish. The Gnomish Language, which uses the Dwarvish script, is renowned for its technical treatises and its catalogs of knowledge about the natural world.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeHalfElf() Race{
  mod := Modifiers{
    StrengthMod:     0,
    DexterityMod:    1,
    ConstitutionMod: 0,
    IntelligenceMod: 1,
    WisdomMod:       0,
    CharismaMod:     2,
  }
  myRace := Race{
    Name : "Half-Elf",
    Image : "halfelf.jpg",
    Description : "Walking in two worlds but belonging to neither, half-elves combine what some say are the best qualities of their elf and human parents",
    Traits : []string{"Age: Half-elves mature at the same rate humans do and reach adulthood around the age of 20. They live much longer than humans, however, often exceeding 180 years.", "Alignment: Half-elves share the chaotic bent of their elven heritage. They value both personal freedom and creative expression, demonstrating neither love of leaders nor desire for followers. They chafe at rules, resent others’ demands, and sometimes prove unreliable, or at least unpredictable.", "Size: Half-elves are about the same size as humans, ranging from 5 to 6 feet tall. Your size is Medium.", "Speed: Your base walking speed is 30 feet.", "Darkvision: Thanks to your elf blood, you have superior vision in dark and dim Conditions. You can see in dim light within 60 feet of you as if it were bright light, and in Darkness as if it were dim light. You can not discern color in Darkness, only shades of gray.",
    "Fey Ancestry: You have advantage on saving throws against being Charmed, and magic can not put you to sleep.", "Languages: You can speak, read, and write Common, Elvish, and one extra Language of your choice.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeHalfOrc() Race{
  mod := Modifiers{
    StrengthMod:     2,
    DexterityMod:    0,
    ConstitutionMod: 1,
    IntelligenceMod: 0,
    WisdomMod:       0,
    CharismaMod:     0,
  }
  myRace := Race{
    Name : "Half-Orc",
    Image : "halforc.jpg",
    Description : "A resuilt of a marriage-sealed alliance between humans and orcs. Some use their human blod to gain an advantage over fellow orcs and rise to be cheiftans, other venture into the world to prove themselves among humans and more civilized races.",
    Traits : []string{"Age: Half-Orcs mature a little faster than humans, reaching adulthood around age 14. They age noticeably faster and rarely live longer than 75 years.", "Alignment: Half-Orcs inherit a tendency toward chaos from their orc parents and are not strongly inclined toward good. Half-Orcs raised among orcs and willing to live out their lives among them are usually evil.", "Size: Half-Orcs are somewhat larger and bulkier than humans, and they range from 5 to well over 6 feet tall. Your size is Medium.",
    "Speed: Your base walking speed is 30 feet.", "Darkvision: Thanks to your orc blood, you have superior vision in dark and dim Conditions. You can see in dim light within 60 feet of you as if it were bright light, and in Darkness as if it were dim light. You can not discern color in Darkness, only shades of gray.", "Menacing: You gain proficiency in the Intimidation skill.", "Relentless Endurance: When you are reduced to 0 hit points but not killed outright, you can drop to 1 hit point instead. You can not use this feature again until you finish a Long Rest.", "Savage Attacks: When you score a critical hit with a melee weapon Attack, you can roll one of the weapon’s damage dice one additional time and add it to the extra damage of the critical hit.", "Languages: You can speak, read, and write Common and Orc. Orc is a harsh, grating Language with hard consonants. It has no script of its own but is written in the Dwarvish script.",
    },
    Bonuses : mod,
  }

  return myRace
}

func makeTiefling() Race{
  mod := Modifiers{
		StrengthMod:     0,
		DexterityMod:    0,
		ConstitutionMod: 0,
		IntelligenceMod: 1,
		WisdomMod:       0,
		CharismaMod:     2,
	}
  myRace := Race{
    Name : "Tiefling",
    Image : "tiefling.jpg",
    Description : "To be greeted with stares and whispers, to suffer violence and insult on the street, to see mistrust and fear in every eye: this is the lot of the tiefling. And to twist the knife, tieflings know that this is because a pact struct generations ago infused the essence of Asmodeus - overlord of the Nine Hells - into their bloodline.",
    Traits : []string{"Age: Tieflings mature at the same rate as humans but live a few years longer.", "Alignment: Tieflings might not have an innate tendency toward evil, but many of them end up there. Evil or not, an independent nature inclines many tieflings toward a chaotic alignment.", "Size: Tieflings are about the same size and build as humans. Your size is Medium.", "Speed: Your base walking speed is 30 feet.", "Darkvision: Thanks to your infernal heritage, you have superior vision in dark and dim Conditions. You can see in dim light within 60 feet of you as if it were bright light, and in Darkness as if it were dim light. You can not discern color in Darkness, only shades of gray.",
    "Hellish Resistance: You have Resistance to fire damage.", "Infernal Legacy. You know the Thaumaturgy cantrip. When you reach 3rd level, you can cast the Hellish Rebuke spell as a 2nd-level spell once with this trait and regain the ability to do so when you finish a Long Rest. When you reach 5th level, you can cast the Darkness spell once with this trait and regain the ability to do so when you finish a Long Rest. Charisma is your spellcasting ability for these Spells.", "Languages: You can speak, read, and write Common and Infernal.",
    },
    Bonuses : mod,
  }

  return myRace
}

func getAllRaces() []Race{
  races := make([]Race, 9)
  races[0] = makeDwarf()
  races[1] = makeElf()
  races[2] = makeHalfling()
  races[3] = makeHuman()
  races[4] = makeDragonborn()
  races[5] = makeGnome()
  races[6] = makeHalfElf()
  races[7] = makeHalfOrc()
  races[8] = makeTiefling()

  return races
}
