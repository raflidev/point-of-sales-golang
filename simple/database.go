package simple

type Database struct {
	Name string
}

type DatabasePostgreSQL Database
type DatabaseMySQL Database

func NewDatabasePostgreSQL() *DatabasePostgreSQL {
	return (*DatabasePostgreSQL)(&Database{Name: "PostgreSQL"})
}

func NewDatabaseMySQL() *DatabaseMySQL {
	return (*DatabaseMySQL)(&Database{Name: "MySQL"})
}

type DatabaseRepository struct {
	DatabasePostgreSQL *DatabasePostgreSQL
	DatabaseMySQL      *DatabaseMySQL
}

func NewDatabaseRepository(postgreSQL *DatabasePostgreSQL, mySQL *DatabaseMySQL) *DatabaseRepository {
	return &DatabaseRepository{
		DatabasePostgreSQL: postgreSQL,
		DatabaseMySQL:      mySQL,
	}
}
