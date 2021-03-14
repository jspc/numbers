package numbers

import (
	"strings"
)

var (
	// units holds a mapping of index -> string for numbers sub 10,000
	// so: in the number 1,234 the unit for '1' (position 0 in the string)
	// is unit[0]. For '2' (position 1) the unit is unit[1]
	//
	// The final unit is empty because there is no unit. We store it, though,
	// to avoid special casing
	units = []string{"천", "백", "십", ""}

	// bigUnits are suffixed to units (once units are all stringified)
	bigUnits = []string{"", "만", "억", "조"}

	// digits is slice of index -> string for individual digits
	digits = []string{"공", "일", "이", "삼", "사", "오", "육", "칠", "팔", "구"}

	// groupSize represents the size of each block in Atoi operations
	groupSize = 4
)

// Itoa accepts an integer and returns a textual representation
func Itoa(i int) string {
	outStrings := []string{}

	if i == 0 {
		return digits[0]
	}

	for i, g := range group(i) {
		outStrings = append([]string{stringify(g, i)}, outStrings...)
	}

	return cleanString(outStrings)
}

// cleanString takes a slice of strings, drops any empty entry, and
// returns the remaining entries joined
func cleanString(s []string) string {
	out := make([]string, 0)
	for _, str := range s {
		if str != "" {
			out = append(out, str)
		}
	}

	return strings.Trim(strings.Join(out, " "), " ")
}

// stringify takes a slice of integers, as taken from group/pad, and
// returns a string representing it
func stringify(slice []int, bigIdx int) string {
	s := strings.Builder{}

	for i, u := range slice {
		if u == 0 {
			continue
		}

		if (u == 1 && i == 3 && bigIdx == 0) || (u > 1) {
			s.WriteString(digits[u])
		}

		unit := units[i]
		if unit != "" {
			s.WriteString(unit)

			// peak ahead; if the rest of slice is just a load of
			// zeroes (if indeed there's anything) then don't bother
			// add a space- there's nothing left to come)
			if sum(slice[i+1:groupSize]) > 0 {
				s.WriteString(" ")
			}
		}
	}

	s.WriteString(bigUnits[bigIdx])

	return s.String()
}

// sum returns a sum of all remaining values in a slice
func sum(i []int) (out int) {
	for _, v := range i {
		out += v
	}

	return
}

// revIntToSlice returns a slice of digits, reversed. Digits are reversed in order
// to make grouping easier, since grouping digits is right-to-left
func revIntToSlice(n int, sequence []int) []int {
	if n != 0 {
		i := n % 10
		sequence = append(sequence, i) // reverse order output
		return revIntToSlice(n/10, sequence)
	}

	return sequence
}

// group takes an integer and returns a slice of slices containing
// each 10 thousand group.
//
// it works, for instance, on the number 712,132 by returning:
// [2,1,3,2], [0,0,7,1]
func group(i int) (chunks [][]int) {
	slice := []int{}
	slice = revIntToSlice(i, slice)

	for i := 0; i < len(slice); i += groupSize {
		end := i + groupSize

		// necessary check to avoid slicing beyond
		// slice capacity
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, reverse(pad(slice[i:end])))
	}

	return chunks
}

// reverse accepts a slice of integers and returns it, reversed
func reverse(s []int) []int {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

// pad accepts a slice of arbitrary length, and pads it to, at minimum,
// be numbers.groupSize long
func pad(slice []int) []int {
	if len(slice) >= groupSize {
		return slice
	}

	s := slice
	for {
		if len(s) >= groupSize {
			return s
		}

		s = append(s, 0)
	}
}
