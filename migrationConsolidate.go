package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	migrationsDir := "./migrations"
	migrationTimestamp := time.Now().Format("20060102150405")
	consolidatedFileName := fmt.Sprintf("%s_consolidated.go", migrationTimestamp)

	files, err := ioutil.ReadDir(migrationsDir)
	if err != nil {
		log.Fatal(err)
	}

	var queries []string

	for _, file := range files {
		if file.IsDir() || file.Name() == "all.go" {
			continue
		}

		filePath := filepath.Join(migrationsDir, file.Name())
		content, err := ioutil.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		query := extractQuery(string(content))
		if query != "" {
			queries = append(queries, query)
		}
	}

	functionName, consolidatedContent := generateConsolidatedFileContent(queries, migrationTimestamp)
	consolidatedFilePath := filepath.Join(migrationsDir, consolidatedFileName)

	err = ioutil.WriteFile(consolidatedFilePath, []byte(consolidatedContent), 0644)
	if err != nil {
		log.Fatal(err)
	}

	err = updateAllMigrationsFile(functionName, migrationTimestamp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Consolidated migration file created successfully:", consolidatedFileName)
}

func extractQuery(content string) string {
	re := regexp.MustCompile(`const query = "(.*?)"`)
	matches := re.FindStringSubmatch(content)
	if len(matches) > 1 {
		return matches[1]
	}
	return ""
}

func generateConsolidatedFileContent(queries []string, migrationTimestamp string) (string, string) {

	var functionName = "ConsolidateMigration" + migrationTimestamp
	var functionQueryName = "QueryMigration" + migrationTimestamp

	consolidatedQuery := "`\n"

	for _, query := range queries {
		consolidatedQuery += query + "\n"
	}

	consolidatedQuery += "`"

	return functionName, fmt.Sprintf(`
package migrations

import "fmt"
import "gofr.dev/pkg/gofr/migration"

func %s() migration.Migrate {
	return migration.Migrate{
		UP: %s,
	}
}

const consolidatedQuery = %s

func %s(d migration.Datasource) error {
	const query = consolidatedQuery
	_, err := d.SQL.Exec(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
`, functionName, functionQueryName, consolidatedQuery, functionQueryName)
}

var fileTemplateAll = `
package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
	}
}
`

func updateAllMigrationsFile(functionName string, migrationTimeStamp string) error {
	fmt.Println("Adding migration entry to all.go file...")

	var allMigrationsFilePath = "./migrations/all.go"

	// Check if the file exists
	if _, err := os.Stat(allMigrationsFilePath); os.IsNotExist(err) {
		// Create the file if it does not exist
		err := os.WriteFile(allMigrationsFilePath, []byte(fileTemplateAll), 0644)
		if err != nil {
			return err
		}
	}

	// Read all.go file
	content, err := os.ReadFile(allMigrationsFilePath)
	if err != nil {
		return err
	}

	// Prepare the new migration entry
	newMigrationEntry := fmt.Sprintf("\t\t%s: %s(),\n", migrationTimeStamp, functionName)

	// Find the position to insert the new migration entry
	lastIndex := strings.LastIndex(string(content), "}")
	if lastIndex == -1 {
		return fmt.Errorf("Failed to find closing bracket in all.go file")
	}

	secondLastIndex := strings.LastIndex(string(content[:lastIndex]), "}")
	if secondLastIndex == -1 {
		return fmt.Errorf("Failed to find second-to-last closing bracket in all.go file")
	}

	// Insert the new migration entry
	newContent := strings.Trim(string(content[:secondLastIndex]), "\t") + newMigrationEntry + "\t" + string(content[secondLastIndex:])

	// Write the updated content back to the file
	err = os.WriteFile(allMigrationsFilePath, []byte(newContent), 0644)
	if err != nil {
		return err
	}

	fmt.Println("Migration entry added successfully")

	return nil
}