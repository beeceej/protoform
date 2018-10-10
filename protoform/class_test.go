package protoform

import (
	"io/ioutil"
	"path/filepath"
	"testing"
)

func Test_safe(t *testing.T) {
	src := Load(filepath.Join("testdata", "A.java"))
	expected := true
	_ = classExtracter.defaultFind(src)
	actual := classExtracter.safe()
	if expected != actual {
		t.Fail()
	}
}

func Test_defaultFind(t *testing.T) {
	src := Load(filepath.Join("testdata", "A.java"))
	expected := "A"
	actual := classExtracter.defaultFind(src)
	if expected != actual {
		t.Fail()
	}
}

func Test_compile(t *testing.T) {
	classExtracter.compile() // should not panic
}

func Test_fnFind(t *testing.T) {
	src := Load(filepath.Join("testdata", "A.java"))
	classExtracter.defaultFind(src)

	expected := "A"

	actual := classExtracter.fnFind(func(m [][]string) string {
		return m[0][1]
	})

	if expected != actual {
		t.Fail()
	}
}

func Load(name string) (s string) {
	var (
		err error
		b   []byte
	)

	if b, err = ioutil.ReadFile(name); err != nil {
		panic(err.Error())
	}

	return string(b)
}
