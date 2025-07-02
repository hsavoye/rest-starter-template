package main

import (
	"log"
	"net/http"
	"serv/db"
	"serv/handlers"
	"time"

	"github.com/gorilla/mux"
)

func middlewareLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s", r.Method, r.URL.Path, time.Since(start))
	})
}

func main() {
	DB, err := db.InitDB()
	if err != nil {
		log.Fatalf("Erreur d'initialisation de la base de données : %v", err)
	}
	defer DB.Close()

	h := &handlers.Handler{DB: DB}

	router := mux.NewRouter()

	router.Use(middlewareLogger)
	
	router.HandleFunc("/", handlers.HomeHandler).Methods(http.MethodGet)
	router.HandleFunc("/articles", h.GetArticles).Methods(http.MethodGet)
	router.HandleFunc("/articles", h.CreateArticle).Methods(http.MethodPost)

	log.Println("Serveur lancé sur http://localhost:8081")
	if err := http.ListenAndServe(":8081", router); err != nil {
		log.Fatalf("Erreur au lancement du serveur: %v", err)
	}
}
