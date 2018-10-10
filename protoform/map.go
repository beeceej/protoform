package protoform

import "regexp"

type mapTypeExtract struct {
	expr *regexp.Regexp

	matches [][]string
}

func (m *mapTypeExtract) compile() {
	m.expr = regexp.MustCompile(`.*Map\<(.+)\>.*`)
}

func (m mapTypeExtract) matchIndex() int {
	return 0
}

func (m mapTypeExtract) groupIndex() int {
	return 1
}

func (m *mapTypeExtract) defaultFind(s string) (match string) {
	if m.expr == nil {
		m.compile()
	}
	m.matches = m.expr.FindAllStringSubmatch(s, -1)
	if m.safe() {
		match = m.matches[m.matchIndex()][m.groupIndex()]
	}

	return match
}

func (m *mapTypeExtract) fnFind(fn findFunc) string {
	return fn(m.matches)
}

func (m mapTypeExtract) safe() (safe bool) {
	if len(m.matches) > m.matchIndex() && len(m.matches[m.matchIndex()]) > m.groupIndex() {
		safe = true
	}
	return safe
}
