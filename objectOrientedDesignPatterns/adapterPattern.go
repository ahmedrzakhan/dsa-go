package main

import (
	"fmt"
)

/**
Adapter Pattern
Implement the Adapter design pattern.

The Adapter is a structural design pattern that allows incompatible interfaces to work together. It wraps an existing class with a new interface so that it becomes compatible with the client's interface.

You are given completed SquareHole, Square and Circle classes. A Square fits into a SquareHole if the Square's side length is less than or equal to the SquareHole's length. A Circle has a radius and a Circle fits into a SquareHole if the Circle's diameter is less than or equal to the SquareHole's length.

Complete the implementation of the CircleToSquareAdapter class such that it adapts a Circle to a Square.

diagram

Example 1:

SquareHole squareHole = new SquareHole(5);

Square square = new Square(5);
squareHole.canFit(square);            // true

Circle circle = new Circle(5);
CircleToSquareAdapter circleAdapter = new CircleToSquareAdapter(circle);
squareHole.canFit(circleAdapter);     // false
Example 2:

SquareHole squareHole = new SquareHole(5);

Square square = new Square(6);
squareHole.canFit(square);            // false

Circle circle = new Circle(2);
CircleToSquareAdapter circleAdapter = new CircleToSquareAdapter(circle);
squareHole.canFit(circleAdapter);     // true
*/

// Fit interface for objects that can fit into the square hole
type Fit interface {
	GetSide() float64
}

// SquareHole struct
type SquareHole struct {
	length float64
}

// NewSquareHole creates a new SquareHole
func NewSquareHole(length float64) *SquareHole {
	return &SquareHole{length: length}
}

// CanFit checks if an object implementing the Fit interface can fit into the hole
func (h *SquareHole) CanFit(fit Fit) bool {
	return fit.GetSide() <= h.length
}

// Square struct
type Square struct {
	side float64
}

// NewSquare creates a new Square
func NewSquare(side float64) *Square {
	return &Square{side: side}
}

// GetSide returns the side length of the square (Fit interface implementation)
func (s *Square) GetSide() float64 {
	return s.side
}

// Circle struct
type Circle struct {
	radius float64
}

// NewCircle creates a new Circle
func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

// CircleToSquareAdapter struct
type CircleToSquareAdapter struct {
	circle *Circle
}

// NewCircleToSquareAdapter creates a new CircleToSquareAdapter
func NewCircleToSquareAdapter(circle *Circle) *CircleToSquareAdapter {
	return &CircleToSquareAdapter{circle: circle}
}

// GetSide adapts the circle to act like a square (Fit interface implementation)
func (adapter *CircleToSquareAdapter) GetSide() float64 {
	// The diameter of the circle is considered the side of the square
	return adapter.circle.radius * 2
}

func main() {
	// Example 1
	squareHole := NewSquareHole(5)

	square := NewSquare(5)
	fmt.Println("Square fits:", squareHole.CanFit(square)) // true

	circle := NewCircle(5)
	circleAdapter := NewCircleToSquareAdapter(circle)
	fmt.Println("Circle fits:", squareHole.CanFit(circleAdapter)) // false

	// Example 2
	squareHole = NewSquareHole(5)

	square = NewSquare(6)
	fmt.Println("Square fits:", squareHole.CanFit(square)) // false

	circle = NewCircle(2)
	circleAdapter = NewCircleToSquareAdapter(circle)
	fmt.Println("Circle fits:", squareHole.CanFit(circleAdapter)) // true
}
