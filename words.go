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
			m := map[string]string{
				"つ": "た", "く": "か", "ゅ": "ゃ", "す": "さ",
				"ぬ": "な", "ふ": "は", "む": "ま", "ゆ": "や",
				"ぐ": "が", "ず": "ざ", "づ": "ざ", "ぶ": "ば",
				"ぷ": "ぱ", "る": "ら"}
			extra := m[lastKana] + "ない"
			return v.addEnd(extra)
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

		if v.kanji == "行く" {
			return Word{"行った", "いった"}
		}

		switch lastKana {
		case "す":
			return v.addEnd("した")
		case "く":
			return v.addEnd("いた")
		case "ぐ":
			return v.addEnd("いだ")
		case "む", "ぶ", "ぬ":
			return v.addEnd("んだ")
		case "る", "う", "つ":
			return v.addEnd("った")
		}

		return v.GetWord()
	}

	return v.GetWord()
}

// PastPolite returns the past polite tense of a Verb.
func (v *Verb) PastPolite() Word {
	switch v.kanji {
	case "する":
		return Word{"しました", "しました"}
	case "くる":
		return Word{"きました", "きました"}
	}

	switch v.Type {
	case "る":
		return v.addEnd("ました")
	case "う":

		lastKana := v.GetLastKana()

		if v.kanji == "行く" {
			return Word{"行きました", "いきました"}
		}

		switch lastKana {
		case "す":
			return v.addEnd("しました")
		case "く":
			return v.addEnd("きました")
		case "ぐ":
			return v.addEnd("ぎました")
		case "む":
			return v.addEnd("みました")
		case "ぶ":
			return v.addEnd("びました")
		case "ぬ":
			return v.addEnd("にました")
		case "る":
			return v.addEnd("りました")
		case "う":
			return v.addEnd("いました")
		case "つ":
			return v.addEnd("ちました")
		}

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
