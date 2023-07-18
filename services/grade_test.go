package services_test

import (
	"go-unit-test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckGrade(t *testing.T) {

	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "grade A", score: 80, expected: "A"},
		{name: "grade B", score: 70, expected: "B"},
		{name: "grade C", score: 60, expected: "C"},
		{name: "grade D", score: 50, expected: "D"},
		{name: "grade F", score: 0, expected: "F"},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			grade := services.CheckGrade(v.score)
			expected := v.expected

			assert.Equal(t, expected, grade)
			// if grade != expected {
			// 	t.Errorf("got %v expected %v", grade, expected)
			// }
		})
	}
}

func BenchmarkCheckGrade(b *testing.B) {

	for i := 0; i < b.N; i++ {
		services.CheckGrade(80)
	}
}
