package main

import "fmt"

//Abstract factory for adidas and nike shirts and tenis
func main() {
	adidasFactory := getSportsFactory("adidas")
	nikeFactory := getSportsFactory("nike")
	adidasShirt := adidasFactory.makeShirt()
	nikeShirt := nikeFactory.makeShirt()
	fmt.Println(adidasShirt.getLogo())
	fmt.Println(nikeShirt.getLogo())
}

//1 crear producto concreto genérico
type shirt struct {
	size int
	logo string
}

//1.1 crear producto concreto a partir del concreto genérico
type adidasShirt struct {
	shirt
}

type nikeShirt struct {
	shirt
}

//2 Crear productos abstractos basado en productos concretos genéricos
type iShirt interface {
	setSize(int)
	setLogo(string)
	getSize() int
	getLogo() string
}

func (s *shirt) setSize(size int) {
	s.size = size
}

func (s *shirt) setLogo(logo string) {
	s.logo = logo
}

func (s *shirt) getSize() int {
	return s.size
}

func (s *shirt) getLogo() string {
	return s.logo
}

//3 crear fabrica abstracta que regresa productos abstractos

type sportsFactory interface {
	makeShirt() iShirt
}

func getSportsFactory(name string) sportsFactory {
	if name == "adidas" {
		return &adidasFactory{}
	}
	if name == "nike" {
		return &nikeFactory{}
	}
	return nil
}

//4 crear la fabrica concreta

type nikeFactory struct {
}

func (n *nikeFactory) makeShirt() iShirt {
	return &nikeShirt{
		shirt: shirt{logo: "nike",
			size: 14,
		},
	}
}

type adidasFactory struct {
}

func (n *adidasFactory) makeShirt() iShirt {
	return &adidasShirt{
		shirt: shirt{logo: "adidas",
			size: 14,
		},
	}
}
