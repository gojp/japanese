package japanese

/* === WORDS === */

type Word struct {
	kanji string
	kana  string
}

// LastKana returns the last kana in a word.
func (w *Word) LastKana() string {
	kanaRune := []rune(w.kana)

	return string(kanaRune[len(kanaRune)-1:])
}

// AllButLast returns all but he last character of a word
func (w *Word) AllButLast() (kanji, kana string) {
	kanjiRune := []rune(w.kanji)
	kanaRune := []rune(w.kana)

	return string(kanjiRune[:len(kanjiRune)-1]), string(kanaRune[:len(kanaRune)-1])
}

/* === VERBS === */

// A Verb is a Japanese verb.
type Verb struct {
	Type string
	Word
}

func (v *Verb) addEnd(end string) Word {
	r, k := v.AllButLast()

	return Word{r + end, k + end}
}

// TeForm returns the te-form of a verb.
func (v *Verb) TeForm() Word {
	switch v.kanji {
	case "する":
		return Word{"して", "して"}
	case "来る":
		return Word{"来て", "きて"}
	case "行く":
		return v.addEnd("って")
	}

	switch v.Type {
	case "る":
		return v.addEnd("て")
	case "う":
		l := v.LastKana()

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
			return v.Word
		}
	}
	return v.Word
}

// Negative returns the negative form of a Verb.
func (v *Verb) Negative() Word {
	switch v.kanji {
	case "する":
		return Word{"しない", "しない"}
	case "来る":
		return Word{"来ない", "こない"}
	case "ある":
		return Word{"ない", "ない"}
	}

	switch v.Type {
	case "る":
		return v.addEnd("ない")
	case "う":
		lastKana := v.LastKana()

		if lastKana == "う" {
			return v.addEnd("わない")
		} else {
			m := map[string]string{
				"つ": "た", "く": "か", "ゅ": "ゃ", "す": "さ",
				"ぬ": "な", "ふ": "は", "む": "ま", "ゆ": "や",
				"ぐ": "が", "ず": "ざ", "づ": "ざ", "ぶ": "ば",
				"ぷ": "ぱ", "る": "ら"}
			extra := m[lastKana] + "ない"
			return v.addEnd(extra)
		}
	}
	return v.Word
}

// NegativePolite returns the negative polite form of a Verb.
func (v *Verb) NegativePolite() Word {
	w := v.Stem()

	return Word{w.kanji + "ません", w.kana + "ません"}
}

// NegativePastPolite returns the negative past polite form of a Verb.
func (v *Verb) NegativePastPolite() Word {
	w := v.Stem()

	return Word{w.kanji + "ませんでした", w.kana + "ませんでした"}
}

// Past returns the past tense of a Verb.
func (v *Verb) Past() Word {
	switch v.kanji {
	case "する":
		return Word{"した", "した"}
	case "来る":
		return Word{"来た", "きた"}
	case "行く":
		return v.addEnd("った")
	}

	switch v.Type {
	case "る":
		return v.addEnd("た")
	case "う":

		lastKana := v.LastKana()

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
	}
	return v.Word
}

// Stem returns the stem of a verb.
func (v *Verb) Stem() Word {
	switch v.kanji {
	case "する":
		return Word{"し", "し"}
	case "来る":
		return Word{"来", "き"}
	}

	r, k := v.AllButLast()
	switch v.Type {
	case "る":
		return Word{r, k}
	case "う":
		lastKana := v.LastKana()

		switch lastKana {
		case "す":
			return v.addEnd("し")
		case "く":
			return v.addEnd("き")
		case "ぐ":
			return v.addEnd("ぎ")
		case "む":
			return v.addEnd("み")
		case "ぶ":
			return v.addEnd("び")
		case "ぬ":
			return v.addEnd("に")
		case "る":
			return v.addEnd("り")
		case "う":
			return v.addEnd("い")
		case "つ":
			return v.addEnd("ち")
		}
	}
	return v.Word
}

// PastPolite returns the past polite tense of a Verb.
func (v *Verb) PastPolite() Word {
	w := v.Stem()

	return Word{w.kanji + "ました", w.kana + "ました"}
}

func (v *Verb) te(end string) Word {
	w := v.TeForm()

	return Word{w.kanji + end, w.kana + end}
}

// Progressive returns the te postive
// form of a Verb.
func (v *Verb) Progressive() Word {
	return v.te("いる")
}

// ProgressiveNegative returns the te negative
// form of a Verb.
func (v *Verb) ProgressiveNegative() Word {
	return v.te("いない")
}

// ProgressivePolite returns the te positive
// polite form of a Verb.
func (v *Verb) ProgressivePolite() Word {
	return v.te("います")
}

// ProgressiveNegativePolite returns the te negative
// polite form of a Verb.
func (v *Verb) ProgressiveNegativePolite() Word {
	return v.te("いません")
}

// ProgressiveShort returns the shortened
// te positive form of a Verb.
func (v *Verb) ProgressiveShort() Word {
	return v.te("る")
}

// ProgressiveShortNegative returns the shortened
// te negative form of a Verb.
func (v *Verb) ProgressiveShortNegative() Word {
	return v.te("ない")
}

/* === ADJECTIVES === */

type Adjective struct {
	Word
}

type NaAdjective struct {
	Adjective
}
