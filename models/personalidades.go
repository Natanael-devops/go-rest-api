package models

type Personalidade struct {
	Nome     string `json:"nome"`
	Historia string `jason:"historia"`
}

var Personalidades []Personalidade
