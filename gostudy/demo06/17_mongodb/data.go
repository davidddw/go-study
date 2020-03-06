package main

import "fmt"

const (
	categoryComputer = "计算机"
	categorySciFi    = "科幻"
	countryChina     = "中国"
	countryAmerica   = "美国"
)

// Book entity
type Book struct {
	ID       string
	Name     string
	Category string
	Weight   int
	Author   AuthorInfo
}

// AuthorInfo author
type AuthorInfo struct {
	Name    string
	Country string
}

// GetID inpletment
func (e Book) GetID() string {
	return e.ID
}

// SetID setter
func (e *Book) SetID(id string) {
	e.ID = id
}

func (e Book) String() string {
	return fmt.Sprintf("BookInfo(id:%s, name:%s, category:%s, author:%s)",
		e.ID, e.Name, e.Category, e.Author.Name)
}

var (
	books = []interface{}{
		&Book{
			ID:       UUID(),
			Name:     "深入理解计算机操作系统",
			Category: categoryComputer,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "兰德尔 E.布莱恩特",
				Country: countryAmerica,
			},
		},
		&Book{
			ID:       UUID(),
			Name:     "深入理解Linux内核",
			Category: categoryComputer,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "博韦，西斯特",
				Country: countryAmerica,
			},
		},
		&Book{
			ID:       UUID(),
			Name:     "三体",
			Category: categorySciFi,
			Weight:   1,
			Author: AuthorInfo{
				Name:    "刘慈欣",
				Country: countryChina,
			},
		},
	}

	book = &Book{
		ID:       UUID(),
		Name:     "Go程序设计语言",
		Category: categoryComputer,
		Weight:   1,
		Author: AuthorInfo{
			Name:    "艾伦A.A.多诺万",
			Country: countryAmerica,
		},
	}
)
