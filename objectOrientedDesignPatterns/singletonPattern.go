package main

import "fmt"

/**
Singleton Pattern
Implement the Singleton design pattern.

The Singleton is a creational design pattern which ensures that at most one instance of a class may exist.

Singleton getInstance() will return the singleton instance.
String getValue() will return the value of the singleton.
void setValue(String value) will update the value of the singleton.
You can assume the singleton will only be used in a single-threaded environment.

Example:

Singleton s = Singleton.getInstance();

s.getValue(); // null

s.setValue("a value string");
s.getValue(); // "a value string"

Singleton s2 = Singleton.getInstance();

s2.getValue(); // "a value string"
*/

// Singleton struct
type Singleton struct {
	value string
}

// instance holds the singleton instance
var instance *Singleton

// getInstance returns the singleton instance
func getInstance() *Singleton {
	if instance == nil {
		instance = &Singleton{}
	}
	return instance
}

// getValue returns the value of the singleton
func (s *Singleton) getValue() string {
	return s.value
}

// setValue updates the value of the singleton
func (s *Singleton) setValue(value string) {
	s.value = value
}

func mainSingleTon() {
	s := getInstance()
	fmt.Println(s.getValue()) // should print an empty string or "null"

	s.setValue("a value string")
	fmt.Println(s.getValue()) // should print "a value string"

	s2 := getInstance()
	fmt.Println(s2.getValue()) // should also print "a value string"
}
