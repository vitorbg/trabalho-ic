package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	EQUAL            = iota //0
	DIFFERENT               //1
	GREATER_OR_EQUAL        //2
	LESS                    //3
)

const TRAINING_DATABASE = true
const RECORD_QUANTITY = 214 //234
const RECORD_QUANTITY_TRAINING = 39

// const TRAINING_DATABASE = false
// const RECORD_QUANTITY = 114
// const RECORD_QUANTITY_TRAINING = 19

const RECORD_SIZE = 35
const RECORD_POS_CLASS = 34
const RECORD_TOTAL_DATABASE = 358

const CROSS_OVER_RATIO = 1

const POS_FAMILY_HISTORY = 10 //11: family history, (0 or 1)
const POS_AGE = 33            //34: Age (linear)

const POPULATION_PARENTS_SIZE = 50
const POPULATION_CHILDREN_SIZE = POPULATION_PARENTS_SIZE * CROSS_OVER_RATIO
const POPULATION_TOTAL_SIZE = POPULATION_PARENTS_SIZE + POPULATION_CHILDREN_SIZE

const ELITISM_QT_PARENTS = 1
const MUTATION_QT_GENE = 10 //30% of 34 genes
const CROSS_OVER_QT = int(POPULATION_CHILDREN_SIZE / 2)

const VECTOR_SIZE = 36
const INDIVIDUAL_SIZE = 34
const POS_EVAL = 34
const POS_ACC_ROULETTE = 35

const TOUR = 3
const NUM_GENERATION = 50
const NUM_EXECUTION = 10

const S1_ROULETTE = 1
const S2_TOURNAMENT = 2
const R1_REINSERTION_ORDERLY = 1
const R2_REINSERTION_PURE_ELITISM = 2

type Individual struct {
	Weight   float32 //0-1
	Operator int     //0-3
	Value    int     //0-3
}

var database [RECORD_QUANTITY][RECORD_SIZE]int
var database_full [RECORD_TOTAL_DATABASE][RECORD_SIZE]int
var population [POPULATION_TOTAL_SIZE][VECTOR_SIZE]Individual
var class_of_execution int
var acc_roulette float32
var n_roulette float32
var aux_roulette float32

var individual_test_1 [VECTOR_SIZE]Individual

func initIndividualTest() {
	//	20: clubbing of the rete ridges >=1
	//  31: perifollicular parakeratosis = 0
	individual_test_1[19].Weight = 0.75
	individual_test_1[19].Operator = 2
	individual_test_1[19].Value = 1
	individual_test_1[30].Weight = 0.75
	individual_test_1[30].Operator = 0
	individual_test_1[30].Value = 0
}

func main() {
	class_of_execution = 1

	initDatabase()

	// initIndividualTest()
	// printIndividual(individual_test_1)
	// calcIndividualFitness(&individual_test_1)
	// printIndividual(individual_test_1)

	fmt.Println("Class of execution: ", class_of_execution)
	for i := 0; i < NUM_EXECUTION; i++ {
		initPopulation()
		calcPopulationFitness()
		for j := 0; j < NUM_GENERATION; j++ {
			//SELECTION AND CROSS OVER PHASE *****************************************
			index := 0
			acc_roulette = 0
			prepareRoulette()
			for k := 0; k < CROSS_OVER_QT; k++ {
				var pos_parent_1, pos_parent_2 int

				pos_parent_1, pos_parent_2 = selectionRoullete()

				child_1, child_2 := calcCrossOverTwoPoints(population[pos_parent_1], population[pos_parent_2])

				population[POPULATION_PARENTS_SIZE+index] = child_1
				population[POPULATION_PARENTS_SIZE+index+1] = child_2
				index = index + 2
			}
			//************************************************************************
			//REINSERTION PHASE ******************************************************
			reinsertionElitismOfTheBetter()
			//************************************************************************
			// printPopulation()

			//MUTATION PHASE *********************************************************
			mutationPopulation()
			//************************************************************************
			calcPopulationFitness()

		}
		// printPopulation()
		optimalIndex := findOptimalIndex()
		fmt.Print("Execution: ", i)
		fmt.Print(" [Optimal Index] ", optimalIndex)
		fmt.Print(" [Optimal Fitness] ", population[optimalIndex][POS_EVAL].Weight)
		fmt.Println(" ")
		fmt.Println(population[optimalIndex])
		printRule(population[optimalIndex])
	}
}

