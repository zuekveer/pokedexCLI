package main

import (
	"time"

	"github.com/zuekveer/courses/projects/pokedexcli/internal/pokeapi"
	"github.com/zuekveer/pokedexCLI/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
