package main

import (
	"unicode/utf8"
	"strings"
	"fmt"
)

func main(){
	str:="my name is billy"
	for i, c := range str{
		if c == utf8.RuneError{
			fmt.Println("1")
		}
		width := utf8.RuneLen(c)
		str = str[i+width:]
		fmt.Println(str)
		fmt.Println(c)
	}
	fmt.Println(strings.Title(str))
}