package main

import (
	"fmt"
	"log"
	"task3/shape"
)

func main() {
	shapes := []shape.Shape{
		inputRectangle(),
		inputCircle(),
	}
	printShareSquare(shapes)
}

func inputRectangle() shape.Shape {
	var length float64
	var weight float64
	fmt.Println("Введите длину и ширину фигуры: ")
	_, err := fmt.Scanln(&length, &weight)
	if err != nil {
		log.Fatal("Ошибка при вводе")
	}
	return shape.NewRectangle(length, weight)
}

func inputCircle() shape.Shape {
	var radius float64
	fmt.Println("Введите радиус окружности: ")
	_, err := fmt.Scanln(&radius)

	if err != nil {
		log.Fatal("Ошибка при вводе")
	}
	return shape.NewCircle(radius)
}

func printShareSquare(shapes []shape.Shape) {
	for _, sh := range shapes {
		switch v := sh.(type) {
		case shape.Rectangle:
			fmt.Println("Площать прямоугольника", v.Area())
		case shape.Circle:
			fmt.Println("Площать окружности", v.Area())
		default:
			fmt.Println("unknown")
		}
	}

}
