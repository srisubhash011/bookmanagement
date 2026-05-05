package servicesmodel

type Book struct {
	BookId          int    `json:"bookid"`
	BookName        string `json:"bookname"`
	BookDescription string `json:"bookdescription"`
}

type BookSearch struct {
	BookName string `json:"bookname"`
}
