package repo

import "fmt"

type Stack struct {
	elements []string
}

func (s *Stack) Push(element string) *string {
	index := len(s.elements) - 1
	s.elements = append(s.elements, element)
	if index < 0 {
		return nil
	} else {
		return &s.elements[index]
	}
}

func (s *Stack) Pop() (*string, error) {
	if len(s.elements) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}
	index := len(s.elements) - 1
	element := s.elements[index]
	s.elements = s.elements[:index]
	return &element, nil
}
