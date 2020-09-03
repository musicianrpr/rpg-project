package characters

func (c *Character) Walk(direction string) {
	switch direction {
	case "up":
		c.Localization.y++
	case "down":
		c.Localization.y--
	case "right":
		c.Localization.x++
	case "left":
		c.Localization.y--
	}
}
