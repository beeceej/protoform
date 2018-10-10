package protoform

import "fmt"

const (
	message = iota
	enum
)

// protobufType is one of message, or enum
type protobufType uint8

// Proto is a struct which contains the components of a protobuf file
type Proto struct {
	FileName   string
	Type       string
	Syntax     string
	Package    string
	Properties []MessageProperty
	Imports    []string
}

// MessageProperty represents a specific field
type MessageProperty struct {
	Type   string
	Name   string
	Number int
}

func (p protobufType) Print() {
	if p == message {
		fmt.Println("message")
	} else if p == enum {
		fmt.Println("enum")
	}
}

func (p protobufType) Sprint() (s string) {
	if p == message {
		s = "message"
	} else if p == enum {
		s = "enum"
	}
	return s
}

// Template returns a templatized version of a protobuf file.
func (p Proto) Template() string {
	return `syntax = "{{- .Syntax }}";
package {{ .Package }};
{{- range .Imports }}
	import {{ . -}};
{{- end }}

message {{ .Type }} {
	{{- range .Properties }}
	{{ .Type }} {{ .Name }} = {{ .Number -}};
	{{- end }}
}`
}
