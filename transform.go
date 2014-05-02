package japanese

import (
	"errors"
	// "log"
)

func SplitEnding(verb string) (root, ending string) {
	ending = IdentifyEnding(verb)
	li := len(verb) - len(ending)
	return verb[:li], verb[li:]
}

// verbify takes the last letter of a word and turns
// it into a -u sound
func verbify(word string) (verb string, err error) {
	if len(word) == 0 {
		return word, nil
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
		return string(wr) + u, nil
	}
	return word, errors.New("Could not change to verb")
}

// DictionaryForm gets the dictionary form of any verb
func DictionaryForm(verb string) (godan string, ichidan string) {
	root, ending := SplitEnding(verb)
	secondRoot, secondEnding := SplitEnding(root)

	var err error
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
			godan, err = verbify(root)
			if err != nil {
				//log.Println("Could not turn into godan verb root.")
			}
		}
	}

	// handle ichidan verbs

	if secondEnding == Ru || secondEnding == U {
		ichidan = root
	} else if secondEnding == Te {
		ichidan = secondRoot + Ru
	} else {
		ichidan = root + Ru
	}

	return godan, ichidan
}