func printRule(individual [VECTOR_SIZE]Individual) {
	for i := 0; i < VECTOR_SIZE-2; i++ {
		if individual[i].Weight >= 0.7 {
			if i == POS_AGE || i == POS_FAMILY_HISTORY {
			} else {
				switch individual[i].Operator {
				case EQUAL:
					fmt.Print("Gene ", i+1)
					fmt.Print(" =  ", individual[i].Value)
					fmt.Println("")
				case DIFFERENT:
					fmt.Print("Gene ", i+1)
					fmt.Print(" !=  ", individual[i].Value)
					fmt.Println("")
				case GREATER_OR_EQUAL:
					fmt.Print("Gene ", i+1)
					fmt.Print(" >=  ", individual[i].Value)
					fmt.Println("")
				case LESS:
					fmt.Print("Gene ", i+1)
					fmt.Print(" <  ", individual[i].Value)
					fmt.Println("")
				}
			}
		}
	}
	fmt.Println(" ")

}

func findOptimalIndex() int {
	var i int
	better_fitness := population[POPULATION_PARENTS_SIZE][POS_EVAL].Weight
	better_index := 0
	for i = 0; i < POPULATION_PARENTS_SIZE; i++ {
		if population[i][POS_EVAL].Weight > better_fitness {
			better_fitness = population[i][POS_EVAL].Weight
			better_index = i
		}
	}

	return better_index
	// fmt.Println("Posicao do melhor: ", better_index)
	// fmt.Println("Posicao do melhor: ", better_fitness)
}

func mutationPopulation() {
	for i := 0; i < POPULATION_TOTAL_SIZE; i++ {
		mutationIndividual(&population[i])
	}
}

func mutationIndividual(individual *[VECTOR_SIZE]Individual) {
	var i int
	rand.Seed(time.Now().UTC().UnixNano())
	for i = 0; i < MUTATION_QT_GENE; i++ { // Mutation on Weight
		gene := rand.Intn(INDIVIDUAL_SIZE)
		individual[gene].Weight = rand.Float32()
		// fmt.Print(" ", gene)
	}
	// fmt.Println(" ")
	for i = 0; i < MUTATION_QT_GENE; i++ { // Mutation on Operator
		gene := rand.Intn(INDIVIDUAL_SIZE)
		individual[gene].Operator = rand.Intn(4)
		// fmt.Print(" ", gene)
	}
	// fmt.Println(" ")
	for i = 0; i < MUTATION_QT_GENE; i++ { // Mutation on Value
		gene := rand.Intn(INDIVIDUAL_SIZE)
		individual[gene].Value = rand.Intn(4)
		// fmt.Print(" ", gene)
	}
	// fmt.Println(" ")
}

func reinsertionElitismOfTheBetter() {
	var i, j int
	better_fitness := population[POPULATION_PARENTS_SIZE][POS_EVAL].Weight
	better_index := 0
	for i = 0; i < POPULATION_PARENTS_SIZE; i++ {
		if population[i][POS_EVAL].Weight > better_fitness {
			better_fitness = population[i][POS_EVAL].Weight
			better_index = i
		}
	}
	population[0] = population[better_index]
	j = 0

	for i := ELITISM_QT_PARENTS; i < POPULATION_PARENTS_SIZE; i++ {
		population[i] = population[POPULATION_PARENTS_SIZE+j]
		j++
	}
	// fmt.Println("Posicao do melhor: ", better_index)
	// fmt.Println("Posicao do melhor: ", better_fitness)
}

func calcCrossOverTwoPoints(individual_1 [VECTOR_SIZE]Individual, individual_2 [VECTOR_SIZE]Individual) ([VECTOR_SIZE]Individual, [VECTOR_SIZE]Individual) {
	child_1 := individual_1
	child_2 := individual_2

	var aux int
	point_one := rand.Intn(INDIVIDUAL_SIZE)
	point_two := rand.Intn(INDIVIDUAL_SIZE)

	// fmt.Println("START CROSS OVER")
	// printIndividual(child_1)
	// printIndividual(child_2)

	if point_two < point_one {
		aux = point_one
		point_one = point_two
		point_two = aux
	}

	for i := point_one; i < point_two; i++ {
		a := child_1[i]
		child_1[i] = child_2[i]
		child_2[i] = a
	}
	// fmt.Println("AFTER CROOS OVER")
	// fmt.Println(point_one)
	// fmt.Println(point_two)
	// printIndividual(child_1)
	// printIndividual(child_2)

	return child_1, child_2
}

