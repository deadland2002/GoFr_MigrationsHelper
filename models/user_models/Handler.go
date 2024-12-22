package user_models

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/text-gofr/Utils"
	"gofr.dev/pkg/gofr"
	gofrHttp "gofr.dev/pkg/gofr/http"
)

func GetUserList(c *gofr.Context) (interface{}, error) {
	rows, err := c.SQL.Query("SELECT * FROM users")
	if err != nil {
		c.Error(err)
		return nil, err
	}
	defer rows.Close()

	var data = make([]User, 0)

	for rows.Next() {
		var row User
		var (
			deletedAt sql.NullString
		)
		if err := rows.Scan(&row.ID, &row.Name, &row.Age, &row.Gender, &row.CreatedAt, &deletedAt); err != nil {
			c.Error(err)
			return nil, err
		}

		err = Utils.UnmarshalNullString(deletedAt, &row.DeletedAt)
		if err != nil {
			c.Error(err)
			return nil, err
		}

		data = append(data, row)
	}

	return data, nil
}

func AddNewUser(c *gofr.Context) (interface{}, error) {
	var template UserTemplate

	err := c.Bind(&template)
	if err != nil {
		return Utils.FormatStructParseError(err)
	}

	err = template.Validate()
	if err != nil {
		return nil, err
	}

	template.ID = uuid.New()

	_, err = template.CreateUser(c)

	if err != nil {
		return nil, err
	}

	data, err := template.FindUser(c)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func EditUser(c *gofr.Context) (interface{}, error) {
	var template UserTemplate

	err := c.Bind(&template)
	if err != nil {
		return Utils.FormatStructParseError(err)
	}

	if template.ID == uuid.Nil {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"id"}}
	}

	err = template.ValidateOptional()
	if err != nil {
		return nil, err
	}

	_, err = template.EditUser(c)

	data, err := template.FindUser(c)
	if err != nil {
		return nil, err
	}

	return data, nil
}
