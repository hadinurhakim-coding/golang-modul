package calculator

import "testing"

func TestAdd(t *testing.T) {
	if got := Add(2, 3); got != 5 {
		t.Fatalf("Add(2,3) = %d; want 5", got)
	}
}

func TestSub(t *testing.T) {
	if got := Sub(10, 4); got != 6 {
		t.Fatalf("Sub(10,4) = %d; want 6", got)
	}
}

func TestMul_TableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{a: 0, b: 5, want: 0},
		{a: 3, b: 4, want: 12},
		{a: -2, b: 3, want: -6},
	}

	for _, tc := range tests {
		got := Mul(tc.a, tc.b)
		if got != tc.want {
			t.Fatalf("Mul(%d,%d) = %d; want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

