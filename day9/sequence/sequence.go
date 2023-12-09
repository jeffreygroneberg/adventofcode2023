package sequence

func CalculateNextInSequence(sequence []int) int {

	var belowSequence []int

	for i := 0; i < len(sequence)-1; i++ {

		belowSequence = append(belowSequence, sequence[i+1]-sequence[i])

	}

	// if belowSeqence contains only 0
	if containsOnlyZeros(belowSequence) {
		return sequence[len(sequence)-1]

	} else {

		lastElement := sequence[len(sequence)-1]

		return lastElement + CalculateNextInSequence(belowSequence)
	}
}

func containsOnlyZeros(slice []int) bool {
	for _, v := range slice {
		if v != 0 {
			return false
		}
	}
	return true
}
