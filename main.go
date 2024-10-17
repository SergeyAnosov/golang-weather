package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Новый проект")
	name := flag.String("name", "Sergey", "Имя пользователя")
	age := flag.Int("age", 36, "Возраст пользователя")
	isAdmin := flag.Bool("isAdmin", true, "Администратора")

	flag.Parse()
	fmt.Println("Имя пользователя: ", *name)
	fmt.Println("Возраст пользователя: ", *age)
	fmt.Println("Является администратором: ", *isAdmin)
}
