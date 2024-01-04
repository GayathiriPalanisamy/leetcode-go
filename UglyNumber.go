package main

import "fmt"

/*
https://leetcode.com/problems/palindrome-permutation/

An ugly number is a positive integer whose prime factors are limited to 2, 3, and 5.

Given an integer n, return true if n is an ugly number.

Example 1:

Input: n = 6
Output: true
Explanation: 6 = 2 × 3
Example 2:

Input: n = 1
Output: true
Explanation: 1 has no prime factors, therefore all of its prime factors are limited to 2, 3, and 5.
Example 3:

Input: n = 14
Output: false
Explanation: 14 is not ugly since it includes the prime factor 7.


Constraints:

-231 <= n <= 231 - 1
*/

func main() {
	fmt.Print(isUgly(27))
}

func isUgly(n int) bool {
	if n < 1 {
		return false
	}
	if n == 1 {
		return true
	}
	prime := []int{2, 3, 5}
	for _, i := range prime {
		for n%i == 0 {
			n = n / i
		}
	}
	return n == 1
}
