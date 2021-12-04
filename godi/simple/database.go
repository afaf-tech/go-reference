package simple

type Database struct {
	Name string
}

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	database := &Database{Name: "PostgreSQL"}

	return (*DatabasePostgreSQL)(database)
}
func NewDatabaseMongoDB() *DatabaseMongoDB {
	database := &Database{Name: "MongoDB"}

	return (*DatabaseMongoDB)(database)
}

// alias
type DatabasePostgreSQL Database
type DatabaseMongoDB Database

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMongoDB    *DatabaseMongoDB
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mongoDB *DatabaseMongoDB) *DatabaseRepository {
	return &DatabaseRepository{DatabasePostgreSQL: postgreSQL, DatabaseMongoDB: mongoDB}
}
