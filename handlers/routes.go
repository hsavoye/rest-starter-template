package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"serv/models"

	_ "github.com/lib/pq"
)

type Handler struct {
	DB *sql.DB
}

// HomeHandler gère la route "/"
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bienvenue sur le serveur Go"))
}

// GetArticles gère la récupération des articles
func (h *Handler) GetArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	rows, err := h.DB.Query("SELECT id, title FROM articles")
	if err != nil {
		log.Printf("Erreur DB : %v", err)
		http.Error(w, "Erreur lors de la récupération des articles", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var articles []models.Article
	for rows.Next() {
		var a models.Article
		if err := rows.Scan(&a.ID, &a.Title); err != nil {
			log.Printf("Erreur scan : %v", err)
			http.Error(w, "Erreur lors du traitement des données", http.StatusInternalServerError)
			return
		}
		articles = append(articles, a)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

// CreateArticle gère la création d'un article
func (h *Handler) CreateArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Méthode non autorisée", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Title string `json:"title"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Title == "" {
		http.Error(w, "Requête invalide : titre requis", http.StatusBadRequest)
		return
	}

	var insertedID int
	err := h.DB.QueryRow("INSERT INTO articles (title) VALUES ($1) RETURNING id", req.Title).Scan(&insertedID)
	if err != nil {
		log.Printf("Erreur DB : %v", err)
		http.Error(w, "Erreur lors de la création de l'article", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Article créé",
		"id":      insertedID,
	})
}
