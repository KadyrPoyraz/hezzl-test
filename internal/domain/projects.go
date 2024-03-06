package domain

type Project struct {
	Id        int64  `json:"id" db:"id"`
	Name      string `json:"name" db:"name"`
	CreatedAt int64  `json:"created_at" db:"created_at"`
}
