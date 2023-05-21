package db

const (
	insertAccountQuery        = "insert into accounts(id, email, password) values ($1,$2,$3)"
	deleteAccountByIDQuery    = "delete from accounts where id=($1)"
	deleteAccountByEmailQuery = "delete from accounts where email=($2)"
	getAccountByIDQuery       = "select * from account where id=($1) limit 1"
	getAccountByEmailQuery    = "select * from account where email=($1) limit 1"
)
