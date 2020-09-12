package utils

import "fmt"

//Sum returns a + b
func Sum(a int, b int) int {
	result := a + b
	fmt.Println(result)
	return result
}

//SuperSum returns total sum
func SuperSum(numbers ...int) int {
	total := 0
	//range는 loop를 적용할 수 있게 해준다.
	//인덱스와 해당값을 받는데 여기서는 인덱스가 필요없기 때문에 _로 받는다.
	for _, number := range numbers {
		total += number
	}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}

	fmt.Println("total sum is ", total)

	return total
}

//Multiply retruns a * b
func Multiply(a, b int) int {
	result := a * b
	fmt.Println(result)
	return result
}
