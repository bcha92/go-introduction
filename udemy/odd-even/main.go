package main

import "fmt"

// Establish ints as a slice type only accepting int(egers)
type ints []int

func main() {
	numbers := ints{} // Establish numbers as empty slice(array/list if you use JS/C-suite/Java or Python respectively)

	// I could just declare numbers as ints{0, 1, 2, ... , 10} but some lazy coding and making good use of append() makes the process far more flexible in my opinion, but I will respect any digression or disagreements to using a for loop to expand the slice.
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	// fmt.Println(numbers) // Console check to make sure numbers correctly output [0 1 2 ... 10]

	// Although it would have been my personal preference to print the output in the previous for loop and end it, but for the sake of this assignment asking for a specific creation of the slice before iterating it for odd/even categorization, I will go along with this.
	for _, num := range numbers {
		if num%2 == 0 { // Using % as a math operation ensures if there is a remainder after dividing by 2 (even numbers always has a remainder of 0)
			fmt.Println(num, "is even")
		} else {
			fmt.Println(num, "is odd")
		}
	}
}
