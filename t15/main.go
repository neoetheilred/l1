package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
	var justString string // Использование глобальной переменной может привести к неопределенному поведению при конкуррентных вычислениях
	func someFunc() {
		v := createHugeString(1 << 10) // Создание огромной строки, само по себе медленно
		// Используются первые 100 символов строки, остальные остаются в памяти, но не используются (string это обертка над []byte)
		justString = v[:100] // Если используются Unicode-символы, то, возможно, мы получим не то, что ожидаем
		// Например len("世界") == 6, а не 2, как можно подумать
		// Запись в глобальную переменную без использования mutex может привести к неопределенному поведению
	}

	func main() {
		someFunc()
	}
*/

func createHugeString(n int) string {
	return strings.Repeat("a", n)
}

var justString string
var justStringMutex sync.Mutex

func someFunc() {
	justStringMutex.Lock()
	justString = createHugeString(100)
	justStringMutex.Unlock()
}

func main() {
	someFunc()
	fmt.Println(justString)
}
