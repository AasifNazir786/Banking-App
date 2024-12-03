package storage

import (
	"Go-GitHub-Projects/Banking-App/models"
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (s *UserRepository) SaveUser(userName, password string) error {

	hashedCode, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	query := `INSERT INTO users 
					(user_name, password)
							VALUES($1, $2)`

	_, err = s.db.Exec(query, userName, string(hashedCode))

	if err != nil {
		return fmt.Errorf("error saving user: %v", err)
	}

	return nil
}

func (s *UserRepository) GetUserByUserName(userName string) (models.AccountHolder, error) {

	query := `SELECT * FROM users
					WHERE user_name = $1`

	var user models.AccountHolder

	err := s.db.QueryRow(query, userName).Scan(&user.UserName, &user.Password)

	if err != nil {

		return models.AccountHolder{}, fmt.Errorf("error parsing user from DB: %v", err)
	}

	return user, nil
}
