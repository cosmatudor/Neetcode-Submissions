type ListTweets struct {
	tweetId int
	userId int
	next *ListTweets
}

type Twitter struct {
    tweets *ListTweets
	network map[int]map[int]struct{}
}


func Constructor() Twitter {
    return Twitter {
		tweets: nil,
		network: make(map[int]map[int]struct{}),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
	if this.tweets == nil {
		this.tweets = &ListTweets{
			tweetId: tweetId,
			userId: userId,
			next: nil,
		}
		return
	}

	newTweet := &ListTweets{
		tweetId: tweetId,
		userId: userId,
		next: this.tweets,
	}
	this.tweets = newTweet
}


func (this *Twitter) GetNewsFeed(userId int) []int {
	allTweetsForUsersFeed := []int{}
	tweet := this.tweets
    for tweet != nil {
		followees := this.network[userId] 
		_, ok := followees[tweet.userId]
		if tweet.userId == userId || ok {
			allTweetsForUsersFeed = append(allTweetsForUsersFeed, tweet.tweetId)
		}

		tweet = tweet.next
	}

	lenTweets := min(10, len(allTweetsForUsersFeed))
	return allTweetsForUsersFeed[:lenTweets]
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
	if _, ok := this.network[followerId]; !ok {
		this.network[followerId] = make(map[int]struct{})
	} 
    this.network[followerId][followeeId] = struct{}{}
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    delete(this.network[followerId], followeeId)
}
