package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

func ReadFile(name string) []byte {
	file, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
	}

	defer file.Close()
	data, err := io.ReadAll(file) // data: []byte
	if err != nil {
		fmt.Println(err)
	}
	return data
}

type JsonDecoder interface {
	ReadJson([]byte) error
	SaveToJson([]byte) error
}

type Base struct {
	Offers
	Employees
}

func (b *Base) AddEmployeeOrOffer() {
	fmt.Println("Type 0 if you want to add Employee or 1 if Offer")
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	text := in.Text()
	if text == "0" {
		fmt.Print("name: ")
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		name := in.Text()

		fmt.Print("age: ")
		in = bufio.NewScanner(os.Stdin)
		in.Scan()
		age, err := strconv.Atoi(in.Text())
		if err != nil {
			fmt.Println("Incorrect value!")
			return
		}

		fmt.Print("education: ")
		in = bufio.NewScanner(os.Stdin)
		in.Scan()
		education := in.Text()

		emp := Employee{name, age, education}
		b.Employees.Employees = append(b.Employees.Employees, emp)
	} else if text == "1" {
		fmt.Print("name: ")
		in := bufio.NewScanner(os.Stdin)
		in.Scan()
		name := in.Text()

		fmt.Print("from (format: 2006-01-02):")
		in = bufio.NewScanner(os.Stdin)
		in.Scan()
		from, err := time.Parse("2006-01-02", in.Text())
		if err != nil {
			fmt.Println("Incorrect value!")
			return
		}

		fmt.Print("to (format: 2006-01-02):")
		in = bufio.NewScanner(os.Stdin)
		in.Scan()
		to, err := time.Parse("2006-01-02", in.Text())
		if err != nil {
			fmt.Println("Incorrect value!")
			return
		}

		fmt.Print("education: ")
		in = bufio.NewScanner(os.Stdin)
		in.Scan()
		education := in.Text()

		off := Offer{name, from, to, education}
		b.Offers.Offers = append(b.Offers.Offers, off)

	}
}

func (b Base) Print() {
	fmt.Println("\t\t\t\tBase:")
	b.Employees.Print()
	b.Offers.Print()
}

func (b Base) PrintOffersForEmployees() {
	for i := range b.Employees.Employees {
		fmt.Println()
		fmt.Println("Offers for:", b.Employees.Employees[i].Name)
		for j := range b.Offers.Offers {
			if b.Offers.Offers[j].Education == b.Employees.Employees[i].Education {
				b.Offers.Offers[j].Print()
				fmt.Println("\t\tcost:", b.Offers.Offers[j].CountCost())
			}
		}
	}
	fmt.Println()
}

func (b *Base) initBase() {
	offers := Offers{[]Offer{}}
	offersData := ReadFile("Offers.json")
	err := offers.ReadJson(offersData)
	if err != nil {
		fmt.Println(err)
	}

	employees := Employees{[]Employee{}}
	employeesData := ReadFile("People.json")
	err = employees.ReadJson(employeesData)
	if err != nil {
		fmt.Println(err)
	}

	b.Offers = offers
	b.Employees = employees
}

func (b Base) SaveToFile() {
	emp := b.Employees.SaveToJson()
	err := os.WriteFile("People.json", emp, 0644)
	if err != nil {
		fmt.Println(err)
	}
	off := b.Offers.SaveToJson()
	err = os.WriteFile("Offers.json", off, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

type Offers struct {
	Offers []Offer `json:"offers"`
}

func (o Offers) Print() {
	fmt.Println("Offers:")
	for i := range o.Offers {
		o.Offers[i].Print()
	}
	fmt.Println()
}

func (o *Offers) ReadJson(b []byte) error {
	err := json.Unmarshal(b, o)
	return err
}

func (o *Offers) SaveToJson() []byte {
	b, err := json.Marshal(o)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

type Offer struct {
	Name      string    `json:"name"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
	Education string    `json:"education"`
}

func (o Offer) Print() {
	fmt.Println("\t", o.Name, ":", o.From, "-", o.To, "|", o.Education)
}

func (o Offer) CountCost() int {
	diff := int(o.To.Sub(o.From).Hours() / 24 / 30)
	if o.Education == "h" {
		return diff * 5000
	} else if o.Education == "e" {
		return diff * 8000
	}
	return 0

}

type Employees struct {
	Employees []Employee `json:"people"`
}

func (e Employees) Print() {
	fmt.Println("Employees:")
	for i := range e.Employees {
		e.Employees[i].Print()
	}
	fmt.Println()
}

func (e *Employees) ReadJson(b []byte) error {
	err := json.Unmarshal(b, e)
	return err
}
func (e *Employees) SaveToJson() []byte {
	b, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

type Employee struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	Education string `json:"education"`
}

func (e Employee) Print() {
	fmt.Println("\t", e.Name, ":", e.Age, "|", e.Education)
}

func main() {
	base := Base{Offers{}, Employees{}}
	base.initBase()
	base.AddEmployeeOrOffer()
	base.Print()
	fmt.Println()
	base.PrintOffersForEmployees()

	base.SaveToFile()
}
