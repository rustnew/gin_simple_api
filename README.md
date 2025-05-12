Voici une documentation prête à être intégrée dans votre `README.md` pour décrire le fonctionnement de votre **`PersonRepository`** dans un projet Go avec une base de données SQL :

---

## 📁 `PersonRepository` — Accès aux données

Le fichier `api/person_repository.go` contient le code de la couche de persistance pour la gestion des entités `Person`. Il fournit une interface directe avec la base de données PostgreSQL.

### 🔧 Dépendances utilisées

* `database/sql` — bibliothèque standard pour les interactions SQL
* `models.Person` — structure représentant une personne
* PostgreSQL — utilisé pour stocker les données

---

## 🧱 Structure

```go
type PersonRepository struct {
    db *sql.DB
}
```

### 🔨 Constructeur

```go
func NewPersonRepository(db *sql.DB) *PersonRepository
```

Crée un nouveau repository en injectant une connexion SQL active (`*sql.DB`).

---

## 📚 Méthodes disponibles

### ➕ `Create`

```go
func (r *PersonRepository) Create(person *models.Person) error
```

Insère une nouvelle personne dans la table `persons`.

* 🔢 Remplit automatiquement l'`ID` du modèle après insertion.
* 🛠️ Paramètres : `*models.Person`
* 📤 Retourne une erreur en cas d'échec.

---

### 📥 `GetAll`

```go
func (r *PersonRepository) GetAll() ([]models.Person, error)
```

Récupère toutes les personnes de la base de données.

* 📥 Retourne une liste de `models.Person`.
* 📤 Retourne une erreur en cas d'échec.

---

### 🔍 `GetByID`

```go
func (r *PersonRepository) GetByID(id int) (*models.Person, error)
```

Récupère une personne selon son `ID`.

* 🧾 Paramètre : `id int`
* 📥 Retourne un pointeur vers `models.Person`
* 📤 Retourne une erreur si aucun résultat ou si échec

---

### ✏️ `Update`

```go
func (r *PersonRepository) Update(person *models.Person) error
```

Met à jour une personne existante selon son `ID`.

* 🧾 Paramètre : `*models.Person` (avec ID rempli)
* 📤 Retourne une erreur en cas d'échec

---

### ❌ `Delete`

```go
func (r *PersonRepository) Delete(id int) error
```

Supprime une personne de la base selon son `ID`.

* 🧾 Paramètre : `id int`
* 📤 Retourne une erreur en cas d’échec

---

## 🗃️ Structure SQL ciblée

Votre table PostgreSQL `persons` devrait avoir la structure suivante :

```sql
CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email TEXT,
    age INT,
    gender TEXT
);
```

---

## 🧪 Exemple d’utilisation

```go
repo := api.NewPersonRepository(db)
err := repo.Create(&models.Person{FirstName: "John", LastName: "Doe", Email: "john@example.com", Age: 30, Gender: "M"})
```

---

Souhaitez-vous aussi que je vous rédige un exemple d'architecture du projet dans le README ?
