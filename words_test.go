package japanese

import (
	"reflect"
	"runtime"
	"testing"
)

type verbTest struct {
	f func() (Word, error)
	w Word
}

var (
	// Ru verbs
	taberu = Verb{"る", Word{"食べる", "たべる"}}
	miru   = Verb{"る", Word{"見る", "みる"}}

	// U verbs
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

	// Exception verbs
	aru  = Verb{"ex", Word{"ある", "ある"}}
	suru = Verb{"ex", Word{"する", "する"}}
	kuru = Verb{"ex", Word{"来る", "くる"}}
)

func functionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func testVerb(t *testing.T, p []verbTest) {
	for _, tt := range p {
		got, err := tt.f()
		if err != nil {
			t.Fatalf("%s(%v) error: %v", functionName(tt.f), tt.w, err)
		}
		if got.Kanji != tt.w.Kanji {
			t.Errorf("%s Kanji = %s, want %s", functionName(tt.f), got.Kanji, tt.w.Kanji)
		}
		if got.Kana != tt.w.Kana {
			t.Errorf("%s Kana = %s, want %s", functionName(tt.f), got.Kana, tt.w.Kana)
		}
	}

}

var progressiveTests = []verbTest{
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

	{suru.Progressive, Word{"している", "している"}},
	{suru.ProgressiveNegative, Word{"していない", "していない"}},
	{suru.ProgressivePolite, Word{"しています", "しています"}},
	{suru.ProgressiveNegativePolite, Word{"していません", "していません"}},
	{suru.ProgressiveShort, Word{"してる", "してる"}},
	{suru.ProgressiveShortNegative, Word{"してない", "してない"}},

	{kuru.Progressive, Word{"来ている", "きている"}},
	{kuru.ProgressiveNegative, Word{"来ていない", "きていない"}},
	{kuru.ProgressivePolite, Word{"来ています", "きています"}},
	{kuru.ProgressiveNegativePolite, Word{"来ていません", "きていません"}},
	{kuru.ProgressiveShort, Word{"来てる", "きてる"}},
	{kuru.ProgressiveShortNegative, Word{"来てない", "きてない"}},
}

func TestProgressive(t *testing.T) {
	testVerb(t, progressiveTests)
}

