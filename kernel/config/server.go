package config

type Server struct {
	Port      int    `json:"port"`
	Templates string `json:"templates"`
	Static    Static `json:"static"`
}

type Static struct {
	Appearance StaticPath `json:"appearance"`
	Assets     StaticPath `json:"assets"`
	Emojis     StaticPath `json:"emojis"`
	Widgets    StaticPath `json:"widgets"`
	Stage      StaticPath `json:"stage"`
	JavaScript StaticPath `json:"javascript"`
	CSS        StaticPath `json:"css"`
}
type StaticPath struct {
	Path     string `json:"path"`
	FilePath string `json:"filepath"`
}
