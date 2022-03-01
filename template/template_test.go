package template

import (
	"bytes"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestGetData(t *testing.T) {
	tests := []struct {
		inputPath string
		want      FileData
	}{
		{"testdata/int.input", FileData{
			PackageName: "colors",
			TypeName:    "Color",
			ConstsType:  "int",
			Consts:      []string{"RED", "GREEN", "BLUE"},
		}},
		{"testdata/string.input", FileData{
			PackageName: "colors",
			TypeName:    "Color",
			ConstsType:  "string",
			Consts:      []string{"RED", "GREEN", "BLUE"},
		}},
	}

	for _, test := range tests {
		_, filename := filepath.Split(test.inputPath)
		testName := filename[:len(filename)-len(filepath.Ext(test.inputPath))]

		t.Run(testName, func(t *testing.T) {
			astFile, err := parser.ParseFile(
				token.NewFileSet(),
				test.inputPath,
				nil,
				0,
			)
			if err != nil {
				t.Fatal("error parse input file: ", err)
			}

			fileData := GetData(astFile)

			if !reflect.DeepEqual(fileData, test.want) {
				t.Errorf("\n==== got:\n%v\n==== want:\n%v\n", fileData, test.want)
			}
		})
	}
}

func TestGenFileContent(t *testing.T) {
	tests := []struct {
		inputPath string
		input     FileData
	}{
		{"testdata/int.input", FileData{
			PackageName: "colors",
			TypeName:    "Color",
			ConstsType:  "int",
			Consts:      []string{"RED", "GREEN", "BLUE"},
		}},
		{"testdata/string.input", FileData{
			PackageName: "colors",
			TypeName:    "Color",
			ConstsType:  "string",
			Consts:      []string{"RED", "GREEN", "BLUE"},
		}},
	}

	for _, test := range tests {
		_, filename := filepath.Split(test.inputPath)
		testName := filename[:len(filename)-len(filepath.Ext(test.inputPath))]

		t.Run(testName, func(t *testing.T) {
			content, err := GenFileContent(test.input)
			if err != nil {
				t.Fatal(err)
			}

			goldenFile := filepath.Join("testdata", testName+".want")
			want, err := os.ReadFile(goldenFile)
			if err != nil {
				t.Fatal("error reading golden file:", err)
			}

			if !bytes.Equal(content.Bytes(), want) {
				t.Errorf("\n==== got:\n%s\n==== want:\n%s\n", content.String(), want)
			}
		})
	}
}

func Test_firstLetter(t *testing.T) {
	var tests = []struct {
		input string
		want  string
	}{
		{"", ""},
		{"Test", "t"},
		{"test", "t"},
		{"t", "t"},
	}

	for _, test := range tests {
		if got := firstLetter(test.input); got != test.want {
			t.Errorf("firstLetter(%q) = %v", test.input, got)
		}
	}
}
