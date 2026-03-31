type MinStack struct {
	storage []int
}

func Constructor() MinStack {
	return MinStack{
		storage: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.storage = append(this.storage, val)
}

func (this *MinStack) Pop() {
	this.storage = this.storage[:len(this.storage)-1]
}

func (this *MinStack) Top() int {
	return this.storage[len(this.storage)-1]
}

func (this *MinStack) GetMin() int {
	res := math.MaxInt32
	for _, x := range this.storage {
		res = min(res, x)
	}

	return res
}
