package util

import "regexp"

func RemoveSpace(bytes []byte) []byte {
	var removeSpace = regexp.MustCompile(`[\s]+`)
	// 除去空格
	content := removeSpace.ReplaceAllString(string(bytes), " ")
	return []byte(content)
}
