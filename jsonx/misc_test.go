package jsonx

import (
	"encoding/json"
	"testing"
)

var jsonStr1 = `{
  "author": "blah@gmail.com",
  "title":  "My Blah",
  "url":    "https://blahblah.io",
  "int": 42,
  "bool": true
}`

func TestUnMarshalMap(t *testing.T) {
	var dest map[string]any // Notice Nil Map does not crash autovivification
	err := json.Unmarshal([]byte(jsonStr1), &dest)
	if err != nil {
		t.Errorf("expected no error but got:%s", err)
	}
	t.Logf("%T %v\n", dest["int"], dest["int"])
	t.Logf("%T %v\n", dest["bool"], dest["bool"])
}
