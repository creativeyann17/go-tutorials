package math

func Add(x, y int) int {
	return x + y
}

func isPrime(value int) bool {
	if value %2==0 || value <2{
		return false
	}
	for i := 3; i< value/2;i+=2 {
		if (value %i ==0) {
			return false
		} 
	}
	return true
}