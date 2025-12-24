package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"pokemon-crud/config"
	"pokemon-crud/models"
	"pokemon-crud/repository"
	"pokemon-crud/services"
	"pokemon-crud/utils"

	"github.com/gorilla/mux"
)

type PokemonHandler struct {
	repo repository.PokemonRepository
}

func NewPokemonHandler(repo repository.PokemonRepository) *PokemonHandler {
	return &PokemonHandler{repo: repo}
}

func (h *PokemonHandler) FetchPokemonFromAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonID := vars["id"]

	if err := utils.ValidatePokemonID(pokemonID); err != nil {
		utils.SendError(w, "Validasi gagal", err, http.StatusBadRequest)
		return
	}

	pokemon, err := services.FetchPokemonFromPokeAPI(pokemonID)
	if err != nil {
		utils.SendError(w, "Gagal mengambil data dari PokeAPI", err, http.StatusNotFound)
		return
	}

	err = h.repo.Create(pokemon)
	if err != nil {
		utils.SendError(w, "Gagal menyimpan ke database", err, http.StatusInternalServerError)
		return
	}

	utils.SendSuccess(w, "Pokemon berhasil disimpan", pokemon, http.StatusCreated)
}

func (h *PokemonHandler) GetAllPokemon(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")

	pageInt := 1
	limitInt := config.AppConfig.DefaultPageSize

	if page != "" {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			pageInt = p
		}
	}

	if limit != "" {
		if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= config.AppConfig.MaxPageSize {
			limitInt = l
		}
	}

	offset := (pageInt - 1) * limitInt

	pokemons, err := h.repo.GetAll(limitInt, offset)
	if err != nil {
		utils.SendError(w, "Gagal mengambil data", err, http.StatusInternalServerError)
		return
	}

	total, err := h.repo.Count()
	if err != nil {
		utils.SendError(w, "Gagal menghitung total data", err, http.StatusInternalServerError)
		return
	}

	totalPages := (total + limitInt - 1) / limitInt

	if len(pokemons) == 0 {
		response := map[string]interface{}{
			"data": []models.Pokemon{},
			"pagination": map[string]interface{}{
				"page":        pageInt,
				"limit":       limitInt,
				"total_data":  total,
				"total_pages": totalPages,
			},
		}
		utils.SendSuccess(w, "Tidak ada pokemon di database", response, http.StatusOK)
		return
	}

	response := map[string]interface{}{
		"data": pokemons,
		"pagination": map[string]interface{}{
			"page":        pageInt,
			"limit":       limitInt,
			"total_data":  total,
			"total_pages": totalPages,
		},
	}
	
	utils.SendSuccess(w, "Berhasil mengambil data pokemon", response, http.StatusOK)
}

func (h *PokemonHandler) GetPokemonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := utils.ValidatePokemonID(id); err != nil {
		utils.SendError(w, "Validasi gagal", err, http.StatusBadRequest)
		return
	}

	p, err := h.repo.GetByID(id)
	if err != nil {
		utils.SendError(w, "Pokemon tidak ditemukan", err, http.StatusNotFound)
		return
	}

	utils.SendSuccess(w, "Berhasil mengambil data pokemon", p, http.StatusOK)
}

func (h *PokemonHandler) UpdatePokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := utils.ValidatePokemonID(id); err != nil {
		utils.SendError(w, "Validasi gagal", err, http.StatusBadRequest)
		return
	}

	var p models.Pokemon
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		utils.SendError(w, "Format JSON tidak valid", err, http.StatusBadRequest)
		return
	}

	if err := utils.ValidatePokemon(&p); err != nil {
		utils.SendError(w, "Validasi data gagal", err, http.StatusBadRequest)
		return
	}

	rowsAffected, err := h.repo.Update(id, &p)
	if err != nil {
		utils.SendError(w, "Gagal update data", err, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		utils.SendError(w, "Pokemon tidak ditemukan", nil, http.StatusNotFound)
		return
	}

	utils.SendSuccess(w, "Pokemon berhasil diupdate", nil, http.StatusOK)
}

func (h *PokemonHandler) DeletePokemon(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := utils.ValidatePokemonID(id); err != nil {
		utils.SendError(w, "Validasi gagal", err, http.StatusBadRequest)
		return
	}

	rowsAffected, err := h.repo.Delete(id)
	if err != nil {
		utils.SendError(w, "Gagal menghapus data", err, http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		utils.SendError(w, "Pokemon tidak ditemukan", nil, http.StatusNotFound)
		return
	}

	utils.SendSuccess(w, "Pokemon berhasil dihapus", nil, http.StatusOK)
}

func (h *PokemonHandler) SearchPokemonByName(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	
	if strings.TrimSpace(name) == "" {
		utils.SendError(w, "Parameter 'name' tidak boleh kosong", nil, http.StatusBadRequest)
		return
	}
	
	pokemons, err := h.repo.SearchByName(name)
	if err != nil {
		utils.SendError(w, "Gagal mencari pokemon", err, http.StatusInternalServerError)
		return
	}
	
	if len(pokemons) == 0 {
		utils.SendSuccess(w, "Tidak ada pokemon yang ditemukan", []models.Pokemon{}, http.StatusOK)
		return
	}
	
	utils.SendSuccess(w, "Berhasil menemukan pokemon", pokemons, http.StatusOK)
}

func (h *PokemonHandler) GetPokemonByType(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pokemonType := vars["type"]
	
	if strings.TrimSpace(pokemonType) == "" {
		utils.SendError(w, "Type tidak boleh kosong", nil, http.StatusBadRequest)
		return
	}
	
	pokemons, err := h.repo.GetByType(pokemonType)
	if err != nil {
		utils.SendError(w, "Gagal mencari pokemon berdasarkan type", err, http.StatusInternalServerError)
		return
	}
	
	if len(pokemons) == 0 {
		utils.SendSuccess(w, "Tidak ada pokemon dengan type tersebut", []models.Pokemon{}, http.StatusOK)
		return
	}
	
	utils.SendSuccess(w, "Berhasil menemukan pokemon", pokemons, http.StatusOK)
}