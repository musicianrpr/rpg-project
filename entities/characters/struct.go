package characters

// Character stores the character stats
type Character struct {
	Name          string
	Health        int
	Mana          int
	Xp            int
	Strength      float64
	Intelligence  float64
	Agility       float64
	Defense       float64
	MagicDefense  float64
	CritChance    float64
	CritDamage    float64
	CommonSense   float64 // the ability to find other ways to do a task/find a rare item/unlock different dialogues
	Weapon        items.Equipable
	Armor         items.Armor
	Boots         items.Boots
	Leggings      items.Leggings
	Helmet        items.Helmet
	StatsBuffs    StatsBuffs
	Localization  Localization
	Inventory     Inventory
	CurrentStatus CurrentStatus
}

//StatsBuffs is the character buffs
type StatsBuffs struct {
	BuffPerCent float64
	BuffFixed   float64
	BuffTime    int // in seconds
}

//Localization is the character localization
type Localization struct {
	X            float64
	Y            float64
	Z            float64
	HitBoxHeight float64
	HitboxWidth  float64
}

//Inventory is the character inventory
type Inventory struct {
	Inventory []items.Items
}

type CurrentStatus struct {
	Health int
	Mana   int
}
