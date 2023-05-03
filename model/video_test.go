package model

import (
	"sort"
	"testing"
	"time"
)

func TestUpdateValues(t *testing.T) {
	type T struct {
		Int    int
		String string    `db:"string"`
		Time   time.Time `db:"time"`
	}
	tests := []struct {
		T    T
		want []values
	}{
		{T{0, "", time.Time{}}, []values{}},
		{T{1, "", time.Time{}}, []values{{value: 1, name: "Int"}}},
		{T{0, "a", time.Time{}}, []values{{value: "a", name: "string"}}},
		{T{0, "", time.Unix(0, 0)}, []values{{value: time.Unix(0, 0), name: "time"}}},
		{T{1, "hello", time.Time{}}, []values{{value: "hello", name: "string"}, {value: 1, name: "Int"}}},
		{T{1, "a", time.Unix(0, 0)}, []values{{value: time.Unix(0, 0), name: "time"}, {value: 1, name: "Int"}, {value: "a", name: "string"}}},
	}

	for _, test := range tests {
		got := updatedValues(test.T)
		if len(got) != len(test.want) {
			t.Errorf("len(updateValues(%q)) = %d but len(want) = %d", test.T, len(got), len(test.want))
			continue
		}
		sort.Slice(got, func(i, j int) bool { return got[i].name < got[j].name })
		sort.Slice(test.want, func(i, j int) bool { return test.want[i].name < test.want[j].name })

		for i := 0; i < len(got); i++ {
			if got[i] != test.want[i] {
				t.Errorf("updateValues(%q)[%d] = %q but want[%d] = %q", test.T, i, got[i], i, test.want[i])
			}
		}
	}
}
