package apiservice

import (
	"database/sql"
	"github.com/Rumpelstiltski1/restapi/store/sqlstore"
	"net/http"
)

func Start(config *Config) error {
	db, err := NewDB(config.DatabaseURL)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)
	srv := newServer(store)
	return http.ListenAndServe(config.BindAddr, srv)
}

func NewDB(databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
