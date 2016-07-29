package main

type Updates struct {
    New []Pokemon
}

func getPokemonIDs(pokemons []Pokemon) []int {
    var ids []int
    for _, pokemon := range pokemons {
        ids = append(ids, pokemon.ID)
    }
    return ids
}

func contains(arr []int, item int) bool {
    for _, value := range arr {
        if value == item {
            return true
        }
    }
    return false
}

func GetUpdates(previous, current []Pokemon) Updates {
    previousIDs := getPokemonIDs(previous)
    var currentIDs []int
    var new []Pokemon
    for _, pokemon := range current {
        // Only if the pokemon hasn't been seen yet
        // And while we're at it, de-dupe the ones with the same UIDs
        if !contains(previousIDs, pokemon.ID) && !contains(currentIDs, pokemon.ID) {
            new = append(new, pokemon)
            currentIDs = append(currentIDs, pokemon.ID)
        }
    }
    return Updates{new}
}
