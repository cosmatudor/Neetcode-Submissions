type WordDictionary struct {
    Val byte
	Children map[byte]*WordDictionary
	IsWord bool
}

func Constructor() WordDictionary {
    return WordDictionary {
		Val: 0,
		Children: make(map[byte]*WordDictionary),
		IsWord: false,
	}
}

func (this *WordDictionary) AddWord(word string)  {
    curr := this
	for i, chRune := range word {
		ch := byte(chRune)
		if _, exist := curr.Children[ch]; !exist {
			curr.Children[ch] = &WordDictionary{
				Val: ch,
				Children: make(map[byte]*WordDictionary),
				IsWord: i == len(word) - 1,
			}
		}

		curr, _ = curr.Children[ch]
	}
	curr.IsWord = true
}

func (this *WordDictionary) Search(word string) bool {
    return search(this, word, 0)
}

func search(this *WordDictionary, word string, index int) bool {
	if index == len(word) {
		return true
	}

	contains := false
	for _, val := range this.Children {
		contains = contains || search(val, word, index + 1)
	}

	ch := word[index]
	_, exist := this.Children[ch]

	return contains && (exist || ch == '.')
}
