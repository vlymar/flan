package main

import (
	"errors"
	"fmt"
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

	cmd_1_annos, prs := commands["cmd_1"]
	if !prs {
		t.Error("cmd_1 not in commands map")
	}
	if len(cmd_1_annos) != 2 {
		t.Error("cmd_1 does not have 2 annotations in commands map")
	}
	if cmd_1_annos[0][0] != "cmd_1 anno1" {
		t.Error("missing annotation1 for cmd_1")
	}
	if cmd_1_annos[0][1] != "cmd_1 example1" {
		t.Error("missing example1 for cmd_1")
	}
	if cmd_1_annos[1][0] != "cmd_1 anno2" {
		t.Error("missing annotation2 for cmd_1")
	}
	if cmd_1_annos[1][1] != "cmd_1 example2" {
		t.Error("missing example2 for cmd_1")
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
	cmd1 := make([][]string, 1)
	cmd1[0] = []string{"a1", "a2"}
	cmd2 := make([][]string, 1)
	cmd2[0] = []string{"b1", "b2"}
	commands := Commands{
		"a": cmd1,
		"b": cmd2,
	}

	if err := WriteFlanFile(commands, tmpfile); err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpfile)

	dat, err := ioutil.ReadFile(tmpfile)
	if err != nil {
		t.Error(err)
	}

	if string(dat) != `{"a":[["a1","a2"]],"b":[["b1","b2"]]}` {
		t.Error("WriteFlanFile wrote unexpected value:", string(dat))
	}
}

func TestStore(t *testing.T) {
	commands := make(Commands)

	Store("ls", "ls -a", "list all files", commands)
	checkStorage("ls", "ls -a", "list all files", commands)

	Store("ls", "ls -l", "list files, showing symlinks", commands)
	checkStorage("ls", "ls -l", "list files, showing symlinks", commands)

	Store("man", "man cat", "see manpage for cat", commands)
	checkStorage("man", "man cat", "see manpage for cat", commands)
}

func checkStorage(cmd, example, anno string, commands Commands) error {
	flannotations := commands[cmd]

	for _, flanno := range flannotations {
		stored_ex := flanno[0]
		if stored_ex == example {
			stored_anno := flanno[1]
			if stored_anno == anno {
				return nil
			}
		}
	}

	return errors.New(fmt.Sprintf("did not find expected flannotation: %v, %v", example, anno))
}
