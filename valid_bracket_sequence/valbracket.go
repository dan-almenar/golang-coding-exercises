package valbracket

import "fmt"

type Node struct {
	Data     rune
	NextNode *Node
}

type Stack struct {
	Head     *Node
	nextNode *Node
	length   uint
}

func newNode(r rune) *Node {
	return &Node{
		Data: r,
	}
}

func newStack() *Stack {
	return &Stack{
		length: 0,
	}
}

func (s *Stack) append(n *Node) {
	previousHead := s.Head
	s.Head = n
	s.nextNode = previousHead
	s.length++
	fmt.Printf("New Head: %q\n", s.Head.Data)
}

func (s *Stack) pop() error {
	if s.Head == nil {
		return fmt.Errorf("stack is empty")
	}
	newHead := s.Head.NextNode
	s.Head = newHead
	s.length--
	return nil
}

func (s *Stack) compare(r rune) bool {
	if s.length == 0 {
		return false
	}
	if s.Head.Data == '(' && r != ')' {
		return false
	}
	if s.Head.Data == '[' && r != ']' {
		return false
	}
	if s.Head.Data == '{' && r != '}' {
		return false
	}
	s.pop()
	return true
}

func ValidBracketSequence(s string) bool {
	brackets := newStack()
	for _, char := range s {
		switch char {
		case '(':
			n := newNode(char)
			brackets.append(n)
		case '[':
			n := newNode(char)
			brackets.append(n)
		case '{':
			n := newNode(char)
			brackets.append(n)
		case ')':
			return brackets.compare(char)
		case ']':
			return brackets.compare(char)
		case '}':
			return brackets.compare(char)
		default:
			continue
		}
	}
	return false
}
