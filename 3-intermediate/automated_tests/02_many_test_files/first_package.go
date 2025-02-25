package first_package

func StudentPass(grades []int, average int)bool{
	if( len(grades) == 0 ){
		return false
	}

	total := 0
	for _, grade := range(grades) {
		total += grade
	}
	studentAverage := float32(total/len(grades))
	return studentAverage >= float32(average)
}
