package japanese

func SplitEnding(verb string) (root, ending string) {
	ending = IdentifyEnding(verb)
	li := len(verb) - len(ending)

	return verb[:li], verb[li:]
}

// verbify takes the last letter of a word and turns
// it into a -u sound
func verbify(word string) (verb string) {
	if len(word) == 0 {
		return word
	}
	conv := map[string]string{
		"い": "う",
		"き": "く",
		"ぎ": "ぐ",
		"ち": "つ",
		"り": "る",
		"み": "む",
		"び": "ぶ",
		"に": "ぬ",
		"し": "す",
	}
	wr := []rune(word)
	last, wr := string(wr[len(wr)-1:]), wr[:len(wr)-1]
	u, ok := conv[string(last)]
	if ok {
		return string(wr) + u
	}
	return word
}

// DictionaryForm gets the dictionary form of any verb
func DictionaryForm(verb string) (godan string, ichidan string) {
	root, ending := SplitEnding(verb)
	secondRoot, secondEnding := SplitEnding(root)

	// handle godan verbs
	switch ending {
	case Ru, U:
		godan = verb
		break
	default:
		if len(root) == 0 && ending == Nai {
			godan = "ある"
		} else if secondEnding == Ru || secondEnding == U {
			godan = root
		} else {
			godan = verbify(root)
		}
	}

	// handle ichidan verbs
	switch secondEnding {
	case Ru, U:
		ichidan = root
	case Te:
		ichidan = secondRoot + Ru
	default:
		ichidan = root + Ru
	}

	return godan, ichidan
}
