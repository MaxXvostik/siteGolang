package model

import (
	"example.com/siteGolang/app/server"
)

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func GetAllUsers() (users []User, err error) {
	query := `SELECT * FROM users`
	rows, err := server.Db.Queryx(query)
	if err != nil {
		return users, err
	}
	defer rows.Close()
	user := User{}
	for rows.Next() {
		err = rows.StructScan(&user)
		if err != nil {
			return users, err
		}
		users = append(users, user)
	}
	return users, err
}

func NewUser(name, surname string) *User {
	return &User{Name: name, Surname: surname}
}

func GetUserById(userId string) (u User, err error) {
	query := `SELECT * FROM users WHERE id = ?`
	err = server.Db.Get(&u, query, userId)
	return
}

func (u *User) Add() (err error) {
	query := `INSERT INTO users (name, surname) VALUES (?, ?)`
	_, err = server.Db.Exec(query, u.Name, u.Surname)
	return
}

func (u *User) Delete() (err error) {
	query := `DELETE FROM users WHERE id = ?`
	_, err = server.Db.Exec(query, u.Id)
	return
}
func (u *User) Update() (err error) {
	query := `UPDATE users SET name = ?, surname = ? WHERE id = ?`
	_, err = server.Db.Exec(query, u.Name, u.Surname, u.Id)
	return
}
