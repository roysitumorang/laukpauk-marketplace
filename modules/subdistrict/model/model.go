package model

type (
	Subdistrict struct {
		ID        int64   `json:"id" db:"id"`
		CityID    int64   `json:"-" db:"city_id"`
		Name      string  `json:"name" db:"name"`
		CreatedBy int64   `json:"-" db:"created_by"`
		CreatedAt string  `json:"-" db:"created_at"`
		UpdatedBy *int64  `json:"-" db:"updated_by"`
		UpdatedAt *string `json:"-" db:"updated_at"`
	}
)
