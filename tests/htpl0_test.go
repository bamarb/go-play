package tests

import (
	"html/template"
	"os"
	"strings"
	"testing"
)

// slurps fixtures that are files in testdata,
// name is the name of the file
func slurp(name string) []byte {
	str, err := os.ReadFile(name)
	if err != nil {
		panic(err)
	}
	return str
}

func TestTemplateBasic(t *testing.T) {
	tpl, err := template.ParseFiles("testdata/tst.tpl")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Defined Templates:%s\n", tpl.DefinedTemplates())
	// Define Template Data
	type Dog struct {
		Name string
		Age  int
	}
	type Test struct {
		HTML     string
		SafeHTML template.HTML
		Title    string
		Path     string
		Dog      Dog
		Map      map[string]string
	}
	data := Test{
		HTML:     "<h1>A Header!</h1>",
		SafeHTML: template.HTML(`<script>alert("hash bang")</script>`),
		Title:    "Ringa Ringa Roses \"\\\" Backslash",
		Path:     "/dashboard/settings?wide=true",
		Dog:      Dog{"Shaggy", 1},
		Map: map[string]string{
			"ringa":  "roses",
			"pocket": "posies",
		},
	}
	err = tpl.Execute(os.Stdout, data)
	if err != nil {
		t.Fatal(err)
	}
}

func TestCompose_basic(t *testing.T) {
	// Parse multiple Templates , the order of the template files matters
	// Execute Executes the "current" template

	tplx, _ := template.ParseFiles("testdata/base.tpl", "testdata/hello.tpl")
	t.Logf("Current Template executed by Execute:%s\n", tplx.Name())
	tpl, err := template.ParseFiles("testdata/hello.tpl", "testdata/base.tpl")
	t.Logf("Current Template:%s\n", tpl.Name())
	if err != nil {
		t.Fatal(err)
	}
	data := map[string]string{
		"name":    "Binga",
		"country": "Bharat",
	}
	strWriter := strings.Builder{}
	err = tpl.Execute(&strWriter, data)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(strWriter.String())
	strWriter.Reset()
	// What happens if a variable used in a template  is non existant
	err = tpl.Execute(&strWriter, map[string]string{"name": "Ringa"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(strWriter.String())
}

func TestTemplate_new(t *testing.T) {
}
