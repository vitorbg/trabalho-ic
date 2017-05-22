package main

import (
	"fmt"
	// "math"
	// "bufio"
	"math/rand"
	// "os"
	"time"
)

const CROSS_OVER_RATIO = 0.8
const MUTATION_RATIO = 0.3
const ELITISM_RATIO = 0.2

const POPULATION_PARENTS_SIZE = 100
const POPULATION_CHILDREN_SIZE = POPULATION_PARENTS_SIZE * CROSS_OVER_RATIO
const POPULATION_TOTAL_SIZE = POPULATION_PARENTS_SIZE + POPULATION_CHILDREN_SIZE

const ELITISM_QT_PARENTS = int(POPULATION_PARENTS_SIZE * ELITISM_RATIO)
const MUTATION_QT_INDIVIDUALS = int(POPULATION_CHILDREN_SIZE * MUTATION_RATIO)
const CROSS_OVER_QT = int(POPULATION_CHILDREN_SIZE / 2)

const VECTOR_SIZE = 13
const INDIVIDUAL_SIZE = 10
const POS_EVAL = 10
const POS_FITNESS = 11
const POS_ACC_ROULLETE = 12

const TOUR = 3
const NUM_GENERATION = 200
const NUM_EXECUTION = 1000

const S1_ROULETTE = 1
const S2_TOURNAMENT = 2
const R1_REINSERTION_ORDERLY = 1
const R2_REINSERTION_PURE_ELITISM = 2

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

var population [POPULATION_TOTAL_SIZE][VECTOR_SIZE]int

