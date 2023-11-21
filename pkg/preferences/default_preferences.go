package preferences

import (
	"changeme/pkg/constants"
)

type DefaultPreferences struct {
	Behavior Behavior `json:"behavior"`
	General  General  `json:"general"`
	Editor   Editor   `json:"editor"`
}

func NewDefaultPreferences() DefaultPreferences {
	return DefaultPreferences{
		Behavior: Behavior{
			AsideWith:    constants.DefaultAsideWith,
			WindowWith:   constants.DefaultWindowWidth,
			WindowHeight: constants.DefaultWindowHeight,
		},
		General: General{
			Theme:       "auto",
			Language:    "auto",
			Font:        "",
			FontSize:    constants.DefaultFontSize,
			CheckUpdate: true,
			SkipVersion: "",
		},
		Editor: Editor{
			Font:     "",
			FontSize: constants.DefaultFontSize,
		},
	}
}

type Behavior struct {
	AsideWith    int `json:"aside_with"`
	WindowWith   int `json:"window_with"`
	WindowHeight int `json:"window_height"`
}

type General struct {
	Theme       string `json:"theme"`
	Language    string `json:"language"`
	Font        string `json:"font"`
	FontSize    int    `json:"font_size"`
	CheckUpdate bool   `json:"check_update"`
	SkipVersion string `json:"skip_version"`
}

type Editor struct {
	Font     string `json:"font"`
	FontSize int    `json:"font_size"`
}
