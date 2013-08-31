package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type JapaneseSuite struct{}

type TestWord struct {
	kanji       string
	kana        string
	english     string
	neg_kanji   string
	neg_kana    string
	neg_english string
}

var _ = Suite(&JapaneseSuite{})

// ru-verbs
var ruVerbs []TestWord = []TestWord{
	TestWord{"見る", "みる", "see", "見ない", "みない", "not see"},
	TestWord{"食べる", "たべる", "eat", "食べない", "たべない", "not eat"},
	TestWord{"寝る", "ねる", "sleep", "寝ない", "ねない", "not sleep"},
	TestWord{"起きる", "おきる", "get up", "起きない", "おきない", "not get up"},
	TestWord{"考える", "かんがえる", "consider", "考えない", "かんがえない", "not consider"},
	TestWord{"教える", "おしえる", "teach", "教えない", "おしえない", "not teach"},
	TestWord{"出る", "でる", "go out", "出ない", "でない", "not go out"},
	TestWord{"着る", "きる", "wear", "着ない", "きない", "not wear"},
	TestWord{"いる", "いる", "be", "いない", "いない", "not be"},
	TestWord{"ある", "ある", "be", "ない", "ない", "not be"},
}

// u-verbs
var uVerbs []TestWord = []TestWord{
	TestWord{"話す", "はなす", "speak", "話さない", "はなさない", "not speak"},
	TestWord{"聞く", "きく", "hear", "聞かない", "きかない", "not hear"},
	TestWord{"泳ぐ", "およぐ", "swim", "泳がない", "およがない", "not swim"},
	TestWord{"遊ぶ", "あそぶ", "play", "遊ばない", "あそばない", "not play"},
	TestWord{"待つ", "まつ", "wait", "待たない", "またない", "not wait"},
	TestWord{"飲む", "のむ", "drink", "飲まない", "のまない", "not drink"},
	TestWord{"買う", "かう", "buy", "買わない", "かわない", "not buy"},
	TestWord{"帰る", "かえる", "return", "帰らない", "かえらない", "not return"},
	TestWord{"死ぬ", "しぬ", "die", "死なない", "しなない", "not die"},
}

// exceptions
var exceptionVerbs []TestWord = []TestWord{
	TestWord{"する", "する", "do", "しない", "しない", "not do"},
	TestWord{"くる", "くる", "come", "こない", "こない", "not come"},
}

func (s *JapaneseSuite) TestNegativeRuVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, verb := range ruVerbs {
		v := RuVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
		c.Check(neg_word.kana, Equals, verb.neg_kana)
	}
}

func (s *JapaneseSuite) TestNegativeUVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, verb := range uVerbs {
		v := UVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
		c.Check(neg_word.kana, Equals, verb.neg_kana)
	}
}

func (s *JapaneseSuite) TestNegativeExceptionVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, verb := range exceptionVerbs {
		v := ExceptionVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
		c.Check(neg_word.kana, Equals, verb.neg_kana)
	}
}
