package DatabaseConnection

import(
	"github.com/alistairfink/Postgres-Go-Driver-POC/Config"
	"github.com/go-pg/pg"
)

func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		User:		Config.Database_Username,
		Password:	Config.Database_Password,
		Database:	Config.Database_Name,
	})

	return db
}

func Close(db *pg.DB) {
	db.Close()
}