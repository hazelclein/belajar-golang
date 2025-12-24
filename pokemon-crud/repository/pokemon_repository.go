package repository

import (
	"database/sql"
	"pokemon-crud/models"
)

type PokemonRepository interface {
	Create(pokemon *models.Pokemon) error
	GetAll(limit, offset int) ([]models.Pokemon, error)
	GetByID(id string) (*models.Pokemon, error)
	Update(id string, pokemon *models.Pokemon) (int64, error)
	Delete(id string) (int64, error)
	SearchByName(name string) ([]models.Pokemon, error)
	GetByType(pokemonType string) ([]models.Pokemon, error)
	Count() (int, error)
}

type pokemonRepository struct {
	db *sql.DB
}

func NewPokemonRepository(db *sql.DB) PokemonRepository {
	return &pokemonRepository{db: db}
}

func (r *pokemonRepository) Create(pokemon *models.Pokemon) error {
	query := `INSERT OR REPLACE INTO pokemon (id, name, height, weight, base_experience, types) 
	          VALUES (?, ?, ?, ?, ?, ?)`
	_, err := r.db.Exec(query, pokemon.ID, pokemon.Name, pokemon.Height,
		pokemon.Weight, pokemon.BaseExperience, pokemon.Types)
	return err
}

func (r *pokemonRepository) GetAll(limit, offset int) ([]models.Pokemon, error) {
	query := "SELECT id, name, height, weight, base_experience, types FROM pokemon LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []models.Pokemon
	for rows.Next() {
		var p models.Pokemon
		err := rows.Scan(&p.ID, &p.Name, &p.Height, &p.Weight, &p.BaseExperience, &p.Types)
		if err != nil {
			continue
		}
		pokemons = append(pokemons, p)
	}
	return pokemons, nil
}

func (r *pokemonRepository) GetByID(id string) (*models.Pokemon, error) {
	var p models.Pokemon
	query := "SELECT id, name, height, weight, base_experience, types FROM pokemon WHERE id = ?"
	err := r.db.QueryRow(query, id).Scan(&p.ID, &p.Name, &p.Height, &p.Weight, &p.BaseExperience, &p.Types)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (r *pokemonRepository) Update(id string, pokemon *models.Pokemon) (int64, error) {
	query := `UPDATE pokemon SET name = ?, height = ?, weight = ?, base_experience = ?, types = ? WHERE id = ?`
	result, err := r.db.Exec(query, pokemon.Name, pokemon.Height, pokemon.Weight,
		pokemon.BaseExperience, pokemon.Types, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *pokemonRepository) Delete(id string) (int64, error) {
	result, err := r.db.Exec("DELETE FROM pokemon WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (r *pokemonRepository) SearchByName(name string) ([]models.Pokemon, error) {
	query := "SELECT id, name, height, weight, base_experience, types FROM pokemon WHERE name LIKE ?"
	rows, err := r.db.Query(query, "%"+name+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []models.Pokemon
	for rows.Next() {
		var p models.Pokemon
		err := rows.Scan(&p.ID, &p.Name, &p.Height, &p.Weight, &p.BaseExperience, &p.Types)
		if err != nil {
			continue
		}
		pokemons = append(pokemons, p)
	}
	return pokemons, nil
}

func (r *pokemonRepository) GetByType(pokemonType string) ([]models.Pokemon, error) {
	query := "SELECT id, name, height, weight, base_experience, types FROM pokemon WHERE types LIKE ?"
	rows, err := r.db.Query(query, "%"+pokemonType+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var pokemons []models.Pokemon
	for rows.Next() {
		var p models.Pokemon
		err := rows.Scan(&p.ID, &p.Name, &p.Height, &p.Weight, &p.BaseExperience, &p.Types)
		if err != nil {
			continue
		}
		pokemons = append(pokemons, p)
	}
	return pokemons, nil
}

func (r *pokemonRepository) Count() (int, error) {
	var total int
	err := r.db.QueryRow("SELECT COUNT(*) FROM pokemon").Scan(&total)
	return total, err
}