func main() {
	selection := S1_ROULETTE
	//  selection := S2_TOURNAMENT
	// reinsertion := R1_REINSERTION_ORDERLY
	reinsertion := R2_REINSERTION_PURE_ELITISM
	fmt.Println("Globals Configurations ----------------------------------------")
	fmt.Println("NUM_EXECUTION: ", NUM_EXECUTION)
	fmt.Println("NUM_GENERATION: ", NUM_GENERATION)
	fmt.Println("POPULATION_PARENTS_SIZE: ", POPULATION_PARENTS_SIZE)
	fmt.Println("POPULATION_CHILDREN_SIZE: ", POPULATION_CHILDREN_SIZE)
	fmt.Println("POPULATION_TOTAL_SIZE: ", POPULATION_TOTAL_SIZE)
	fmt.Println("CROSS_OVER_RATIO: ", CROSS_OVER_RATIO)
	fmt.Println("CROSS_OVER_QT: ", CROSS_OVER_QT)
	fmt.Println("MUTATION_RATIO: ", MUTATION_RATIO)
	fmt.Println("MUTATION_QT_INDIVIDUALS: ", MUTATION_QT_INDIVIDUALS)
	if selection == S1_ROULETTE {
		fmt.Println("SELECTION METHOD: ROULETTE")
	}
	if selection == S2_TOURNAMENT {
		fmt.Println("SELECTION METHOD: TOURNAMENT - TOUR OF ", TOUR)
	}
	if reinsertion == R1_REINSERTION_ORDERLY {
		fmt.Println("REINSERTION METHOD: ORDELY")
	}
	if reinsertion == R2_REINSERTION_PURE_ELITISM {
		fmt.Println("REINSERTION METHOD: PURE ELITISM")
		fmt.Println("ELITISM_RATIO: ", ELITISM_RATIO)
		fmt.Println("ELITISM_QT_PARENTS: ", ELITISM_QT_PARENTS)
	}
	fmt.Println("---------------------------------------------------------------")

	var qt_generation int
	qt_find_optimal := 0
	timeGeneration := time.Now()
	for i := 0; i < NUM_EXECUTION; i++ {
		initPopulation()
		calcPopulationFitness()
		//printPopulation()
		for j := 0; j < NUM_GENERATION; j++ {
			// fmt.Println("GENERATION ", j)
			//SELECTION AND CROSS OVER PHASE *****************************************
			index := 0
			for k := 0; k < CROSS_OVER_QT; k++ {
				var pos_parent_1, pos_parent_2 int
				if selection == S1_ROULETTE {
					pos_parent_1, pos_parent_2 = selectionRoullete()
				}
				if selection == S2_TOURNAMENT {
					pos_parent_1, pos_parent_2 = selectionTournament()
				}

				child_1, child_2 := calcCrossOverCycle(population[pos_parent_1], population[pos_parent_2])

				population[POPULATION_PARENTS_SIZE+index] = child_1
				population[POPULATION_PARENTS_SIZE+index+1] = child_2
				index = index + 2
			}
			//************************************************************************
			//printPopulation()
			//REINSERTION PHASE ******************************************************
			if reinsertion == R1_REINSERTION_ORDERLY {
				reinsertionOrderly()
			}
			if reinsertion == R2_REINSERTION_PURE_ELITISM {
				reinsertionPureElitism()
			}
			//************************************************************************
			//MUTATION PHASE *********************************************************
			mutation()
			//************************************************************************
			// pos_optimal, ok := calcPopulationFitness()
			_, ok := calcPopulationFitness()
			if ok {
				// fmt.Print("Execution: [", i)
				// fmt.Print("] Generation: [", j)
				// fmt.Println("] Solution: ", population[pos_optimal])
				qt_generation = qt_generation + j
				qt_find_optimal++
				j = NUM_GENERATION + 2
			}
			// printPopulation()
		}
	}

	time2 := time.Now()
	time2.Sub(timeGeneration)

	fmt.Println("Metrics -------------------------------------------------------")
	// fmt.Println("Qt Generations Total: ", qt_generation)
	fmt.Println("Generation Average: ", qt_generation/NUM_EXECUTION)
	fmt.Println("Convertion Ratio: ", float32(qt_find_optimal)/float32(NUM_EXECUTION))
	fmt.Println("Execution Total Time: ", time2.Sub(timeGeneration))
	fmt.Println("Execution Generation Time: ", (time2.Sub(timeGeneration))/NUM_EXECUTION)
	fmt.Println("---------------------------------------------------------------")

	//	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

func initPopulation() {
	//rand.Seed(1)
	rand.Seed(time.Now().UTC().UnixNano())

	var i, j, pos, temp int

	for i = 0; i < POPULATION_PARENTS_SIZE; i++ {
		for j = 0; j < 10; j++ {
			population[i][j] = j
		}
	}
	for i = 0; i < POPULATION_TOTAL_SIZE; i++ {
		for j = 0; j < INDIVIDUAL_SIZE; j++ {
			pos = rand.Intn(INDIVIDUAL_SIZE)
			temp = population[i][j]
			population[i][j] = population[i][pos]
			population[i][pos] = temp
		}
	}
}

func validateIndividual(individual [VECTOR_SIZE]int) bool {
	var vector [INDIVIDUAL_SIZE]int
	for i := 0; i < INDIVIDUAL_SIZE; i++ {
		vector[individual[i]]++
		if vector[individual[i]] > 1 {
			return false
		}
	}

	return true
}

func calcCrossOverCycle(individual_1 [VECTOR_SIZE]int, individual_2 [VECTOR_SIZE]int) ([VECTOR_SIZE]int, [VECTOR_SIZE]int) {
	child_1 := individual_1
	child_2 := individual_2

	var val, aux int
	pos_ini := rand.Intn(INDIVIDUAL_SIZE)
	// fmt.Println("BEFORE CROSS OVER ")
	// fmt.Println("child_1: ", child_1)
	// fmt.Println("child_2: ", child_2)
	// fmt.Println("pos_ini: ", pos_ini)

	val = child_2[pos_ini]

	aux = child_1[pos_ini]
	child_1[pos_ini] = child_2[pos_ini]
	child_2[pos_ini] = aux
	find := true
	for find == true {
		if validateIndividual(child_1) {
			find = false
		} else {
			for i := 0; i < INDIVIDUAL_SIZE; i++ {
				if child_1[i] == val && i != pos_ini {
					aux = child_1[i]
					child_1[i] = child_2[i]
					child_2[i] = aux
					val = child_1[i]
					pos_ini = i
				}
			}
		}
	}
	// fmt.Println("AFTER CROSS OVER ")
	// fmt.Println("child_1: ", child_1)
	// fmt.Println("child_2: ", child_2)
	return child_1, child_2
}

func selectionRoullete() (int, int) {
	QuickSort(&population, true)
	population[0][POS_ACC_ROULLETE] = population[0][POS_FITNESS]

	for i := 1; i < POPULATION_PARENTS_SIZE; i++ {
		population[i][POS_ACC_ROULLETE] = population[i-1][POS_ACC_ROULLETE] + population[i][POS_FITNESS]
	}

	sorted := rand.Intn(population[POPULATION_PARENTS_SIZE-1][POS_ACC_ROULLETE])
	// fmt.Println("Sorted_1: ", sorted)
	var pos_winner_1 int
	var pos_winner_2 int

	for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
		if sorted < population[i][POS_ACC_ROULLETE] {
			pos_winner_1 = i
			i = POPULATION_PARENTS_SIZE + 2
		}
	}
	sorted = rand.Intn(population[POPULATION_PARENTS_SIZE-1][POS_ACC_ROULLETE])
	// fmt.Println("Sorted_2: ", sorted)
	for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
		if sorted < population[i][POS_ACC_ROULLETE] {
			pos_winner_2 = i
			i = POPULATION_PARENTS_SIZE + 2
		}
	}
	// fmt.Println("pos_winner_1: ", pos_winner_1)
	// fmt.Println("pos_winner_2: ", pos_winner_2)

	return pos_winner_1, pos_winner_2
}

