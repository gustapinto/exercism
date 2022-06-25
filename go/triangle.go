package triangle

type Kind string

const (
	NaT = "NaT" // not a triangle
	Equ = "Equ" // equilateral
	Iso = "Iso" // isosceles
	Sca = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {
	isAnySideZeroOrLess := (a <= 0) || (b <= 0) || (c <= 0)
	isSumOfAnySidesGreaterThanOther := (a+b >= c) && (a+c >= b) && (b+c >= a)
	isAValidTriangle := !isAnySideZeroOrLess && isSumOfAnySidesGreaterThanOther

	if !isAValidTriangle {
		return NaT
	}

	if a == b && a == c && b == c {
		return Equ
	}

	if a != b && a != c && b != c {
		return Sca
	}

	return Iso
}
