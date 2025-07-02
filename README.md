# Go Articles API

Ce projet est une API REST écrite en Go, utilisant le routeur Gorilla Mux et une base de données PostgreSQL. Elle permet de gérer une collection d’articles via des opérations de base (GET, POST). C’est un bon point de départ pour construire une API modulaire en Go.

## Fonctionnalités

- API REST en Go
- Récupération et création d’articles
- Routage avec Gorilla Mux
- Connexion PostgreSQL
- Configuration par fichier `.env`
- Architecture claire et modulaire 

## Structure du projet

```
serv/
├── db/
    └── db.go          → Connexion à la db 
├── handlers/
    └── routes.go      → Fonctions des routes HTTP (Get,Post) 
├── models/            
    └── article.go     → Structure des articles
├── .env               → Variables pour accéder db
├── go.mod    
└── main.go            → Point entrée de l'app
````

## Configuration
**Attention à bien configurer votre BDD PostgreSQL dans pgAdmin4 (Port, User, Password).** 

Crée un fichier `.env` à la racine du projet avec les variables suivantes : (en fonction de votre bdd Postgres)

```env
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=postgres
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable
````

Ces variables seront chargées automatiquement au lancement de l’application.

## Installation

1. Clone le dépôt :

   ```bash
   git clone https://github.com/hsavoye/rest-starter-template.git
   ```
   ```bash
   cd go-articles-api
   ```

2. Installe les dépendances :

   ```bash
   go mod tidy
   ```

## Lancer le serveur

```bash
go run main.go
```

Le serveur démarrera sur `http://localhost:8081`

## Endpoints disponibles

| Méthode | URL       | Description                |
| ------: | --------- | -------------------------- |
|     GET | `/`       | Message d’accueil          |
|     GET | `/get`    | Récupère tous les articles |
|    POST | `/create` | Crée un nouvel article     |

## Test rapide avec curl

```bash
# Créer un article (Windows)
curl -X POST http://localhost:8081/create -H "Content-Type: application/json" -d "{\"title\": \"Nouveau contenu\"}"

# Lister les articles (Windows)
curl http://localhost:8081/articles   
```

## Dépendances principales

* [Go](https://golang.org/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [PostgreSQL](https://www.postgresql.org/)
* [joho/godotenv](https://github.com/joho/godotenv)

