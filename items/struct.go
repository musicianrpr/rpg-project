package items

import (
	"reflect"

	"github.com/musicianrpr/rpg-project/tree/0.0.0/characters"
)

// Item means all type of items
type Item interface {
	Pick()
}

// CommonItem characteristics
type CommonItem struct {
	Name        string
	Quantity    int
	Weigth      float64
	Price       int // copper coins
	Description string
}

// HealItem is an item that heals the player
type HealItem struct {
	CommonItem
	Heal int
}

// BuffItem is an item that buffs the player
type BuffItem struct {
	CommonItem
	Buff characters.StatsBuffs
}

type Weapon struct {
	CommonItem
	Damage      float64
	MagicDamage float64
	AttackSpeed float64
}

// Armor is an equipable armor
type Armor struct {
	CommonItem
	Defense float64
}

// Boots is an equipable pair of boots
type Boots struct {
	CommonItem
	Defense   float64
	MoveSpeed float64
}

// Leggings is an equipable legging
type Leggings struct {
	CommonItem
	Defense float64
}

// Helmet is an equipable helmet
type Helmet struct {
	CommonItem
	Defense float64
}

// Consumable means all type of consumables
type Consumable interface {
	UseItem()
}

type Equipable interface {
	EquipItem()
}

func (c *characters.Character) Pick(item Item) {
	c.Inventory.Inventory = append(c.Inventory.Inventory, item)
}

func (c *characters.Character) EquipItem(item Equipable) {
	switch reflect.TypeOf(item) {
	case Weapon:
		c.Weapon = item
	case Armor:
		c.Armor = item
	}
}
