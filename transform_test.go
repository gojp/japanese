package japanese

import "testing"

func TestSplitEnding(t *testing.T) {
	tests := []struct {
		verb   string
		root   string
		ending string
	}{
		{"ある", "あ", "る"},
		{"あります", "あり", "ます"},
		{"ない", "", "ない"},
		{"ありません", "あり", "ません"},
	}

	for _, tt := range tests {
		root, ending := SplitEnding(tt.verb)
		if root != tt.root {
			t.Errorf("SplitEnding(%q) root = %q, want %q", tt.verb, root, tt.root)
		}
		if ending != tt.ending {
			t.Errorf("SplitEnding(%q) end = %q, want %q", tt.verb, ending, tt.ending)
		}
	}
}

func TestDictionaryFormBasicIchidans(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", "る"},
		{"たべます", "たべる"},
		{"おしえます", "おしえる"},
		{"います", "いる"},
		{"食べます", "食べる"},
	}

	for _, tt := range tests {
		_, got := DictionaryForm(tt.in)
		if got != tt.want {
			t.Errorf("DictionaryForm(%q) ichidan = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestDictionaryFormBasicGodans(t *testing.T) {
	tests := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"いく", "いく"},
		{"ききます", "きく"},
		{"ぬぎます", "ぬぐ"},
		{"うたいます", "うたう"},
	}

	for _, tt := range tests {
		got, _ := DictionaryForm(tt.in)
		if got != tt.want {
			t.Errorf("DictionaryForm(%q) godan = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestDictionaryFormAru(t *testing.T) {
	verbs := []string{"ある", "あります", "ない", "ありません"}
	for _, v := range verbs {
		if got, _ := DictionaryForm(v); got != "ある" {
			t.Errorf("DictionaryForm(%q) godan = %q, want %q", v, got, "ある")
		}
	}
}

func TestDictionaryFormIchidanAkeru(t *testing.T) {
	verbs := []string{"あける", "あけない", "あけます", "あけません", "あけよう", "あけるだろう", "あけないだろう", "あけましょう", "あけるでしょう", "あけないでしょう", "あけろ", "あけるな", "あけてください", "あけないでください", "あけた", "あけなかった",
		"あけました", "あけませんでした", "あけたろう", "あけただろう", "あけなかっただろう", "あけたでしょう", "あけなかったでしょう"} //"あけている", "あけています", "あけていません", "あけていた", "あけていました", "あけていませんでした", "あければ", "あけなければ", "あけませば", "あけますれば", "あけませんなら", "あけたら", "あけなかったら", "あけましたら", "あけませんでしたら", "あけられる", "あけられない", "あけられます", "あけられません", "あけさせる", "あけさせない", "あけさせます", "あけさせません"}

	for _, v := range verbs {
		if _, got := DictionaryForm(v); got != "あける" {
			t.Errorf("DictionaryForm(%q) ichidan = %q, want %q", v, got, "あける")
		}
	}
}

func TestDictionaryFormIchidan(t *testing.T) {
	verbs := []string{"たべます", "たべました", "たべた"}
	for _, v := range verbs {
		if _, got := DictionaryForm(v); got != "たべる" {
			t.Errorf("DictionaryForm(%q) ichidan = %q, want %q", v, got, "たべる")
		}
	}

	verbs = []string{"開ける", "開けて", "開けます", "開けない", "開けません"}
	for _, v := range verbs {
		if _, got := DictionaryForm(v); got != "開ける" {
			t.Errorf("DictionaryForm(%q) ichidan = %q, want %q", v, got, "開ける")
		}
	}
}
