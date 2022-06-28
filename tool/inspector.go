package tool

import "strings"

var illegalWords = []string{
	"习进", "反gong", "操", "NM", "nm", "sb", "SB",
	"<script","select *","SELECT *","ORDER BY","order by",
	"and if","AND IF","SLEEP()","sleep()","=",
}

func IllegalWordsInspect(content string) bool {
	for _, v := range illegalWords {
		if strings.Contains(content, v) {
			return false
		} 
	}
	return true
}