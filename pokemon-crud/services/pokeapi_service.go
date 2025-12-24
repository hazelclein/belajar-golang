package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"pokemon-crud/config"
	"pokemon-crud/models"
	"strings"
)

func FetchPokemonFromPokeAPI(pokemonID string) (*models.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", config.AppConfig.PokeAPIBaseURL, pokemonID)
	
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Pokemon tidak ditemukan di PokeAPI")
	}

	var pokeAPIResp models.PokeAPIResponse
	err = json.NewDecoder(resp.Body).Decode(&pokeAPIResp)
	if err != nil {
		return nil, err
	}

	var types []string
	for _, t := range pokeAPIResp.Types {
		types = append(types, t.Type.Name)
	}
	typesStr := strings.Join(types, ", ")

	pokemon := &models.Pokemon{
		ID:             pokeAPIResp.ID,
		Name:           pokeAPIResp.Name,
		Height:         pokeAPIResp.Height,
		Weight:         pokeAPIResp.Weight,
		BaseExperience: pokeAPIResp.BaseExperience,
		Types:          typesStr,
	}

	return pokemon, nil
}