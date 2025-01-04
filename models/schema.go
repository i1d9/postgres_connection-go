package models

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Client struct {
	ID            int64
	Name          string
	Redirect_uri  string
	Client_id     string
	Client_secret string
}

type Session struct {
	ID        int64
	UserID    int64
	IPAddress string
	UserAgent string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt int64
}

type User struct {
	ID                   int        `json:"id"`
	Email                *string    `json:"email"`
	Password             *string    `json:"password"`
	Verified             *bool      `json:"verified"`
	Active               *bool      `json:"active"`
	FirstName            *string    `json:"first_name"`
	LastName             *string    `json:"last_name"`
	Surname              *string    `json:"surname"`
	MobileNumber         *string    `json:"mobile_number"`
	MobileVerified       *bool      `json:"mobile_verified"`
	Gender               *string    `json:"gender"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`
	TermsAccepted        *bool      `json:"terms_accepted"`
	CreatedAt            *time.Time `json:"created_at"`
	UpdatedAt            *time.Time `json:"updated_at"`
	LastLoginAt          *time.Time `json:"last_login_at,omitempty"`
	RequirePasswordReset *bool      `json:"require_password_reset"`
	PasswordExpiry       *time.Time `json:"password_expiry"`
	DeviceID             *string    `json:"device_id"`
}

type Token struct {
	Id         int64
	Token_type string // access or refresh or authorization code
	User_id    int64
	Client_id  int64
	Token      string
	Expires_at int64
}


func DeleteUser(pool *pgxpool.Pool, id int) error {
	ctx := context.Background()

	query := `
		DELETE FROM users
		WHERE id = $1
	`

	_, err := pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("query execution failed: %w", err)
	}

	return nil
}
func SearchUsers(pool *pgxpool.Pool, term string) ([]User, error) {
	ctx := context.Background()

	query := `
		SELECT
			id, email, password, verified, active,
			first_name, last_name, surname, mobile_number, mobile_verified,
			gender, deleted_at, terms_accepted, created_at,
			updated_at, last_login_at, require_password_reset, password_expiry,
			device_id
		FROM users
		WHERE
			first_name ILIKE '%' || $1 || '%' OR
			last_name ILIKE '%' || $1 || '%' OR
			surname ILIKE '%' || $1 || '%' OR
			email ILIKE '%' || $1 || '%' OR
			mobile_number ILIKE '%' || $1 || '%'
	`

	rows, err := pool.Query(ctx, query, term)

	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.Email, &user.Password,
			&user.Verified, &user.Active, &user.FirstName, &user.LastName, &user.Surname,
			&user.MobileNumber, &user.MobileVerified, &user.Gender,
			&user.DeletedAt, &user.TermsAccepted, &user.CreatedAt,
			&user.UpdatedAt, &user.LastLoginAt, &user.RequirePasswordReset,
			&user.PasswordExpiry, &user.DeviceID)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration error: %w", rows.Err())
	}

	return users, nil
}

func GetAllUsers(pool *pgxpool.Pool) ([]User, error) {
	ctx := context.Background()

	query := `
		SELECT
			id, email, password, verified, active,
			first_name, last_name, surname, mobile_number, mobile_verified,
			gender, deleted_at, terms_accepted, created_at,
			updated_at, last_login_at, require_password_reset, password_expiry,
			device_id
		FROM users
	`

	rows, err := pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("query execution failed: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID, &user.Email, &user.Password,
			&user.Verified, &user.Active, &user.FirstName, &user.LastName, &user.Surname,
			&user.MobileNumber, &user.MobileVerified, &user.Gender,
			&user.DeletedAt, &user.TermsAccepted, &user.CreatedAt,
			&user.UpdatedAt, &user.LastLoginAt, &user.RequirePasswordReset,
			&user.PasswordExpiry, &user.DeviceID)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		users = append(users, user)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows iteration error: %w", rows.Err())
	}

	return users, nil
}

func GetClients() []Client {

	var clients []Client

	return clients
}


func CreateUser(pool *pgxpool.Pool, user User) error {
	ctx := context.Background()

	query := `
		INSERT INTO users (
			first_name, last_name, surname, email, mobile_number, created_at, updated_at
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
	`

	_, err := pool.Exec(ctx, query, user.FirstName, user.LastName, user.Surname, user.Email, user.MobileNumber, time.Now(), time.Now())
	if err != nil {
		return fmt.Errorf("query execution failed: %w", err)
	}

	return nil

}




// nullableString handles *string and returns a readable string
func NullableString(s *string) string {
	if s == nil {
		return "<nil>"
	}
	return *s
}

// nullableBool handles *bool and returns a readable string
func NullableBool(b *bool) string {
	if b == nil {
		return "<nil>"
	}
	return fmt.Sprintf("%t", *b)
}

// nullableTime handles *time.Time and returns a readable string
func NullableTime(t *time.Time) string {
	if t == nil {
		return "<nil>"
	}
	return t.Format("2006-01-02 15:04:05")
}
