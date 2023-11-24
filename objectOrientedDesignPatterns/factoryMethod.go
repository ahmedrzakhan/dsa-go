package main

import "fmt"

/**
Factory Method Pattern
Implement the Factory Method design pattern.

The Factory Method is a creational design pattern that provides an interface for creating objects in a superclass but allows subclasses to alter the type of objects that will be created.

You are given code that includes a few vehicles types and their respective factories. Complete the factory method implementation such that each factory returns the correct vehicle.

Example:

VehicleFactory carFactory = new CarFactory();
VehicleFactory truckFactory = new TruckFactory();
VehicleFactory bikeFactory = new BikeFactory();

Vehicle myCar = carFactory.createVehicle();
Vehicle myTruck = truckFactory.createVehicle();
Vehicle myBike = bikeFactory.createVehicle();

myCar.getType();   // "Car"
myTruck.getType(); // "Truck"
myBike.getType();  // "Bike"

*/

// Vehicle interface
type Vehicle interface {
	GetType() string
}

// Car struct
type Car struct{}

// GetType for Car
func (c Car) GetType() string {
	return "Car"
}

// Truck struct
type Truck struct{}

// GetType for Truck
func (t Truck) GetType() string {
	return "Truck"
}

// Bike struct
type Bike struct{}

// GetType for Bike
func (b Bike) GetType() string {
	return "Bike"
}

// VehicleFactory interface
type VehicleFactory interface {
	CreateVehicle() Vehicle
}

// CarFactory struct
type CarFactory struct{}

// CreateVehicle for CarFactory
func (cf CarFactory) CreateVehicle() Vehicle {
	return Car{}
}

// TruckFactory struct
type TruckFactory struct{}

// CreateVehicle for TruckFactory
func (tf TruckFactory) CreateVehicle() Vehicle {
	return Truck{}
}

// BikeFactory struct
type BikeFactory struct{}

// CreateVehicle for BikeFactory
func (bf BikeFactory) CreateVehicle() Vehicle {
	return Bike{}
}

// main function
func mainFactory() {
	var carFactory VehicleFactory = CarFactory{}
	var truckFactory VehicleFactory = TruckFactory{}
	var bikeFactory VehicleFactory = BikeFactory{}

	myCar := carFactory.CreateVehicle()
	myTruck := truckFactory.CreateVehicle()
	myBike := bikeFactory.CreateVehicle()

	fmt.Println(myCar.GetType())   // "Car"
	fmt.Println(myTruck.GetType()) // "Truck"
	fmt.Println(myBike.GetType())  // "Bike"
}

/**
Dependency Inversion Principle (DIP)
The DIP is about high-level modules not depending on low-level modules, but both depending on abstractions. Additionally, abstractions should not depend upon details, but details should depend upon abstractions.

In the context of the Go program:

Abstractions (Interfaces): The VehicleFactory and Vehicle are interfaces representing high-level abstractions.

VehicleFactory is an abstraction for factories that create vehicles.
Vehicle is an abstraction for different types of vehicles.
High-Level and Low-Level Modules:

The high-level modules in this context are the client code or any part of the system that needs a vehicle. They depend on the Vehicle interface, not on the specific types of vehicles (like Car, Truck, or Bike).
The low-level modules are the concrete implementations of the Vehicle interface (Car, Truck, Bike) and the VehicleFactory interface (CarFactory, TruckFactory, BikeFactory).
Inversion of Control: The factories (CarFactory, TruckFactory, BikeFactory) are responsible for instantiating the concrete vehicle types. This control over instantiation is inverted from the client to the factory, which decides what type of object to create.

Open/Closed Principle
This principle states that software entities should be open for extension, but closed for modification.

Extensibility:

If you want to add another vehicle type, say, Bus, you can do so by adding a new Bus struct that implements the Vehicle interface and a BusFactory struct that implements the VehicleFactory interface. This extends the functionality without modifying existing code.
Closed for Modification:

The existing code does not need to change when a new type of vehicle is added. You don't need to alter the Vehicle or VehicleFactory interfaces or any of the existing vehicle types (Car, Truck, Bike) or their factories.
In summary, the Go program demonstrates the Dependency Inversion Principle through its use of interfaces and factory methods, ensuring that high-level modules depend on abstractions rather than concrete implementations. The Open/Closed Principle is reflected in the design that allows new types of vehicles and their factories to be added without altering existing code.

*/
