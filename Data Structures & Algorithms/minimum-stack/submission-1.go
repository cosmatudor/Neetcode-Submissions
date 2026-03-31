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
	
	currMin := this.minStorage[len(this.minStorage)-1]
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
	return this.minStorage = this.minStorage[len(this.minStorage)-1]
}
