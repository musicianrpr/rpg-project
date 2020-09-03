package characters

func (c *character) Walk(direction string) {
	switch direction {
	case "up":
		c.localization.y++
	case "down":
		c.localization.y--
	case "right":
		c.localization.x++
	case "left":
		c.localization.y--
	}
}
