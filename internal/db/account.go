package db

import "database/sql"

type Account struct {
	ID       int
	Email    string
	Password string
}

func insertAccount(account Account, db *sql.DB) error {
	_, err := db.Exec(insertAccountQuery, account.ID, account.Email, account.Password)
	if err != nil {
		return err
	}
	return nil
}
func deleteAccountByEmail(email string, db *sql.DB) error {
	_, err := db.Exec(deleteAccountByEmailQuery, email)
	if err != nil {
		return err
	}
	return nil
}
func deleteAccountByID(id int, db *sql.DB) error {
	_, err := db.Exec(deleteAccountByIDQuery, id)
	if err != nil {
		return err
	}
	return nil
}
func getAccountsByID(id int, db *sql.DB) ([]*Account, error) {
	accounts := []*Account{}
	rows, err := db.Query(getAccountsByIDQuery, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		acc := &Account{}
		rows.Scan(acc.ID, acc.Email, acc.Password)
		accounts = append(accounts, acc)
	}
	return accounts, nil
}
func getAccountsByEmail(email string, db *sql.DB) ([]*Account, error) {
	accounts := []*Account{}
	rows, err := db.Query(getAccountsByEmailQuery, email)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		acc := &Account{}
		rows.Scan(acc.ID, acc.Email, acc.Password)
		accounts = append(accounts, acc)
	}
	return accounts, nil
}