func random(min, max float32) float32 {
	return rand.Float32()*(max-min) + min
}

func prepareRoulette() {

	n_roulette = (float32(1) / float32(150))
	aux_roulette = 0

	for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
		acc_roulette = acc_roulette + population[i][POS_EVAL].Weight
	}

	population[0][POS_ACC_ROULETTE].Weight = population[0][POS_EVAL].Weight / acc_roulette

	for i := 1; i < POPULATION_PARENTS_SIZE; i++ {
		population[i][POS_ACC_ROULETTE].Weight = (population[i][POS_EVAL].Weight / acc_roulette) + population[i-1][POS_ACC_ROULETTE].Weight
	}
}

func selectionRoullete() (int, int) {
	var pos_winner_1 int
	var pos_winner_2 int
	//var pos_winner_3 int
	var sorted float32
	//var different bool

	//for i := 0; i < TOUR+10; i++ {
	//sorted = random(aux_roulette, n_roulette)

	sorted = random(0, acc_roulette)
	// fmt.Println("sorted_1: ", sorted)

	for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
		if sorted < population[i][POS_ACC_ROULETTE].Weight {
			pos_winner_1 = i
			i = POPULATION_PARENTS_SIZE + 2
		}
	}

	sorted = random(0, acc_roulette)
	// fmt.Println("sorted_2: ", sorted)
	for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
		if sorted < population[i][POS_ACC_ROULETTE].Weight {
			pos_winner_2 = i
			i = POPULATION_PARENTS_SIZE + 2
		}
	}

	for pos_winner_2 == pos_winner_1 {
		sorted = random(0, acc_roulette)
		// fmt.Println("sorted_2: ", sorted)
		for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
			if sorted < population[i][POS_ACC_ROULETTE].Weight {
				pos_winner_2 = i
				i = POPULATION_PARENTS_SIZE + 2
			}
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
		if population[candidate_pos][POS_EVAL].Weight < population[winner_pos][POS_EVAL].Weight {
			winner_pos = candidate_pos
		}
	}
	// fmt.Println("winner: ",winner)
	return winner_pos
}

func calcPopulationFitness() (int, bool) {
	for i := 0; i < POPULATION_TOTAL_SIZE; i++ {
		// fmt.Println("Individual ", i)
		calcIndividualFitness(&population[i])
		// fmt.Println(" ")
	}
	return -1, false
}

func calcIndividualFitness(individual *[VECTOR_SIZE]Individual) {
	var i, k int
	var fp, fn, tn, tp float32
	var sp, se, eval float32
	var fail_A, fail_C bool

	/*Falsos positivos (Fp), falsos negativos (Fn),
	verdadeiros negativos (Tn) e verdadeiros positivos (Tp)*/

	for k = 0; k < RECORD_QUANTITY; k++ {
		fail_A = false
		fail_C = false
		for i = 0; i < RECORD_SIZE-1; i++ {
			if individual[i].Weight >= 0.7 {
				// fmt.Print("  ", individual[i].Weight)
				// fmt.Print("  ", individual[i].Operator)
				// fmt.Print("  ", individual[i].Value)
				// fmt.Println(" ")
				if i == POS_AGE || i == POS_FAMILY_HISTORY {
				} else {
					switch individual[i].Operator {
					case EQUAL:
						if individual[i].Value == database[k][i] {
						} else {
							fail_A = true
						}
					case DIFFERENT:
						if individual[i].Value != database[k][i] {
						} else {
							fail_A = true
						}
					case GREATER_OR_EQUAL:
						if individual[i].Value >= database[k][i] {
						} else {
							fail_A = true
						}
					case LESS:
						if individual[i].Value < database[k][i] {
						} else {
							fail_A = true
						}
					}
				}
			}
		}
		if class_of_execution == database[k][RECORD_POS_CLASS] { // C
		} else { // ¬C
			fail_C = true
		}
		/*SE A  ENTAO C  = TP
		SE A  ENTAO ¬C = FP
		SE ¬A ENTAO C  = FN
		SE ¬A ENTAO ¬C = TN*/
		if fail_A && fail_C {
			tn++
		} else {
			if fail_A {
				fn++
			} else {
				if fail_C {
					fp++
				} else {
					tp++
				}
			}
		}
	}

	// Fidelis
	// se = tp / (tp +fn)       sp = tn / (tn + fp)

	se = tp / (tp + fn)
	sp = tn / (tn + fp)

	// se = tp + 1/(tp+fn+1)
	// sp = tn + 1/(tn+fp+1)

	if se == 0 {
		se = 0.1
	}
	if sp == 0 {
		se = 0.1
	}
	// Clay
	// sp = tp / (tp + fp)      se = tn / (tn + fn)
	// sp = tp / (tp + fp)
	// se = tn / (tn + fn)
	eval = sp * se
	individual[POS_EVAL].Weight = eval
	/*
		fmt.Print("TN: ", tn)
		fmt.Print("    FP: ", fp)
		fmt.Print("    TN+FP: ", tn+fp)
		fmt.Println(" ")

		fmt.Print("TP: ", tp)
		fmt.Print("    FN: ", fn)
		fmt.Print("    TP+FN: ", tp+fn)
		fmt.Println(" ")

		fmt.Print("SP: ", sp)
		fmt.Print("    SE: ", se)
		fmt.Println(" ")
		fmt.Print("EVAL: ", eval)
		fmt.Println(" ")*/
}

