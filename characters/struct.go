package characters

type Character struct {
	name         string
	health       int
	mana         int
	xp           int
	strength     float64
	intelligence float64
	agility      float64
	defense      float64
	magicDefense float64
	critChance   float64
	critDamage   float64
	commonSense  float64 // the ability to find other ways to do a task/find a rare item/unlock different dialogues
	statsBuffs   statsBuffs
	localization localization
	inventory    inventory
}

type statsBuffs struct {
	buffPerCent float64
	buffFixed   float64
	buffTime    int // in seconds
}

type localization struct {
	x            float64
	y            float64
	z            float64
	hitBoxHeight float64
	hitboxWidth  float64
}

type inventory struct {
	inventory []items.Items
}
