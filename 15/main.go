/*
 * К каким негативным последствиям может привести данный фрагмент кода, 
 * и как это исправить? 
 * Приведите корректный пример реализации.
 *
 * var justString string
 * func someFunc() {
 *   v := createHugeString(1 << 10)
 *   justString = v[:100]
 * }
 *
 * func main() {
 *   someFunc()
 * }
 */

package main

import (
	"fmt"
	"strings"
)

/*
 * Проблемы:
 * 1 - На строку v ссылается переменная justString путем получения среза, 
 * поэтому garbage collector не будет ее очищать. Это бесполезно и дорого;
 * 2 - Нецелесообразное использование глобальной переменной;
 * 3 - someFunc работает корректно только для ASCII-символов.
 *
 * Решения:
 * 1 - Получать копию первых 100 символов, чтобы "отпустить" длинную строку v;
 * 2 - Добавить someFunc возвращаемое значение и создавать justString локально;
 * 3 - Использовать руны.
 */

func someFunc(v string) string {
	res := []rune(v[:100]) // Путем преобразования в слайс рун, создается новая копия
	return string(res)
}

func createHugeString(length int) string {
	return strings.Repeat("x", length)
}

func main() {
	v := createHugeString(1 << 10)

	justString := someFunc(v)

	fmt.Println(justString)
}
