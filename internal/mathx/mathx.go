package mathx

import "errors"

// Sum возвращает сумму двух чисел
func Sum(a, b int) int { return a + b }

// Divide выполняет деление и возвращает ошибку при делении на ноль
func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("divide by zero")
	}
	return a / b, nil
}