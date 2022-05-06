package app

func tokenize(s string) (tokens []string) {
	bstr := []byte(s)
	bstr = append(bstr, ' ')

	var (
		prevChar byte //previos char
		token    []byte
		split    = true
	)

	for _, char := range bstr {
		if char == ' ' && split {
			tokens = append(tokens, string(token))
			token = []byte{}
		} else {
			switch char {
			case '"':
				if prevChar == '\\' {
					token = append(token, char)
				} else {
					split = !split
				}
			case '\\':
				if prevChar == '\\' {
					token = append(token, char)
				}
			default:
				token = append(token, char)
			}
		}
		prevChar = char
	}
	return
}

type tObj struct {
	Method string
	Args   []string
}

func tokensToObject(s []string) *tObj {
	var t tObj
	if len(s) >= 2 {
		t.Method = s[0]
		t.Args = append(t.Args, s[1:]...)
		return &t
	}
	return nil
}
