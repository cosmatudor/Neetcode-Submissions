type PrefixTree struct {
	Val byte
	Children map[byte]*PrefixTree
	IsWord bool
}

func Constructor() PrefixTree {
    return PrefixTree{
		Val: 0,
		Children: make(map[byte]*PrefixTree),
		IsWord: false,
	}
}

func (this *PrefixTree) Insert(word string) {
	curr := this
	for i, chRune := range word {
		ch := byte(chRune)
		if _, exist := curr.Children[ch]; !exist {
			curr.Children[ch] = &PrefixTree{
				Val: ch,
				Children: make(map[byte]*PrefixTree),
				IsWord: i == len(word) - 1,
			}
		}

		curr, _ = curr.Children[ch]
	}
	curr.IsWord = true
}

func (this *PrefixTree) Search(word string) bool {
	curr := this
	for _, chRune := range word {
		ch := byte(chRune)
		if _, exist := curr.Children[ch]; !exist {
			return false
		}
		curr, _ = curr.Children[ch]
	}

	return curr.IsWord 
}

func (this *PrefixTree) StartsWith(prefix string) bool {
	curr := this
	for _, chRune := range prefix {
		ch := byte(chRune)
		if _, exist := curr.Children[ch]; !exist {
			return false
		}
		curr, _ = curr.Children[ch]
	}

	return true
}
