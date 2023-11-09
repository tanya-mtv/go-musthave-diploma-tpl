package repository

import (
	"github.com/jmoiron/sqlx"
)

const (
	usersTable = "users"
)

func NewPostgresDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, login varchar(50), password_hash varchar(255), salt varchar(255) not null,
            UNIQUE (login));`)

	if err != nil {
		return db, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS orders (id serial PRIMARY KEY, number VARCHAR(100) not null unique, status varchar(50),
            user_id int references users (id) on delete cascade not null,
                upload_date timestamp   DEFAULT now(), update_date timestamp  without time zone ) ;`)

	if err != nil {
		return db, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS balance (id serial PRIMARY KEY, number VARCHAR(100) not null, sum double precision not null DEFAULT 0,
                user_id int references users (id) on delete cascade not null,
                processed timestamp   DEFAULT now()) ;`)

	if err != nil {
		return db, err
	}

	return db, nil
}
