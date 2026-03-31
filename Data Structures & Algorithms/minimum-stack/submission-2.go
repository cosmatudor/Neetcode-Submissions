type MinStack struct {
	storage []int
	minStorage []int
}

func Constructor() MinStack {
	return MinStack{
		storage: []int{},
		minStorage: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.storage = append(this.storage, val)
	
	var currMin int
	if len(this.minStorage) > 0 {
		currMin = this.minStorage[len(this.minStorage)-1]
	} else {
		currMin = math.MaxInt32
	}
	newMin := min(currMin, val)
	this.minStorage = append(this.minStorage, newMin)
}

func (this *MinStack) Pop() {
	this.storage = this.storage[:len(this.storage)-1]
	this.minStorage = this.minStorage[:len(this.minStorage)-1]
}

func (this *MinStack) Top() int {
	return this.storage[len(this.storage)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStorage[len(this.minStorage)-1]
}
