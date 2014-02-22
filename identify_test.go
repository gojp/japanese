package japanese

import "testing"

func wantXandGotY(t *testing.T, got interface{}, want interface{}) {
	if got != want {
		t.Errorf("Expected %v, but got %v", got, want)
	}
}

func wantXgotYgivenZ(t *testing.T, got interface{}, want interface{}, give interface{}) {
	if got != want {
		t.Errorf("Expected f(%v) = %v, but got %v", give, got, want)
	}
}

func TestIdentifyFormPresentIndicative(t *testing.T) {
	verbs := []string{"ある", "あります", "ない", "ありません"}
	for _, v := range verbs {
		wantXandGotY(t, PresentIndicative, IdentifyForm(v))
	}
}

func TestIdentifyFormPastIndicative(t *testing.T) {
	verbs := []string{"あった", "なかった", "ありました", "ありませんでした"}
	for _, v := range verbs {
		wantXandGotY(t, PastIndicative, IdentifyForm(v))
	}
}

func TestIdentifyPositivity(t *testing.T) {
	verbs := []string{"ない", "ません", "なかった", "ませんでした"}
	for _, v := range verbs {
		wantXandGotY(t, Negative, IdentifyPositivity(v))
	}

	verbs = []string{"ある", "あります", "あった", "ありました"}
	for _, v := range verbs {
		wantXandGotY(t, Positive, IdentifyPositivity(v))
	}
}
