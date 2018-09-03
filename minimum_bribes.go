package koans

import (
	"errors"
)

/*

Source: Hacker Rank (https://www.hackerrank.com/challenges/new-year-chaos/)

It's New Year's Day and everyone's in line for the Wonderland rollercoaster ride!
There are a number of people queued up, and each person wears a sticker indicating
their initial position in the queue. Initial positions increment by `1` from `1` at
the front of the line to `n` at the back.

Any person in the queue can bribe the person directly in front of them to swap positions.
If two people swap positions, they still wear the same sticker denoting their original places
in line. One person can bribe at most two others. For example, if `n = 8` and `Person 5` bribes
`Person 4`, the queue will look like this: `1,2,3,5,4,6,7,8`.

Fascinated by this chaotic queue, you decide you must know the minimum number of bribes that
took place to get the queue into its current state!

*/

func MinimumBribes(q []int) (int, error) {
	var swaps int

	for i, n := range q {
		if (n - 1 - i) > 2 {
			return -1, errors.New("Too chaotic")
		}
	}

	swapped := true

	for swapped {
		swapped = false
		for i := 0; i < len(q)-1; i++ {
			if q[i] > q[i+1] {
				tmp := q[i+1]

				q[i+1] = q[i]
				q[i] = tmp

				swaps++
				swapped = true
			}
		}
	}

	return swaps, nil
}
