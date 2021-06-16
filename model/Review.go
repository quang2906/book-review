package model

type Review struct {
	Id      int64  `json:"id"`
	BookId  int64  `json:"bookId"`
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
}
