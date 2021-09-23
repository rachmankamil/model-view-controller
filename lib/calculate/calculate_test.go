package calculate_test

import (
	"testing"

	"github.com/rachmankamil/kampus-merdeka-b/lib/calculate"
)

func TestAddition(t *testing.T) {
	t.Run("valid test", func(t *testing.T) {
		param1 := 10
		param2 := -10
		result := calculate.Addition(param1, param2)

		if result != 0 {
			t.Error("Expected 0")
		}
	})
}
