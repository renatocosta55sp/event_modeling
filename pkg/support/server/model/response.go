package response

import "github.com/google/uuid"

type List struct {
	Pagination Pagination `json:"pagination"`
	Data       any        `json:"data"`
}

type Pagination struct {
	Total int64 `json:"total"`
}

type Created struct {
	ID uuid.UUID `json:"id"`
}

type CreatedInt64 struct {
	ID int64 `json:"id"`
}

type Exception struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}