var negativeTests = []verbTest{
	{taberu.Negative, Word{"食べない", "たべない"}},
	{taberu.NegativePast, Word{"食べなかった", "たべなかった"}},
	{taberu.NegativePolite, Word{"食べません", "たべません"}},
	{taberu.NegativePastPolite, Word{"食べませんでした", "たべませんでした"}},
	{miru.Negative, Word{"見ない", "みない"}},
	{miru.NegativePast, Word{"見なかった", "みなかった"}},
	{miru.NegativePolite, Word{"見ません", "みません"}},
	{miru.NegativePastPolite, Word{"見ませんでした", "みませんでした"}},
	{hanasu.Negative, Word{"話さない", "はなさない"}},
	{hanasu.NegativePast, Word{"話さなかった", "はなさなかった"}},
	{hanasu.NegativePolite, Word{"話しません", "はなしません"}},
	{hanasu.NegativePastPolite, Word{"話しませんでした", "はなしませんでした"}},
	{kaku.Negative, Word{"書かない", "かかない"}},
	{kaku.NegativePast, Word{"書かなかった", "かかなかった"}},
	{kaku.NegativePolite, Word{"書きません", "かきません"}},
	{kaku.NegativePastPolite, Word{"書きませんでした", "かきませんでした"}},
	{oyogu.Negative, Word{"泳がない", "およがない"}},
	{oyogu.NegativePast, Word{"泳がなかった", "およがなかった"}},
	{oyogu.NegativePolite, Word{"泳ぎません", "およぎません"}},
	{oyogu.NegativePastPolite, Word{"泳ぎませんでした", "およぎませんでした"}},
	{nomu.Negative, Word{"飲まない", "のまない"}},
	{nomu.NegativePast, Word{"飲まなかった", "のまなかった"}},
	{nomu.NegativePolite, Word{"飲みません", "のみません"}},
	{nomu.NegativePastPolite, Word{"飲みませんでした", "のみませんでした"}},
	{asobu.Negative, Word{"遊ばない", "あそばない"}},
	{asobu.NegativePast, Word{"遊ばなかった", "あそばなかった"}},
	{asobu.NegativePolite, Word{"遊びません", "あそびません"}},
	{asobu.NegativePastPolite, Word{"遊びませんでした", "あそびませんでした"}},
	{shinu.Negative, Word{"死なない", "しなない"}},
	{shinu.NegativePast, Word{"死ななかった", "しななかった"}},
	{shinu.NegativePolite, Word{"死にません", "しにません"}},
	{shinu.NegativePastPolite, Word{"死にませんでした", "しにませんでした"}},
	{kiru.Negative, Word{"切らない", "きらない"}},
	{kiru.NegativePast, Word{"切らなかった", "きらなかった"}},
	{kiru.NegativePolite, Word{"切りません", "きりません"}},
	{kiru.NegativePastPolite, Word{"切りませんでした", "きりませんでした"}},
	{kau.Negative, Word{"買わない", "かわない"}},
	{kau.NegativePast, Word{"買わなかった", "かわなかった"}},
	{kau.NegativePolite, Word{"買いません", "かいません"}},
	{kau.NegativePastPolite, Word{"買いませんでした", "かいませんでした"}},
	{motsu.Negative, Word{"持たない", "もたない"}},
	{motsu.NegativePast, Word{"持たなかった", "もたなかった"}},
	{motsu.NegativePolite, Word{"持ちません", "もちません"}},
	{motsu.NegativePastPolite, Word{"持ちませんでした", "もちませんでした"}},
	{iku.Negative, Word{"行かない", "いかない"}},
	{iku.NegativePast, Word{"行かなかった", "いかなかった"}},
	{iku.NegativePolite, Word{"行きません", "いきません"}},
	{iku.NegativePastPolite, Word{"行きませんでした", "いきませんでした"}},
	{aru.Negative, Word{"ない", "ない"}},
	{aru.NegativePast, Word{"なかった", "なかった"}},
	{aru.NegativePolite, Word{"ありません", "ありません"}},
	{aru.NegativePastPolite, Word{"ありませんでした", "ありませんでした"}},
	{suru.Negative, Word{"しない", "しない"}},
	{suru.NegativePast, Word{"しなかった", "しなかった"}},
	{suru.NegativePolite, Word{"しません", "しません"}},
	{suru.NegativePastPolite, Word{"しませんでした", "しませんでした"}},
	{kuru.Negative, Word{"来ない", "こない"}},
	{kuru.NegativePast, Word{"来なかった", "こなかった"}},
	{kuru.NegativePolite, Word{"来ません", "きません"}},
	{kuru.NegativePastPolite, Word{"来ませんでした", "きませんでした"}},
}

func TestNegative(t *testing.T) {
	testVerb(t, negativeTests)
}

var pastTests = []verbTest{
	{taberu.Past, Word{"食べた", "たべた"}},
	{taberu.PastPolite, Word{"食べました", "たべました"}},
	{miru.Past, Word{"見た", "みた"}},
	{miru.PastPolite, Word{"見ました", "みました"}},
	{hanasu.Past, Word{"話した", "はなした"}},
	{hanasu.PastPolite, Word{"話しました", "はなしました"}},
	{kaku.Past, Word{"書いた", "かいた"}},
	{kaku.PastPolite, Word{"書きました", "かきました"}},
	{oyogu.Past, Word{"泳いだ", "およいだ"}},
	{oyogu.PastPolite, Word{"泳ぎました", "およぎました"}},
	{nomu.Past, Word{"飲んだ", "のんだ"}},
	{nomu.PastPolite, Word{"飲みました", "のみました"}},
	{asobu.Past, Word{"遊んだ", "あそんだ"}},
	{asobu.PastPolite, Word{"遊びました", "あそびました"}},
	{shinu.Past, Word{"死んだ", "しんだ"}},
	{shinu.PastPolite, Word{"死にました", "しにました"}},
	{kiru.Past, Word{"切った", "きった"}},
	{kiru.PastPolite, Word{"切りました", "きりました"}},
	{kau.Past, Word{"買った", "かった"}},
	{kau.PastPolite, Word{"買いました", "かいました"}},
	{motsu.Past, Word{"持った", "もった"}},
	{motsu.PastPolite, Word{"持ちました", "もちました"}},
	{iku.Past, Word{"行った", "いった"}},
	{iku.PastPolite, Word{"行きました", "いきました"}},
	{aru.Past, Word{"あった", "あった"}},
	{aru.PastPolite, Word{"ありました", "ありました"}},
	{suru.Past, Word{"した", "した"}},
	{suru.PastPolite, Word{"しました", "しました"}},
	{kuru.Past, Word{"来た", "きた"}},
	{kuru.PastPolite, Word{"来ました", "きました"}},
}

