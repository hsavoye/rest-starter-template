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

.
├── db/
│   └── db.go          # Connexion à la base de données
├── handlers/
│   └── routes.go      # Définition des routes et logique HTTP
├── main.go            # Point d’entrée de l’application
├── .env               # Variables d’environnement (non versionnées)
├── go.mod             # Dépendances du projet
└── README.md

````

## Configuration

Crée un fichier `.env` à la racine du projet avec les variables suivantes :

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
   git clone https://github.com/tonpseudo/go-articles-api.git
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

Exemple d’appel POST :

```json
POST /create
{
  "title": "Mon premier article"
}
```

## Test rapide avec curl

```bash
# Créer un article
curl -X POST http://localhost:8081/create \
  -H "Content-Type: application/json" \
  -d '{"title": "Un article en Go"}'

# Lister les articles
curl http://localhost:8081/get
```

## Dépendances principales

* [Go](https://golang.org/)
* [Gorilla Mux](https://github.com/gorilla/mux)
* [PostgreSQL](https://www.postgresql.org/)
* [joho/godotenv](https://github.com/joho/godotenv)

