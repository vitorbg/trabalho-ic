package main

import (
	"fmt"
	//"math"
	"math/rand"
	"time"
)

const POPULATION_SIZE = 10
const INDIVIDUAL_SIZE = 12
const POS_FITNESS = 10

const (
	D = iota // 0
	E        // 1
	M        // 2
	N
	O
	R
	S
	Y
)

var currentPopulation [POPULATION_SIZE][INDIVIDUAL_SIZE]int
var nextPopulation [POPULATION_SIZE][INDIVIDUAL_SIZE]int

func main() {

	fmt.Println("Galoo")
	initPopulation()
	printPopulation()
	fmt.Println("Olha o fitnes")
	calcIndividualFitness(&currentPopulation[0])
	calcPopulationFitness()
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

func calcPopulationFitness() {

	for i := 0; i < POPULATION_SIZE; i++ {
		for j := 0; j < INDIVIDUAL_SIZE; j++ {
			calcIndividualFitness(&currentPopulation[i])
		}
	}

}

func calcIndividualFitness(individual *[INDIVIDUAL_SIZE]int) {

	var send, more, money int
	send += individual[S] * 1000
	send += individual[E] * 100
	send += individual[N] * 10
	send += individual[D]

	more += individual[M] * 1000
	more += individual[O] * 100
	more += individual[R] * 10
	more += individual[E]

	money += individual[M] * 10000
	money += individual[O] * 1000
	money += individual[N] * 100
	money += individual[E] * 10
	money += individual[Y]

	fitness := (send + more) - money
	if fitness < 0 {
		fitness = fitness * (-1)
	}
	individual[10] = fitness

	fmt.Println(individual)
	fmt.Println("SEND: ", send)
	fmt.Println("MORE: ", more)
	fmt.Println("MONEY: ", money)
	fmt.Println("SEND+MORE: ", send+more)
	fmt.Println("(SEND+MORE)-MONEY: ", fitness)
	fmt.Println("D:", individual[D])
	fmt.Println("E:", individual[E])
	fmt.Println("M:", individual[M])
	fmt.Println("N:", individual[N])
	fmt.Println("O:", individual[O])
	fmt.Println("R:", individual[R])
	fmt.Println("S:", individual[S])
	fmt.Println("Y:", individual[Y])
}
