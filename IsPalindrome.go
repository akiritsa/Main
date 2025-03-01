package main

import "fmt"

func main() {
	chislo := -123456

	fmt.Println("Это число <", chislo, "> полиндром?:", isPalindrome(chislo))

}

func isPalindrome(x int) bool {
	//	if x < 0 {
	//		return false
	//	}
	xstr := fmt.Sprintf("%d", x)
	xstrmass := []rune(xstr)
	i := 0
	j := len(xstrmass) - 1
	for i <= j {
		if xstrmass[i] != xstrmass[j] {
			return false
		}
		i++
		j--
	}
	return true
}
