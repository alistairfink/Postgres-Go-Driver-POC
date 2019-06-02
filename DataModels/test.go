package DataModels

import (
	"github.com/go-pg/pg"
)

type Test struct {
	key int
}

func GetTest(db *pg.DB, key int) (Result *Test) {
	_, err := db.QueryOne(Result, `SELECT * FROM test WHERE test = ?`, key)
	if err != nil {
		panic(err)
	}

	return Result
}

func GetALLTest(db *pg.DB) (Results *[]Test) {
	_, err := db.Query(Result, `SELECT * FROM test`)
	if err != nil {
		panic(err)
	}

	return Results
}

func InsertTest(db *pg.DB, object *Test) (Inserted bool) {
	err := db.Insert(object)
	if err != nil {
		panic(err)
	}

	return true
}

