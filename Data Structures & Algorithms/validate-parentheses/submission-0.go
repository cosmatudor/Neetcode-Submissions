type stack struct{
    values []byte
}

func (s *stack) push(x byte) {
    s.values = append(s.values, x)
}

func (s *stack) pop() byte {
    top := s.values[len(s.values)-1]
    s.values = s.values[:len(s.values)-1]

    return top
}

func (s *stack) size() int {
    return len(s.values)
}

func isValid(ss string) bool {
    if len(ss) == 0 {
        return true
    }

    s := stack{}

    for _, br := range ss {
        if string(br) == "(" || string(br) == "[" || string(br) == "{" {
            s.push(byte(br))
        } else {
            top := s.pop()

            shouldContinue := false
            switch string(br) {
            case ")":
                if string(top) == "(" {
                    shouldContinue = true
                }
            case "]":
                if string(top) == "[" {
                    shouldContinue = true
                }
            case "}":
                if string(top) == "{" {
                    shouldContinue = true
                }
            }

            if !shouldContinue {
                return false
            }
        }
    }

    if s.size() > 0 {
        return false
    }

    return true
}
