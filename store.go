package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const flanFile = ".flan"

type commands map[string][][]string

func store(cmd, example, anno string, cmds commands) {
	flannotations, prs := cmds[cmd]
	if prs {
		flannotations = append(flannotations, []string{example, anno})
	} else {
		flannotations = make([][]string, 1)
		flannotations[0] = []string{example, anno}
	}
	cmds[cmd] = flannotations
}

func readFlanFile() (commands, error) {
	flanPath, err := flanPath()
	if err != nil {
		return nil, err
	}

	cmds, err := readFlanFileFromPath(flanPath)
	if err != nil {
		return nil, err
	}

	return cmds, nil
}

func readFlanFileFromPath(path string) (commands, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return make(commands), nil
		}
		return nil, err
	}

	var cmds commands
	if err := json.Unmarshal(dat, &cmds); err != nil {
		return nil, err
	}
	return cmds, nil
}

func writeFlanFile(cmds commands) error {
	flanPath, err := flanPath()
	if err != nil {
		return err
	}
	if err := writeFlanFileToPath(cmds, flanPath); err != nil {
		return err
	}
	return nil
}

func writeFlanFileToPath(cmds commands, path string) error {
	b, err := json.Marshal(cmds)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, b, 0644); err != nil {
		return err
	}
	return nil
}

func flanPath() (string, error) {
	usr, err := user.Current()
	if err != nil {
		return "", err
	}

	p := path.Join(usr.HomeDir, flanFile)

	return p, nil
}
