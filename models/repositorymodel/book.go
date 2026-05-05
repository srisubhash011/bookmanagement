package repositorymodel

type Book struct {
	BookId          int    `json:"bookid"`
	BookName        string `json:"bookname"`
	BookDescription string `json:"bookdescription"`
	BookStatus      int    `json:"bookstatus"`
}
