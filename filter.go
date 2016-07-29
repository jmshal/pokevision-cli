package main

const (
    FILTER_ONLY string = "only"
    FILTER_IGNORE string = "ignore"
)

type Filter struct {
    filterType string
    pokemonIds []int
}

func (filter *Filter) IsAllowed(id int) bool {
    found := false
    for _, pokemonId := range filter.pokemonIds {
        if pokemonId == id {
            found = true
            break
        }
    }
    return found == (filter.filterType == FILTER_ONLY)
}

func NewFilter(filterType string, pokemonIds []int) Filter {
    return Filter{filterType, pokemonIds}
}

func NewFilterFromPokedexNames(filterType string, pokedex Pokedex, names []string) (Filter, error) {
    ids := make([]int, len(names))
    for index, name := range names {
        pokedexPokemon, err := pokedex.GetByName(name)
        if err != nil {
            return Filter{}, err
        }
        ids[index] = pokedexPokemon.Index
    }
    return NewFilter(filterType, ids), nil
}
