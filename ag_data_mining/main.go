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
	EQUAL = iota
	DIFFERENT
	GREATER_OR_EQUAL
	LESS
)

const RECORD_SIZE = 35
const RECORD_QUANTITY = 100
const TRAINING_DATABASE = true

const CROSS_OVER_RATIO = 1
const MUTATION_RATIO = 0.3

//const ELITISM_RATIO = 0.2

const POPULATION_PARENTS_SIZE = 50
const POPULATION_CHILDREN_SIZE = POPULATION_PARENTS_SIZE * CROSS_OVER_RATIO
const POPULATION_TOTAL_SIZE = POPULATION_PARENTS_SIZE + POPULATION_CHILDREN_SIZE

const ELITISM_QT_PARENTS = 1
const MUTATION_QT_INDIVIDUALS = int(POPULATION_CHILDREN_SIZE * MUTATION_RATIO)
const CROSS_OVER_QT = int(POPULATION_CHILDREN_SIZE / 2)

const VECTOR_SIZE = 36
const INDIVIDUAL_SIZE = 34
const POS_EVAL = 34
const POS_ACC_ROULETTE = 35

const TOUR = 3
const NUM_GENERATION = 5
const NUM_EXECUTION = 1000

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
var population [POPULATION_TOTAL_SIZE][VECTOR_SIZE]Individual
var class_of_execution int
var acc_roulette float32
var n_roulette float32
var aux_roulette float32

/*Os operadores genéticos utilizados foram os seguintes: torneio
estocástico, de tamanho 3, como método de seleção; crossover
de  dois  pontos,  com  probabilidade  de  100%  e  elitismo  como
estratégica   de   reprodução,   preservando   apenas   o   melhor
indivíduo.
Para  a  mutação,  utilizou-se  três  operadores  de  mutação,  sendo
cada  um  deles  associado  aos  segmentos  existentes  em  cada
gene  do  cromossomo.
Assim,  aplicou-se  uma  taxa de  mutação
de 30%, para cada tipo de mutação separadamente.*/

func main() {
	class_of_execution = 1

	initDatabase()
	initPopulation()
	calcPopulationFitness()
	// printPopulation()
	index := 0
	//SELECTION AND CROSS OVER PHASE *****************************************

	for k := 0; k < CROSS_OVER_QT; k++ {
		//var pos_parent_1, pos_parent_2 int

		pos_parent_1, pos_parent_2 := selectionTournament()

		child_1, child_2 := calcCrossOverTwoPoints(population[pos_parent_1], population[pos_parent_2])

		population[POPULATION_PARENTS_SIZE+index] = child_1
		population[POPULATION_PARENTS_SIZE+index+1] = child_2
		index = index + 2
	}
	//************************************************************************

	prepareRoulette()
	printPopulation()
	selectionRoullete()

	// printPopulation()
	//reinsertionElitismOfTheBetter()

	fmt.Println("ACC roulette: ", acc_roulette)
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
	fmt.Println("Posicao do melhor: ", better_index)
	fmt.Println("Posicao do melhor: ", better_fitness)
}

func calcCrossOverTwoPoints(individual_1 [VECTOR_SIZE]Individual, individual_2 [VECTOR_SIZE]Individual) ([VECTOR_SIZE]Individual, [VECTOR_SIZE]Individual) {
	child_1 := individual_1
	child_2 := individual_2

	var aux int
	point_one := rand.Intn(INDIVIDUAL_SIZE - 1)
	point_two := rand.Intn(INDIVIDUAL_SIZE - 1)

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
	printIndividual(child_1)
	printIndividual(child_2)

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
	var sorted float32

	for i := 0; i < TOUR+10; i++ {
		sorted = random(aux_roulette, n_roulette)

		fmt.Println(fmt.Sprintf("Range[%f - %f]", aux_roulette, n_roulette))
		fmt.Println("sorted: ", sorted)

		aux_roulette = n_roulette
		n_roulette = aux_roulette + n_roulette

		for i := 0; i < POPULATION_PARENTS_SIZE; i++ {
			if sorted < population[i][POS_ACC_ROULETTE].Weight {
				pos_winner_1 = i
				i = POPULATION_PARENTS_SIZE + 2
			}
		}
		fmt.Println("pos_winner_1: ", pos_winner_1)
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
		calcIndividualFitness(&population[i])
	}
	return -1, false
}

func calcIndividualFitness(individual *[VECTOR_SIZE]Individual) {
	var i, k int
	var fp, fn, tn, tp float32
	var sp, se, eval float32
	/*Falsos positivos (Fp), falsos negativos (Fn),
	verdadeiros negativos (Tn) e verdadeiros positivos (Tp)*/

	for k = 0; k < RECORD_QUANTITY; k++ {
		for i = 0; i < RECORD_SIZE-2; i++ { //Less 2, Age and Class
			/*SE A  ENTAO C  = TP
			SE A  ENTAO ¬C = FP
			SE ¬A ENTAO C  = FN
			SE ¬A ENTAO ¬C = TN*/
			if individual[i].Weight >= 0.7 {
				// fmt.Print("  ", individual[i].Weight)
				// fmt.Print("  ", individual[i].Operator)
				// fmt.Print("  ", individual[i].Value)
				// fmt.Println(" ")

				switch individual[i].Operator {

				case EQUAL:
					if individual[i].Value == database[k][i] { // A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							tp++
						} else { // ¬C
							fp++
						}
					} else { // ¬A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							fn++
						} else { // ¬C
							tn++
						}
					}
				case DIFFERENT:
					if individual[i].Value == database[k][i] { // A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							tp++
						} else { // ¬C
							fp++
						}
					} else { // ¬A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							fn++
						} else { // ¬C
							tn++
						}
					}
				case GREATER_OR_EQUAL:
					if individual[i].Value == database[k][i] { // A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							tp++
						} else { // ¬C
							fp++
						}
					} else { // ¬A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							fn++
						} else { // ¬C
							tn++
						}
					}
				case LESS:
					if individual[i].Value == database[k][i] { // A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							tp++
						} else { // ¬C
							fp++
						}
					} else { // ¬A
						if class_of_execution == database[k][RECORD_SIZE-1] { // C
							fn++
						} else { // ¬C
							tn++
						}
					}
				}
			}
		}
	}
	sp = tp / (tp + fp)
	se = tn / (tn + fn)
	eval = sp * se
	individual[POS_EVAL].Weight = eval
	// fmt.Println("TP: ", tp)
	// fmt.Println("TN: ", tn)
	// fmt.Println("FP: ", fp)
	// fmt.Println("FN: ", fn)
	// fmt.Println("SP: ", sp)
	// fmt.Println("SE: ", se)
	// fmt.Println("EVAL: ", eval)
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
	//for i := 0; i < POPULATION_TOTAL_SIZE; i++ {
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
	// 366 / 6 = 61 records for each class
	// training = 244 records - 40 records for each class
	// teste = 122 records - 20 records fo each class
	//var class_1, class_2, class_3, class_4, class_5, class_6 int
	if file, err := os.Open("dermatology.data"); err == nil {

		defer file.Close()

		scanner := bufio.NewScanner(file)
		i := 0
		//for scanner.Scan() {
		for k := 0; k < RECORD_QUANTITY; k++ {
			scanner.Scan()
			result := strings.Split(scanner.Text(), ",")
			//			class, _ := strconv.Atoi(result[RECORD_SIZE-1])

			for j := 0; j < RECORD_SIZE; j++ {
				database[i][j], _ = strconv.Atoi(result[j])
			}
			i++
		}
		if err = scanner.Err(); err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(err)
	}
}
