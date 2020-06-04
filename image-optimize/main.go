package main

import "fmt"

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/
func main() {

	var largestPalindrome = 0
	var m = 0
	var n = 0

	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			k := i * j
			if palindrome(k) && largestPalindrome < k {
				largestPalindrome = k
				m = i
				n = j
			}
		}
	}

	fmt.Printf("%d is product of %d,%d\r\n", largestPalindrome, m, n)

}

func palindrome(i int) bool {
	var reverseNumber = 0
	var temp = i

	for temp > 0 {
		var remainder = temp % 10
		reverseNumber = (reverseNumber * 10) + remainder
		temp = temp / 10
	}

	if i == reverseNumber {
		return true
	}

	return false

}
