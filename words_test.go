package main

import (
	. "launchpad.net/gocheck"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type JapaneseSuite struct{}

type TestWord struct {
	root     string
	negative string
}

var _ = Suite(&JapaneseSuite{})

// ru-verbs
var ruVerbs []TestWord = []TestWord{
	TestWord{"見る", "見ない"},
	TestWord{"食べる", "食べない"},
	TestWord{"寝る", "寝ない"},
	TestWord{"起きる", "起きない"},
	TestWord{"考える", "考えない"},
	TestWord{"教える", "教えない"},
	TestWord{"出る", "出ない"},
	TestWord{"着る", "着ない"},
	TestWord{"いる", "いない"},
	TestWord{"ある", "ない"},
}

// u-verbs
var uVerbs []TestWord = []TestWord{
	TestWord{"話す", "話さない"},
	TestWord{"聞く", "聞かない"},
	TestWord{"泳ぐ", "泳がない"},
	TestWord{"遊ぶ", "遊ばない"},
	TestWord{"待つ", "待たない"},
	TestWord{"飲む", "飲まない"},
	TestWord{"買う", "買わない"},
	TestWord{"帰る", "帰らない"},
	TestWord{"死ぬ", "死なない"},
}

// exceptions
var exceptionVerbs []TestWord = []TestWord{
	TestWord{"する", "しない"},
	TestWord{"くる", "こない"},
}

func (s *JapaneseSuite) TestNegativeRuVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, ruVerb := range ruVerbs {
		v := RuVerb{Verb{Word{ruVerb.root, "", ""}}}
		c.Check(v.Negative(), Equals, ruVerb.negative)
	}
}

func (s *JapaneseSuite) TestNegativeUVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, uVerb := range uVerbs {
		v := UVerb{Verb{Word{uVerb.root, "", ""}}}
		c.Check(v.Negative(), Equals, uVerb.negative)
	}
}

func (s *JapaneseSuite) TestNegativeExceptionVerbs(c *C) {
	// check that ru-verbs get the correct negative suffix
	for _, exceptionVerb := range exceptionVerbs {
		v := ExceptionVerb{Verb{Word{exceptionVerb.root, "", ""}}}
		c.Check(v.Negative(), Equals, exceptionVerb.negative)
	}
}
