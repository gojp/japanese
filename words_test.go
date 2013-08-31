package japanese

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type JapaneseSuite struct {
	RuVerbs        []RuVerb
	UVerbs         []UVerb
	ExceptionVerbs []ExceptionVerb
}

type TestWord struct {
	kanji   string
	kana    string
	english string

	neg_kanji   string
	neg_kana    string
	neg_english string

	past_kanji   string
	past_kana    string
	past_english string
}

var _ = Suite(&JapaneseSuite{})

// ru-verbs
var ruVerbs []TestWord = []TestWord{
	TestWord{"見る", "みる", "see", "見ない", "みない", "not see", "見た", "みた", "did see"},
	TestWord{"食べる", "たべる", "eat", "食べない", "たべない", "not eat", "食べた", "たべた", "did eat"},
	TestWord{"寝る", "ねる", "sleep", "寝ない", "ねない", "not sleep", "寝た", "ねた", "did sleep"},
	TestWord{"起きる", "おきる", "get up", "起きない", "おきない", "not get up", "起きた", "おきた", "did get up"},
	TestWord{"考える", "かんがえる", "consider", "考えない", "かんがえない", "not consider", "考えた", "かんがえた", "did consider"},
	TestWord{"教える", "おしえる", "teach", "教えない", "おしえない", "not teach", "教えた", "おしえた", "did teach"},
	TestWord{"出る", "でる", "go out", "出ない", "でない", "not go out", "出た", "でた", "did go out"},
	TestWord{"着る", "きる", "wear", "着ない", "きない", "not wear", "着た", "きた", "did wear"},
	TestWord{"いる", "いる", "exist", "いない", "いない", "not exist", "いた", "いた", "did exist"},
	TestWord{"ある", "ある", "exist", "ない", "ない", "not exist", "あた", "あた", "did exist"},
}

// u-verbs
var uVerbs []TestWord = []TestWord{
	TestWord{"話す", "はなす", "speak", "話さない", "はなさない", "not speak", "話した", "はなした", "did speak"},
	TestWord{"聞く", "きく", "hear", "聞かない", "きかない", "not hear", "聞いた", "きいた", "did hear"},
	TestWord{"泳ぐ", "およぐ", "swim", "泳がない", "およがない", "not swim", "泳いた", "およいた", "did swim"},
	TestWord{"遊ぶ", "あそぶ", "play", "遊ばない", "あそばない", "not play", "遊んだ", "あそんだ", "did play"},
	TestWord{"待つ", "まつ", "wait", "待たない", "またない", "not wait", "待った", "まった", "did wait"},
	TestWord{"飲む", "のむ", "drink", "飲まない", "のまない", "not drink", "飲んだ", "のんだ", "did drink"},
	TestWord{"買う", "かう", "buy", "買わない", "かわない", "not buy", "買った", "かった", "did buy"},
	TestWord{"帰る", "かえる", "return", "帰らない", "かえらない", "not return", "帰った", "かえった", "did return"},
	TestWord{"死ぬ", "しぬ", "die", "死なない", "しなない", "not die", "死んだ", "しんだ", "did die"},
}

// exceptions
var exceptionVerbs []TestWord = []TestWord{
	TestWord{"する", "する", "do", "しない", "しない", "not do", "した", "した", "did do"},
	TestWord{"くる", "くる", "come", "こない", "こない", "not come", "きた", "きた", "did come"},
}

func (s *JapaneseSuite) SetUpSuite(c *C) {
	s.RuVerbs = []RuVerb{}
	for _, verb := range ruVerbs {
		v := RuVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		s.RuVerbs = append(s.RuVerbs, v)
	}
	s.UVerbs = []UVerb{}
	for _, verb := range uVerbs {
		v := UVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		s.UVerbs = append(s.UVerbs, v)
	}
	s.ExceptionVerbs = []ExceptionVerb{}
	for _, verb := range exceptionVerbs {
		v := ExceptionVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		s.ExceptionVerbs = append(s.ExceptionVerbs, v)
	}
}

func (s *JapaneseSuite) TestNegativeRuVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for i := range s.RuVerbs {
		v := s.RuVerbs[i]
		t := ruVerbs[i]
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, t.neg_kanji)
		c.Check(neg_word.kana, Equals, t.neg_kana)
		c.Check(neg_word.english, Equals, t.neg_english)
	}
}

func (s *JapaneseSuite) TestPastRuVerbs(c *C) {
	// check that ru-verbs get the correct past suffix
	for i := range s.RuVerbs {
		v := s.RuVerbs[i]
		t := ruVerbs[i]
		past_word := v.Past()
		c.Check(past_word.kanji, Equals, t.past_kanji)
		c.Check(past_word.kana, Equals, t.past_kana)
		c.Check(past_word.english, Equals, t.past_english)
	}
}

func (s *JapaneseSuite) TestNegativeUVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, verb := range uVerbs {
		v := UVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
		c.Check(neg_word.kana, Equals, verb.neg_kana)
		c.Check(neg_word.english, Equals, verb.neg_english)
	}
}

func (s *JapaneseSuite) TestPastUVerbs(c *C) {
	// check that ru-verbs get the correct past suffix
	for i := range s.UVerbs {
		v := s.UVerbs[i]
		t := uVerbs[i]
		past_word := v.Past()
		c.Check(past_word.kanji, Equals, t.past_kanji)
		c.Check(past_word.kana, Equals, t.past_kana)
		c.Check(past_word.english, Equals, t.past_english)
	}
}

func (s *JapaneseSuite) TestNegativeExceptionVerbs(c *C) {
	// check that exception-verbs get the correct negative suffix
	for _, verb := range exceptionVerbs {
		v := ExceptionVerb{Verb{Word{verb.kanji, verb.kana, verb.english}}}
		neg_word := v.Negative()
		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
		c.Check(neg_word.kana, Equals, verb.neg_kana)
	}
}

func (s *JapaneseSuite) TestPastExceptionVerbs(c *C) {
	// check that exception-verbs get the correct past suffix
	for i := range s.ExceptionVerbs {
		v := s.ExceptionVerbs[i]
		t := exceptionVerbs[i]
		past_word := v.Past()
		c.Check(past_word.kanji, Equals, t.past_kanji)
		c.Check(past_word.kana, Equals, t.past_kana)
		c.Check(past_word.english, Equals, t.past_english)
	}
}
