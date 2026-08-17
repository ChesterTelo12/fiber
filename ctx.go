package fiber

// ... existing code ...

func (c *Ctx) release() {
	c.route = nil
	c.app = nil
	c.fasthttp = nil
	c.matched = false
	c.path = ""
	c.method = ""
	c.methodINT = 0
	c.indexRoute = -1
	c.indexHandler = 0
	c.values = nil
}

// ... existing code ...