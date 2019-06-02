package main

import(
	"github.com/alistairfink/Postgres-Go-Driver-POC/DataModels"
	"github.com/alistairfink/Postgres-Go-Driver-POC/DatabaseConnection"
)

func main() {
	db := DatabaseConnection.Connect()
	defer db.Close()

	max := 0

	println("Get All:")
	tests := DataModels.GetAllTests(db)
	for _, val := range *tests {
		println(val.Key)
		if max < val.Key {
			max = val.Key
		}
	} 

	println("\nGet One:")
	test := DataModels.GetTest(db, 1)
	println(test.Key)

	insertTest := DataModels.Test{
		Key: max + 1,
	}
	println("Insert: ", insertTest.Key)
	DataModels.InsertTest(db, &insertTest)

	println("\nGet All:")
	tests = DataModels.GetAllTests(db)
	for _, val := range *tests {
		println(val.Key)
		if max < val.Key {
			max = val.Key
		}
	} 
}