package japanese

import "fmt"

type Word struct {
	Kanji string
	Kana  string
}

// LastKana returns the last kana in a word.
func (w *Word) LastKana() string {
	k := []rune(w.Kana)

	return string(k[len(k)-1:])
}

func (v *Verb) addEnd(end string) Word {
	r, k := v.AllButLast()

	return Word{r + end, k + end}
}

// AllButLast returns the kanji and kana readings of
// a word without the last character.
func (w *Word) AllButLast() (kanji, kana string) {
	kr := []rune(w.Kanji)
	kar := []rune(w.Kana)

	return string(kr[:len(kr)-1]), string(kar[:len(kar)-1])
}

// A Verb is a Japanese verb.
type Verb struct {
	Type string
	Word
}

func (v *Verb) stem(u map[string]string) (w Word, err error) {
	switch v.Kanji {
	case "する":
		return Word{"し", "し"}, nil
	}

	r, k := v.AllButLast()
	switch v.Type {
	case "る":
		return Word{r, k}, nil
	case "う":
		lastKana := v.LastKana()

		if val, ok := u[lastKana]; !ok {
			return w, fmt.Errorf("Attempt to get stem of U verb with invalid ending %s. Valid endings: %v", lastKana, u)
		} else {
			return v.addEnd(val), nil
		}
	}

	return w, fmt.Errorf("Attempt to get stem of verb with invalid ending.")
}

// Stem returns the stem of a verb.
// (食べる: 食べ; 言う:   言い)
func (v *Verb) Stem() (w Word, err error) {
	switch v.Kanji {
	case "来る":
		return Word{"来", "き"}, nil
	}

	m := map[string]string{
		"す": "し",
		"く": "き",
		"ぐ": "ぎ",
		"む": "み",
		"ぶ": "び",
		"ぬ": "に",
		"る": "り",
		"う": "い",
		"つ": "ち",
	}

	word, err := v.stem(m)
	if err != nil {
		return w, err
	}

	return word, nil
}

// ShortStem returns the short stem of a verb.
// (食べる: 食べ; 言う:   言わ)
func (v *Verb) ShortStem() (w Word, err error) {
	switch v.Kanji {
	case "来る":
		return Word{"来", "こ"}, nil
	}

	m := map[string]string{
		"す": "さ",
		"く": "か",
		"ぐ": "が",
		"む": "ま",
		"ぶ": "ば",
		"ぬ": "な",
		"る": "ら",
		"う": "わ",
		"つ": "た",
	}

	word, err := v.stem(m)
	if err != nil {
		return w, err
	}

	return word, nil
}

// TeForm returns the te-form of a verb.
func (v *Verb) TeForm() (word Word, err error) {
	switch v.Kanji {
	case "する":
		return Word{"して", "して"}, nil
	case "来る":
		return Word{"来て", "きて"}, nil
	case "行く":
		return v.addEnd("って"), nil
	}

	switch v.Type {
	case "る":
		return v.addEnd("て"), nil
	case "う":
		l := v.LastKana()

		switch l {
		case "す":
			return v.addEnd("して"), nil
		case "く":
			return v.addEnd("いて"), nil
		case "ぐ":
			return v.addEnd("いで"), nil
		case "む", "ぶ", "ぬ":
			return v.addEnd("んで"), nil
		case "る", "う", "つ":
			return v.addEnd("って"), nil
		default:
			return word, fmt.Errorf("Attempt to get te form of verb with invalid ending %s.", l)
		}
	}
	return word, fmt.Errorf("Attempt to get te form of verb with invalid ending")
}

// Negative returns the negative form of a Verb.
func (v *Verb) Negative() (word Word, err error) {
	switch v.Kanji {
	case "ある":
		return Word{"ない", "ない"}, nil
	}

	w, err := v.ShortStem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "ない", w.Kana + "ない"}, nil
}

// NegativePast returns the negative past form of a Verb.
func (v *Verb) NegativePast() (word Word, err error) {
	switch v.Kanji {
	case "ある":
		return Word{"なかった", "なかった"}, nil
	}

	w, err := v.ShortStem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "なかった", w.Kana + "なかった"}, nil
}

// NegativePolite returns the negative polite form of a Verb.
func (v *Verb) NegativePolite() (word Word, err error) {
	w, err := v.Stem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "ません", w.Kana + "ません"}, nil
}

// NegativePastPolite returns the negative past polite form of a Verb.
func (v *Verb) NegativePastPolite() (word Word, err error) {
	w, err := v.Stem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "ませんでした", w.Kana + "ませんでした"}, nil
}

