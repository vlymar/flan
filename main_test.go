package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

const testFlanFile = "testdata/flanfixture.json"

func TestReadFlanFileFromPath(t *testing.T) {
	commands, err := readFlanFileFromPath(testFlanFile)
	if err != nil {
		t.Error(err)
	}

	if len(commands) != 2 {
		t.Error("expected 2 commands, got", len(commands))
	}

	cmd1Annos, prs := commands["cmd_1"]
	if !prs {
		t.Error("cmd_1 not in commands map")
	}
	if len(cmd1Annos) != 2 {
		t.Error("cmd_1 does not have 2 annotations in commands map")
	}
	if cmd1Annos[0][0] != "cmd_1 anno1" {
		t.Error("missing annotation1 for cmd_1")
	}
	if cmd1Annos[0][1] != "cmd_1 example1" {
		t.Error("missing example1 for cmd_1")
	}
	if cmd1Annos[1][0] != "cmd_1 anno2" {
		t.Error("missing annotation2 for cmd_1")
	}
	if cmd1Annos[1][1] != "cmd_1 example2" {
		t.Error("missing example2 for cmd_1")
	}

	commands, err = readFlanFileFromPath("non_existant_file.xyz")
	if err != nil {
		t.Error(err)
	}

	if len(commands) != 0 {
		t.Error(`readFlanFileFromPath() did't return empty Commands map for
		        non existant flonfile`)
	}
}

func TestWriteFlanFile(t *testing.T) {
	tmpfile := "testdata/tmpfile.json"
	cmd1 := make([][]string, 1)
	cmd1[0] = []string{"a1", "a2"}
	cmd2 := make([][]string, 1)
	cmd2[0] = []string{"b1", "b2"}
	cmds := commands{
		"a": cmd1,
		"b": cmd2,
	}

	if err := writeFlanFileToPath(cmds, tmpfile); err != nil {
		t.Error(err)
	}
	defer os.Remove(tmpfile)

	dat, err := ioutil.ReadFile(tmpfile)
	if err != nil {
		t.Error(err)
	}

	if string(dat) != `{"a":[["a1","a2"]],"b":[["b1","b2"]]}` {
		t.Error("WriteFlanFileToPath wrote unexpected value:", string(dat))
	}
}

func TestStore(t *testing.T) {
	cmds := make(commands)

	store("ls", "ls -a", "list all files", cmds)
	checkStorage("ls", "ls -a", "list all files", cmds)

	store("ls", "ls -l", "list files, showing symlinks", cmds)
	checkStorage("ls", "ls -l", "list files, showing symlinks", cmds)

	store("man", "man cat", "see manpage for cat", cmds)
	checkStorage("man", "man cat", "see manpage for cat", cmds)
}

func checkStorage(cmd, example, anno string, cmds commands) error {
	flannotations := cmds[cmd]

	for _, flanno := range flannotations {
		storedEx := flanno[0]
		if storedEx == example {
			storedAnno := flanno[1]
			if storedAnno == anno {
				return nil
			}
		}
	}

	return fmt.Errorf("did not find expected flannotation: %v, %v", example, anno)
}
