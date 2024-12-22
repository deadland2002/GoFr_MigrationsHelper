package migrations

import (
	"fmt"
	"gofr.dev/pkg/gofr/migration"
)

func AddColumnUseridPosts() migration.Migrate {
	return migration.Migrate{
		UP: addColumnUseridPostsQueryFunction,
	}
}

func addColumnUseridPostsQueryFunction(d migration.Datasource) error {
	const query = "ALTER TABLE POSTS ADD COLUMN owner_id UUID REFERENCES USERS(ID);"
	_, err := d.SQL.Exec(query)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	return nil
}
