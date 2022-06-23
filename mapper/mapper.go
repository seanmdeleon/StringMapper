/*
	Part 2 of the assignment
	Create a mapper package that implements part 1 using interfaces

	For simplicity I am including the SkipString implementation in this file instead of creating a SkipString package seperately

*/
package mapper

import (
	"fmt"
	"strings"

	"github.com/seanmdeleon/StringMapper/utils"
)

// Interface describes the list of methods required to implement the Interface interface
type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

// MapString takes in an Interface and maps the string given the underlying implementation
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

/*
	SkipString holds data for capitalizing characters within AlphanumericString at a given frequency
	SkipString is a structre that implements Interface and Stringer interfaces
	Keeping the struct fields non-exported for fun!
*/
type SkipString struct {
	frequency          int
	count              int
	alphanumericString string
	runeSlice          []rune
}

// Verify that the SkipString struct implements Interface interface
var _ Interface = (*SkipString)(nil)

// Verify that the SkipString struct implements the Stringer interface
var _ fmt.Stringer = (*SkipString)(nil)

// NewSkipString is the constructor for a SkipString
// freq defines the frequence of how many alphanumeric values to skip
func NewSkipString(freq int, s string) *SkipString {
	return &SkipString{
		frequency:          freq,
		count:              0,
		alphanumericString: s,
	}
}

// TransformRune transforms a rune (char) to uppercase if the value at pos is lowercase and the count == the frequency
func (s *SkipString) TransformRune(pos int) {

	if utils.IsAlphanumericRune(s.runeSlice[pos]) {
		s.count++
		if s.count == s.frequency {
			// it's time to uppercase
			// if the rune at pos is a number, nothing will happen
			s.runeSlice[pos] = utils.CapitalizeUnicode(s.runeSlice[pos])

			s.count = 0
		}
	}
}

// GetValueAsRuneSlice castes a string to a slice of rune and stores it in the SkipString structure.
// For this implementation of Interface, it also takes the liberty of lowercasing all characters in the string to begin with
// This lowercasing using the strings package does not alter the underlying AlphanumericString
func (s *SkipString) GetValueAsRuneSlice() []rune {
	s.runeSlice = []rune(strings.ToLower(s.alphanumericString))
	return s.runeSlice
}

// String pretty prints a SkipString
func (s *SkipString) String() string {
	return fmt.Sprintf("SkipString('%s', %d) --> '%s'", s.alphanumericString, s.frequency, string(s.runeSlice))
}

// AreStringsEqual is a bonus method that checks if expected is equal to the underlying SkipString
// This is used for testing and is not part of the Interface interface
func (s *SkipString) AreStringsEqual(expected string) bool {
	return expected == string(s.runeSlice)
}
