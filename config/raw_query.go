package config

const (
	SelectTaskList       = `SELECT id, title, content, author_id, created_at, updated_at FROM tasks ORDER BY created_at DESC LIMIT $1 OFFSET $2`
	InsertTask           = `INSERT INTO tasks (title, content, author_id) VALUES ($1, $2, $3) RETURNING id, created_at`
	SelectTaskByAuthorID = `SELECT id, title, content, author_id, created_at, updated_at FROM tasks WHERE author_id = $1 ORDER BY created_at DESC`

	SelectAuthorByID          = `SELECT id, name, email, role, created_at, updated_at FROM authors WHERE id = $1`
	SelectAuthorListRoleUser  = `SELECT id, name, email, created_at, updated_at FROM authors WHERE id = $1 ORDER BY created_at DESC `
	SelectAuthorListRoleAdmin = `SELECT id, name, email, created_at, updated_at FROM authors ORDER BY created_at DESC `
	SelectAuthorWithTask      = `SELECT t.id, t.title, t.content, t.created_at, t.updated_at FROM authors a JOIN tasks t on a.id = t.author_id
WHERE a.id = $1;`
)
