package main

import "fmt"

type PasswordEntry struct{
	Site string
	Login string
	Password string
}

type PasswordManager struct{
	Passwords map[string]PasswordEntry
}


func main() {

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

if uu==1{
	addPasword()
}

}

func(p *PasswordManager) addPasword(ff PasswordEntry){
	var entry PasswordEntry
	
	fmt.Println("добавить пароль")
	var t string
	fmt.Scan(t)
	entry.Password=t
	
}



