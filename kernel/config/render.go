package config

type Render struct {
	Mode int  `json:"mode"`
	Path Path `json:"path"`
	File File `json:"file"`
}

type File struct {
	Style  Style  `json:"style"`
	Script Script `json:"script"`
}

type Style struct {
	Font    string `json:"font"`
	Base    string `json:"base"`
	Publish string `json:"publish"`
	Light   Theme  `json:"light"`
	Dark    Theme  `json:"dark"`
}

type Theme struct {
	Default string `json:"default"`
	Theme   string `json:"theme"`
	Custom  string `json:"custom"`
}

type Script struct {
	Icon    string `json:"icon"`
	Emoji   string `json:"emoji"`
	Lute    string `json:"lute"`
	HTML    string `json:"html"`
	Protyle string `json:"protyle"`
	Config  string `json:"config"`
	Publish string `json:"publish"`
}

type Path struct {
	Protyle string `json:"protyle"`
}
