package japanese

import "testing"

func TestSplitEnding(t *testing.T) {
	verbs := []string{"ある", "あります", "ない", "ありません"}
	roots := []string{"あ", "あり", "", "あり"}
	ends := []string{"る", "ます", "ない", "ません"}
	for i, v := range verbs {
		root, ending := SplitEnding(v)
		wantXandGotY(t, roots[i], root)
		wantXandGotY(t, ends[i], ending)
	}
}

func TestDictionaryFormBasicIchidans(t *testing.T) {
	// ichidan verbs
	verbs := []string{"たべます", "おしえます", "います"}
	dicts := []string{"たべる", "おしえる", "いる"}
	for i, v := range verbs {
		_, ichidan := DictionaryForm(v)
		wantXgotYgivenZ(t, dicts[i], ichidan, v)
	}
}

func TestDictionaryFormBasicGodans(t *testing.T) {
	// godan verbs
	verbs := []string{"いく", "ききます", "ぬぎます", "うたいます"}
	dicts := []string{"いく", "きく", "ぬぐ", "うたう"}
	for i, v := range verbs {
		godan, _ := DictionaryForm(v)
		wantXgotYgivenZ(t, dicts[i], godan, v)
	}
}

func TestDictionaryFormAru(t *testing.T) {
	verbs := []string{"ある", "あります", "ない", "ありません"}
	for _, v := range verbs {
		godan, _ := DictionaryForm(v)
		wantXandGotY(t, "ある", godan)
	}
}

func TestDictionaryFormIchidan(t *testing.T) {
	verbs := []string{"たべます", "たべました", "たべた"}
	for _, v := range verbs {
		_, ichidan := DictionaryForm(v)
		wantXandGotY(t, "たべる", ichidan)
	}

	verbs = []string{"開ける", "開けて", "開けます", "開けない", "開けません"}
	for _, v := range verbs {
		_, ichidan := DictionaryForm(v)
		wantXandGotY(t, "開ける", ichidan)
	}
}

func TestDictionaryFormIchidanAkeru(t *testing.T) {
	verbs := []string{"あける", "あけない", "あけます", "あけません", "あけよう", "あけるだろう", "あけないだろう", "あけましょう", "あけるでしょう", "あけないでしょう", "あけろ", "あけるな", "あけてください", "あけないでください", "あけた", "あけなかった",
		"あけました", "あけませんでした", "あけたろう", "あけただろう", "あけなかっただろう", "あけたでしょう", "あけなかったでしょう"} //"あけている", "あけています", "あけていません", "あけていた", "あけていました", "あけていませんでした", "あければ", "あけなければ", "あけませば", "あけますれば", "あけませんなら", "あけたら", "あけなかったら", "あけましたら", "あけませんでしたら", "あけられる", "あけられない", "あけられます", "あけられません", "あけさせる", "あけさせない", "あけさせます", "あけさせません"}

	for _, v := range verbs {
		_, ichidan := DictionaryForm(v)
		wantXgotYgivenZ(t, "あける", ichidan, v)
	}
}

func TestDictionaryFormGodan(t *testing.T) {
	verbs := []string{"たべます", "たべました", "たべた"}
	for _, v := range verbs {
		_, ichidan := DictionaryForm(v)
		wantXandGotY(t, "たべる", ichidan)
	}

	verbs = []string{"開ける", "開けて", "開けます", "開けない", "開けません"}
	for _, v := range verbs {
		_, ichidan := DictionaryForm(v)
		wantXandGotY(t, "開ける", ichidan)
	}
}
