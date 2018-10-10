package protoform

import "regexp"

type constructorExtract struct {
	expr    *regexp.Regexp
	matches [][]string
}

func (c *constructorExtract) compile() {
	c.expr = regexp.MustCompile(`(public|private|protected) .+(\((.+ .+,|\n|.+ .+\)|.+ .+,)+)`)
}

func (c constructorExtract) matchIndex() int {
	return 0
}

func (c constructorExtract) groupIndex() int {
	return 2
}

func (c *constructorExtract) defaultFind(s string) (match string) {
	if c.expr == nil {
		c.compile()
	}
	c.matches = c.expr.FindAllStringSubmatch(s, -1)
	if c.safe() {
		match = c.matches[c.matchIndex()][c.groupIndex()]
	}

	return match
}

func (c *constructorExtract) fnFind(fn findFunc) string {
	return fn(c.matches)
}

func (c constructorExtract) safe() (safe bool) {
	if len(c.matches) > c.matchIndex() && len(c.matches[c.matchIndex()]) > c.groupIndex() {
		safe = true
	}
	return safe
}
