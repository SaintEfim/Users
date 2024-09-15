package psql

const retrieveAllUsers = `SELECT id, name FROM users`
const retrieveOneById = `SELECT id, name FROM users WHERE id = $1`
const createUser = `INSERT INTO users (name) VALUES ($1)`
const deleteUser = `DELETE FROM users WHERE id = $1`
const updateUser = `UPDATE users SET name = $1 WHERE id = $2`