func TestPast(t *testing.T) {
	testVerb(t, pastTests)
}

var potentialTests = []verbTest{
	{taberu.Potential, Word{"食べられる", "たべられる"}},
	{taberu.PotentialNegative, Word{"食べられない", "たべられない"}},
	{taberu.PotentialPolite, Word{"食べられます", "たべられます"}},
	{taberu.PotentialNegativePolite, Word{"食べられません", "たべられません"}},
	{miru.Potential, Word{"見られる", "みられる"}},
	{miru.PotentialNegative, Word{"見られない", "みられない"}},
	{miru.PotentialPolite, Word{"見られます", "みられます"}},
	{miru.PotentialNegativePolite, Word{"見られません", "みられません"}},
	{hanasu.Potential, Word{"話せる", "はなせる"}},
	{hanasu.PotentialNegative, Word{"話せない", "はなせない"}},
	{hanasu.PotentialPolite, Word{"話せます", "はなせます"}},
	{hanasu.PotentialNegativePolite, Word{"話せません", "はなせません"}},
	{kaku.Potential, Word{"書ける", "かける"}},
	{kaku.PotentialNegative, Word{"書けない", "かけない"}},
	{kaku.PotentialPolite, Word{"書けます", "かけます"}},
	{kaku.PotentialNegativePolite, Word{"書けません", "かけません"}},
	{oyogu.Potential, Word{"泳げる", "およげる"}},
	{oyogu.PotentialNegative, Word{"泳げない", "およげない"}},
	{oyogu.PotentialPolite, Word{"泳げます", "およげます"}},
	{oyogu.PotentialNegativePolite, Word{"泳げません", "およげません"}},
	{nomu.Potential, Word{"飲める", "のめる"}},
	{nomu.PotentialNegative, Word{"飲めない", "のめない"}},
	{nomu.PotentialPolite, Word{"飲めます", "のめます"}},
	{nomu.PotentialNegativePolite, Word{"飲めません", "のめません"}},
	{asobu.Potential, Word{"遊べる", "あそべる"}},
	{asobu.PotentialNegative, Word{"遊べない", "あそべない"}},
	{asobu.PotentialPolite, Word{"遊べます", "あそべます"}},
	{asobu.PotentialNegativePolite, Word{"遊べません", "あそべません"}},
	{shinu.Potential, Word{"死ねる", "しねる"}},
	{shinu.PotentialNegative, Word{"死ねない", "しねない"}},
	{shinu.PotentialPolite, Word{"死ねます", "しねます"}},
	{shinu.PotentialNegativePolite, Word{"死ねません", "しねません"}},
	{kiru.Potential, Word{"切れる", "きれる"}},
	{kiru.PotentialNegative, Word{"切れない", "きれない"}},
	{kiru.PotentialPolite, Word{"切れます", "きれます"}},
	{kiru.PotentialNegativePolite, Word{"切れません", "きれません"}},
	{kau.Potential, Word{"買える", "かえる"}},
	{kau.PotentialNegative, Word{"買えない", "かえない"}},
	{kau.PotentialPolite, Word{"買えます", "かえます"}},
	{kau.PotentialNegativePolite, Word{"買えません", "かえません"}},
	{motsu.Potential, Word{"持てる", "もてる"}},
	{motsu.PotentialNegative, Word{"持てない", "もてない"}},
	{motsu.PotentialPolite, Word{"持てます", "もてます"}},
	{motsu.PotentialNegativePolite, Word{"持てません", "もてません"}},
	{iku.Potential, Word{"行ける", "いける"}},
	{iku.PotentialNegative, Word{"行けない", "いけない"}},
	{iku.PotentialPolite, Word{"行けます", "いけます"}},
	{iku.PotentialNegativePolite, Word{"行けません", "いけません"}},
	{suru.Potential, Word{"できる", "できる"}},
	{suru.PotentialNegative, Word{"できない", "できない"}},
	{suru.PotentialPolite, Word{"できます", "できます"}},
	{suru.PotentialNegativePolite, Word{"できません", "できません"}},
	{kuru.Potential, Word{"来られる", "こられる"}},
	{kuru.PotentialNegative, Word{"来られない", "こられない"}},
	{kuru.PotentialPolite, Word{"来られます", "こられます"}},
	{kuru.PotentialNegativePolite, Word{"来られません", "こられません"}},
}

func TestPotential(t *testing.T) {
	testVerb(t, potentialTests)
}
