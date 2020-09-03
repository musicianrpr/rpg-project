package items

import "\workspaces\myrpg\characters"

type Item interface {
	Pick()
}

type commonItem struct {
	name        string
	quantity    int
	weigth      float64
	price       int // copper coins
	description string
}

type healItem struct {
	commonItem
	heal int
}

type armor struct {
	commonItem
	defense float64
}

type consumable interface {
	UseItem()
}

type equipable interface {
}

func (c *) Pick() {

}
