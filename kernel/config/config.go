package config

import (
	"github.com/BurntSushi/toml"
)

var (
	C *Config
)

type Config struct {
	Server Server `json:"server"`
	Siyuan Siyuan `json:"siyuan"`
	Render Render `json:"render"`
}

func init() {
}

func LoadConfigFile(path string) error {
	C = &Config{}
	if _, err := toml.DecodeFile(path, C); err != nil {
		return err
	}
	return nil
}

func LoadDefaultConfig() {
	C = &Config{
		Server: Server{
			Debug:     false,
			Port:      8080,
			Templates: "./../app/templates/*.html",
			Index: Index{
				Paths: []string{
					"/",
				},
				URL: "/block?id=20200812220555-lj3enxa",
			},
			Static: Static{
				Appearance: StaticPath{
					Path:     "/appearance",
					FilePath: "./../app/static/appearance",
				},
				Assets: StaticPath{
					Path:     "/assets",
					FilePath: "./../app/static/assets",
				},
				Emojis: StaticPath{
					Path:     "/emojis",
					FilePath: "./../app/static/emojis",
				},
				Widgets: StaticPath{
					Path:     "/widgets",
					FilePath: "./../app/static/widgets",
				},
				Stage: StaticPath{
					Path:     "/stage",
					FilePath: "./../app/static/stage",
				},
				JavaScript: StaticPath{
					Path:     "/js",
					FilePath: "./../app/src/js",
				},
				CSS: StaticPath{
					Path:     "/css",
					FilePath: "./../app/src/css",
				},
			},
		},
		Siyuan: Siyuan{
			Server:  "http://localhost:6806",
			Token:   "",
			Timeout: 10000,
			Retry:   3,
			Publish: Publish{
				Access: Access{
					Name:      "custom-publish-access",
					Public:    Attribute{Value: "public"},
					Protected: Attribute{Value: "protected"},
					Private:   Attribute{Value: "private"},
				},
			},
		},
		Render: Render{
			Appearance: Appearance{
				Mode:                0,
				CodeBlockThemeLight: "atom-one-light",
				CodeBlockThemeDark:  "atom-one-dark",
			},
			Editor: Editor{
				FontSize:                   16,
				CodeLineWrap:               true,
				CodeLigatures:              true,
				CodeSyntaxHighlightLineNum: true,
				PlantUMLServePath:          "https://www.plantuml.com/plantuml/svg/~1",
			},
			Path: Path{
				Protyle: "/stage/protyle",
			},
			File: File{
				Style: Style{
					Base:    "/stage/build/export/base.css",
					Publish: "/css/publish.css",
					Light: Theme{
						Default: "/appearance/themes/daylight/theme.css",
						Theme:   "/appearance/themes/daylight/theme.css",
						Custom:  "/appearance/themes/daylight/custom.css",
					},
					Dark: Theme{
						Default: "/appearance/themes/midnight/theme.css",
						Theme:   "/appearance/themes/midnight/theme.css",
						Custom:  "/appearance/themes/midnight/custom.css",
					},
				},
				Script: Script{
					Icon:    "/appearance/icons/material/icon.js",
					Emoji:   "/appearance/emojis/twitter-emoji.js",
					Lute:    "/stage/protyle/js/lute/lute.min.js",
					HTML:    "/stage/protyle/js/protyle-html.js",
					Protyle: "/stage/build/export/protyle-method.js",
					Config:  "/js/config.js",
					Publish: "/js/publish.js",
				},
			},
		},
	}
}
