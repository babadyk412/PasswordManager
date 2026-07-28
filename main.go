package main

import (
	"fmt"
)

type PasswordEntry struct {
	Site     string
	Login    string
	Password string
}

type PasswordManager struct {
	Passwords map[string]PasswordEntry
}

func main() {

	pas := PasswordManager{
		Passwords: make(map[string]PasswordEntry),
	}
	for {
		fmt.Println("======================")
		fmt.Println("Password Manager")
		fmt.Println("======================")

		fmt.Println("1 добавить пароль")
		fmt.Println("2 найти пароль")
		fmt.Println("3 показать все сайты")
		fmt.Println("4 удалить пароль")
		fmt.Println("5 сгенерировать пароль")
		fmt.Println("6 выход")

		fmt.Println("======================")
		fmt.Println("что делать")
		var uu int
		fmt.Scan(&uu)

		if uu == 1 {
			pas.addPasword()
		}
		if uu == 2 {
			pas.findPasword()
		}
		if uu == 3 {
			pas.allSpisok()
		}
		if uu == 4 {
			pas.deletePasword()
		}
		if uu==6{
			break
		}
	}
}

func (p *PasswordManager) addPasword() {
	var entry PasswordEntry

	fmt.Println("добавить сайт")
	var site string
	fmt.Scan(&site)
	entry.Site = site

	fmt.Println("добавить логин")
	var login string
	fmt.Scan(&login)
	entry.Login = login

	fmt.Println("добавить пароль")
	var pasword string
	fmt.Scan(&pasword)
	entry.Password = pasword
	p.Passwords[site] = entry

}
func (p *PasswordManager) findPasword() {
	fmt.Println("какой сайт?")
	var y string
	fmt.Scan(&y)
	fmt.Println(p.Passwords[y].Password)
}

func (p *PasswordManager) allSpisok() {
	for _, i := range p.Passwords {
		fmt.Println("сайт", i.Site)
		fmt.Println("логин", i.Login)
		fmt.Println("пароль", i.Password)
		fmt.Println("======================")
	}
}

func (p *PasswordManager) deletePasword() {
	fmt.Println("какой сайт  удалить")
	var a string
	fmt.Scan(&a)
	delete(p.Passwords, a)

}
