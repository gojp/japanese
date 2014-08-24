package japanese

import "testing"

var tests = []struct {
	in     string
	form   int
	ending string
}{
	{"ある", PresentIndicative, Ru},
	{"あります", PresentIndicative, Masu},
	{"ない", PresentIndicative, Nai},
	{"ありません", PresentIndicative, Masen},
	{"あった", PastIndicative, Ta},
	{"なかった", PastIndicative, Nakatta},
	{"ありました", PastIndicative, Mashita},
	{"ありませんでした", PastIndicative, MasenDeshita},
	{"行っている", PresentProgressive, Teiru},
	{"行っています", PresentProgressive, Teimasu},
	{"知ってる", PresentProgressive, Teru},
	{"知ってます", PresentProgressive, Temasu},
	{"行っていた", PastProgressive, Teita},
	{"行っていました", PastProgressive, Teimashita},
	{"知ってた", PastProgressive, Teta},
	{"知ってました", PastProgressive, Temashita},
	{"すれば", Provisional, Reba},
	{"使えば", Provisional, Eba},
	{"焼けば", Provisional, Keba},
	{"泳げば", Provisional, Geba},
	{"示せば", Provisional, Seba},
	{"待てば", Provisional, Teba},
	{"死ねば", Provisional, Neba},
	{"呼べば", Provisional, Beba},
	{"読めば", Provisional, Meba},
	{"走れば", Provisional, Reba},
	{"あれば", Provisional, Reba},
	{"見れば", Provisional, Reba},
	{"寒ければ", Provisional, Kereba},
	{"簡単であれば", Provisional, Deareba},
	{"行かなければ", Provisional, Nakereba},
	{"foo", Unknown, ""},
}

func TestIdentifyForm(t *testing.T) {
	for _, tt := range tests {
		if got := IdentifyForm(tt.in); got != tt.form {
			t.Errorf("IdentifyForm(%q) = %d, want %d", tt.in, got, tt.form)
		}
	}
}

func TestIdentifyEnding(t *testing.T) {
	for _, tt := range tests {
		if got := IdentifyEnding(tt.in); got != tt.ending {
			t.Errorf("IdentifyEnding(%q) = %q, want %q", tt.in, got, tt.ending)
		}
	}
}

func TestIdentifyPositivity(t *testing.T) {
	tests := []struct {
		in   string
		want int
	}{
		{"ある", Positive},
		{"あります", Positive},
		{"あった", Positive},
		{"ありました", Positive},
		{"ない", Negative},
		{"ません", Negative},
		{"なかった", Negative},
		{"ませんでした", Negative},
	}
	for _, tt := range tests {
		if got := IdentifyPositivity(tt.in); got != tt.want {
			t.Errorf("IdentifyPositivity(%q) = %d, want %d", tt.in, got, tt.want)
		}
	}
}
