package goTest

type Mensagem struct {
	IdClient       int
	Comando        rune
	sequenceNumber int
}

type Client struct {
	Id       int
	Anterior rune
}