func initPopulation() {
	//rand.Seed(1)
	rand.Seed(time.Now().UTC().UnixNano())
	var i, j int

	for i = 0; i < POPULATION_PARENTS_SIZE; i++ {
		for j = 0; j < INDIVIDUAL_SIZE; j++ {
			var individual Individual
			individual.Weight = rand.Float32()
			individual.Operator = rand.Intn(4)
			individual.Value = rand.Intn(4)
			population[i][j] = individual
		}
	}
}

func printPopulation() {
	for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
		fmt.Print("Individuo [", i)
		fmt.Println("]: ", population[i])
	}
}

func printIndividual(individual [VECTOR_SIZE]Individual) {
	fmt.Println("Individual ")
	for i := 0; i < VECTOR_SIZE; i++ {
		fmt.Print(" [pos ", i)
		fmt.Print("  ", individual[i])
		fmt.Print(" ]")
	}
	fmt.Println(" ")
}

func initDatabase() {
	// training = 234 records - 39 records for each class
	// teste = 114 records - 19 records fo each class
	var class_1, class_2, class_3, class_4, class_5, class_6 int
	var i, j, k int
	if file, err := os.Open("dermatology.data"); err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		i = 0

		for k = 0; k < RECORD_TOTAL_DATABASE; k++ {
			scanner.Scan()
			result := strings.Split(scanner.Text(), ",")
			for j = 0; j < RECORD_SIZE; j++ {
				database_full[i][j], _ = strconv.Atoi(result[j])
			}
			i++
		}

		if TRAINING_DATABASE {
			k = 0
			i = 0
		} else {
			k = RECORD_TOTAL_DATABASE - 1
			i = 0
		}
		for i < RECORD_QUANTITY {
			// fmt.Println(k)
			// fmt.Println(i)
			// fmt.Println(class_1)
			// fmt.Println(class_2)
			// fmt.Println(class_3)
			// fmt.Println(class_4)
			// fmt.Println(class_5)
			// fmt.Println(class_6)
			switch database_full[k][RECORD_POS_CLASS] {
			case 1:
				if class_1 < RECORD_QUANTITY_TRAINING {
					database[i] = database_full[k]
					i++
					class_1++
				}
			case 2:
				if class_2 < RECORD_QUANTITY_TRAINING {
					database[i] = database_full[k]
					i++
					class_2++
				}
			case 3:
				if class_3 < RECORD_QUANTITY_TRAINING {
					database[i] = database_full[k]
					i++
					class_3++
				}
			case 4:
				if class_4 < RECORD_QUANTITY_TRAINING {
					database[i] = database_full[k]
					i++
					class_4++
				}
			case 5:
				if class_5 < RECORD_QUANTITY_TRAINING {
					database[i] = database_full[k]
					i++
					class_5++
				}
			case 6:
				if class_6 < 19 {
					database[i] = database_full[k]
					i++
					class_6++
				}
			}
			if TRAINING_DATABASE {
				k++
			} else {
				k--
			}
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
