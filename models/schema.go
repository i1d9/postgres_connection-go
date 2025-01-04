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
	ID         int64
	UserID     int64
	IPAddress  string
	UserAgent  string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ExpiresAt  int64
}

type User struct {
	ID                   int        `json:"id"`
	Email                *string     `json:"email"`
	Password             *string     `json:"password"`
	Verified             *bool       `json:"verified"`
	Active               *bool       `json:"active"`
	FirstName            *string     `json:"first_name"`
	LastName             *string     `json:"last_name"`
	Surname              *string     `json:"surname"`
	MobileNumber         *string     `json:"mobile_number"`
	MobileVerified       *bool       `json:"mobile_verified"`
	Gender               *string     `json:"gender"`
	DeletedAt            *time.Time `json:"deleted_at,omitempty"`
	TermsAccepted        *bool       `json:"terms_accepted"`
	CreatedAt            *time.Time  `json:"created_at"`
	UpdatedAt            *time.Time  `json:"updated_at"`
	LastLoginAt          *time.Time `json:"last_login_at,omitempty"`
	RequirePasswordReset *bool       `json:"require_password_reset"`
	PasswordExpiry       *time.Time  `json:"password_expiry"`
	DeviceID             *string     `json:"device_id"`
	
}

type Token struct {
	Id         int64
	Token_type string // access or refresh or authorization code
	User_id    int64
	Client_id  int64
	Token      string
	Expires_at int64
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
			&user.ID,  &user.Email, &user.Password,
			&user.Verified, &user.Active, &user.FirstName, &user.LastName, &user.Surname,
			&user.MobileNumber, &user.MobileVerified, &user.Gender,
			&user.DeletedAt, &user.TermsAccepted,  &user.CreatedAt,
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

// func InitiateDatabase(dbpool pgxpool.Pool, ctx context.Context) {

// 	_, err := dbpool.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS users (id serial PRIMARY KEY, first_name text, last_name text, email text, password text, active boolean)")
// 	if err != nil {
// 	}

// }
// func createUser(dbpool *pgxpool.Pool, ctx context.Context, user User) (User, error) {

// 	dbpool.Query(ctx.Background(), "CREATE TABLE IF NOT EXISTS greeting (greeting text)")

// 	_, err := dbpool.Exec(ctx.Background(), "INSERT INTO greeting (greeting) VALUES ('Hello, world!')")
// 	if err != nil {
// 		fmt.Fprintf(os.Stderr, "Insert failed: %v\n", err)
// 		os.Exit(1)
// 	}

// 	return user, nil

// }

func createClient(client Client) (Client, error) {
	return client, nil
}

func createToken(token Token) (Token, error) {
	return token, nil
}

func createSession(session Session) (Session, error) {
	return session, nil
}

func getUserByEmail(email string) (User, error) {
	return User{}, nil
}

func getUserById(id int64) (User, error) {
	return User{}, nil
}

func getClientById(id int64) (Client, error) {
	return Client{}, nil
}

func getClientByClientId(client_id string) (Client, error) {

	return Client{}, nil
}

func getTokenByToken(token string, token_type string) (Token, error) {
	return Token{}, nil
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