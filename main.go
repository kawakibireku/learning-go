package main

import "fmt"

func main() {
	/*Try to print line using fmt ( format ) function*/
	fmt.Println("Hello, World!")

	/*Try to declare string variable ( var {name_variable} {type_data} ) but golang is picky, it doesn't want to have a variable hanging around*/
	var whatToSay string
	/*So I need to fill out the variable with some string*/
	whatToSay = "Goodbye, cruel world"
	/*looks like golang still being picky
	(it's still have the red underscore thingy on the whatToSay variable) because I still not using it,
	so I used the fmt function to print line the variable */
	fmt.Println(whatToSay)
}
