package psql

const (
	retrieveAllUsers = `SELECT id, name FROM users`
	retrieveOneById  = `SELECT id, name FROM users WHERE id = $1`
	createUser       = `INSERT INTO users (name) VALUES ($1)`
	deleteUser       = `DELETE FROM users WHERE id = $1`
	updateUser       = `UPDATE users SET name = $1 WHERE id = $2`
)
