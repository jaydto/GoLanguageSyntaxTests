package arraysgo

import "fmt"

func LoopsInArray(){
	s := make([]string, 5)
	s[0] = "q"
	s[1] = "2"

	primes := []int{1, 2, 3, 4, 5}
	s = append(s, "1", "2")
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(s[0:4])


	fmt.Println(primes[1:4])
	fmt.Println(len(primes))
	var twoDimensions [2][3]int
	var fourDimentsion [2][3][4][5]int

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			for z := 0; z < 4; z++ {
				for k := 0; k < 5; k++ {
					fourDimentsion[i][j][z][k] = i + j + k + z
				}
			}
		}
	}

	fmt.Println(fourDimentsion)

	for i := range fourDimentsion {
		for j := range fourDimentsion {
			for k := range fourDimentsion {
				for z := range fourDimentsion {
					fourDimentsion[i][j][k][z] = i + j + k + z
				}
			}
		}
	}

	fmt.Println(fourDimentsion)

	x, y, z := 2, 3, 5
	threeDimentsions := make([][][]int, x)

	for i := range threeDimentsions {
		threeDimentsions[i] = make([][]int, y)

		for j := range threeDimentsions {
			threeDimentsions[i][j] = make([]int, z)
		}

	}

	fmt.Println(threeDimentsions)

	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {

			twoDimensions[i][j] = i + j
		}
	}

	fmt.Println(twoDimensions)

}