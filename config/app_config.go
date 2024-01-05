package config

const (
	ApiGroup = "/api/v1"

	AuthorGetList = "/authors/"
	AuthorGetById = "/authors/:id"
	AuthorPost    = "/authors/"
	AuthorDelete  = "/authors/:id"
	AuthorUpdate  = "/authors"

	TaskGetList     = "/tasks"
	TaskGetById     = "/tasks/:id"
	TaskPost        = "/tasks"
	TaskDelete      = "/tasks/:id"
	TaskUpdate      = "/tasks"
	TaskGetByAuthor = "/tasks/:author"
)
