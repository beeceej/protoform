package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
	"unicode"

	"github.com/beeceej/protoform/protoform"
)

func main() {
	inFile := flag.String("in-file", "", "Input File")
	pkg := flag.String("package", "", "Package of proto")
	out := flag.String("out-file", "", "Output File")

	flag.Parse()

	b, err := ioutil.ReadFile(*inFile)
	if err != nil {
		panic(err.Error())
	}
	pb := protoform.Parser{
		Package: *pkg,
		OutFile: *out,
		Syntax:  "proto3",
	}.Parse(string(b))

	snakeCaseFileName := properCaseToSnakeCase(pb.FileName)
	var (
		f *os.File
	)

	f, err = os.Create(filepath.Join(*out, *pkg, fmt.Sprintf("%s.proto", snakeCaseFileName)))
	if err != nil {
		if err = os.MkdirAll(filepath.Join(*out, *pkg), 0777); err != nil {
			panic(err.Error())
		}

		if f, err = os.Create(filepath.Join(*out, *pkg, fmt.Sprintf("%s.proto", snakeCaseFileName))); err != nil {
			panic(err.Error())
		}
	}
	defer f.Close()
	t := template.Must(template.New("tmpl").Parse(pb.Template()))
	t.Execute(f, pb)
}

func properCaseToSnakeCase(name string) (newName string) {
	for i, r := range name {
		if unicode.IsUpper(r) && i != 0 {
			newName = fmt.Sprintf("%s_%s", newName, string(unicode.ToLower(r)))
		} else {
			newName = fmt.Sprintf("%s%s", newName, string(unicode.ToLower(r)))
		}
	}
	return newName
}
