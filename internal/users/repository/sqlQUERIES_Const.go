package repository

// SQL Querise
const (
	// SignUPQuery is
	SignUPQuery = `INSERT INTO Users (
					username, email, password) VALUES (
					$1, $2, $3) RETURNING user_id`
)
