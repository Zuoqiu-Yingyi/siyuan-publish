package utils

import "regexp"

func GetStyleFileName(html string) (file_name string) {
	re := regexp.MustCompile(`<link href="(base\.[0-9a-f]{20}\.css)" rel="stylesheet">`)
	result := re.FindStringSubmatch(html)
	if len(result) > 1 {
		file_name = result[1]
	} else {
		panic("base.hash.css file not found")
	}
	return
}
