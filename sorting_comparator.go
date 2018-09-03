package koans

/*

Source: Hacker Rank (https://www.hackerrank.com/challenges/ctci-comparator-sorting/)

In this challenge, you'll create a comparator and use it to sort an array.
The Player class is provided in the editor below. It has two fields:

1. `name`: a string.
2. `score`: an integer.

Given an array of `n` Player objects, write a comparator that sorts them
in order of decreasing score. If 2 or more players have the same score,
sort those players alphabetically ascending by name. To do this, you must
create a Checker class that implements the Comparator interface, then write an int
compare(Player a, Player b) method implementing the Comparator.compare(T o1, T o2)
method. In short, when sorting in ascending order, a comparator function returns
`-1` if `a < b`, `0` if `a = b` , and `1` if `a > b`.

*/

type Person struct {
	Name  string
	Score int
}

type SortedPeople []Person

func (s SortedPeople) Len() int {
	return len(s)
}

func (s SortedPeople) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortedPeople) Less(i, j int) bool {
	if s[i].Score == s[j].Score {
		return s[i].Name > s[j].Name
	}
	return s[i].Score < s[j].Score
}
