package repository

import (
	"database/sql"
	"gotodo-app/config"
	"gotodo-app/model"
	"gotodo-app/shared/shared_model"
	"log"
	"math"
)

// TODO:
/**
1. Interface (v) -> kontrak (CRUD), mudahkan unit testing
3. Struct (v) -> data acces (db)
2. Method (v) -> implementasi dari interface
3. Function (v) -> constructor (gerbang penghubung)
*/

type TaskRepository interface {
	List(page, size int) ([]model.Task, shared_model.Paging, error)
	Create(payload model.Task) (model.Task, error)
	GetByAuthor(authorId string) ([]model.Task, error)
}

type taskRepository struct {
	db *sql.DB
}

func (t *taskRepository) GetByAuthor(authorId string) ([]model.Task, error) {
	var tasks []model.Task
	rows, err := t.db.Query(config.SelectTaskByAuthorID, authorId)
	if err != nil {
		log.Println("taskRepository.GetByAuthor.Query:", err.Error())
		return nil, err
	}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Content,
			&task.AuthorId,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("taskRepository.GetByAuthor.Rows:", err.Error())
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (t *taskRepository) List(page, size int) ([]model.Task, shared_model.Paging, error) {
	var tasks []model.Task
	offset := (page - 1) * size
	rows, err := t.db.Query(config.SelectTaskList, size, offset)
	if err != nil {
		log.Println("taskRepository.Query:", err.Error())
		return nil, shared_model.Paging{}, err
	}
	for rows.Next() {
		var task model.Task
		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Content,
			&task.AuthorId,
			&task.CreatedAt,
			&task.UpdatedAt,
		)
		if err != nil {
			log.Println("taskRepository.Rows.Next():", err.Error())
			return nil, shared_model.Paging{}, err
		}

		tasks = append(tasks, task)
	}

	totalRows := 0
	if err := t.db.QueryRow("SELECT COUNT(*) FROM tasks").Scan(&totalRows); err != nil {
		return nil, shared_model.Paging{}, err
	}

	paging := shared_model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return tasks, paging, nil
}

func (t *taskRepository) Create(payload model.Task) (model.Task, error) {
	var task model.Task
	err := t.db.QueryRow(config.InsertTask,
		payload.Title,
		payload.Content,
		payload.AuthorId).Scan(
		&task.ID,
		&task.CreatedAt,
	)
	if err != nil {
		log.Println("taskRepository.QueryRow:", err.Error())
		return model.Task{}, err
	}
	task.Title = payload.Title
	task.Content = payload.Content
	task.AuthorId = payload.AuthorId
	task.UpdatedAt = payload.UpdatedAt
	return task, nil
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{db: db}
}
