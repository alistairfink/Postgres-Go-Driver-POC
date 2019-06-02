package DataModels

import (
	"github.com/go-pg/pg"
)

type Test struct {
	tableName struct{} `sql:"test"`
	Key int `sql:"test,pk"`
}

func GetTest(db *pg.DB, key int) (*Test) {
	Result := Test{
		Key: key,
	}

	err := db.Select(&Result)
	if err != nil {
		panic(err)
	}

	return &Result
}

func GetAllTests(db *pg.DB) (*[]Test) {
	Results := []Test{}
	err := db.Model(&Results).Select()
	if err != nil {
		panic(err)
	}

	return &Results
}

func InsertTest(db *pg.DB, object *Test) (Inserted bool) {
	err := db.Insert(object)
	if err != nil {
		panic(err)
	}

	return true
}

