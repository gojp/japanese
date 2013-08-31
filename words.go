package main

import "fmt"

/* === WORDS === */

type Word struct {
	kanji   string
	kana    string
	english string
}

func (w *Word) Print() {
	fmt.Println(w.kana, w.kanji)
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

func (v *RuVerb) Negative() Word {
	// one exception for case "ある":
	if v.kanji == "ある" {
		return Word{"ない", "ない", "not " + v.english}
	}
	// drop the る and attach ない
	restOfKanji, restOfKana := v.GetAllButLast()
	return Word{restOfKanji + "ない", restOfKana + "ない", "not " + v.english}
}

func (v *RuVerb) Past() Word {
	// drop the る and attach た
	restOfKanji, restOfKana := v.GetAllButLast()
	return Word{restOfKanji + "た", restOfKana + "た", "did " + v.english}
}

type UVerb struct {
	Verb
}

func (v *UVerb) Negative() Word {
	lastCharacter := v.GetLastKana()
	restOfKanji, restOfKana := v.GetAllButLast()

	// if verb ends in う, replace う with わない
	if lastCharacter == "う" {
		extra := "わない"
		return Word{restOfKanji + extra, restOfKana + extra, "not " + v.english}
		// otherwise replace with the -a equivalent
	} else {
		original := []string{"つ", "く", "ゅ", "す", "ぬ", "ふ", "む", "ゆ", "ぐ", "ず", "づ", "ぶ", "ぷ", "る"}
		replace := []string{"た", "か", "ゃ", "さ", "な", "は", "ま", "や", "が", "ざ", "ざ", "ば", "ぱ", "ら"}
		for i, o := range original {
			if o == lastCharacter {
				extra := replace[i] + "ない"
				return Word{restOfKanji + extra, restOfKana + extra, "not " + v.english}
			}
		}
	}
	// return original word if all else fails
	return v.GetWord()
}

func (v *UVerb) Past() Word {
	/*
		Get the past-tense form of an U-verb
	*/

	lastCharacter := v.GetLastKana()
	restOfKanji, restOfKana := v.GetAllButLast()

	// handle exceptions first:
	if v.kanji == "する" {
		return Word{"した", "した", "did" + v.english}
	}
	if v.kanji == "くる" {
		return Word{"きた", "きた", "did" + v.english}
	}
	// 行く is only an exception for this rule
	if v.kanji == "行く" {
		return Word{"行った", "いった", "did" + v.english}
	}

	switch lastCharacter {
	case "す":
		return Word{restOfKanji + "した", restOfKana + "した", "did " + v.english}
	case "く":
	case "ぐ":
		return Word{restOfKanji + "いた", restOfKana + "いた", "did " + v.english}
	case "む":
	case "ぶ":
	case "ぬ":
		return Word{restOfKanji + "んだ", restOfKana + "んだ", "did " + v.english}
	case "る":
	case "う":
	case "つ":
		return Word{restOfKanji + "った", restOfKana + "った", "did " + v.english}
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
		return Word{"しない", "しない", "not do"}
	case "くる":
		return Word{"こない", "こない", "not come"}
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

func main() {
	ruverb := RuVerb{Verb{Word{"食べる", "たべる", "eat"}}}
	uverb := UVerb{Verb{Word{"買う", "かう", "buy"}}}
	uverb2 := UVerb{Verb{Word{"待つ", "まつ", "wait"}}}

	fmt.Println(ruverb.Negative(), "食べない")
	fmt.Println(uverb.Negative(), "買わない")
	fmt.Println(uverb2.Negative(), "待たない")
}
