package goTest

type Mensagem struct {
	IdClient           int
	NovaPosX, NovaPosY int
	PosX, PosY         int
	Anterior           rune
	Comando            rune
}

type Client struct {
	Id         int
	PosX, PosY int
}
