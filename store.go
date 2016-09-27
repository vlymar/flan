package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"os/user"
	"path"
)

const flanFile = ".flan"

type Commands map[string][][]string

func Store(cmd, example, anno string, commands Commands) {
	flannotations, prs := commands[cmd]
	if prs {
		flannotations = append(flannotations, []string{example, anno})
	} else {
		flannotations = make([][]string, 1)
		flannotations[0] = []string{example, anno}
	}
	commands[cmd] = flannotations
}

func ReadFlanFile(path string) (Commands, error) {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return make(Commands), nil
		}
		return nil, err
	}

	var commands Commands
	if err := json.Unmarshal(dat, &commands); err != nil {
		return nil, err
	}
	return commands, nil
}

func WriteFlanFile(commands Commands, path string) error {
	b, err := json.Marshal(commands)
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
