package goTest

type Mensagem struct {
	IdClient       int
	Comando        rune
	SequenceNumber int
}

type Client struct {
	Id             int
	PosX, PosY     int
	Anterior       rune
	SequenceNumber int
}
