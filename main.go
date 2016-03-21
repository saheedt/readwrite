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

//==============================================================================
// If you are connecting to a mgo database, you don't need to always redial
// else you are creating multiple connections, a mongo session can be used
// multiple times as many as you want. Hence I will only create it once then
// get the session to do what i want in the other functions.

// ses provides internal global session for our database operation,we will
// create new sessions of this one, once we have successfully connected.
var ses *mgo.Session

func getSession() *mgo.Session {
	if ses != nil {
		// We do a copy because we want to have a unique session from what
		// another might be using but we still want to multiplex on the same
		// connection underneath it, its much cleaner.
		return ses.Copy()
	}

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}

	// We are setting the mod of operation to Monotonic to ensure we get
	// assurances that our writes and reads are synchornized and we not
	// cause inconsistencies.
	// NOTE: Never forget to do this.
	session.SetMode(mgo.Monotonic, true)

	ses = session
	return session
}

//==============================================================================

// GetPeople Returns all the person(s) in the Person collection.
func GetPeople() []Person {

	session := getSession()
	defer session.Close()

	connect := session.DB("people").C("People")

	var persons []Person

	// TODO: Fix your problem here, why did you not check if there was an error
	// and why did you not log it.
	err := connect.Find(bson.M{}).All(&persons)

	fmt.Println("All :", &persons)

	return persons
}

// AddPerson writes new person instance into DB collection.
func AddPerson(name string, age, phone int) {

	session := getSession()
	defer session.Close()

	connect := session.DB("people").C("Peoples")

	err := connect.Insert(&Person{name, age, phone})

	// TODO: Do you really think a panic is appropriate here, if am running this
	// on production server, my whole app will die because on db query failed,
	// those that make sense to you? O_O
	if err != nil {
		panic(err)
	}

	fmt.Println("New person(s) added..")
}

// UpdatePerson updates an already existing record in DB.
// ToDO
func UpdatePerson(name string, age, phone int) {
	session := getSession()
	defer session.Close()

	// TODO: why did you not run your update, learn to work with it
	// getting used to this operation is important.
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

	// Get a session.
	session := getSession()
	defer session.Close()

	connect := session.DB("people").C("Peoples")

	// NOTE: why not do a
	// var person Person.
	// When declaring variables you do plan to initialize with a value
	// use the declaration not the assignative form.
	persons := &Person{}

	err := connect.Find(bson.M{"name": name}).One(persons)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Person Found: ", persons)

	return persons

}

// RemovePerson removes person(s) from the person collection
func RemovePerson(name string) {
	// Get a session.
	session := getSession()
	defer session.Close()

	connect := session.DB("people").C("Peoples")

	err := connect.Remove(bson.M{"name": name})
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

	//TODO: What about Delete and Update. Its called CRUD for a reason:
	// CRUD = CREATE READ UPDATE DELETE.
	// I want to see all this operations running.
}
