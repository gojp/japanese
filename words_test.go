package japanese

import (
	"reflect"
	"runtime"
	"testing"
)

// type TestWord struct {
// 	kanji string
// 	kana  string

// 	neg_kanji string
// 	neg_kana  string

// 	past_kanji string
// 	past_kana  string

// 	long_kanji string
// 	long_kana  string

// 	long_neg_kanji string
// 	long_neg_kana  string

// 	long_past_kanji string
// 	long_past_kana  string
// }

// var _ = Suite(&JapaneseSuite{})

// // ru-verbs
// var ruVerbs []TestWord = []TestWord{
// 	TestWord{"見る", "みる", "見ない", "みない", "見た", "みた", "見ます", "みます", "見ません", "みません", "見ました", "みました"},
// 	TestWord{"食べる", "たべる", "食べない", "たべない", "食べた", "たべた", "食べます", "たべます", "食べません", "たべません", "食べました", "たべました"},
// 	TestWord{"寝る", "ねる", "寝ない", "ねない", "寝た", "ねた", "寝ます", "ねます", "寝ません", "ねません", "寝ました", "ねました"},
// 	TestWord{"起きる", "おきる", "起きない", "おきない", "起きた", "おきた", "起きます", "おきます", "起きません", "おきません", "起きました", "おきました"},
// 	TestWord{"考える", "かんがえる", "考えない", "かんがえない", "考えた", "かんがえた", "考えます", "かんがえます", "考えません", "かんがえません", "考えました", "かんがえました"},
// 	TestWord{"教える", "おしえる", "教えない", "おしえない", "教えた", "おしえた", "教えます", "おしえます", "教えません", "おしえません", "教えました", "おしえました"},
// 	TestWord{"出る", "でる", "出ない", "でない", "出た", "でた", "出ます", "でます", "出ません", "でません", "出ました", "でました"},
// 	TestWord{"着る", "きる", "着ない", "きない", "着た", "きた", "着ます", "きます", "着ません", "きません", "着ました", "きました"},
// 	TestWord{"いる", "いる", "いない", "いない", "いた", "いた", "います", "います", "いません", "いません", "いました", "いました"},
// }

// // u-verbs
// var uVerbs []TestWord = []TestWord{
// 	TestWord{"話す", "はなす", "話さない", "はなさない", "話した", "はなした", "話します", "はなします", "話しません", "はなしません", "話しました", "はなしました"},
// 	TestWord{"聞く", "きく", "聞かない", "きかない", "聞いた", "きいた", "聞きます", "ききます", "聞きません", "ききません", "聞きました", "ききました"},
// 	TestWord{"泳ぐ", "およぐ", "泳がない", "およがない", "泳いた", "およいた", "泳ぎます", "およぎます", "泳ぎません", "およぎません", "泳ぎました", "およぎました"},
// 	TestWord{"遊ぶ", "あそぶ", "遊ばない", "あそばない", "遊んだ", "あそんだ", "遊びます", "あそびます", "遊びません", "あそびません", "遊びました", "あそびました"},
// 	TestWord{"待つ", "まつ", "待たない", "またない", "待った", "まった", "待ちます", "まちます", "待ちません", "まちません", "待ちました", "待ちました"},
// 	TestWord{"飲む", "のむ", "飲まない", "のまない", "飲んだ", "のんだ", "飲みます", "のみます", "飲みません", "のみません", "飲みました", "のみました"},
// 	TestWord{"買う", "かう", "買わない", "かわない", "買った", "かった", "買います", "かいます", "買いません", "かいません", "買いました", "かいました"},
// 	TestWord{"帰る", "かえる", "帰らない", "かえらない", "帰った", "かえった", "帰ります", "かえります", "帰りません", "かえりません", "帰りました", "かえりました"},
// 	TestWord{"死ぬ", "しぬ", "死なない", "しなない", "死んだ", "しんだ", "死にます", "しにます", "死にません", "しにません", "死にました", "しにました"},
// 	TestWord{"ある", "ある", "ない", "ない", "あった", "あった", "あります", "あります", "ありません", "ありません", "ありました", "ありました"},
// }

// // exceptions
// var exceptionVerbs []TestWord = []TestWord{
// 	TestWord{"する", "する", "しない", "しない", "した", "した", "します", "します", "しません", "しません", "しました", "しました"},
// 	TestWord{"くる", "くる", "こない", "こない", "きた", "きた", "きます", "きます", "きません", "きません", "きました", "きました"},
// }

