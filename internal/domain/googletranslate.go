package domain

type (
	GoogleLanguage struct {
		Language string `json:"language"`
	}

	GoogleDetector struct {
		IsReliable bool `json:"isReliable"`
		Confidence int  `json:"confidence"`
		GoogleLanguage
	}

	GoogleTranslate struct {
		Text string `json:"translatedText"`
	}
)
