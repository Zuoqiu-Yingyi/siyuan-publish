package utils

import "github.com/88250/lute/parse"

func IAL2Map(IAL string) map[string]string {
	return parse.IAL2Map(parse.Tokens2IAL([]byte(IAL)))
}