func selectionTournament() (int, int) {
	pos_parent_1 := getParentTour()
	pos_parent_2 := getParentTour()

	return pos_parent_1, pos_parent_2
}

func getParentTour() int {
	var winner_pos int
	var candidate_pos int

	winner_pos = rand.Intn(POPULATION_PARENTS_SIZE)

	for i := 1; i < TOUR; i++ {
		candidate_pos = rand.Intn(POPULATION_PARENTS_SIZE)
		// fmt.Println("n sorted: ", n, population[n])
		if population[candidate_pos][POS_EVAL] < population[winner_pos][POS_EVAL] {
			winner_pos = candidate_pos
		}
	}
	// fmt.Println("winner: ",winner)
	return winner_pos
}

func reinsertionOrderly() {
	QuickSort(&population, false)
}

func reinsertionPureElitism() {
	QuickSort(&population, true)
	j := 0
	for i := ELITISM_QT_PARENTS; i < POPULATION_PARENTS_SIZE; i++ {
		population[i] = population[POPULATION_PARENTS_SIZE+j]
		j++
	}
}

func mutation() {
	for i := 0; i < int(MUTATION_QT_INDIVIDUALS); i++ {
		pos_individual := rand.Intn(POPULATION_PARENTS_SIZE)
		pos_1 := rand.Intn(INDIVIDUAL_SIZE)
		pos_2 := rand.Intn(INDIVIDUAL_SIZE)
		// fmt.Println("Mutation Individual: ", pos_individual+POPULATION_PARENTS_SIZE)
		// fmt.Println("Pos_1: ", pos_1)
		// fmt.Println("Pos_2: ", pos_2)
		// fmt.Println("Individual before mutation: ", population[pos_individual])
		aux := population[pos_individual][pos_1]
		population[pos_individual][pos_1] = population[pos_individual][pos_2]
		population[pos_individual][pos_2] = aux
		// fmt.Println("Individual after mutation: ", population[pos_individual])
	}
}

func printPopulation() {
	for i := 0; i < POPULATION_TOTAL_SIZE; i++ {
		fmt.Print("Individuo [", i)
		fmt.Println("]: ", population[i])
	}
}

func calcPopulationFitness() (int, bool) {
	for i := 0; i < POPULATION_TOTAL_SIZE; i++ {
		calcIndividualFitness(&population[i])
		if population[i][POS_EVAL] == 0 {
			return i, true
		}
	}
	return -1, false
}

func calcIndividualFitness(individual *[VECTOR_SIZE]int) {

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
	individual[POS_EVAL] = fitness
	individual[POS_FITNESS] = 100000 - fitness

	// fmt.Println(individual)
	// fmt.Println("SEND: ", send)
	// fmt.Println("MORE: ", more)
	// fmt.Println("MONEY: ", money)
	// fmt.Println("SEND+MORE: ", send+more)
	// fmt.Println("(SEND+MORE)-MONEY: ", fitness)
	// fmt.Println("D:", individual[D])
	// fmt.Println("E:", individual[E])
	// fmt.Println("M:", individual[M])
	// fmt.Println("N:", individual[N])
	// fmt.Println("O:", individual[O])
	// fmt.Println("R:", individual[R])
	// fmt.Println("S:", individual[S])
	// fmt.Println("Y:", individual[Y])
}

func QuickSort(values *[POPULATION_TOTAL_SIZE][VECTOR_SIZE]int, elitism bool) {
	if elitism {
		sort(values, 0, POPULATION_PARENTS_SIZE-1)
	} else {
		sort(values, 0, POPULATION_TOTAL_SIZE-1)
	}
	// sort(values, 0, len(values)-1)
}

func sort(values *[POPULATION_TOTAL_SIZE][VECTOR_SIZE]int, l int, r int) {
	if l >= r {
		return
	}

	pivot := values[l]
	i := l + 1

	for j := l; j <= r; j++ {
		if pivot[POS_EVAL] > values[j][POS_EVAL] {
			values[i], values[j] = values[j], values[i]
			i++
		}
	}

	values[l], values[i-1] = values[i-1], pivot

	sort(values, l, i-2)
	sort(values, i, r)
}
