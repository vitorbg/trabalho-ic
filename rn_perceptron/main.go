package main

import (
	"fmt"
	"github.com/vitorbg/trabalho-ic/rn_perceptron/display"
)

func main() {

	d1 := []int{0, 1}
	d2 := [2][]int{{1, 0}, {0, 1}}
	d3 := [6][]int{{1, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 0, 0}, {0, 0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 1, 0}, {0, 0, 0, 0, 0, 1}}
	n := display.GetNumbers()
	n_zero_distorted := display.GetNumbersZeroDistorted()
	n_one_distorted := display.GetNumbersOneDistorted()
	n_two_distorted := display.GetNumbersTwoDistorted()
	n_three_distorted := display.GetNumbersThreeDistorted()
	n_four_distorted := display.GetNumbersFourDistorted()
	n_five_distorted := display.GetNumbersFiveDistorted()
	n_letters := display.GetLetters()

	//Perceptron One ***************************************************************
	w1_init := display.GetWeightZero()
	// w1_init := display.GetWeightRandom()

	fmt.Println("PERCEPTRON ONE ************************************************")
	fmt.Println("INPUT WEIGHT:")
	fmt.Print(w1_init)
	fmt.Println("")

	epoch, w1 := perceptronOne(n, w1_init, d1)

	fmt.Println("Perceptron One Weight Vector")
	fmt.Println("Epochs of Training: ", epoch)
	printDisplay(w1)

	for i := 0; i < 2; i++ {
		fmt.Println("INPUT")
		printDisplay(n[i])
		fmt.Println("OUTPUT: ", y(n[i], w1))
	}
	fmt.Println("One Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_one_distorted[i])
		fmt.Println("OUTPUT: ", y(n_one_distorted[i], w1))
	}
	fmt.Println("Zero Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_zero_distorted[i])
		fmt.Println("OUTPUT: ", y(n_zero_distorted[i], w1))
	}
	fmt.Println("Numbers Examples")
	for i := 2; i < 6; i++ {
		fmt.Println("INPUT")
		printDisplay(n[i])
		fmt.Println("OUTPUT: ", y(n[i], w1))
	}
	//******************************************************************************
	//Perceptron Two ***************************************************************
	// w2_init := display.GetWeightZeroTwoNeurons()
	w2_init := display.GetWeightRandomTwoNeurons()

	fmt.Println("PERCEPTRON TWO ************************************************")
	fmt.Println("INPUT WEIGHT:")
	fmt.Print(w2_init[0])
	fmt.Println("")
	fmt.Print(w2_init[1])
	fmt.Println("")

	epoch, w2 := perceptronTwo(n, w2_init, d2)

	fmt.Println("Perceptron Two Weight Vectors")
	fmt.Println("Epochs of Training: ", epoch)
	fmt.Println("Neuron 1")
	printDisplay(w2[0])
	fmt.Println("Neuron 2")
	printDisplay(w2[1])

	for i := 0; i < 2; i++ {
		fmt.Println("INPUT")
		printDisplay(n[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n[i], w2[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n[i], w2[1]))
	}
	fmt.Println("Zero Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_zero_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_zero_distorted[i], w2[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_zero_distorted[i], w2[1]))
	}
	fmt.Println("One Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_one_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_one_distorted[i], w2[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_one_distorted[i], w2[1]))
	}

	fmt.Println("Numbers Examples")
	for i := 2; i < 6; i++ {
		fmt.Println("INPUT")
		printDisplay(n[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n[i], w2[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n[i], w2[1]))
	}
	//******************************************************************************
	//Perceptron Three ***************************************************************
	w3_init := display.GetWeightZeroSixNeurons()
	// w3_init := display.GetWeightRandomSixNeurons()

	fmt.Println("PERCEPTRON THREE **********************************************")
	fmt.Println("INPUT WEIGHT:")
	fmt.Print(w3_init[0])
	fmt.Println("")
	fmt.Print(w3_init[1])
	fmt.Println("")
	fmt.Print(w3_init[2])
	fmt.Println("")
	fmt.Print(w3_init[3])
	fmt.Println("")
	fmt.Print(w3_init[4])
	fmt.Println("")
	fmt.Print(w3_init[5])
	fmt.Println("")

	epoch, w3 := perceptronThree(n, w3_init, d3)

	fmt.Println("Perceptron Two Weight Vectors")
	fmt.Println("Epochs of Training: ", epoch)
	fmt.Println("Neuron 1")
	printDisplay(w3[0])
	fmt.Println("Neuron 2")
	printDisplay(w3[1])
	fmt.Println("Neuron 3")
	printDisplay(w3[2])
	fmt.Println("Neuron 4")
	printDisplay(w3[3])
	fmt.Println("Neuron 5")
	printDisplay(w3[4])
	fmt.Println("Neuron 6")
	printDisplay(w3[5])
	for i := 0; i < 6; i++ {
		fmt.Println("INPUT")
		printDisplay(n[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n[i], w3[5]))

	}
	fmt.Println("Zero Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_zero_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_zero_distorted[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_zero_distorted[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n_zero_distorted[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n_zero_distorted[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n_zero_distorted[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n_zero_distorted[i], w3[5]))
	}
	fmt.Println("One Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_one_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_one_distorted[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_one_distorted[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n_one_distorted[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n_one_distorted[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n_one_distorted[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n_one_distorted[i], w3[5]))
	}
	fmt.Println("Two Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_two_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_two_distorted[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_two_distorted[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n_two_distorted[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n_two_distorted[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n_two_distorted[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n_two_distorted[i], w3[5]))
	}
	fmt.Println("Three Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_three_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_three_distorted[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_three_distorted[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n_three_distorted[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n_three_distorted[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n_three_distorted[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n_three_distorted[i], w3[5]))
	}
	fmt.Println("Four Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_four_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_four_distorted[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_four_distorted[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n_four_distorted[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n_four_distorted[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n_four_distorted[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n_four_distorted[i], w3[5]))
	}
	fmt.Println("Five Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_five_distorted[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_five_distorted[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_five_distorted[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n_five_distorted[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n_five_distorted[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n_five_distorted[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n_five_distorted[i], w3[5]))
	}
	fmt.Println("Letters Examples")
	for i := 0; i < 6; i++ {
		fmt.Println("INPUT")
		printDisplay(n_letters[i])
		fmt.Println("OUTPUT NEURON 1: ", y(n_letters[i], w3[0]))
		fmt.Println("OUTPUT NEURON 2: ", y(n_letters[i], w3[1]))
		fmt.Println("OUTPUT NEURON 3: ", y(n_letters[i], w3[2]))
		fmt.Println("OUTPUT NEURON 4: ", y(n_letters[i], w3[3]))
		fmt.Println("OUTPUT NEURON 5: ", y(n_letters[i], w3[4]))
		fmt.Println("OUTPUT NEURON 6: ", y(n_letters[i], w3[5]))
	}
	//******************************************************************************
}

func y(x []int, w []int) int {
	var u int
	for i := 0; i < 31; i++ {
		u += x[i] * w[i]
	}
	a := degree(u)

	return a
}

func degree(x int) int {
	if x > 0 {
		return 1
	} else {
		return 0
	}
}

func perceptronOne(n [6][]int, w []int, d []int) (int, []int) {
	var err bool
	var e int
	var epoch int
	err = true
	epoch = 0

	for err {
		for i := 0; i < 2; i++ {
			y := y(n[i], w)
			e = (d[i] - y)
			if d[i] != y {
				for j := 0; j < 31; j++ {
					w[j] += e * n[i][j]
				}
				err = true
			} else {
				err = false
			}
		}
		epoch++
	}

	return epoch, w
}

func perceptronTwo(n [6][]int, w [2][]int, d [2][]int) (int, [2][]int) {
	var err bool
	var e int
	var epoch int

	err = true
	epoch = 0

	for err {
		for i := 0; i < 2; i++ {

			for neur := 0; neur < 2; neur++ {
				y := y(n[i], w[neur])
				e = (d[i][neur] - y)
				if d[i][neur] != y {
					for j := 0; j < 31; j++ {
						w[neur][j] += e * n[i][j]
					}
					err = true
				} else {
					err = false
				}
			}
		}
		epoch++
	}

	return epoch, w
}

func perceptronThree(n [6][]int, w [6][]int, d [6][]int) (int, [6][]int) {
	var err bool
	var e int
	var epoch int

	err = true
	epoch = 0

	for err {
		for i := 0; i < 6; i++ {
			for neur := 0; neur < 6; neur++ {
				y := y(n[i], w[neur])
				e = (d[i][neur] - y)
				if d[i][neur] != y {
					for j := 0; j < 31; j++ {
						w[neur][j] += e * n[i][j]
					}
					err = true
				} else {
					err = false
				}
			}
		}
		epoch++
	}

	return epoch, w
}

func printDisplay(v []int) {
	fmt.Println(fmt.Sprintf("%d ", v[0]))
	for i := 1; i < 31; i++ {
		fmt.Print(fmt.Sprintf("%d ", v[i]))
		if i%5 == 0 {
			fmt.Println("")
		}
	}
}
