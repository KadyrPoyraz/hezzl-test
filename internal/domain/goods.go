package domain

type Good struct {
	ID          int64  `json:"id" db:"id"`
	ProjectID   int64  `json:"project_id" db:"project_id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description,omitempty" db:"description"`
	Priority    int    `json:"priority,omitempty" db:"priority"`
	Removed     bool   `json:"removed,omitempty" db:"removed"`
	CreatedAt   int64  `json:"created_at" db:"created_at"`
}
