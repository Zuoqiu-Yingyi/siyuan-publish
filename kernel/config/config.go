package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

const (
	DEFAULT_DIR_MODE = os.ModeDir | 0764 // 默认目录权限
)

var (
	C *Config
)

type Config struct {
	Debug    bool     `json:"debug"`
	I18n     I18n     `json:i18n""`
	Database Database `json:"database"`
	Server   Server   `json:"server"`
	Siyuan   Siyuan   `json:"siyuan"`
	Render   Render   `json:"render"`
}

func Init(path string) {
	if err := LoadConfigFile(path); err != nil {
		fmt.Printf("load config file error: %v\n", err)

		/* 加载默认配置 */
		LoadDefaultConfig()
	}
	if C.Debug {
		if str, err := json.MarshalIndent(C, "", "    "); err == nil {
			fmt.Println(string(str))
		}
	}
}

func LoadConfigFile(path string) error {
	LoadDefaultConfig()
	if _, err := toml.DecodeFile(path, C); err != nil {
		return err
	}
	return nil
}

func LoadDefaultConfig() {
	C = &Config{
		Debug: false,
		I18n: I18n{
			Directory: "./app/locales",
			Default:   "zh-Hans",
		},
		Database: Database{
			Debug:  false,
			Reset:  true,
			SQLite: "./temp/publish.db",
		},
		Server: Server{
			Debug:     false,
			Port:      80,
			Logs:      "./temp/logs/",
			Templates: "./app/templates/*.html",
			Mode: Mode{
				Page: "dynamic",
				File: "dynamic",
			},
			Index: Index{
				Paths: []string{
					"/",
				},
				URL:  "/block?id=20200812220555-lj3enxa",
				Icon: "#icon-1f4d4",
			},
			Static: Static{
				Reset: false,
				Path:  "./temp/static/",
				Appearance: StaticPath{
					Path:     "/appearance",
					FilePath: "./temp/static/appearance/",
				},
				Assets: StaticPath{
					Path:     "/assets",
					FilePath: "./temp/static/assets/",
				},
				Emojis: StaticPath{
					Path:     "/emojis",
					FilePath: "./temp/static/emojis/",
				},
				Export: StaticPath{
					Path:     "/export",
					FilePath: "./temp/static/export/",
				},
				Stage: StaticPath{
					Path:     "/stage",
					FilePath: "./temp/static/stage/",
				},
				Widgets: StaticPath{
					Path:     "/widgets",
					FilePath: "./temp/static/widgets/",
				},
				Favicon: StaticPath{
					Path:     "/facicon.ico",
					FilePath: "./app/src/facicon.ico",
				},
				JavaScript: StaticPath{
					Path:     "/js",
					FilePath: "./app/src/js/",
				},
				CSS: StaticPath{
					Path:     "/css",
					FilePath: "./app/src/css/",
				},
			},
		},
		Siyuan: Siyuan{
			Debug:   false,
			Server:  "http://localhost:6806",
			Token:   "",
			Timeout: 10000,
			Retry:   3,
			Publish: Publish{
				Access: Access{
					Name:      "custom-publish-access",
					Default:   "private",
					Public:    Attribute{Value: "public"},
					Protected: Attribute{Value: "protected"},
					Private:   Attribute{Value: "private"},
				},
			},
		},
		Render: Render{
			Appearance: Appearance{
				Mode:                2,
				CodeBlockThemeLight: "atom-one-light",
				CodeBlockThemeDark:  "atom-one-dark",
			},
			Editor: Editor{
				CodeLigatures:              true,
				CodeLineWrap:               true,
				CodeSyntaxHighlightLineNum: true,
				FontFamily:                 []string{},
				FontSize:                   16,
				FullWidth:                  true,
				KatexMacros:                "{}",
				PlantUMLServePath:          "https://www.plantuml.com/plantuml/svg/~1",

				Contenteditable: false,
				Spellcheck:      false,
			},
			Plugin: Plugin{
				Load: []string{
					"url",
					"dom",
					"selected",
					"icon",
					"reg",
				},
				Before: []string{
					"theme",
					"font",
					"title",
					"breadcrumb",
					"edit-state",
					"link",
					"locate",
					"popover",
				},
				After: []string{},
			},
			Popover: Popover{
				Timeout: 1000,
				Width:   "33vmax",
				Height:  "50vmin",
			},
			Path: Path{
				Protyle: "/stage/protyle",
				Plugins: "/js/plugins",
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
					Plugin:  "/js/plugin.js",
					Before:  "/js/before.js",
					Render:  "/js/render.js",
					After:   "/js/after.js",
				},
			},
		},
	}
}
