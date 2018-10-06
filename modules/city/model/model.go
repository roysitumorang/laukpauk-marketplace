package model

type (
	City struct {
		ID         int64   `json:"id" db:"id"`
		ProvinceID int64   `json:"-" db:"province_id"`
		Type       string  `json:"type" db:"type"`
		Name       string  `json:"name" db:"name"`
		CreatedBy  int64   `json:"-" db:"created_by"`
		CreatedAt  string  `json:"-" db:"created_at"`
		UpdatedBy  *int64  `json:"-" db:"updated_by"`
		UpdatedAt  *string `json:"-" db:"updated_at"`
	}
)
