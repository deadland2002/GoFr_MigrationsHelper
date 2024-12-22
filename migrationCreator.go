package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"
)

func main() {
	var migrationName string
	var migrationTimeStamp string

	// Get current timestamp
	migrationTimeStamp = time.Now().Format("20060102150405")

	// Prompt user for migration name
	fmt.Print("Enter Migration Name: ")
	fmt.Scanf("%s", &migrationName)

	functionName, err := createMigrationFile(migrationTimeStamp, migrationName)
	if err != nil {
		log.Fatal(err)
	}

	err = addFunctionToMigrations(functionName, migrationTimeStamp)
	if err != nil {
		log.Fatal(err)
	}
}

func formatNameForFunction(name string) string {
	words := strings.Split(name, "_")
	var camelCaseName string

	camelCaseName = ""

	for i := 0; i < len(words); i++ {
		words[i] = string(unicode.ToUpper(rune(words[i][0]))) + strings.ToLower(words[i][1:])
		camelCaseName += words[i]
	}

	return camelCaseName
}

func formatNameForQueryFunction(name string) string {
	words := strings.Split(name, "_")
	camelCaseName := words[0]

	for i := 1; i < len(words); i++ {
		words[i] = string(unicode.ToUpper(rune(words[i][0]))) + strings.ToLower(words[i][1:])
		camelCaseName += words[i]
	}

	return camelCaseName
}

var fileTemplate = `
package migrations

import "gofr.dev/pkg/gofr/migration"

func %s() migration.Migrate {
	return migration.Migrate{
		UP: %s,
	}
}

func %s(d migration.Datasource) error {
	const query = ""
	_, err := d.SQL.Exec(query)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	return nil
}
`

func createMigrationFile(migrationTimeStamp string, migrationName string) (string, error) {
	fmt.Println("Creating migration file...")
	var fileName = fmt.Sprintf("%s_%s.go", migrationTimeStamp, migrationName)
	var functionName = formatNameForFunction(migrationName)
	var queryFunctionName = string(unicode.ToLower(rune(functionName[0]))) + functionName[1:] + "QueryFunction"

	formattedFileContent := fmt.Sprintf(fileTemplate, functionName, queryFunctionName, queryFunctionName)
	formattedFilePath := "./migrations/" + fileName

	file, err := os.OpenFile(formattedFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}

	defer file.Close()

	_, err = file.Write([]byte(formattedFileContent))
	if err != nil {
		return "", err
	}

	fmt.Println("Migration file created successfully")

	file.Close()

	return functionName, nil
}

var allFileTemplate = `
package migrations

import (
	"gofr.dev/pkg/gofr/migration"
)

func All() map[int64]migration.Migrate {
	return map[int64]migration.Migrate{
	}
}
`

func addFunctionToMigrations(functionName string, migrationTimeStamp string) error {
	fmt.Println("Adding migration entry to all.go file...")

	var allMigrationsFilePath = "./migrations/all.go"

	// Check if the file exists
	if _, err := os.Stat(allMigrationsFilePath); os.IsNotExist(err) {
		// Create the file if it does not exist
		err := os.WriteFile(allMigrationsFilePath, []byte(allFileTemplate), 0644)
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
