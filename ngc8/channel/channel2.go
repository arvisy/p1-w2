package channel

func SumSquare(num int, result chan int) {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i * i
	}

	result <- sum
	close(result)
}

func SquareSum(num int, result chan int) {
	sum := 0
	for i := 1; i <= num; i++ {
		sum += i
	}

	result <- sum * sum
	close(result)
}
