package repository

import (
	"database/sql"
	"gotodo-app/config"
	"gotodo-app/model"
	"log"
)

type AuthorRepository interface {
	List(author string) ([]model.Author, error)
	Get(id string) (model.Author, error)
	GetByEmail(email string) (model.Author, error)
}

type authorRepository struct {
	db *sql.DB
}

func (a *authorRepository) GetByEmail(email string) (model.Author, error) {
	var author model.Author
	err := a.db.QueryRow(`SELECT id, email, password, role FROM authors WHERE email = $1`, email).Scan(&author.ID, &author.Email, &author.Password, &author.Role)
	if err != nil {
		log.Println("authorRepository.GetByEmail:", err.Error())
		return model.Author{}, err
	}
	return author, nil
}

func (a *authorRepository) List(author string) ([]model.Author, error) {
	var authors []model.Author
	var rows *sql.Rows
	var err error
	// cek get by id
	aut, _ := a.Get(author)
	if aut.Role == "admin" {
		// query admin
		rows, err = a.db.Query(config.SelectAuthorListRoleAdmin)
	} else {
		// query user
		rows, err = a.db.Query(config.SelectAuthorListRoleUser, author)
	}

	if err != nil {
		log.Println("authorRepository.List.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var author model.Author
		err := rows.Scan(&author.ID, &author.Name, &author.Email, &author.CreatedAt, &author.UpdatedAt)
		if err != nil {
			log.Println("authorRepository.List.rows.Next:", err.Error())
			return nil, err
		}
		tasksRows, err := a.db.Query(config.SelectAuthorWithTask, author.ID)
		if err != nil {
			log.Println("authorRepository.Get.Query:", err.Error())
			return nil, err
		}
		for tasksRows.Next() {
			var task model.Task
			err := tasksRows.Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt)
			if err != nil {
				log.Println("authorRepository.tasksRows.Next():", err.Error())
				return nil, err
			}
			author.Tasks = append(author.Tasks, task)
		}
		authors = append(authors, author)
	}
	return authors, nil
}

func (a *authorRepository) Get(id string) (model.Author, error) {
	var author model.Author
	err := a.db.QueryRow(config.SelectAuthorByID, id).Scan(
		&author.ID,
		&author.Name,
		&author.Email,
		&author.Role,
		&author.CreatedAt,
		&author.UpdatedAt,
	)
	if err != nil {
		log.Println("authorRepository.Get.QueryRow:", err.Error())
		return model.Author{}, err
	}
	rows, err := a.db.Query(config.SelectAuthorWithTask, id)
	if err != nil {
		log.Println("authorRepository.Get.Query:", err.Error())
		return model.Author{}, err
	}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Content, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("authorRepository.rows.Next():", err.Error())
			return model.Author{}, err
		}
		author.Tasks = append(author.Tasks, task)
	}
	return author, nil
}

func NewAuthorRepository(db *sql.DB) AuthorRepository {
	return &authorRepository{db: db}
}
