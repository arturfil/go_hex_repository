package user

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/arturfil/yt_ecomm/types"
)

type Store struct {
    db *sql.DB
}

func NewStore(db *sql.DB) *Store {
    return &Store{db: db}
}

func (s *Store) GetUserByEmail(email string) (*types.User, error) {
    rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
    if err != nil {
        return nil, err
    }

    var u *types.User = &types.User{}
    for rows.Next() {
        u, err = scanRowIntoUser(rows)
        if err != nil {
            return nil, err
        }
    }

    if u.ID == 0 {
        return nil, fmt.Errorf("user not found")
    }

    return u, nil
}


func (s *Store) GetUserByID(id int) (*types.User, error) {
    rows, err := s.db.Query("SELECT * FROM users WHERE id = ?", id)
    if err != nil {
        return nil, err
    }

    var user *types.User
    for rows.Next() {
        user, err = scanRowIntoUser(rows)
        if err != nil {
            return nil, err
        }
    }

    if user.ID == 0 {
        return nil, fmt.Errorf("user not found")
    }

    return user, nil
}

func (s *Store) CreateUser(user types.User) error {
    query := `
        INSERT INTO users (first_name, last_name, email, password)
        VALUES (?, ?, ?, ?)
    `
    _, err := s.db.Exec(
        query, 
        user.FirstName,
        user.LastName,
        user.Email,
        user.Password,
    ) 
    if err != nil {
        return err
    }

    return nil
}

func scanRowIntoUser(rows *sql.Rows) (*types.User, error) {
    user := &types.User{}

    err := rows.Scan(
        &user.ID,
        &user.FirstName,
        &user.LastName,
        &user.Email,
        &user.Password,
        time.Now(),
        time.Now(),
    )
    if err != nil {
        return nil, err
    }

    return user, nil
}
