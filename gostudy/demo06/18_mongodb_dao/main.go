package main

import (
	"time"
)

// Article model
type Article struct {
	Name  string `validate:"required,min=1,max=200,excludesall=!/#?@&+="`
	Title string `validate:"required,min=1,max=100"`
	// article type name
	Type        string `validate:"required,min=1,max=100,excludesall=!/#?@&+="`
	Description string
	Content     string `validate:"required,min=1"`
	Sort        int    `validate:"min=-99999,max=99999"`
	// previous article
	Prev string
	// next article
	Next string
	Good int
	Top  int
	// article tags
	Tags      []string
	Hits      int
	Author    string
	Chapters  []*Article
	CreatedAt time.Time
}

var articleDao ArticleDao

var article = Article{
	Name:        "西游记",
	Title:       "西游记",
	Type:        "2",
	Description: "good one",
	Content:     "good study",
	Sort:        1,
	Prev:        "水浒传",
	Next:        "红楼梦",
	Good:        2,
	Top:         3,
	Tags:        []string{"神话", "冒险"},
	Hits:        2,
	Author:      "吴承恩",
}

var article1 = Article{
	Name:        "西游记",
	Title:       "西游记",
	Type:        "3",
	Description: "good one",
	Content:     "good study",
	Sort:        1,
	Prev:        "水浒传",
	Next:        "红楼梦",
	Good:        2,
	Top:         3,
	Tags:        []string{"神话", "冒险"},
	Hits:        2,
	Author:      "吴承恩123",
}

func main() {
	defer cancel()
	//articleDao.Insert("", &article)
	//articleDao.UpdateTypeName("2", "3")
	articleDao.UpdateArticle("3", "西游记", &article1)
}
