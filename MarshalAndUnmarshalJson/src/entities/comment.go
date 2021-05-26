package entities

import "fmt"

type Comment struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func (comment Comment) ToString() string {
	
	return fmt.Sprintf("Title:%s\nContent:%s",comment.Title,comment.Content)

}