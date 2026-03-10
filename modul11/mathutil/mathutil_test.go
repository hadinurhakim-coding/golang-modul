package mathutil

import "testing"

func TestSum_TableDriven(t *testing.T) {
	tests := []struct {
		nama  string
		input []int
		want  int
	}{
		{nama: "kosong", input: []int{}, want: 0},
		{nama: "satu angka", input: []int{5}, want: 5},
		{nama: "banyak angka", input: []int{1, 2, 3}, want: 6},
		{nama: "angka negatif", input: []int{-2, 5, -3}, want: 0},
	}

	for _, tc := range tests {
		t.Run(tc.nama, func(t *testing.T) {
			got := Sum(tc.input)
			if got != tc.want {
				t.Fatalf("Sum(%v) = %d; want %d", tc.input, got, tc.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	t.Run("normal", func(t *testing.T) {
		got, err := Divide(10, 2)
		if err != nil {
			t.Fatalf("expected nil error, got %v", err)
		}
		if got != 5 {
			t.Fatalf("Divide(10,2) = %v; want 5", got)
		}
	})

	t.Run("bagi nol", func(t *testing.T) {
		_, err := Divide(10, 0)
		if err == nil {
			t.Fatalf("expected error, got nil")
		}
	})
}

