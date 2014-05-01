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
	Word
}

type RuVerb struct {
	Verb
}

func (v *RuVerb) TeForm() Word {
	r, k := v.GetAllButLast()

	return Word{r + "て", k + "て"}
}

func (v *RuVerb) Negative() Word {
	// drop the る and attach ない
	restOfKanji, restOfKana := v.GetAllButLast()
	return Word{restOfKanji + "ない", restOfKana + "ない"}
}

func (v *RuVerb) Past() Word {
	// drop the る and attach た
	restOfKanji, restOfKana := v.GetAllButLast()
	return Word{restOfKanji + "た", restOfKana + "た"}
}

func (v *RuVerb) progressive(end string) Word {
	w := v.TeForm()

	return Word{w.kanji + end, w.kana + end}
}

// Progressive returns the progressive postive
// form of a RuVerb.
func (v *RuVerb) Progressive() Word {
	return v.progressive("いる")
}

// ProgressiveNegative returns the progressive negative
// form of a RuVerb.
func (v *RuVerb) ProgressiveNegative() Word {
	return v.progressive("いない")
}

// ProgressivePolite returns the progressive positive
// polite form of a RuVerb.
func (v *RuVerb) ProgressivePolite() Word {
	return v.progressive("います")
}

// ProgressiveNegativePolite returns the progressive negative
// polite form of a RuVerb.
func (v *RuVerb) ProgressiveNegativePolite() Word {
	return v.progressive("いません")
}

// ProgressiveShort returns the shortened
// progressive positive form of a RuVerb.
func (v *RuVerb) ProgressiveShort() Word {
	return v.progressive("る")
}

// ProgressiveShortNegative returns the shortened
// progressive negative form of a RuVerb.
func (v *RuVerb) ProgressiveShortNegative() Word {
	return v.progressive("ない")
}

type UVerb struct {
	Verb
}

func (v *UVerb) Negative() Word {
	lastCharacter := v.GetLastKana()
	restOfKanji, restOfKana := v.GetAllButLast()

	// one exception for case "ある":
	if v.kanji == "ある" {
		return Word{"ない", "ない"}
	}

	// if verb ends in う, replace う with わない
	if lastCharacter == "う" {
		extra := "わない"
		return Word{restOfKanji + extra, restOfKana + extra}
		// otherwise replace with the -a equivalent
	} else {
		original := []string{"つ", "く", "ゅ", "す", "ぬ", "ふ", "む", "ゆ", "ぐ", "ず", "づ", "ぶ", "ぷ", "る"}
		replace := []string{"た", "か", "ゃ", "さ", "な", "は", "ま", "や", "が", "ざ", "ざ", "ば", "ぱ", "ら"}
		for i, o := range original {
			if o == lastCharacter {
				extra := replace[i] + "ない"
				return Word{restOfKanji + extra, restOfKana + extra}
			}
		}
	}
	// return original word if all else fails
	return v.GetWord()
}

// Past gets the past-tense form of an U-verb
func (v *UVerb) Past() Word {
	lastCharacter := v.GetLastKana()
	restOfKanji, restOfKana := v.GetAllButLast()

	// 行く is only an exception for this rule
	if v.kanji == "行く" {
		return Word{"行った", "いった"}
	}

	switch lastCharacter {
	case "す":
		return Word{restOfKanji + "した", restOfKana + "した"}
	case "く", "ぐ":
		return Word{restOfKanji + "いた", restOfKana + "いた"}
	case "む", "ぶ", "ぬ":
		return Word{restOfKanji + "んだ", restOfKana + "んだ"}
	case "る", "う", "つ":
		return Word{restOfKanji + "った", restOfKana + "った"}
	}

	// otherwise we just return the same word, because we don't know what to do:
	return v.GetWord()
}

type ExceptionVerb struct {
	Verb
}

func (v *ExceptionVerb) Negative() Word {
	switch v.kanji {
	case "する":
		return Word{"しない", "しない"}
	case "くる":
		return Word{"こない", "こない"}
	}
	return v.GetWord()
}

func (v *ExceptionVerb) Past() Word {
	if v.kanji == "する" {
		return Word{"した", "した"}
	}
	if v.kanji == "くる" {
		return Word{"きた", "きた"}
	}
	return v.GetWord()
}

/* === ADJECTIVES === */

type Adjective struct {
	Word
}

type NaAdjective struct {
	Adjective
}
