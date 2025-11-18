package main

import "fmt"

// Создайте массив, содержащий дни недели, и срез, оставшихся дней недели от 
// сегодняшней даты. Выведете информацию о длине и ёмкости массива и среза.
func main() {
	daysOfWeek := [7]string{
		"Понедельник", 
		"Вторник", 
		"Среда", 
		"Четверг", 
		"Пятница", 
		"Суббота", 
		"Воскресенье",
	}

	remainingDays := daysOfWeek[2:]

	fmt.Println(len(remainingDays), cap(remainingDays))
	fmt.Println(remainingDays)
}