package models

import (
	"database/sql"
	"time"
)

type Models struct {
	DB DBModel
}

func NewModel(db *sql.DB) Models {
	return Models{
		DB: DBModel{DB: db},
	}
}

type Member struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Debut_date time.Time `json:"debut_date"`
	Birth_date time.Time `json:"birth_date"`
	MemberGen  string    `json:"generation"`
}

type Generation struct {
	ID   int    `json:"id"`
	Name string `json:"generation_name"`
}

type MemberGeneration struct {
	ID         int        `json:"id"`
	Mem_id     int        `json:"mem_id"`
	Gen_id     int        `json:"-"`
	Generation Generation `json:"generation"`
}
