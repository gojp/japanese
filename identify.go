package japanese

import (
	"strings"

	"github.com/gojp/kana"
)

// possible verb forms
const (
	TeForm = iota
	Infinitive
	PresentIndicative
	Presumptive
	Imperative
	PastIndicative
	PastPresumptive
	PresentProgressive
	PastProgressive
	Provisional
	Conditional
	Potential
	Causative
	Unknown
)

// politeness levels
const (
	Plain = iota
	Polite
)

// positive / negative
const (
	Positive = iota
	Negative
)

// verb classes
const (
	Ichidan = iota
	Godan
	Irregular
)

// verb endings
const (
	Te         = "て"
	Teiru      = "ている"
	Teimasu    = "ています"
	Teru       = "てる"
	Temasu     = "てます"
	Teita      = "ていた"
	Teimashita = "ていました"
	Teta       = "てた"
	Temashita  = "てました"

	Ru    = "る"
	U     = "う"
	Masu  = "ます"
	Nai   = "ない"
	Masen = "ません"

	Ta           = "た"
	Mashita      = "ました"
	Katta        = "かった"
	Nakatta      = "なかった"
	MasenDeshita = "ませんでした"

	You       = "よう"
	Darou     = "だろう"
	Naidarou  = "ないだろう"
	Mashou    = "ましょう"
	Deshou    = "でしょう"
	Naideshou = "ないでしょう"

	Ro           = "ろ"
	Kudasai      = "ください" // te-kudasai
	Naidekudasai = "ないでください"
	Na           = "な"

	Tarou         = "たろう"
	Tadarou       = "ただろう"
	Nakattadarou  = "なかっただろう"
	Tadeshou      = "たでしょう"
	Nakattadeshou = "なかったでしょう"

	Reba     = "れば"
	Eba      = "えば"
	Keba     = "けば"
	Geba     = "げば"
	Seba     = "せば"
	Teba     = "てば"
	Neba     = "ねば"
	Beba     = "べば"
	Meba     = "めば"
	Kereba   = "ければ"
	Deareba  = "であれば"
	Nakereba = "なければ"

	Rareru = "られる"
	Eru    = "える"
	Keru   = "ける"
	Geru   = "げる"
	Seru   = "せる"
	// Teru   = "てる"
	Neru = "ねる"
	Beru = "べる"
	Meru = "める"
	Reru = "れる"
)

type FormEnding struct {
	form    int
	endings []string
}

var (
	PresentIndicativeEnding  = FormEnding{PresentIndicative, []string{Ru, U, Masu, Nai, Masen}}
	PresumptiveEnding        = FormEnding{Presumptive, []string{You, Darou, Naidarou, Mashou, Deshou, Naideshou}}
	ImperativeEnding         = FormEnding{Imperative, []string{Ro, Kudasai, Naidekudasai, Na}}
	PastIndicativeEnding     = FormEnding{PastIndicative, []string{Ta, Mashita, Katta, Nakatta, MasenDeshita}}
	PastPresumptiveEnding    = FormEnding{PastPresumptive, []string{Tarou, Tadarou, Nakattadarou, Tadeshou, Nakattadeshou}}
	PresentProgressiveEnding = FormEnding{PresentProgressive, []string{Teiru, Teimasu, Teru, Temasu}}
	PastProgressiveEnding    = FormEnding{PastProgressive, []string{Teita, Teimashita, Teta, Temashita}}
	ProvisionalEnding        = FormEnding{Provisional, []string{Reba, Eba, Keba, Geba, Seba, Teba, Neba, Beba, Meba, Kereba, Deareba, Nakereba}}
	ConditionalEnding        = FormEnding{Conditional, []string{Rareru, Eru, Seru, Teru, Neru, Beru, Meru, Reru}}
	PotentialEnding          = FormEnding{}
	CausativeEnding          = FormEnding{}
	InfinitiveEnding         = FormEnding{}
	TeFormEnding             = FormEnding{TeForm, []string{Te}}
)

var EndingForms = []FormEnding{
	TeFormEnding,
	InfinitiveEnding,
	PresentIndicativeEnding,
	PresumptiveEnding,
	ImperativeEnding,
	PastIndicativeEnding,
	PastPresumptiveEnding,
	PresentProgressiveEnding,
	PastProgressiveEnding,
	ProvisionalEnding,
	ConditionalEnding,
	PotentialEnding,
	CausativeEnding,
}

// hasAnySuffix
func hasAnySuffix(s string, endings ...string) bool {
	for _, e := range endings {
		if strings.HasSuffix(s, e) {
			return true
		}
	}
	return false
}

// IdentifyEnding gets the longest ending for the verb
// from the list of known verb endings.
func IdentifyEnding(verb string) (ending string) {
	// Should make this a trie at some point, but... MVP
	for _, ef := range EndingForms {
		for _, e := range ef.endings {
			if strings.HasSuffix(verb, e) && len(e) > len(ending) {
				ending = e
			}
		}
	}
	return ending
}

// IdentifyForm tries to identify the verb form.
// It returns the integer form constant
func IdentifyForm(verb string) (form int) {
	switch {
	// TODO
	//case hasAnySuffix(verb, Keru, Geru) && isIchidan(verb):
	//	return PresentIndicative
	case hasAnySuffix(verb, Reru, Eru, Geru, Seru, Neru, Beru, Meru):
		return Conditional
	case hasAnySuffix(verb, Teru):
		if !kana.IsKanji(strings.Split(verb, Teru)[0]) {
			return PresentProgressive
		}
		return Conditional
	case hasAnySuffix(verb, Temasu):
		if !kana.IsKanji(strings.Split(verb, Temasu)[0]) {
			return PresentProgressive
		}
		return Conditional
	case hasAnySuffix(verb, Teiru, Teimasu):
		return PresentProgressive
	case hasAnySuffix(verb, Teita, Teimashita, Teta, Temashita):
		return PastProgressive
	case hasAnySuffix(verb, Ru, U, Masu, Nai, Masen):
		return PresentIndicative
	case hasAnySuffix(verb, Ta, Mashita, Nakatta, MasenDeshita):
		return PastIndicative
	case hasAnySuffix(verb, Reba, Eba, Keba, Geba, Seba, Teba, Neba, Beba, Meba, Kereba, Deareba, Nakereba):
		return Provisional
	default:
		return Unknown
	}

	return Unknown
}

// IdentifyPositivity determines whether the verb is
// in affirmative or negative form.
func IdentifyPositivity(verb string) (positivity int) {
	if hasAnySuffix(verb, Nai, Masen, Nakatta, MasenDeshita) {
		return Negative
	}
	return Positive
}

// // IdentifyGroup tries to identify the verb class
// func IdentifyGroup(verb string) (class int) {
// 	// TODO
// 	return class
// }
