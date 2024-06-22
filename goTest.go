package goTest

import "github.com/nsf/termbox-go"

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

type Elemento struct {
	simbolo     rune
	cor         termbox.Attribute
	corFundo    termbox.Attribute
	tangivel    bool
	interagivel bool
}
