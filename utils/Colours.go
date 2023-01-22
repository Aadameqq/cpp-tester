package utils

var Colours = struct {
	Reset  string
	Yellow string
	Green  string
	Red    string
	Blue   string
	Purple string
}{
	Reset:  "\033[0m",
	Blue:   "\033[36m",
	Green:  "\033[32m",
	Red:    "\033[31m",
	Yellow: "\033[33m",
	Purple: "\033[35m",
}
