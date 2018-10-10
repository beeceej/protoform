package protoform

import "regexp"

type classExtract struct {
	expr    *regexp.Regexp
	matches [][]string
}

func (c *classExtract) compile() {
	c.expr = regexp.MustCompile(`.* class (.+?) (extends|).*(implements|).*{`)
}

func (c classExtract) matchIndex() int {
	return 0
}

func (c classExtract) groupIndex() int {
	return 1
}

func (c *classExtract) defaultFind(s string) (match string) {
	if c.expr == nil {
		c.compile()
	}
	c.matches = c.expr.FindAllStringSubmatch(s, -1)
	if c.safe() {
		match = c.matches[c.matchIndex()][c.groupIndex()]
	}

	return match
}

func (c *classExtract) fnFind(fn findFunc) string {
	return fn(c.matches)
}

func (c *classExtract) safe() (safe bool) {
	if len(c.matches) > c.matchIndex() && len(c.matches[c.matchIndex()]) > c.groupIndex() {
		safe = true
	}
	return safe
}
