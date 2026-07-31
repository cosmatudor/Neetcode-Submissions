type MaxHeap []*ListTweets
func (h MaxHeap) Len() int { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	return h[i].count > h[j].count
}
func (h MaxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *MaxHeap) Push(x any) {
    *h = append(*h, x.(*ListTweets))
}
func (h *MaxHeap) Pop() any {
	n := len(*h)
	val := (*h)[n-1]
	*h = (*h)[:n-1]
	return val
}

type ListTweets struct {
	tweetId int
	userId int
	count int
	next *ListTweets
}

type Twitter struct {
	counter int
    tweets map[int]*ListTweets
	network map[int]map[int]struct{}
}


func Constructor() Twitter {
    return Twitter {
		counter: 0,
		tweets: make(map[int]*ListTweets),
		network: make(map[int]map[int]struct{}),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	this.counter++
	if this.tweets[userId] == nil {
		this.tweets[userId] = &ListTweets{
			tweetId: tweetId,
			userId: userId,
			count: this.counter,
			next: nil,
		}
		return
	}

	newTweet := &ListTweets{
		tweetId: tweetId,
		userId: userId,
		count: this.counter,
		next: this.tweets[userId],
	}
	this.tweets[userId] = newTweet
	fmt.Println("AICI", this.tweets[userId])
	fmt.Println("AICI", this.tweets[userId].next)
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	allTweetsForUsersFeed := []int{}
	h := MaxHeap{}
	heap.Init(&h)

	if this.tweets[userId] != nil {
		heap.Push(&h, this.tweets[userId])
	}
	followees := this.network[userId]
	fmt.Println(followees)
	for followee := range followees {
		if this.tweets[followee] != nil {
			heap.Push(&h, this.tweets[followee])
		}
	}

	for len(allTweetsForUsersFeed) != 10 && h.Len() != 0 {
		tweet := heap.Pop(&h).(*ListTweets)
		allTweetsForUsersFeed = append(allTweetsForUsersFeed, tweet.tweetId)
		if tweet.next != nil {
			heap.Push(&h, tweet.next)
		}
	}
	
	return allTweetsForUsersFeed
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if _, ok := this.network[followerId]; !ok {
		this.network[followerId] = make(map[int]struct{})
	} 
	if followerId == followeeId {
		return
	}
    this.network[followerId][followeeId] = struct{}{}
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.network[followerId], followeeId)
}
