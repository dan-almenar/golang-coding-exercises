package main

import (
	"fmt"

	valbracket "github.com/dan-almenar/golang/coding-exercises/valid_bracket_sequence"
)

func main() {
	trueString := "klsgheklsckghkl(kdflkmvsk)"
	test := valbracket.ValidBracketSequence(trueString)
	fmt.Println(test)

	falseString := "lsrhklsdfk[sdkfjlsdksdl)fsdf"
	test = valbracket.ValidBracketSequence(falseString)
	fmt.Println(test)
}
