/*
	Part 1 of the assignment
	Part 2 implementation is in mapper/mapper.go

	Run either part using this main function
*/
package main

import (
	"fmt"
	"strings"

	"github.com/seanmdeleon/StringMapper/mapper"
	"github.com/seanmdeleon/StringMapper/utils"
)

func main() {

	fmt.Printf("---------- PART 1 testing ------------\n\n")
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("Aspiration.com"))
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("A..spir..ation.com"))
	fmt.Println(CapitalizeEveryThirdAlphanumericChar("A12sp34ira5555....tion.com"))
	fmt.Printf("\n\n")

	fmt.Printf("---------- PART 2 testing ------------\n\n")
	s1 := mapper.NewSkipString(3, "Aspiration.com")
	mapper.MapString(s1)
	fmt.Println(s1)
	fmt.Printf("TestPassed: %t\n\n", s1.AreStringsEqual("asPirAtiOn.cOm"))

	s2 := mapper.NewSkipString(1, "1234567890")
	mapper.MapString(s2)
	fmt.Println(s2)
	fmt.Printf("TestPassed: %t\n\n", s2.AreStringsEqual("1234567890"))

	s3 := mapper.NewSkipString(5, "A12sp34ira5555....tion.com")
	mapper.MapString(s3)
	fmt.Println(s3)
	fmt.Printf("TestPassed: %t\n\n", s3.AreStringsEqual("a12sP34irA5555....Tion.cOm"))

	s4 := mapper.NewSkipString(500, "Aspiration.com")
	mapper.MapString(s4)
	fmt.Println(s4)
	fmt.Printf("TestPassed: %t\n\n", s4.AreStringsEqual("aspiration.com"))

	s5 := mapper.NewSkipString(7, "Aspiration.com is a company that I am interviewing for")
	mapper.MapString(s5)
	fmt.Println(s5)
	fmt.Printf("TestPassed: %t\n\n", s5.AreStringsEqual("aspiraTion.com Is a compAny that I am inteRviewinG for"))
}

// CapitalizeEveryThirdAlphanumericChar capitalizes every third alphanumeric char, and lowercases everyother char
// Cannot capitalize/lowercase numbers, but a number is included in the count
func CapitalizeEveryThirdAlphanumericChar(s string) string {

	// lowercase everything in the beginning
	s = strings.ToLower(s)

	// go strings are immutable so change this string to a rune slice
	runeSlice := []rune(s)

	count := 0
	for i := 0; i < len(runeSlice); i++ {
		if utils.IsAlphanumericRune(runeSlice[i]) {
			count++
			if count == 3 {
				// Capitalize an alphanumeric Unicode. If a number, nothing happens but we include it in the count
				runeSlice[i] = utils.CapitalizeUnicode(runeSlice[i])

				// reset the count
				count = 0
			}
		}
	}

	// return the updated charSlice as a string
	return string(runeSlice)
}
