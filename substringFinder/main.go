package main

import (
	"flag"
	"fmt"
)

func main() {

	str := flag.String("str", "", "Основная строка для поиска")
	substr := flag.String("substr", "", "Подстрока для поиска")
	flag.Parse()

	if *str == "" || *substr == "" {
		fmt.Println("Необходимо указать значения флагов --str и --substr!")
		return
	}

	fmt.Println("Результат поиска в основной строке:", substringFinder(*str, *substr))
}

// Функция, которая запускает поиск подстроки substr в основной строке s.
// Для анализа и поиска совпадений использует дополнительную функцию matchingStrings,
// вложенную в цикле.
func substringFinder(s, subs string) bool {

	runeS := []rune(s)
	runeSubs := []rune(subs)

	lenS := len(runeS)
	lenSubs := len(runeSubs)

	for i := 0; i <= lenS-lenSubs; i++ {
		if matchingStrings(runeS[i:i+lenSubs], runeSubs) {
			return true
		}
	}
	return false
}

// Функция, которая сравнивает передаваемую часть строки s из функции matchingStrings
// на соответствие с подстрокой subs по значениям индексов обоих.
// Если даже первая итерация выявляет разность значений элементов, возвращается false.
func matchingStrings(s, subs []rune) bool {
	for i := range subs {
		if s[i] != subs[i] {
			return false
		}
	}
	return true
}
