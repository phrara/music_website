package tool

import (
	"fmt"
	"math"
	"strconv"
)

func DurationTransform(ms string) string {
	duration, _ := strconv.ParseFloat(ms, 64) 
	duration = duration / 60000
	integer := math.Ceil(duration) - 1
	second := (duration - integer) * 60
	minutes := int(integer)
	if second > 9 {
		return fmt.Sprintf("%d:%d", minutes, int(math.Ceil(second)))
	} else {
		return fmt.Sprintf("%d:0%d", minutes, int(math.Ceil(second)))
	}
}