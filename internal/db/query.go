package db

const (
	insertAccountQuery        = "insert into accounts(id, email, password) values ($1,$2,$3)"
	deleteAccountByIDQuery    = "delete from accounts where id=($1)"
	deleteAccountByEmailQuery = "delete from accounts where email=($2)"
	getAccountsByIDQuery      = "select * from account where id=($1)"
	getAccountsByEmailQuery   = "select * from account where email=($1)"
	updateAccountByIDQuery    = "update accounts set email=($2), password=($3) where id=($1)"
	updateAccountByEmailQuery = "update accounts set id=($1), password=($3), where email=($2)"
)
