package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

// Объявляем файлы в которых будут храниться данные приложения и файл для записи логов
var fpath = "~/dat"

//var lpath = "~/log"

// Структура которая будет объединять в себе пароль, логин и ресурс, к которому логин и пароль привязаны
type info struct {
	resource string
	login    string
	pwd      string
}

// Мастер пароль по которому будет осуществляться доступ к основным функциям приложения
const masterPass = "SoMePaSs1234"

// Добавление новой записи в файл
func AddPass(inf info) error {
	file, err := os.Open(fpath)
	if err != nil {
		// Запись информации об ошибки в Лог файл
		return err
	}
	defer file.Close()

	// Создаём строку следующего формата "***resource***login***password\n" и добавляем её в файл
	str := "***" + inf.resource + "***" + inf.login + "***" + inf.pwd + "\n"
	_, err = file.WriteString(str)

	if err != nil {
		// Запись информации об ошибки в Лог файл
		return err
	}
	return nil
}

// Поиск пароля по названию ресурса
func FindPass(rn string) (info, error) {
	file, err := os.Open(fpath)
	if err != nil {
		// Запись информации об ошибки в Лог файл
		return info{"", "", ""}, err
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		_, err := file.Read(data)
		if err == io.EOF { // если конец файла
			fmt.Println("Password not found")
			break // выходим из цикла
		}

		s := strings.Trim(string(data), "***")

		if string(s[0]) == rn {
			return info{string(s[0]), string(s[1]), string(s[2])}, nil
		}
	}

	return info{}, errors.New("problem while finding password")
}

// Проверка правильности Мастер Пароля
func CheckMP() (bool, error) {
	var userInput string                      // user input - Введённая пользователем строка
	_, err := fmt.Fscan(os.Stdin, &userInput) // получаем userInput
	// Проверка того что пароль введённый пользователем успешно получен
	if err != nil {
		fmt.Println("The problem with scanning appeared")
		//f, _ := os.OpenFile(lpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		return false, errors.New("problem with scanning user input")
	}
	// Начало обработки пароля пользователя
	if userInput == masterPass {
		// Пишем в Лог файл что введён правильный пароль
		return true, nil
	} else {
		fmt.Println("Master Password is incorrect\n " +
			"Try one more time")
		return false, nil
	}
}

func main() {

	// Цикл проверки пароля
	for {
		b, err := CheckMP()

		if err != nil {
			// Запись в лог файл сообщения об ошибке
		}

		if b { // если пароли совпадают выходим из цикла проверки пароля
			break
		}
	}

	userChoice := "" // Здесь будет храниться команда пользователя
	for userChoice != "EXIT" {

		fmt.Println("Type one of command below:\n" +
			"n to add new password\n" +
			"f to find password\n" +
			"exit - to exit program" +
			"Your command: ")
		_, err := fmt.Fscan(os.Stdin, &userChoice)

		// Обработка ошибки вызванной системной комнандой
		if err != nil {
			// Запись ошибки в журнал
			fmt.Println("The problem with scanning appeared")
		}

		strings.ToUpper(userChoice)
		switch userChoice {
		case "N":
			{
				var res, log, pwd string
				fmt.Print("Write Resource: ")
				fmt.Fscan(os.Stdin, &res)
				fmt.Println()
				fmt.Print("Write Login: ")
				fmt.Fscan(os.Stdin, &log)
				fmt.Println()
				fmt.Print("Write Password: ")
				fmt.Fscan(os.Stdin, &pwd)
				fmt.Println()

				var a info = info{res, log, pwd}
				err := AddPass(a)
				if err != nil {
					// Запись ошибки в журнал
					fmt.Println("The problem with adding new password appeared")
				} else {
					fmt.Println("Password added successfully")
				}
			}

		case "F":
			{
				var res string
				fmt.Print("Write Resource: ")
				fmt.Fscan(os.Stdin, &res)
				fmt.Println()
				inf, err := FindPass(res)
				if err != nil {
					// Запись ошибки в журнал
					fmt.Println("The problem with finding password appeared")
				}
				fmt.Printf("Resouce: %s\n Login: %s\n Password: %s\n", inf.resource, inf.login, inf.pwd)
			}

		default:
			{
				fmt.Printf(" %s - does not exist\n", userChoice+
					"Try another time!")
			}
		}
	}

}

// Функции для реализации потом
// Изменение имеющегося пароля
// Удаление пароля
