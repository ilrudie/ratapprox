package ratapprox

//Fraction defines a new type representing a fraction with a numerator and a denominator stored as indipendent values
type Fraction struct {
	N int64
	D int64
}

//Perform a single iteration of "naive" fraction addition to get a new rational number guaranteed to be between low and high
func farey(low, high Fraction) Fraction {
	return Fraction{low.N + high.N, low.D + high.D}
}

//Just convert the Fraction to a float64 for easy comparison
func (input Fraction) divFraction() float64 {
	return (float64(input.N) / float64(input.D))
}

/*
RatApprox calculates the rational representation (or best effort approximation bound by interations) of the provided float64 "input"
Loop:
	Uses farey to get a new midpoint rational between high a low.
	Compares mid to the input and sets mid to the low or high bound for the next round or returns if the same.
Returns the best guess if loop never terminated by reaching a rational value equal to the input
*/
func RatApprox(input float64) Fraction {
	var mid Fraction
	low := Fraction{0, 1}
	high := Fraction{1, 1}
	for i := 0; i < 1000; i++ {
		mid = farey(low, high)
		var comp float64 = mid.divFraction()
		if comp == input {
			return mid
		} else if comp < input {
			low = mid
		} else {
			high = mid
		}
	}
	return mid
}
