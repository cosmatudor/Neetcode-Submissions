func carFleet(target int, position []int, speed []int) int {
	stack := []float64{}
	type Car struct {
    	pos, spd int
	}

	cars := make([]Car, len(position))
	for i := range position {
		cars[i] = Car{position[i], speed[i]}
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	for i := 0; i < len(cars); i++ {
		time := float64(target - cars[i].pos) / float64(cars[i].spd)
		if len(stack) == 0 || time > stack[len(stack)-1] {
            stack = append(stack, time)
        }
	}

	return len(stack)

}

// 0 1 2 3 4 ...
// 1 3 5 7 9 11
// 4 7 10 
// 7 8 9 10

// 7 4 1 0
// 3 2 4 10
