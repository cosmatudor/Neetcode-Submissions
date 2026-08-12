type TimeMap struct {
    keyToTimeToVal map[string][]Pair
}

type Pair struct {
    timestamp int
    value     string
}
func Constructor() TimeMap {
	return TimeMap {
		keyToTimeToVal: make(map[string][]Pair),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.keyToTimeToVal[key] = append(this.keyToTimeToVal[key], Pair{timestamp, value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	temp, _ := this.keyToTimeToVal[key]

	left, right := 0, len(temp)-1
	for left <= right {
		mid := (left + right) / 2
		if temp[mid].timestamp <= timestamp {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if right < 0 {
		return ""
	}
	return this.keyToTimeToVal[key][right].value
}

