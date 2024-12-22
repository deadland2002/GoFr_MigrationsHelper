package Utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"path"
	"regexp"
	"strings"
)

func JoinPaths(paths ...string) string {
	return path.Join(paths...)
}

func UnmarshalNullString(nullString sql.NullString, v any) error {
	if nullString.Valid && nullString.String != "" {
		if err := json.Unmarshal([]byte(nullString.String), v); err != nil {
			return err
		}
	}

	return nil
}

func FormatStructParseError(err error) (string, error) {
	errMsg := err.Error()
	if strings.Contains(errMsg, "cannot unmarshal") {
		re := regexp.MustCompile(`cannot unmarshal .* into Go struct field (\w+\.\w+) of type (\w+)`)
		matches := re.FindStringSubmatch(errMsg)
		if len(matches) == 3 {
			fieldParts := strings.Split(matches[1], ".")
			field := fieldParts[len(fieldParts)-1]
			fieldType := matches[2]
			return fmt.Sprintf("Field %s should be of type %s", field, fieldType), nil
		}
	}
	return "", err
}
