package main

import (
	"encoding/json"
)

type Command string
type Flannotation []string
type Store map[Command]Flannotation
