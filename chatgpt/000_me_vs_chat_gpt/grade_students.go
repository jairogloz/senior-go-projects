package chatgpt

func GradingStudents2(grades []int32) []int32 {
	// Write your code here
	// 73 - 75
	for i, g := range grades {
		if g < 38 {
			continue
		}
		// 73 - (3) + 5 = 75
		nextMultiple := g - (g % 5) + 5
		if nextMultiple-g < 3 {
			grades[i] = nextMultiple
		}
	}

	return grades

}

func GradingStudents1(grades []int32) []int32 {
	// Write your code here
	roundedGrades := make([]int32, len(grades))

	for i, grade := range grades {
		if grade < 38 {
			// If the grade is less than 38, don't round it
			roundedGrades[i] = grade
		} else {
			// Calculate the next multiple of 5
			// 73 / 5 = 15*5 = 75
			nextMultipleOfFive := ((grade / 5) + 1) * 5
			if nextMultipleOfFive-grade < 3 {
				// If the difference is less than 3, round up
				roundedGrades[i] = nextMultipleOfFive
			} else {
				// Otherwise, keep the original grade
				roundedGrades[i] = grade
			}
		}
	}

	return roundedGrades

}
