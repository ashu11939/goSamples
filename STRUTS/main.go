package main 

import "fmt"

type contactInfo struct {
	email string
	zip int
}

type person struct {
	firstName string
	lastName string
	contact contactInfo
}

func main() {
	alex := person{
		firstName:"Alex", 
		lastName:"Anderson",
		contact: contactInfo {
			email: "alex@gmail.com",
			zip: 411014,
		},
	}
	var ashu person //GO assigns a zero value by default
	ashu.firstName = "Ashutosh"
	ashu.lastName = "Kumar Gupta"
	ashu.contact.email = "ashutosh11939@gmail.com"
	ashu.contact.zip = 412207
	fmt.Println(alex)
	ashu.print()

	ashuPointer := &ashu //to a pointer 
	println("memory address : ", ashuPointer)

	ashu.updateName("Gunjan")
	ashu.print()
}

func (p person) print() {
	fmt.Println(p)
}

//Receiver : expecting a pointer as input
//GO benefits : Automatically turn a value to pointer if we receive by *type
func (p *person) updateName(newFirstName string) {
	//This is used to manipulate the value of pointer
	p.firstName = newFirstName
}

/**
	Turn a value into address --> &value
	Turn a pointer into value --> *pointer
*/