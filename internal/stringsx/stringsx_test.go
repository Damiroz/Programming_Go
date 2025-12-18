package stringsx

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestClip(t *testing.T) {
	tests := []struct {
		name string
		s    string
		max  int
		want string
	}{
		{"empty string", "", 5, ""},
		{"max is zero", "hello", 0, ""},
		{"max is negative", "hello", -1, ""},
		{"max equals len", "go", 2, "go"},
		{"max greater than len", "go", 10, "go"},
		{"normal clip", "fullstack", 4, "full"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, Clip(tt.s, tt.max))
		})
	}
}