package api

type Response struct {
	Request Request
	Book    Book `xml:"book"`
}
