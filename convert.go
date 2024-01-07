package weight

func KgToLb(n float64) LB {
	return LB(n * kgToLbRatio)
}
