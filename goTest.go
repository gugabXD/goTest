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
	Anterior       Elemento
	SequenceNumber int
	Montado        bool
	Key            bool
	Estado         GameState
}

type Elemento struct {
	Simbolo     rune
	Cor         termbox.Attribute
	CorFundo    termbox.Attribute
	Tangivel    bool
	Interagivel bool
}

type GameState struct {
	Mapa     [][]Elemento
	Msg      string
	GameOver bool
	Win      bool
}
