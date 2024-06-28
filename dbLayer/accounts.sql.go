// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: accounts.sql

package dbLayer

import (
	"context"
)

const createAccount = `-- name: CreateAccount :exec
INSERT INTO Accounts (
    userName, passwdHash, powerLevel, firstName, lastName, email
) VALUES (
    $1, $2, $3, $4, $5, $6
)
`

type CreateAccountParams struct {
	Username   string
	Passwdhash string
	Powerlevel int32
	Firstname  string
	Lastname   string
	Email      string
}

func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParams) error {
	_, err := q.db.Exec(ctx, createAccount,
		arg.Username,
		arg.Passwdhash,
		arg.Powerlevel,
		arg.Firstname,
		arg.Lastname,
		arg.Email,
	)
	return err
}

const deleteAccount = `-- name: DeleteAccount :exec
UPDATE Accounts
SET valid = False
WHERE userName = $1 AND valid = True
`

func (q *Queries) DeleteAccount(ctx context.Context, username string) error {
	_, err := q.db.Exec(ctx, deleteAccount, username)
	return err
}

const retrieveAccount = `-- name: RetrieveAccount :one
SELECT username, passwdhash, powerlevel, firstname, lastname, email, valid FROM Accounts
WHERE userName = $1 AND valid = True
`

func (q *Queries) RetrieveAccount(ctx context.Context, username string) (Account, error) {
	row := q.db.QueryRow(ctx, retrieveAccount, username)
	var i Account
	err := row.Scan(
		&i.Username,
		&i.Passwdhash,
		&i.Powerlevel,
		&i.Firstname,
		&i.Lastname,
		&i.Email,
		&i.Valid,
	)
	return i, err
}

const updateAccountPowerLevel = `-- name: UpdateAccountPowerLevel :exec
UPDATE Accounts
SET powerLevel = $2
WHERE userName = $1 AND valid = True
`

type UpdateAccountPowerLevelParams struct {
	Username   string
	Powerlevel int32
}

func (q *Queries) UpdateAccountPowerLevel(ctx context.Context, arg UpdateAccountPowerLevelParams) error {
	_, err := q.db.Exec(ctx, updateAccountPowerLevel, arg.Username, arg.Powerlevel)
	return err
}

const updatePasswdHash = `-- name: UpdatePasswdHash :exec
UPDATE Accounts
SET passwdHash = $2
WHERE userName = $1 AND valid = True
`

type UpdatePasswdHashParams struct {
	Username   string
	Passwdhash string
}

func (q *Queries) UpdatePasswdHash(ctx context.Context, arg UpdatePasswdHashParams) error {
	_, err := q.db.Exec(ctx, updatePasswdHash, arg.Username, arg.Passwdhash)
	return err
}
