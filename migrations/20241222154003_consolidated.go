package migrations

import "fmt"
import "gofr.dev/pkg/gofr/migration"

func ConsolidateMigration20241222154003() migration.Migrate {
	return migration.Migrate{
		UP: QueryMigration20241222154003,
	}
}

const consolidatedQuery = `
CREATE TABLE USERS (ID UUID NOT NULL PRIMARY KEY, NAME VARCHAR(256) NOT NULL);
ALTER TABLE users ADD COLUMN age int not null;
CREATE TABLE POSTS (id UUID PRIMARY KEY, title VARCHAR(256) NOT NULL, body TEXT NOT NULL, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP);
ALTER TABLE POSTS ADD COLUMN deleted_at TIMESTAMP;
`

func QueryMigration20241222154003(d migration.Datasource) error {
	const query = consolidatedQuery
	_, err := d.SQL.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
