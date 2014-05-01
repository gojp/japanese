package japanese

/* === WORDS === */

type Word struct {
	kanji string
	kana  string
}

func (w *Word) GetWord() Word {
	return *w
}

func (w *Word) GetRoot() string {
	return w.kana
}

func (w *Word) GetLastKana() string {
	// get the last kana character in a word
	kanaRune := []rune(w.kana)
	return string(kanaRune[len(kanaRune)-1:])
}

func (w *Word) GetAllButLast() (kanji, kana string) {
	// get all but the last character of a word
	kanjiRune := []rune(w.kanji)
	kanaRune := []rune(w.kana)
	return string(kanjiRune[:len(kanjiRune)-1]), string(kanaRune[:len(kanaRune)-1])
}

/* === VERBS === */

type Verb struct {
	Type string
	Word
}

func (v *Verb) addEnd(end string) Word {
	r, k := v.GetAllButLast()

	return Word{r + end, k + end}
}

func (v *Verb) TeForm() Word {
	switch v.Type {
	case "る":
		return v.addEnd("て")
	case "う":
		l := v.GetLastKana()

		if v.kanji == "行く" {
			return v.addEnd("って")
		}

		switch l {
		case "す":
			return v.addEnd("して")
		case "く":
			return v.addEnd("いて")
		case "ぐ":
			return v.addEnd("いで")
		case "む", "ぶ", "ぬ":
			return v.addEnd("んで")
		case "る", "う", "つ":
			return v.addEnd("って")
		default:
			return v.GetWord()
		}
	}
	return v.GetWord()
}

func (v *Verb) Negative() Word {
	switch v.kanji {
	case "する":
		return Word{"しない", "しない"}
	case "くる":
		return Word{"こない", "こない"}
	}

	switch v.Type {
	case "る":
		return v.addEnd("ない")
	case "う":
		lastKana := v.GetLastKana()

		// one exception for case "ある":
		if v.kanji == "ある" {
			return Word{"ない", "ない"}
		}

		// if verb ends in う, replace う with わない
		if lastKana == "う" {
			return v.addEnd("わない")
			// otherwise replace with the -a equivalent
		} else {
			original := []string{"つ", "く", "ゅ", "す", "ぬ", "ふ", "む", "ゆ", "ぐ", "ず", "づ", "ぶ", "ぷ", "る"}
			replace := []string{"た", "か", "ゃ", "さ", "な", "は", "ま", "や", "が", "ざ", "ざ", "ば", "ぱ", "ら"}
			for i, o := range original {
				if o == lastKana {
					extra := replace[i] + "ない"
					return v.addEnd(extra)
				}
			}
		}
		// return original word if all else fails
		return v.GetWord()
	}
	return v.GetWord()
}

// Past returns the past tense of a Verb.
func (v *Verb) Past() Word {
	switch v.kanji {
	case "する":
		return Word{"した", "した"}
	case "くる":
		return Word{"きた", "きた"}
	}

	switch v.Type {
	case "る":
		return v.addEnd("た")
	case "う":

		lastKana := v.GetLastKana()

		// 行く is only an exception for this rule
		if v.kanji == "行く" {
			return Word{"行った", "いった"}
		}

		switch lastKana {
		case "す":
			return v.addEnd("した")
		case "く", "ぐ":
			return v.addEnd("いた")
		case "む", "ぶ", "ぬ":
			return v.addEnd("んだ")
		case "る", "う", "つ":
			return v.addEnd("った")
		}

		// otherwise we just return the same word, because we don't know what to do:
		return v.GetWord()
	}
	return v.GetWord()
}

func (v *Verb) progressive(end string) Word {
	w := v.TeForm()

	return Word{w.kanji + end, w.kana + end}
}

// Progressive returns the progressive postive
// form of a Verb.
func (v *Verb) Progressive() Word {
	return v.progressive("いる")
}

// ProgressiveNegative returns the progressive negative
// form of a Verb.
func (v *Verb) ProgressiveNegative() Word {
	return v.progressive("いない")
}

// ProgressivePolite returns the progressive positive
// polite form of a Verb.
func (v *Verb) ProgressivePolite() Word {
	return v.progressive("います")
}

// ProgressiveNegativePolite returns the progressive negative
// polite form of a Verb.
func (v *Verb) ProgressiveNegativePolite() Word {
	return v.progressive("いません")
}

// ProgressiveShort returns the shortened
// progressive positive form of a Verb.
func (v *Verb) ProgressiveShort() Word {
	return v.progressive("る")
}

// ProgressiveShortNegative returns the shortened
// progressive negative form of a Verb.
func (v *Verb) ProgressiveShortNegative() Word {
	return v.progressive("ない")
}

/* === ADJECTIVES === */

type Adjective struct {
	Word
}

type NaAdjective struct {
	Adjective
}
