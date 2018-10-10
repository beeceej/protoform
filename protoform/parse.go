package protoform

import (
	"strings"
)

var classExtracter = classExtract{}
var constructorExtracter = constructorExtract{}
var listTypeExtracter = listTypeExtract{}
var mapTypeExtracter = mapTypeExtract{}

func init() {
	classExtracter.compile()
	constructorExtracter.compile()
	listTypeExtracter.compile()
	mapTypeExtracter.compile()
}

type findFunc func(m [][]string) string

// Parser is a struct which holds information necessary to parse and output a Proto Object
type Parser struct {
	Syntax  string
	Package string
	OutFile string
}

// Parse takes a raw java string and returns an in memory Proto definition
func (p Parser) Parse(raw string) *Proto {
	className := classExtracter.defaultFind(raw)
	rawFieldProperties := constructorExtracter.defaultFind(raw)

	fields := strings.Split(strings.TrimSpace(strings.Replace(rawFieldProperties, "(", "", -1)), ",")
	var propertyParts []string
	properties := []MessageProperty{}
	imports := []string{}
	var skipNext bool
	for i, f := range fields {
		if !skipNext {
			if strings.Contains(strings.ToLower(f), "map<") {
				// fmt.Println("Fields ", len(fields), "i: ", i)
				propertyParts = strings.Split(strings.TrimSpace(fields[i]+","+fields[i+1]), " ") //rejoin the split map
				skipNext = true
			} else {
				propertyParts = strings.Split(strings.TrimSpace(f), " ")
			}
			rawType := propertyParts[0]

			rawName := propertyParts[1]
			var protoType string

			if parameterType := parameterizedTypeToProto(rawType); parameterType != "" {
				protoType = parameterType
			} else {
				if strings.Contains(rawType, "Integer") || strings.Contains(rawType, "int") {
					protoType = "int32"
				} else if strings.Contains(rawType, "Map") {
					protoType = "map"
				} else if strings.Contains(rawType, "Long") || strings.Contains(rawType, "long") {
					protoType = "int64"
				} else if strings.Contains(rawType, "String") {
					protoType = "string"
				} else if strings.Contains(rawType, "Bool") || strings.Contains(rawType, "boolean") {
					protoType = "bool"
				} else if strings.Contains(rawType, "DateTime") {
					dateTimeImport := `import "google/protobuf/timestamp.proto";`
					protoType = "google.protobuf.Timestamp"
					var alreadyImported bool
					for _, impt := range imports {
						if impt == dateTimeImport {
							alreadyImported = true
						}
					}
					if !alreadyImported {
						imports = append(imports, dateTimeImport)
					}
				} else {
					protoType = rawType
				}
			}
			protoName := strings.Replace(rawName, ")", "", -1)
			properties = append(properties, MessageProperty{
				Type:   protoType,
				Name:   protoName,
				Number: i + 1,
			})
		} else {
			skipNext = false
		}
	}
	return &Proto{
		FileName:   className,
		Type:       "message",
		Properties: properties,
		Imports:    imports,
		Package:    p.Package,
		Syntax:     p.Syntax,
	}
}

func parameterizedTypeToProto(rawType string) (proto string) {
	rawType = strings.TrimSpace(rawType)
	if proto = listTypeExtracter.defaultFind(rawType); proto != "" {
		return proto
	}
	if proto = mapTypeExtracter.defaultFind(rawType); proto != "" {
		return proto
	}

	return proto
}
