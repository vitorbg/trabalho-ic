package main

import (
	"fmt"
	"github.com/vitorbg/trabalho-ic/rn_perceptron/display"
)



func main() {

	d1 := []int{0, 1}
	n := display.GetNumbers()
	n_zero_distorted := display.GetNumbersZeroDistorted()
	n_one_distorted := display.GetNumbersOneDistorted()

	w := []int{0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0,
		0, 0, 0, 0, 0}

	//Perceptron One ***************************************************************
	w2 := perceptronOne(n, w, d1)

	fmt.Println("Perceptron One Weight Vector")
	printDisplay(w2)

	for i := 0; i < 2; i++ {
		fmt.Println("INPUT")
		printDisplay(n[i])
		fmt.Println("OUTPUT: ", y(n[i], w2))
	}
	fmt.Println("One Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_one_distorted[i])
		fmt.Println("OUTPUT: ", y(n_one_distorted[i], w2))
	}
	fmt.Println("Zero Distorted Examples")
	for i := 0; i < 10; i++ {
		fmt.Println("INPUT")
		printDisplay(n_zero_distorted[i])
		fmt.Println("OUTPUT: ", y(n_zero_distorted[i], w2))
	}
	fmt.Println("Numbers Examples")
	for i := 2; i < 6; i++ {
		fmt.Println("INPUT")
		printDisplay(n[i])
		fmt.Println("OUTPUT: ", y(n[i], w2))
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

func perceptronOne(n [6][]int, w []int, d []int) []int {
	var err bool
	var e int
	err = true
	for err {
		for i := 0; i < 2; i++ {
			y := y(n[i], w)
			e = (d[i] - y)
			if d[i] != y {
				// fmt.Println("DEU RUIM")
				for j := 0; j < 31; j++ {
					w[j] += e * n[i][j]
				}
				err = true
			} else {
				err = false
			}
		}
	}

	return w
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

/*

int x [6][31] = {{1,
		  0,1,1,1,0,
		  1,0,0,0,1,
		  1,0,0,0,1,
		  1,0,0,0,1,
		  1,0,0,0,1,
		  0,1,1,1,0},

		 {1,
		  0,0,1,0,0,
		  0,1,1,0,0,
		  0,0,1,0,0,
		  0,0,1,0,0,
		  0,0,1,0,0,
		  1,1,1,1,1},

		 {1,
		  0,1,1,1,0,
		  0,1,0,1,0,
		  0,0,0,1,0,
		  0,0,1,0,0,
		  0,1,0,0,0,
		  1,1,1,1,1},

		 {1,
		  0,1,1,1,1,
		  0,0,0,0,1,
		  0,1,1,1,1,
		  0,0,0,0,1,
		  0,0,0,0,1,
		  0,1,1,1,1},

		 {1,
		  1,0,0,1,0,
		  1,0,0,1,0,
		  1,0,0,1,0,
		  1,1,1,1,0,
		  0,0,0,1,0,
		  0,0,0,1,0},

		 {1,
		  0,1,1,1,1,
		  0,1,0,0,0,
		  0,1,1,1,1,
		  0,0,0,0,1,
		  0,0,0,0,1,
		  0,1,1,1,1}
};

*/
