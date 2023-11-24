/**
Builder Pattern
Implement the Builder design pattern.

The Builder is a creational design pattern that allows the construction of complex objects step by step. It is useful when an object has several components and it's desirable to construct the object with various configurations.

Your task is to complete the implementation of the MealBuilder class that incrementally builds a Meal object. A Meal consists of a cost, an indication of whether it's a takeout meal, a main dish, and a drink.

Example:

MealBuilder builder = new MealBuilder();
Meal myMeal = builder.addCost(15.99)
                     .addTakeOut(true)
                     .addMainCourse("Burger")
                     .addDrink("Coke")
                     .build();

myMeal.getCost();       // Cost: 15.99
myMeal.getTakeOut();    // TakeOut: true
myMeal.getMain();       // Main: "Burger"
myMeal.getDrink();      // Drink: "Coke"
Notice that each method returns the MealBuilder instance, allowing for chaining of method calls.

*/

package main

import "fmt"

// Meal struct representing the product to be built
type Meal struct {
	cost       float64
	takeOut    bool
	mainCourse string
	drink      string
}

// MealBuilder struct for building a Meal
type MealBuilder struct {
	cost       float64
	takeOut    bool
	mainCourse string
	drink      string
}

// NewMealBuilder creates a new MealBuilder instance
func NewMealBuilder() *MealBuilder {
	return &MealBuilder{}
}

// AddCost sets the cost of the meal
func (b *MealBuilder) AddCost(cost float64) *MealBuilder {
	b.cost = cost
	return b
}

// AddTakeOut sets whether the meal is takeout
func (b *MealBuilder) AddTakeOut(takeOut bool) *MealBuilder {
	b.takeOut = takeOut
	return b
}

// AddMainCourse sets the main course of the meal
func (b *MealBuilder) AddMainCourse(mainCourse string) *MealBuilder {
	b.mainCourse = mainCourse
	return b
}

// AddDrink sets the drink of the meal
func (b *MealBuilder) AddDrink(drink string) *MealBuilder {
	b.drink = drink
	return b
}

// Build finalizes the building process and returns a Meal
func (b *MealBuilder) Build() Meal {
	return Meal{
		cost:       b.cost,
		takeOut:    b.takeOut,
		mainCourse: b.mainCourse,
		drink:      b.drink,
	}
}

func mainBuilderPattern() {
	builder := NewMealBuilder()
	myMeal := builder.AddCost(15.99).
		AddTakeOut(true).
		AddMainCourse("Burger").
		AddDrink("Coke").
		Build()

	fmt.Printf("Cost: %.2f\n", myMeal.cost)
	fmt.Printf("TakeOut: %v\n", myMeal.takeOut)
	fmt.Printf("Main: %s\n", myMeal.mainCourse)
	fmt.Printf("Drink: %s\n", myMeal.drink)
}
