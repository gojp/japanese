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

func (w *Word) GetRoot() string {
	return w.kana
}

/* === VERBS === */

type Verb struct {
	Word
}

type RuVerb struct {
	Verb
}

func (v *RuVerb) Negative() string {
	// drop the る and attach ない
	kanjiRune := []rune(v.kanji)
	// one exception for case "ある":
	if v.kanji == "ある" {
		return "ない"
	}
	return string(kanjiRune[:len(kanjiRune)-1]) + "ない"
}

type UVerb struct {
	Verb
}

func (v *UVerb) Negative() string {
	kanjiRune := []rune(v.kanji)
	lastCharacter := string(kanjiRune[len(kanjiRune)-1:])
	restOfWord := string(kanjiRune[:len(kanjiRune)-1])

	// if verb ends in う, replace う with わない
	if lastCharacter == "う" {
		return restOfWord + "わない"
		// otherwise replace with the -a equivalent
	} else {
		original := []string{"つ", "く", "ゅ", "す", "ぬ", "ふ", "む", "ゆ", "ぐ", "ず", "づ", "ぶ", "ぷ", "る"}
		replace := []string{"た", "か", "ゃ", "さ", "な", "は", "ま", "や", "が", "ざ", "ざ", "ば", "ぱ", "ら"}
		for i, o := range original {
			if o == lastCharacter {
				return restOfWord + replace[i] + "ない"
			}
		}
	}
	// return original word if all else fails
	return v.kanji
}

type ExceptionVerb struct {
	Verb
}

func (v *ExceptionVerb) Negative() string {
	switch v.kanji {
	case "する":
		return "しない"
	case "くる":
		return "こない"
	}
	return v.kanji
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
