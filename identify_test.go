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