// Past returns the past tense of a Verb.
func (v *Verb) Past() (word Word, err error) {
	switch v.Kanji {
	case "する":
		return Word{"した", "した"}, nil
	case "来る":
		return Word{"来た", "きた"}, nil
	case "行く":
		return v.addEnd("った"), nil
	}

	switch v.Type {
	case "る":
		return v.addEnd("た"), nil
	case "う":

		lastKana := v.LastKana()

		switch lastKana {
		case "す":
			return v.addEnd("した"), nil
		case "く":
			return v.addEnd("いた"), nil
		case "ぐ":
			return v.addEnd("いだ"), nil
		case "む", "ぶ", "ぬ":
			return v.addEnd("んだ"), nil
		case "る", "う", "つ":
			return v.addEnd("った"), nil
		}
	}
	return v.Word, fmt.Errorf("Attempt to get past tense of non-verb")
}

// PastPolite returns the past polite tense of a Verb.
func (v *Verb) PastPolite() (word Word, err error) {
	w, err := v.Stem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "ました", w.Kana + "ました"}, nil
}

func (v *Verb) te(end string) (word Word, err error) {
	w, err := v.TeForm()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + end, w.Kana + end}, nil
}

// Progressive returns the te postive
// form of a Verb.
func (v *Verb) Progressive() (word Word, err error) {
	w, err := v.te("いる")
	if err != nil {
		return word, err
	}
	return w, nil
}

// ProgressiveNegative returns the te negative
// form of a Verb.
func (v *Verb) ProgressiveNegative() (word Word, err error) {
	w, err := v.te("いない")
	if err != nil {
		return word, err
	}
	return w, nil
}

// ProgressivePolite returns the te positive
// polite form of a Verb.
func (v *Verb) ProgressivePolite() (word Word, err error) {
	w, err := v.te("います")
	if err != nil {
		return word, err
	}
	return w, nil
}

// ProgressiveNegativePolite returns the te negative
// polite form of a Verb.
func (v *Verb) ProgressiveNegativePolite() (word Word, err error) {
	w, err := v.te("いません")
	if err != nil {
		return word, err
	}
	return w, nil
}

// ProgressiveShort returns the shortened
// te positive form of a Verb.
func (v *Verb) ProgressiveShort() (word Word, err error) {
	w, err := v.te("る")
	if err != nil {
		return word, err
	}
	return w, nil
}

// ProgressiveShortNegative returns the shortened
// te negative form of a Verb.
func (v *Verb) ProgressiveShortNegative() (word Word, err error) {
	w, err := v.te("ない")
	if err != nil {
		return word, err
	}
	return w, nil
}

// PotentialStem returns the potential stem of a verb.
func (v *Verb) PotentialStem() (word Word, err error) {
	switch v.Kanji {
	case "する":
		return Word{"でき", "でき"}, nil
	case "来る":
		return Word{"来られ", "こられ"}, nil
	}
	switch v.Type {
	case "る":
		return v.addEnd("られ"), nil
	case "う":
		lastKana := v.LastKana()

		switch lastKana {
		case "す":
			return v.addEnd("せ"), nil
		case "く":
			return v.addEnd("け"), nil
		case "ぐ":
			return v.addEnd("げ"), nil
		case "む":
			return v.addEnd("め"), nil
		case "ぶ":
			return v.addEnd("べ"), nil
		case "ぬ":
			return v.addEnd("ね"), nil
		case "る":
			return v.addEnd("れ"), nil
		case "う":
			return v.addEnd("え"), nil
		case "つ":
			return v.addEnd("て"), nil
		}
	}
	return word, fmt.Errorf("Attempt to get potential stem of verb with invalid ending.")
}

func (v *Verb) Potential() (word Word, err error) {
	w, err := v.PotentialStem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "る", w.Kana + "る"}, nil
}

func (v *Verb) PotentialNegative() (word Word, err error) {
	w, err := v.PotentialStem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "ない", w.Kana + "ない"}, nil
}

func (v *Verb) PotentialPolite() (word Word, err error) {
	w, err := v.PotentialStem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "ます", w.Kana + "ます"}, nil
}

func (v *Verb) PotentialNegativePolite() (word Word, err error) {
	w, err := v.PotentialStem()
	if err != nil {
		return word, err
	}

	return Word{w.Kanji + "ません", w.Kana + "ません"}, nil
}

type Adjective struct {
	Word
}

type NaAdjective struct {
	Adjective
}
