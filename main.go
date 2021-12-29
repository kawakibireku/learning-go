package main

import "fmt"

func main() {
	/*Try to print line using fmt ( format ) function*/
	fmt.Println("Hello, World!")

	/*Try to declare string variable ( var {name_variable} {type_data} ) but golang is picky, it doesn't want to have a variable hanging around*/
	var whatToSay string
	var i int
	/*So I need to fill out the variable with some string*/
	whatToSay = "Goodbye, cruel world"
	/*looks like golang still being picky
	(it's still have the red underscore thingy on the whatToSay variable) because I still not using it,
	so I used the fmt function to print line the variable */
	fmt.Println(whatToSay)

	i = 1
	fmt.Println("i is set to", i)

	whatWasSaid, theOtherThingThatWhatSaid := saySomething()

	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWhatSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
