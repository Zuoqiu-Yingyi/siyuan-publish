package locale

import (
	"fmt"

	"github.com/kataras/i18n"
)

var (
	L    *i18n.I18n
	Lang string
)

func Init(directory string, langCode string) {
	if l, err := i18n.New(i18n.Glob(fmt.Sprintf("%s/*/*", directory))); err != nil {
		panic(err)
	} else {
		l.SetDefault(langCode)
		L = l
		Lang = langCode
	}
}

func T(format string, args ...interface{}) string {
	return L.Tr(Lang, format, args...)
}
