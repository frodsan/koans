package koans

import (
	"sort"
	"testing"
)

func TestSortingComparator(t *testing.T) {
	people := []Person{
		{Name: "amy", Score: 100},
		{Name: "david", Score: 100},
		{Name: "heraldo", Score: 50},
		{Name: "aakansha", Score: 75},
		{Name: "aleksa", Score: 150},
	}

	sort.Sort(sort.Reverse(SortedPeople(people)))

	expected := []string{"aleksa", "amy", "david", "aakansha", "heraldo"}

	for i, n := range expected {
		if n != people[i].Name {
			t.Errorf("Expected %s, got %s", n, people[i].Name)
		}
	}
}
