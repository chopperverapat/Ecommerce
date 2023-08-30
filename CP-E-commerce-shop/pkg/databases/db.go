package databases

import (
	"cpshop/config"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func ConnectDB(c config.Idb) *sqlx.DB {
	db, err := sqlx.Connect("pgx", c.Url())
	if err != nil {
		log.Fatalln(err)
	}
	db.DB.SetMaxOpenConns(c.MaxOpenConns())
	return db

}
