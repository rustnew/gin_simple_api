package api 


import (
	"api-go/models"
	"database/sql"
)


type PersonRepository struct {
	db *sql.DB
}

func NewPersonRepository(db *sql.DB) *PersonRepository {
	return &PersonRepository{db: db}
}


func (r *PersonRepository) Create(person *models.Person) error {
	query := `
	INSERT INTO persons (first_name, last_name, email, age, gender)
	VALUES ($1, $2, $3, $4, $5)
	RETURNING id
	`

	err := r.db.QueryRow(
		query,
		person.FirstName,
		person.LastName,
		person.Email,
		person.Age,
		person.Gender,
	).Scan(&person.ID)

	return err
}


func (r *PersonRepository) GetAll() ([]models.Person, error) {
	query := `SELECT id, first_name, last_name, email, age, gender FROM persons`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []models.Person
	for rows.Next() {
		var p models.Person
		err := rows.Scan(&p.ID, &p.FirstName, &p.LastName, &p.Email, &p.Age, &p.Gender)
		if err != nil {
			return nil, err
		}
		persons = append(persons, p)
	}

	return persons, nil
}


func (r *PersonRepository) GetByID(id int) (*models.Person, error) {
	query := `
	SELECT id, first_name, last_name, email, age, gender 
	FROM persons 
	WHERE id = $1
	`

	var person models.Person
	err := r.db.QueryRow(query, id).Scan(
		&person.ID,
		&person.FirstName,
		&person.LastName,
		&person.Email,
		&person.Age,
		&person.Gender,
	)

	if err != nil {
		return nil, err
	}

	return &person, nil
}


func (r *PersonRepository) Update(person *models.Person) error {
	query := `
	UPDATE persons 
	SET first_name = $1, last_name = $2, email = $3, age = $4, gender = $5
	WHERE id = $6
	`

	_, err := r.db.Exec(
		query,
		person.FirstName,
		person.LastName,
		person.Email,
		person.Age,
		person.Gender,
		person.ID,
	)

	return err
}


func (r *PersonRepository) Delete(id int) error {
	query := `DELETE FROM persons WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}