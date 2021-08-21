# Golang Coding Exercises

Hi!

Here i'll be uploading Golang exercises from coding interviews I've found -and possibly had failed- (specially in said cases, since this is more than anything else a place for me to learn from my mistakes).

## List of Exercises:
### - **Validate Brackets Sequence**:
First exercise I worked with is the *'Validate Brackets Sequence'*, which essentially consist on verifying if a closing bracket ")", "]", "}" corresponds to the last openning one found on a given string and returning a boolean in either case.

### Example:
Given the string "([{}])" would return true whereas the string "([}]]" would return false, since ']' is not paired with '{'.

### My implementation:
As you may see in the valbracket package I defined two strucks:
- Node, with a Data field holding a rune and a nextNode field that holds the pointer to a node.
- Stack (Head: \*Node, nextNode: \*Node, lenght: uint)
- - The Stack struckt has three methods: **append**, which adds a new Head, and **pop** that deletes the current Head, setting the stack Head with the previous Head nextNode. Finally, the **compare** method will test if a given rune matches the value of the Head.Data, returning a boolean.

The exported **ValidBracketSequence** function receives a string and traverses it: if an openning bracket is found it will be appended to the brackets Stack -defined at the top of the function-. If a closing bracket is found, the bracket.compare() method is called. That way the string only has to be traversed once, meaning **the ValidBracketSequence function has a Big O(n) or linear time complexity**.  

Note: Each exercise will be a separated Golang package, while the *main.go* file will remain as a playground for showing purpuses.

Thanks!
#### Dan Almenar
