package main


import (
	"os/exec"
)

func OutputToNotificationCenter(pokemon PokemonMeta, config Config) error {

    exec.Command("notify-send", "-i", "", pokemon.Name + " located near " + pokemon.Location, "Expires in " + HumanTime(pokemon.ExpiresAt)).Run()
    return nil
}
