package utils

import (
	"errors"
	"pokemon-crud/models"
	"strings"
)

// V Pokemon ID
func ValidatePokemonID(id string) error {
	if id == "" {
		return errors.New("Pokemon ID tidak boleh kosong")
	}
	
	if id == "0" {
		return errors.New("Pokemon ID harus lebih dari 0")
	}
	
	return nil
}

// V Pokemon data
func ValidatePokemon(pokemon *models.Pokemon) error {
	// V nama
	if strings.TrimSpace(pokemon.Name) == "" {
		return errors.New("Nama pokemon tidak boleh kosong")
	}
	
	if len(pokemon.Name) < 2 {
		return errors.New("Nama pokemon minimal 2 karakter")
	}
	
	// V height
	if pokemon.Height < 0 {
		return errors.New("Height tidak boleh negatif")
	}
	
	// V weight
	if pokemon.Weight < 0 {
		return errors.New("Weight tidak boleh negatif")
	}
	
	// V base experience
	if pokemon.BaseExperience < 0 {
		return errors.New("Base experience tidak boleh negatif")
	}
	
	// V types
	if strings.TrimSpace(pokemon.Types) == "" {
		return errors.New("Types tidak boleh kosong")
	}
	
	return nil
}