// func (s *JapaneseSuite) SetUpSuite(c *C) {
// 	s.RuVerbs = []RuVerb{}
// 	for _, verb := range ruVerbs {
// 		v := RuVerb{Verb{Word{verb.kanji, verb.kana}}}
// 		s.RuVerbs = append(s.RuVerbs, v)
// 	}
// 	s.UVerbs = []UVerb{}
// 	for _, verb := range uVerbs {
// 		v := UVerb{Verb{Word{verb.kanji, verb.kana}}}
// 		s.UVerbs = append(s.UVerbs, v)
// 	}
// 	s.ExceptionVerbs = []ExceptionVerb{}
// 	for _, verb := range exceptionVerbs {
// 		v := ExceptionVerb{Verb{Word{verb.kanji, verb.kana}}}
// 		s.ExceptionVerbs = append(s.ExceptionVerbs, v)
// 	}
// }

// func (s *JapaneseSuite) TestNegativeRuVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for i := range s.RuVerbs {
// 		v := s.RuVerbs[i]
// 		t := ruVerbs[i]
// 		neg_word := v.Negative()
// 		c.Check(neg_word.kanji, Equals, t.neg_kanji)
// 		c.Check(neg_word.kana, Equals, t.neg_kana)
// 	}
// }

// func (s *JapaneseSuite) TestPastRuVerbs(c *C) {
// 	// check that ru-verbs get the correct past suffix
// 	for i := range s.RuVerbs {
// 		v := s.RuVerbs[i]
// 		t := ruVerbs[i]
// 		past_word := v.Past()
// 		c.Check(past_word.kanji, Equals, t.past_kanji)
// 		c.Check(past_word.kana, Equals, t.past_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongRuVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for i := range s.RuVerbs {
// 		v := s.RuVerbs[i]
// 		t := ruVerbs[i]
// 		long_word := v.Long()
// 		c.Check(long_word.kanji, Equals, t.long_kanji)
// 		c.Check(long_word.kana, Equals, t.long_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongNegativeRuVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for i := range s.RuVerbs {
// 		v := s.RuVerbs[i]
// 		t := ruVerbs[i]
// 		long_neg_word := v.LongNegative()
// 		c.Check(long_neg_word.kanji, Equals, t.long_neg_kanji)
// 		c.Check(long_neg_word.kana, Equals, t.long_neg_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongPastRuVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for i := range s.RuVerbs {
// 		v := s.RuVerbs[i]
// 		t := ruVerbs[i]
// 		long_past_word := v.LongPast()
// 		c.Check(long_past_word.kanji, Equals, t.long_past_kanji)
// 		c.Check(long_past_word.kana, Equals, t.long_past_kana)
// 	}
// }

// func (s *JapaneseSuite) TestNegativeUVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for _, verb := range uVerbs {
// 		v := UVerb{Verb{Word{verb.kanji, verb.kana}}}
// 		neg_word := v.Negative()
// 		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
// 		c.Check(neg_word.kana, Equals, verb.neg_kana)
// 	}
// }

// func (s *JapaneseSuite) TestPastUVerbs(c *C) {
// 	// check that ru-verbs get the correct past suffix
// 	for i := range s.UVerbs {
// 		v := s.UVerbs[i]
// 		t := uVerbs[i]
// 		past_word := v.Past()
// 		c.Check(past_word.kanji, Equals, t.past_kanji)
// 		c.Check(past_word.kana, Equals, t.past_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongUVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for i := range s.UVerbs {
// 		v := s.UVerbs[i]
// 		t := uVerbs[i]
// 		long_word := v.Long()
// 		c.Check(long_word.kanji, Equals, t.long_kanji)
// 		c.Check(long_word.kana, Equals, t.long_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongNegativeUVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for i := range s.UVerbs {
// 		v := s.UVerbs[i]
// 		t := uVerbs[i]
// 		long_neg_word := v.LongNegative()
// 		c.Check(long_neg_word.kanji, Equals, t.long_neg_kanji)
// 		c.Check(long_neg_word.kana, Equals, t.long_neg_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongPastUVerbs(c *C) {
// 	// check that ru-verbs get the correct negative suffix
// 	for i := range s.UVerbs {
// 		v := s.UVerbs[i]
// 		t := uVerbs[i]
// 		long_past_word := v.LongPast()
// 		c.Check(long_past_word.kanji, Equals, t.long_past_kanji)
// 		c.Check(long_past_word.kana, Equals, t.long_past_kana)
// 	}
// }

// func (s *JapaneseSuite) TestNegativeExceptionVerbs(c *C) {
// 	// check that exception-verbs get the correct negative suffix
// 	for _, verb := range exceptionVerbs {
// 		v := ExceptionVerb{Verb{Word{verb.kanji, verb.kana}}}
// 		neg_word := v.Negative()
// 		c.Check(neg_word.kanji, Equals, verb.neg_kanji)
// 		c.Check(neg_word.kana, Equals, verb.neg_kana)
// 	}
// }

