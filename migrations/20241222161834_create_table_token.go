package migrations

import (
	"fmt"
	"gofr.dev/pkg/gofr/migration"
)

func CreateTableToken() migration.Migrate {
	return migration.Migrate{
		UP: createTableTokenQueryFunction,
	}
}

func createTableTokenQueryFunction(d migration.Datasource) error {
	const query = "CREATE TABLE TOKENS (ID SERIAL PRIMARY KEY, TOKEN TEXT NOT NULL, USER_ID INT NOT NULL REFERENCES USERS(ID), CREATED_AT TIMESTAMP NOT NULL, EXPIRES_AT TIMESTAMP NOT NULL);"
	_, err := d.SQL.Exec(query)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	return nil
}
