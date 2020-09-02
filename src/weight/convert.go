package weight

func KgToLb(k KG) LB {
	return LB(k * kgToLbRatio)
}
