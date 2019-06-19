package main

import "fmt"

//A method is nothing more than a FUNC attached to a TYPE
//Metóda nie je nič iné ako FUNC pripojené k TYPU
type Person struct{
	name string
	surname string
}

type PersonalInfo struct{
	person Person
	info string
}

func(p PersonalInfo) giveInfo(){
	fmt.Println("Personal infos are : ", p.person.name, p.person.surname, p.info)
}

func(person Person) getPersonInfo(){
	fmt.Println("Person info is : ", person.name, person.surname)
}


func main(){
	pi := PersonalInfo{
		person:Person{
			"Roman",
			"Ceresnak",
		},
		info : "It is dush",
	}

	person := Person{
		"Roman",
		"Nomeseuru",
	}


	pi.giveInfo()
	person.getPersonInfo()
}