// func (s *JapaneseSuite) TestPastExceptionVerbs(c *C) {
// 	// check that exception-verbs get the correct past suffix
// 	for i := range s.ExceptionVerbs {
// 		v := s.ExceptionVerbs[i]
// 		t := exceptionVerbs[i]
// 		past_word := v.Past()
// 		c.Check(past_word.kanji, Equals, t.past_kanji)
// 		c.Check(past_word.kana, Equals, t.past_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongExceptionVerbs(c *C) {
// 	for i := range s.ExceptionVerbs {
// 		v := s.ExceptionVerbs[i]
// 		t := exceptionVerbs[i]
// 		long_word := v.Long()
// 		c.Check(long_word.kanji, Equals, t.long_kanji)
// 		c.Check(long_word.kana, Equals, t.long_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongNegativeExceptionVerbs(c *C) {
// 	for i := range s.ExceptionVerbs {
// 		v := s.ExceptionVerbs[i]
// 		t := exceptionVerbs[i]
// 		long_neg_word := v.LongNegative()
// 		c.Check(long_neg_word.kanji, Equals, t.long_neg_kanji)
// 		c.Check(long_neg_word.kana, Equals, t.long_neg_kana)
// 	}
// }

// func (s *JapaneseSuite) TestLongPastExceptionVerbs(c *C) {
// 	for i := range s.ExceptionVerbs {
// 		v := s.ExceptionVerbs[i]
// 		t := exceptionVerbs[i]
// 		long_past_word := v.LongPast()
// 		c.Check(long_past_word.kanji, Equals, t.long_past_kanji)
// 		c.Check(long_past_word.kana, Equals, t.long_past_kana)
// 	}
// }

type progressiveTest struct {
	f func() Word
	w Word
}

var (
	taberu = Verb{"る", Word{"食べる", "たべる"}}
	miru   = Verb{"る", Word{"見る", "みる"}}
)

var ruProgressiveTests = []progressiveTest{
	{taberu.Progressive, Word{"食べている", "たべている"}},
	{taberu.ProgressiveNegative, Word{"食べていない", "たべていない"}},
	{taberu.ProgressivePolite, Word{"食べています", "たべています"}},
	{taberu.ProgressiveNegativePolite, Word{"食べていません", "たべていません"}},
	{taberu.ProgressiveShort, Word{"食べてる", "たべてる"}},
	{taberu.ProgressiveShortNegative, Word{"食べてない", "たべてない"}},

	{miru.Progressive, Word{"見ている", "みている"}},
	{miru.ProgressiveNegative, Word{"見ていない", "みていない"}},
	{miru.ProgressivePolite, Word{"見ています", "みています"}},
	{miru.ProgressiveNegativePolite, Word{"見ていません", "みていません"}},
	{miru.ProgressiveShort, Word{"見てる", "みてる"}},
	{miru.ProgressiveShortNegative, Word{"見てない", "みてない"}},
}

func functionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func testProgressive(t *testing.T, p []progressiveTest) {
	for _, tt := range p {
		if got := tt.f(); got.kanji != tt.w.kanji {
			t.Errorf("%s kanji = %s, want %s", functionName(tt.f), got.kanji, tt.w.kanji)
		}
		if got := tt.f(); got.kana != tt.w.kana {
			t.Errorf("%s kana = %s, want %s", functionName(tt.f), got.kana, tt.w.kana)
		}
	}

}

func TestRuProgressive(t *testing.T) {
	testProgressive(t, ruProgressiveTests)
}

var (
	hanasu = Verb{"う", Word{"話す", "はなす"}}
	kaku   = Verb{"う", Word{"書く", "かく"}}
	oyogu  = Verb{"う", Word{"泳ぐ", "およぐ"}}
	nomu   = Verb{"う", Word{"飲む", "のむ"}}
	asobu  = Verb{"う", Word{"遊ぶ", "あそぶ"}}
	shinu  = Verb{"う", Word{"死ぬ", "しぬ"}}
	kiru   = Verb{"う", Word{"切る", "きる"}}
	kau    = Verb{"う", Word{"買う", "かう"}}
	motsu  = Verb{"う", Word{"持つ", "もつ"}}
	iku    = Verb{"う", Word{"行く", "いく"}}
)

