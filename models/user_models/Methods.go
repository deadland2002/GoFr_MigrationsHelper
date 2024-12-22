package user_models

import (
	"database/sql"
	"github.com/text-gofr/Utils"
	"gofr.dev/pkg/gofr"
	gofrHttp "gofr.dev/pkg/gofr/http"
	"strings"
)

func (u *UserTemplate) CreateUser(c *gofr.Context) (interface{}, error) {
	query := "INSERT INTO users (id, name , age , gender) VALUES (?, ? , ? , ?)"
	data, err := c.SQL.Exec(query, u.ID, u.Name, u.Age, u.Gender)

	return data, err
}

func (u *UserTemplate) EditUser(c *gofr.Context) (interface{}, error) {
	var fields []string
	var args []interface{}

	if u.Name != "" {
		fields = append(fields, "name = ?")
		args = append(args, u.Name)
	}
	if u.Age > 0 {
		fields = append(fields, "age = ?")
		args = append(args, u.Age)
	}
	if u.Gender == "MALE" || u.Gender == "FEMALE" {
		fields = append(fields, "gender = ?")
		args = append(args, u.Gender)
	}

	if len(fields) == 0 {
		return nil, gofrHttp.ErrorMissingParam{Params: []string{"name", "age", "gender"}}
	}

	args = append(args, u.ID)
	query := "UPDATE users SET " + strings.Join(fields, ", ") + " WHERE id = ?"
	_, err := c.SQL.Exec(query, args...)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (u *UserTemplate) FindUser(c *gofr.Context) (interface{}, error) {
	query := "SELECT id, name, age , gender , created_at, deleted_at FROM users WHERE id = ?"
	user := User{}
	var (
		deletedAt sql.NullString
	)
	err := c.SQL.QueryRow(query, u.ID).Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.CreatedAt, &deletedAt)
	if err != nil {
		return nil, err
	}

	err = Utils.UnmarshalNullString(deletedAt, &user.DeletedAt)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (t *UserTemplate) Validate() error {
	params := make([]string, 0)
	if t.Name == "" {
		params = append(params, "name")
	}

	if t.Age <= 0 {
		params = append(params, "age")
	}

	if t.Gender != "FEMALE" && t.Gender != "MALE" {
		params = append(params, "gender")
	}

	println(t.Age, t.Name, t.Gender)

	if len(params) > 0 {
		return gofrHttp.ErrorMissingParam{Params: params}
	}

	return nil
}

func (t *UserTemplate) ValidateOptional() error {
	params := make([]string, 0)

	if t.Gender != "FEMALE" && t.Gender != "MALE" && t.Gender != "" {
		params = append(params, "gender")
	}

	if len(params) > 0 {
		return gofrHttp.ErrorInvalidParam{Params: params}
	}

	return nil
}
