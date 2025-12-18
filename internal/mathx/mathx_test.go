package mathx

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Табличный тест для суммы с использованием стандартного пакета testing
func TestSum_Table(t *testing.T) {
	cases := []struct{ a, b, want int }{
		{2, 3, 5},
		{10, -5, 5},
		{0, 0, 0},
	}
	for _, c := range cases {
		got := Sum(c.a, c.b)
		if got != c.want {
			t.Fatalf("Sum(%d,%d)=%d; want %d", c.a, c.b, got, c.want)
		}
	}
}

// Тест деления с использованием testify
func TestDivide_OkAndError(t *testing.T) {
	// Успешный кейс
	got, err := Divide(10, 2)
	require.NoError(t, err) // require останавливает тест при ошибке
	assert.Equal(t, 5, got)

	// Кейс с ошибкой
	_, err = Divide(10, 0)
	assert.Error(t, err)
	assert.Equal(t, "divide by zero", err.Error())
}

// Бенчмарк для функции Sum
func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Sum(123, 456)
	}
}