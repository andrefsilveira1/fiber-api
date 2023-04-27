package repositories

import (
	"api/src/models"
	"database/sql"
	"fmt"
)

type users struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *users {
	return &users{db}
}

func (u users) CreateUser(user *models.User) (uint64, error) {
	statement, erro := u.db.Prepare("insert into users (name, email, password) values (?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(user.Name, user.Email, user.Password)
	if erro != nil {
		return 0, erro
	}

	ID, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(ID), nil
}

func (repositorie users) Search(name string) ([]models.User, error) {
	name = fmt.Sprintf("%%%s%%", name)

	lines, erro := repositorie.db.Query(
		"select id, name, email, createdAt from users where nome LIKE ? or email LIKE ?", name, name,
	)

	if erro != nil {
		return nil, erro
	}
	defer lines.Close()

	var users []models.User

	for lines.Next() {
		var user models.User
		if erro = lines.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.CreatedAt,
		); erro != nil {
			return nil, erro
		}
		users = append(users, user)
	}
	return users, nil
}

func (repositorie users) Delete(Id uint64) error {
	statement, erro := repositorie.db.Prepare("delete from users where id = ?")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(Id); erro != nil {
		return erro
	}

	return nil
}
