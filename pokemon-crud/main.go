package main

import (
	"log"
	"net/http"

	"pokemon-crud/config"
	"pokemon-crud/handlers"
	"pokemon-crud/repository"
	"pokemon-crud/utils"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	defer config.DB.Close()

	pokemonRepo := repository.NewPokemonRepository(config.DB)

	pokemonHandler := handlers.NewPokemonHandler(pokemonRepo)

	r := mux.NewRouter()

	// Routes
	r.HandleFunc("/pokemon/fetch/{id}", pokemonHandler.FetchPokemonFromAPI).Methods("POST")
	r.HandleFunc("/pokemon", pokemonHandler.GetAllPokemon).Methods("GET")
	r.HandleFunc("/pokemon/search", pokemonHandler.SearchPokemonByName).Methods("GET")
	r.HandleFunc("/pokemon/type/{type}", pokemonHandler.GetPokemonByType).Methods("GET")
	r.HandleFunc("/pokemon/{id}", pokemonHandler.GetPokemonByID).Methods("GET")
	r.HandleFunc("/pokemon/{id}", pokemonHandler.UpdatePokemon).Methods("PUT")
	r.HandleFunc("/pokemon/{id}", pokemonHandler.DeletePokemon).Methods("DELETE")

	r.Use(utils.LoggingMiddleware)

	log.Println("========================================")
	log.Println("Pokemon CRUD API Server")
	log.Println("Server berjalan di http://localhost:8080")
	log.Println("========================================")

	log.Fatal(http.ListenAndServe(":8080", r))
}
