package protoform

import "regexp"

type listTypeExtract struct {
	expr *regexp.Regexp

	matches [][]string
}

func (l *listTypeExtract) compile() {
	l.expr = regexp.MustCompile(`.*List\<(.+)\>.*`)
}

func (l listTypeExtract) matchIndex() int {
	return 0
}

func (l listTypeExtract) groupIndex() int {
	return 1
}

func (l *listTypeExtract) defaultFind(s string) (match string) {
	if l.expr == nil {
		l.compile()
	}
	l.matches = l.expr.FindAllStringSubmatch(s, -1)
	if l.safe() {
		match = l.matches[l.matchIndex()][l.groupIndex()]
	}

	return match
}

func (l *listTypeExtract) fnFind(fn findFunc) string {
	return fn(l.matches)
}

func (l listTypeExtract) safe() (safe bool) {
	if len(l.matches) > l.matchIndex() && len(l.matches[l.matchIndex()]) > l.groupIndex() {
		safe = true
	}
	return safe
}
