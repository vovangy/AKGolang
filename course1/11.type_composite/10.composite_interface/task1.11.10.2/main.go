package main

func Operate(f func(xs ...interface{}) interface{}, i ...interface{}) interface{} {
	return f(i...)
}

func Concat(xs ...interface{}) interface{} {
	result := ""
	for _, x := range xs {
		str, _ := x.(string)
		//if !ok {
		//	log.Fatalf("Concat: all arguments must be of type string, got %T", x)
		//}
		result += str
	}
	return result
}

func Sum(xs ...interface{}) interface{} {
	var intSum int
	var floatSum float64
	var isFloat bool

	for _, x := range xs {
		switch v := x.(type) {
		case int:
			intSum += v
		case float64:
			floatSum += v
			isFloat = true
			//default:
			//	log.Fatalf("Sum: arguments must be of type int or float64, got %T", x)
		}
	}

	if isFloat {
		return float64(intSum) + floatSum
	}
	return intSum
}
