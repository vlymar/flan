package main

import (
	"io/ioutil"
	"os"
	"testing"
)

const testFlanFile = "testdata/flanfixture.json"

func TestReadFlanFile(t *testing.T) {
	commands, err := ReadFlanFile(testFlanFile)
	if err != nil {
		t.Error(err)
	}

	if len(commands) != 2 {
		t.Error("expected 2 commands, got", len(commands))
	}

	cmd_1_anno, prs := commands["cmd_1"]
	if !prs {
		t.Error("cmd_1 not in commands map")
	}
	if cmd_1_anno[0] != "cmd_1 anno" {
		t.Error("missing annotation for cmd_1")
	}
	if cmd_1_anno[1] != "cmd_1 example" {
		t.Error("missing example for cmd_1")
	}

	commands, err = ReadFlanFile("non_existant_file.xyz")
	if err != nil {
		t.Error(err)
	}

	if len(commands) != 0 {
		t.Error(`ReadFlanFile did't return empty Commands map for
		        non existant flonfile`)
	}
}

func TestWriteFlanFile(t *testing.T) {
	tmpfile := "testdata/tmpfile.json"
	commands := Commands{
		"a": []string{"a1", "a2"},
		"b": []string{"b1", "b2"},
	}

	if err := WriteFlanFile(commands, tmpfile); err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpfile)

	dat, err := ioutil.ReadFile(tmpfile)
	if err != nil {
		t.Error(err)
	}

	if string(dat) != `{"a":["a1","a2"],"b":["b1","b2"]}` {
		t.Error("WriteFlanFile wrote unexpected value:", string(dat))
	}
}
