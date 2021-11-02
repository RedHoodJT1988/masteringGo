package main

import (
	"fmt"
	s "strings"
	"unicode"
)

var f = fmt.Printf

func main() {
	upper := s.ToUpper("Hello there!")
	f("To Upper: %s\n", upper)
	f("To Lower: %s\n", s.ToLower("Hello THERE"))

	f("%s\n", s.Title("tHis will be A title!"))

	f("EqualFold: %v\n", s.EqualFold("Jonathan", "JONAthan"))
	f("EqualFold: %v\n", s.EqualFold("Jonathan", "JONAthan"))

	f("Prefix: %v\n", s.HasPrefix("Jonathan", "Jo"))
	f("Prefix: %v\n", s.HasPrefix("Jonathan", "jo"))
	f("Suufix: %v\n", s.HasSuffix("Jonathan", "an"))
	f("Suffix: %v\n", s.HasSuffix("Jonathan", "AN"))

	f("Index: %v\n", s.Index("Jonathan", "at"))
	f("Index: %v\n", s.Index("Jonathan", "At"))
	f("Count: %v\n", s.Count("Jonathan", "a"))
	f("Count: %v\n", s.Count("Jonathan", "A"))
	f("Repeat: %s\n", s.Repeat("ab", 5))

	f("TrimSpace: %s\n", s.TrimSpace(" \tThis is a line. \n"))
	f("TrimLeft: %s", s.TrimLeft(" \tThis is a\t line. \n", "\n\t"))
	f("TrimRight: %s\n", s.TrimRight(" \tThis is a\t line. \n", "\n\t"))

	f("Compare: %v\n", s.Compare("Jonathan", "JONATHAN"))
	f("Compare: %v\n", s.Compare("Jonathan", "Jonathan"))
	f("Compare: %v\n", s.Compare("JONATHAN", "JONathan"))

	t := s.Fields("This is a string!")
	f("Fields: %v\n", len(t))
	t = s.Fields("ThisIs a\tstring!")
	f("Fields: %v\n", len(t))

	f("%s\n", s.Split("abcd efg", ""))
	f("%s\n", s.Replace("abcd efg", "", "_", -1))
	f("%s\n", s.Replace("abcd efg", "", "_", 4))
	f("%\n", s.Replace("abcd efg", "", "_", 2))

	lines := []string{"Line 1", "Line 2", "Line 3"}
	f("Join: %s\n", s.Join(lines, "+++"))

	f("SplitAfter: %s\n", s.SplitAfter("123++432++", "++"))

	trimFunction := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	f("TrimFunc: %s\n", s.TrimFunc("123 abc ABC \t .", trimFunction))
}
