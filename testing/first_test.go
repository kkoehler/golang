package testing

import "testing"

func TestAdd(testing *testing.T) {

	result := Add(10, 10)

	if result != 20 {
		testing.Errorf("Wrong result. Got %d but wanted 20", result)
	}

}

func TestTableAdd(t *testing.T) {

	tables := []struct {
		name  string
		first int
		y     int
		n     int
	}{
		{"one and one", 1, 1, 2},
		{"one and two", 1, 2, 3},
		{"two and two", 2, 2, 4},
		{"five and two", 5, 2, 7},
	}

	for _, table := range tables {

		t.Run(table.name, func(t *testing.T) {
			total := Add(table.first, table.y)
			if total != table.n {
				t.Errorf("%v: Sum of (%d+%d) was incorrect, got: %d, want: %d.", table.name, table.first, table.y, total, table.n)
			}
		})

	}

}

func BenchmarkAdd(b *testing.B) {

	for n := 0; n < b.N; n++ {
		Add(10, 10)
	}

}
