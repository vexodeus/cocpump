package db

const (
	insertAccountQuery        = "insert into accounts(id, email, password) values ($1,$2,$3)"
	deleteAccountByIDQuery    = "delete from accounts where id=($1)"
	deleteAccountByEmailQuery = "delete from accounts where email=($2)"
)
