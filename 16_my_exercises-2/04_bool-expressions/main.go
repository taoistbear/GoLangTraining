/*
4. What's the value of the expression: (true && false) || (false && true) ||
!(false && false)?
*/

/*
At first pass this will express itself to True as the first two statments in
() each evaluate to false yet they are linked via the || (or) opeartor making
the last statement (NOT FALSE) equal true, thus the whole statement is ture.
If these were linked any other way even with one && (and) operator this would
be false. I will write a program to prove my statemetns.
*/

package main

import "fmt"

func main() {
	var first bool = true && false
	var second bool = false && true
	var third bool = !(false && !true)
	if first || second || third {
		fmt.Println("I was right, this was evaluated as TRUE!")
	} else {
		fmt.Println("I was wrong, this was evaluated as FALSE!")
	}
}

/*
In order to get his to run i had to get creative with my expressions, particularly
the thrid expression. !(false & flase) is the same as !(false && !true), the
compiler would not let me run this any other way. It wouldnt even let me run the
if statement with just the straight expressions thus I hid them in variables.
*/
