package storage

import (
	"database/sql"
	"fmt"
	// "log"

	_ "github.com/lib/pq"
	"github.com/soufianiso/boxd/types"
	pb "user/api/proto/user"
)

type Storage struct {
	db *sql.DB
}



type Store interface {
	CreateUser(*pb.RegisterUserRequest) error
	GetUserByEmail(string) (*types.User, error)
}

func NewUserStorage(db *sql.DB) *Storage{
	return &Storage{db: db}
}


func(s *Storage) CreateUser(user *pb.RegisterUserRequest) error{
	_, err := s.db.Exec("INSERT INTO users (first_name, last_name, email, password) VALUES ($1,$2,$3,$4)", 
			user.GetFirstName(), 
			user.GetLastName(), 
			user.GetEmail(), 
			user.GetPassword())
	if err != nil{
		return err
	}

	return nil
}

func (s *Storage) GetUserByEmail(email string) (*types.User, error) {
	rows, err := s.db.Query("SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}

	u := new(types.User)
	for rows.Next() {
		u, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	if u.ID == 0 {
		return nil, fmt.Errorf("email not found")
	}

	return u, nil
}


func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}
