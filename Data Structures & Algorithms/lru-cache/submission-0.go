type LRUCache struct {
    cap int
	storage map[int]*Node
	startNode *Node
	endNode *Node
}

type Node struct {
    key  int
    val  int
    next *Node
    prev *Node
}

func Constructor(capacity int) LRUCache {
	start := &Node{}
	end := &Node{}
	start.next = end
	end.prev = start
	
    return LRUCache {
		cap: capacity,
		storage: make(map[int]*Node),
		startNode: start,
		endNode: end,
	}
}

func (this *LRUCache) Get(key int) int {
    if node, ok := this.storage[key]; ok {
		prev := node.prev
		next := node.next
		prev.next = next
		next.prev = prev

		this.endNode.prev.next = node
		node.prev = this.endNode.prev
		node.next = this.endNode
		this.endNode.prev = node

		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.storage[key]; ok {
		this.Get(key)
		this.storage[key].val = value
		return
	}

	node := &Node{
		key: key,
		val: value,
	}
	this.storage[key] = node
	
	lastNode := this.endNode.prev
	lastNode.next = node
	node.prev = lastNode
	node.next = this.endNode
	this.endNode.prev = node

	if len(this.storage) > this.cap {
		evictNode := this.startNode.next
		this.startNode.next = evictNode.next
		evictNode.next.prev = this.startNode
		delete(this.storage, evictNode.key)
	}
}