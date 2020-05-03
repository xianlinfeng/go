package fund

func Grow(arr []int) int{
	mul := 1
	for i := range arr{
		mul *= i
	}
	return mul
}