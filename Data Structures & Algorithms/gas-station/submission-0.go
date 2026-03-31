func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)

	var start int
	maxVolume := 0
    for i := 0; i < n; i++ {
		if gas[i] - cost[i] > maxVolume {
			maxVolume = gas[i] - cost[i]
			start = i
		}
	}

	if maxVolume == 0 {
		return -1
	}

	fmt.Println("Ceva")

	volume := maxVolume
	i := (start+1) % n
	for volume > 0 && i != start{
		nextStation := (i+1) % n
		volume = volume - cost[i] + gas[i]
		i = nextStation
	}

	if i == start {
		return start
	}
	return -1
}
