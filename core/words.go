package core

var (
	Words = []Word{
		{
			rusPresent: Combination{
				I:    "люблю",
				You:  "любишь",
				They: "любят",
				He:   "любит",
				She:  "любит",
			},
			rusPast: Combination{
				I:    "любил",
				You:  "любил",
				They: "любили",
				He:   "любил",
				She:  "любила",
			},
			rusFuture: Combination{
				I:    "полюблю",
				You:  "полюбишь",
				They: "полюбят",
				He:   "полюбит",
				She:  "полюбит",
			},

			engBase:    "love",
			eng3Person: "loves",
			engPast:    "loved",
		},
		{
			rusPresent: Combination{
				I:    "знаю",
				You:  "знаешь",
				They: "знают",
				He:   "знает",
				She:  "знает",
			},
			rusPast: Combination{
				I:    "знал",
				You:  "знал",
				They: "знали",
				He:   "знал",
				She:  "знала",
			},
			rusFuture: Combination{
				I:    "узнаю",
				You:  "узнаешь",
				They: "узнают",
				He:   "узнает",
				She:  "узнает",
			},

			engBase:    "know",
			eng3Person: "knows",
			engPast:    "knew",
		},
		{
			rusPresent: Combination{
				I:    "иду",
				You:  "идешь",
				They: "идут",
				He:   "идет",
				She:  "идет",
			},
			rusPast: Combination{
				I:    "шел",
				You:  "шел",
				They: "шли",
				He:   "шел",
				She:  "шла",
			},
			rusFuture: Combination{
				I:    "пойду",
				You:  "пойдешь",
				They: "пойдут",
				He:   "пойдет",
				She:  "пойдет",
			},

			engBase:    "go",
			eng3Person: "goes",
			engPast:    "went",
		},
	}
)
