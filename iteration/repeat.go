package repeat

func Repeat(s string, times int) string {
	repeated_s := ""
	for i := 0; i < times; i++ {
		repeated_s += s
	}
	return repeated_s
}
