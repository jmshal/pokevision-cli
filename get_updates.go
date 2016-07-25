package main

type Updates struct {
    New []Pokemon
}

func getPokemonUIDs(pokemons []Pokemon) []string {
    var ids []string
    for _, pokemon := range pokemons {
        ids = append(ids, pokemon.UID)
    }
    return ids
}

func contains(arr []string, item string) bool {
    for _, value := range arr {
        if value == item {
            return true
        }
    }
    return false
}

func GetUpdates(previous, current []Pokemon) Updates {
    previousUIDs := getPokemonUIDs(previous)
    var currentUIDs []string
    var new []Pokemon
    for _, pokemon := range current {
        // Only if the pokemon hasn't been seen yet
        // And while we're at it, de-dupe the ones with the same UIDs
        if !contains(previousUIDs, pokemon.UID) && !contains(currentUIDs, pokemon.UID) {
            new = append(new, pokemon)
            currentUIDs = append(currentUIDs, pokemon.UID)
        }
    }
    return Updates{new}
}
