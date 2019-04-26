package service

type StringFlipper interface {
	FlipString(val string) string
}

type StringFlipperImpl struct {
}

func (sf StringFlipperImpl) FlipString(val string) string {

	a := []rune(val)
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return string(a)

}
