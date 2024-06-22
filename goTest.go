package goTest

type Mensagem struct {
	IdClient int
	Comando  rune
}

type Client struct {
	Id       int
	Anterior rune
}
