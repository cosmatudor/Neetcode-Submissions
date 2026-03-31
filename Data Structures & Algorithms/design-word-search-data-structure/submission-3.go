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
    return dfs(this, word, 0)
}

func dfs(node *WordDictionary, word string, index int) bool {
    if index == len(word) {
        return node.IsWord
    }

    char := word[index]

    if char == '.' {
        for _, child := range node.Children {
            if dfs(child, word, index+1) {
                return true
            }
        }
        return false
    }

    if child, exist := node.Children[char]; exist {
        return dfs(child, word, index+1)
    }

    return false
}