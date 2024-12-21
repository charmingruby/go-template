package mysql_repository

import (
	"database/sql"
	"go-template/internal/identity/core/model"
	"go-template/internal/shared/custom_err"

	"github.com/jmoiron/sqlx"
)

const (
	createUser    = "create user"
	findUserByID  = "find user by id"
	findManyUsers = "find many users"
	updateUser    = "update user"
	deleteUser    = "delete user"
)

func userQueries() map[string]string {
	return map[string]string{
		createUser:    `INSERT INTO user (id, name, email) VALUES (?, ?, ?)`,
		findUserByID:  `SELECT id, name, email FROM user WHERE id = ?`,
		findManyUsers: `SELECT id, name, email FROM user`,
		updateUser:    `UPDATE user SET name = ?, email = ? WHERE id = ?`,
		deleteUser:    `DELETE FROM user WHERE id = ?`,
	}
}

func NewUserRepository(db *sql.DB) (*UserRepository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	dbSqlx := sqlx.NewDb(db, "mysql")

	for queryName, statement := range userQueries() {
		stmt, err := dbSqlx.Preparex(statement)
		if err != nil {
			return nil,
				custom_err.NewPreparationErr(queryName, "user", err)
		}

		stmts[queryName] = stmt
	}

	return &UserRepository{
		db:    dbSqlx,
		stmts: stmts,
	}, nil
}

type UserRepository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *UserRepository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			custom_err.NewStatementNotPreparedErr(queryName, "user")
	}

	return stmt, nil
}

func (r *UserRepository) Create(u model.User) error {
	stmt, err := r.statement(createUser)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(u.ID, u.Name, u.Email); err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	stmt, err := r.statement(findManyUsers)
	if err != nil {
		return nil, err
	}

	var users []model.User

	if err := stmt.Select(&users); err != nil {
		return nil, err
	}

	return users, nil
}

func (r *UserRepository) FindOne(id string) (*model.User, error) {
	stmt, err := r.statement(findUserByID)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := stmt.Get(&user, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Update(name, email, id string) (*model.User, error) {
	stmt, err := r.statement(updateUser)
	if err != nil {
		return nil, err
	}

	var user model.User
	if err := stmt.Get(&user, name, email, id); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(id string) error {
	stmt, err := r.statement(deleteUser)
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(id); err != nil {
		return err
	}

	return nil
}
