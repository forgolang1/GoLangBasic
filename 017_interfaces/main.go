package main

import "fmt"

/*

 Полиморфизм
1 - Возможность объектов с одинаковой спецификацией иметь различную реализацию.

2 - Способность обьекта использовать методы производного класса
	который не существует на момент создания базового.

3 - Это свойство классов иметь одинаковые методы,
	которые будут работать по-разному в контексте объектов.

 */

type Document struct {
	Date string
	Number string
	NumberOfPages int

}

type PersonCard struct {
	Date string
	FirstName string
	LastName string
	Age int
}

type PrintInterface interface {
	CheckData()
}

func (d *Document) CheckData() {
	//реализация функции интерфейса может быть произвольной для объекта структуры
	//главное соответсвие сигнатуре функции интерфейса
	if d.Date != "" {
		fmt.Println("Date in the doc - correct")
	} else {
		fmt.Println("Date in the doc - not correct")
	}
}

func (p *PersonCard) CheckData() {
	if p.Date != "" {
		fmt.Println("Date in the person card - correct")
	} else {
		fmt.Println("Date in the person card - not correct")
	}
}


func main() {

	doc := new(Document)
	pcard := new(PersonCard)

	doc.Date = "1.10.2018"
	doc.NumberOfPages = 5
	doc.Number = "A - 100"

	pcard.Date = "1.10.2018"
	pcard.Age = 21
	pcard.FirstName = "User"
	pcard.LastName = "Test"

	sl := []PrintInterface{doc, pcard}

	PrintAnyDoc(sl)

}

func PrintAnyDoc(anyDoc []PrintInterface) {
	for _, v := range anyDoc {
		fmt.Println(v)
	}
}


