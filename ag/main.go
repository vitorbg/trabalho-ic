package main

import (
	"fmt"
	//"github.com/vitorbg/trabalho-ic/ag/domain"
	"math/rand"
	"time"
)

const POPULATION_SIZE = 10
const INDIVIDUAL_SIZE = 12
const POS_FITNESS = 10


var currentPopulation [POPULATION_SIZE][11]int
var nextPopulation [POPULATION_SIZE][11]int

func main() {

	fmt.Println("Galoo")
	initPopulation()
	printPopulation()


}

func initPopulation() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(rand.Intn(10))

	var i, j, pos, temp int

	for i = 0; i < POPULATION_SIZE; i++ {
		for j = 0; j < 10; j++ {
			currentPopulation[i][j] = j
		}
	}

	for i = 0; i < POPULATION_SIZE; i++ {
		for j = 0; j < 8; j++ {
			pos = rand.Intn(10)
			temp = currentPopulation[i][j]
			currentPopulation[i][j] = currentPopulation[i][pos]
			currentPopulation[i][pos] = temp
		}
	}
}

func printPopulation() {
	for i := 0; i < 10; i++ {
		fmt.Println("Individuo ", i)
		fmt.Println(currentPopulation[i])
	}
}
