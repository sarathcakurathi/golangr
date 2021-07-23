package main

import (
	"fmt"
	"strings"
)

func main() {
	var str1 = "String example"
	fmt.Println(str1)

	// Observe a % char printed, why?
	fmt.Printf(str1)

	fmt.Println()
	var str2 = "Mein name ist\nSarath"
	fmt.Println(str2)

	// Observe a % char printed again, why?
	fmt.Print(str2)
	fmt.Println("")

	// String index
	str3 := "Test indexing"
	fmt.Println("-------- Test indexing --------")
	for k, v := range str3 {
		fmt.Printf("k: %d, v: %c == %d\n", k, v, v)
	}
	fmt.Println(str3[2])
	fmt.Println("Address: ", &str3)

	/* Note: Index cannot be negative; str3[-5] won't work
	Index prints ASCII code
	Wont work - fmt.Println("Address Index: ", &str3[0]) - Getting address of a byte in a string is illegal
	*/

	// String splicing
	str := "sar" + "ath"
	str += " chandra"
	fmt.Println(str)
	fmt.Println("Byte length of str: ", len(str))

	// String comparison
	// The general comparison operators (==,!=, <, <=,>=, >) implement the comparison of strings by byte comparison in memory.
	str1_test := "sar"
	str2_test := "sara"
	fmt.Println(str1_test == str2_test, str1_test >= str2_test, str1_test <= str2_test)

	// String join
	s := strings.Join([]string{str1_test, str2_test}, ", ")
	fmt.Println(s)
}
