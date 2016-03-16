package main

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Person struc to define new users.
type Person struct {
	Name  string
	Age   int
	phone int
}

// GetPeople Returns all the person(s) in the Person collection.
func GetPeople() []Person {

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	connect := session.DB("people").C("People")

	var persons []Person

	err = connect.Find(bson.M{}).All(&persons)

	fmt.Println("All :", &persons)

	return persons
}

// AddPerson writes new person instance into DB collection.
func AddPerson(name string, age, phone int) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	connect := session.DB("people").C("Peoples")

	err = connect.Insert(&Person{name, age, phone})
	if err != nil {
		panic(err)
	}
	fmt.Println("New person(s) added..")
}

// UpdatePerson updates an already existing record in DB.
// ToDO
func UpdatePerson(name string, age, phone int) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	//connect := session.DB("people").C("Peoples")

	/*if name == nil {
		err = connect.Update(bson.M{"name": name}, bson.M{"age": age, "phone": phone})
		if err != nil {
			log.Fatal(err)
		}
	} //if ends*/
}

// FindPerson returns a slice of the match person(s) from the person collection.
func FindPerson(name string) *Person {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	connect := session.DB("people").C("Peoples")

	persons := &Person{}

	err = connect.Find(bson.M{"name": name}).One(persons)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Person Found: ", persons)

	return persons

}

// RemovePerson removes person(s) from the person collection
func RemovePerson(name string) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	defer session.Close()

	connect := session.DB("people").C("Peoples")

	err = connect.Remove(bson.M{"name": name})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Person(s) removed..")
}

func main() {

	name := "saheed"
	//age := 27
	//phone := 802222222281

	//AddPerson(name, age, phone)

	FindPerson(name)
	GetPeople()
}