var uProgressiveTests = []progressiveTest{
	{hanasu.Progressive, Word{"話している", "はなしている"}},
	{hanasu.ProgressiveNegative, Word{"話していない", "はなしていない"}},
	{hanasu.ProgressivePolite, Word{"話しています", "はなしています"}},
	{hanasu.ProgressiveNegativePolite, Word{"話していません", "はなしていません"}},
	{hanasu.ProgressiveShort, Word{"話してる", "はなしてる"}},
	{hanasu.ProgressiveShortNegative, Word{"話してない", "はなしてない"}},

	{kaku.Progressive, Word{"書いている", "かいている"}},
	{kaku.ProgressiveNegative, Word{"書いていない", "かいていない"}},
	{kaku.ProgressivePolite, Word{"書いています", "かいています"}},
	{kaku.ProgressiveNegativePolite, Word{"書いていません", "かいていません"}},
	{kaku.ProgressiveShort, Word{"書いてる", "かいてる"}},
	{kaku.ProgressiveShortNegative, Word{"書いてない", "かいてない"}},

	{oyogu.Progressive, Word{"泳いでいる", "およいでいる"}},
	{oyogu.ProgressiveNegative, Word{"泳いでいない", "およいでいない"}},
	{oyogu.ProgressivePolite, Word{"泳いでいます", "およいでいます"}},
	{oyogu.ProgressiveNegativePolite, Word{"泳いでいません", "およいでいません"}},
	{oyogu.ProgressiveShort, Word{"泳いでる", "およいでる"}},
	{oyogu.ProgressiveShortNegative, Word{"泳いでない", "およいでない"}},

	{nomu.Progressive, Word{"飲んでいる", "のんでいる"}},
	{nomu.ProgressiveNegative, Word{"飲んでいない", "のんでいない"}},
	{nomu.ProgressivePolite, Word{"飲んでいます", "のんでいます"}},
	{nomu.ProgressiveNegativePolite, Word{"飲んでいません", "のんでいません"}},
	{nomu.ProgressiveShort, Word{"飲んでる", "のんでる"}},
	{nomu.ProgressiveShortNegative, Word{"飲んでない", "のんでない"}},

	{asobu.Progressive, Word{"遊んでいる", "あそんでいる"}},
	{asobu.ProgressiveNegative, Word{"遊んでいない", "あそんでいない"}},
	{asobu.ProgressivePolite, Word{"遊んでいます", "あそんでいます"}},
	{asobu.ProgressiveNegativePolite, Word{"遊んでいません", "あそんでいません"}},
	{asobu.ProgressiveShort, Word{"遊んでる", "あそんでる"}},
	{asobu.ProgressiveShortNegative, Word{"遊んでない", "あそんでない"}},

	{shinu.Progressive, Word{"死んでいる", "しんでいる"}},
	{shinu.ProgressiveNegative, Word{"死んでいない", "しんでいない"}},
	{shinu.ProgressivePolite, Word{"死んでいます", "しんでいます"}},
	{shinu.ProgressiveNegativePolite, Word{"死んでいません", "しんでいません"}},
	{shinu.ProgressiveShort, Word{"死んでる", "しんでる"}},
	{shinu.ProgressiveShortNegative, Word{"死んでない", "しんでない"}},

	{kiru.Progressive, Word{"切っている", "きっている"}},
	{kiru.ProgressiveNegative, Word{"切っていない", "きっていない"}},
	{kiru.ProgressivePolite, Word{"切っています", "きっています"}},
	{kiru.ProgressiveNegativePolite, Word{"切っていません", "きっていません"}},
	{kiru.ProgressiveShort, Word{"切ってる", "きってる"}},
	{kiru.ProgressiveShortNegative, Word{"切ってない", "きってない"}},

	{kau.Progressive, Word{"買っている", "かっている"}},
	{kau.ProgressiveNegative, Word{"買っていない", "かっていない"}},
	{kau.ProgressivePolite, Word{"買っています", "かっています"}},
	{kau.ProgressiveNegativePolite, Word{"買っていません", "かっていません"}},
	{kau.ProgressiveShort, Word{"買ってる", "かってる"}},
	{kau.ProgressiveShortNegative, Word{"買ってない", "かってない"}},

	{motsu.Progressive, Word{"持っている", "もっている"}},
	{motsu.ProgressiveNegative, Word{"持っていない", "もっていない"}},
	{motsu.ProgressivePolite, Word{"持っています", "もっています"}},
	{motsu.ProgressiveNegativePolite, Word{"持っていません", "もっていません"}},
	{motsu.ProgressiveShort, Word{"持ってる", "もってる"}},
	{motsu.ProgressiveShortNegative, Word{"持ってない", "もってない"}},

	{iku.Progressive, Word{"行っている", "いっている"}},
	{iku.ProgressiveNegative, Word{"行っていない", "いっていない"}},
	{iku.ProgressivePolite, Word{"行っています", "いっています"}},
	{iku.ProgressiveNegativePolite, Word{"行っていません", "いっていません"}},
	{iku.ProgressiveShort, Word{"行ってる", "いってる"}},
	{iku.ProgressiveShortNegative, Word{"行ってない", "いってない"}},
}

func TestUProgressive(t *testing.T) {
	testProgressive(t, uProgressiveTests)
}
