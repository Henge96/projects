package repo

import "database/sql"

type Nachinka struct {
	db *sql.DB
}

func NewNachinka(db *sql.DB) Nachinka {
	return Nachinka{db: db}
}
