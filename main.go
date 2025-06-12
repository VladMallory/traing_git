package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func writeMessage(x int) string {
	var messages []string
	filename := "messages.json"

	// Чтение файла, если он существует
	if data, err := os.ReadFile(filename); err == nil {
		json.Unmarshal(data, &messages)
	}

	// Ввод новой строки с поддержкой пробелов
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите сообщение: ")
	input, _ := reader.ReadString('\n') // считывает всю строку до Enter

	// Добавляем сообщение (убираем символ перевода строки)
	messages = append(messages, input[:len(input)-1])

	// Сохраняем обратно
	data, _ := json.MarshalIndent(messages, "пиздааа", "  ")
	os.WriteFile(filename, data, 0644)

	// Вывод всех сообщений с индексами
	fmt.Println("Все ошибки:")

	var result string
	for i, msg := range messages {
		i += 1
		result = result + fmt.Sprintf("[%d]: %s\n", i, msg)
	}
	return result
}

func viewMessage(x int) string {
	//открытие файла
	file, err := os.Open("messages.json")
	if err != nil {
		fmt.Println("Ошибка в viewMessage")
	}

	//чтение файла и вывод файла
	readFile, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка readFile")
	}

	var messages []string
	json.Unmarshal(readFile, &messages)

	var result string
	for i, msg := range messages {
		i += 1
		result = result + fmt.Sprintf("[%d]: %s\n", i, msg)
	}

	defer file.Close()
	return result
}

func deleteMessage(x int) string {
	file, err := os.Open("messages.json")
	if err != nil {
		fmt.Println("Ошибка в deleteMessage")
	}

	//открытие файла
	readFile, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Ошибка в deleteMessage")
	}

	//парсинг json
	var messages []string
	json.Unmarshal(readFile, &messages)

	//спрашиваем число
	var indexInput int
	fmt.Println("Введи номер который нужно удалить")
	fmt.Scanln(&indexInput)

	element := messages[indexInput]
	fmt.Println(element)
	//messages[indexInput] = messages[len(messages)-1]
	return "Сообщение удалено"
}

func main() {
	fmt.Println("Вы хотите посмотреть ошибки или записать?")
	fmt.Printf("1. Посмотреть\n2. Записать\n3. Удалить\n")

	var choice int
	fmt.Scanln(&choice)

	//условие
	if choice == 1 {
		resultViewMessage := viewMessage(choice)
		fmt.Println(resultViewMessage)

	} else if choice == 2 {
		resultWriteMessage := writeMessage(choice)
		fmt.Println(resultWriteMessage)
	} else if choice == 3 {
		var input int
		input += 1
		resultDeleteMessage := deleteMessage(choice)
		fmt.Println(resultDeleteMessage)
	}
}
