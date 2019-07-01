package main

import (
	"errors"
	"fmt"
)

// Pages contains a slice of strings that will hold the information
type Pages struct {
	//Elements holds all of the information
	Elements []string
	// PageSize specifies the number of elements that a single page will hold
	PageSize int
	current  int
}

// FirstPage returns the elements on the first page and sets the current page to first
func (p *Pages) FirstPage() []string {
	p.current = 1
	return p.Elements[:p.PageSize]
}

// LastPage returns the elements on the last page and sets the current page to last
func (p *Pages) LastPage() []string {
	p.current = len(p.Elements) / p.PageSize
	residue := len(p.Elements) % p.PageSize
	if residue != 0 {
		p.current++
		return p.Elements[len(p.Elements)-residue:]
	}
	return p.Elements[len(p.Elements)-p.PageSize:]
}

// GetCurrentPageNumber returns the current page that we are on
func (p *Pages) GetCurrentPageNumber() int {
	return p.current
}

// HasNext checks if there is a page after the one we are currently on
func (p Pages) HasNext() bool {
	i := len(p.Elements) + p.PageSize - (p.current+1)*p.PageSize
	return i > 0
}

// HasPrevious checks if there is a page before the one we are currently on
func (p Pages) HasPrevious() bool {
	i := len(p.Elements) - (p.current+1)*(p.PageSize+1)
	return i <= 0
}

// Next returns the page after the one we are currently on
func (p *Pages) Next() ([]string, error) {
	if p.HasNext() != true {
		return nil, errors.New("There is no next page")
	}
	if p.current == len(p.Elements)/p.PageSize {
		return p.LastPage(), nil
	}
	s := p.Elements[p.current*p.PageSize : p.current*p.PageSize+p.PageSize]
	return s, nil
}

// Previous returns the page before the one we are currently on
func (p *Pages) Previous() ([]string, error) {
	if p.HasPrevious() != true {
		return nil, errors.New("There is no previous page")
	}
	if p.current == 2 {
		return p.FirstPage(), nil
	}
	s := p.Elements[p.current*p.PageSize-p.PageSize*2 : p.current*p.PageSize-p.PageSize]
	return s, nil
}

func main() {
	s := Pages{[]string{"book", "work", "cat", "house", "phone", "case", "cup", "mouse", "watch", "clock", "dog"}, 3, 4}
	//fmt.Println(s)
	//fmt.Println(s.FirstPage())
	//fmt.Println(s)
	//fmt.Println(s.LastPage())
	//fmt.Println(s)
	fmt.Println(s.GetCurrentPageNumber())
	fmt.Println(s.HasNext())
	fmt.Println(s.HasPrevious())
	f, err := s.Next()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(f)
	}
	m, err1 := s.Previous()
	if err1 != nil {
		fmt.Println(err1)
	} else {
		fmt.Println(m)
	}
}
