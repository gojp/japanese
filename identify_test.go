package japanese

import "testing"

func TestIdentifyForm(t *testing.T) {
	tests := []struct {
		give string
		want int
	}{
		{"ある", PresentIndicative},
		{"あります", PresentIndicative},
		{"ない", PresentIndicative},
		{"ありません", PresentIndicative},
		{"あった", PastIndicative},
		{"なかった", PastIndicative},
		{"ありました", PastIndicative},
		{"ありませんでした", PastIndicative},
		{"行っている", PresentProgressive},
		{"行っています", PresentProgressive},
		{"知ってる", PresentProgressive},
		{"知ってます", PresentProgressive},
		{"行っていた", PastProgressive},
		{"行っていました", PastProgressive},
		{"知ってた", PastProgressive},
		{"知ってました", PastProgressive},
		{"すれば", Provisional},
		{"使えば", Provisional},
		{"焼けば", Provisional},
		{"泳げば", Provisional},
		{"示せば", Provisional},
		{"待てば", Provisional},
		{"死ねば", Provisional},
		{"呼べば", Provisional},
		{"読めば", Provisional},
		{"走れば", Provisional},
		{"あれば", Provisional},
		{"見れば", Provisional},
		{"寒ければ", Provisional},
		{"簡単であれば", Provisional},
		{"行かなければ", Provisional},
		{"foo", Unknown},
	}

	for _, tt := range tests {
		if got := IdentifyForm(tt.give); got != tt.want {
			t.Errorf("IdentifyForm(%q) = %d, want %d", tt.give, got, tt.want)
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
