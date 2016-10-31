/*
Their names are matched by this regular expression: Find(All)?(String)?(Submatch)?(Index)?

Find()  FindAll()  FindAllSubmatch()
*/

package main

import (
	"fmt"
	"regexp"
)

var digitRegxp = regexp.MustCompile(`(\d+)\D+(\d+)`)
var myExp = regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)

// embed regexp.Regexp in a new type so we can extend it
type myRegexp struct {
	*regexp.Regexp
}

func (r *myRegexp) FindStringSubmatchMap(s string) map[string]string {
	captures := make(map[string]string)

	match := r.FindStringSubmatch(s)
	if match == nil {
		return captures
	}

	for i, name := range r.SubexpNames() {
		// Ignore the whole regexp match and unnamed groups
		if i == 0 || name == "" {
			continue
		}

		captures[name] = match[i]

	}
	return captures
}

var selfExp = myRegexp{regexp.MustCompile(`(?P<first>\d+)\.(\d+).(?P<second>\d+)`)}
var exp2 = myRegexp{regexp.MustCompile(`(?P<first>\d+)\.(\D+).(?P<second>\d+)`)}

func main() {
	s := "1000abcd123"
	//fmt.Println(digitRegxp.FindSubmatch(s))
	fmt.Println(regexp.MustCompile(`\d+`).FindString(s))
	fmt.Println(regexp.MustCompile(`\d+`).FindAllString(s, -1))
	fmt.Println(digitRegxp.FindStringSubmatch(s))
	fmt.Printf("%+v\n", myExp.FindStringSubmatch("1234.5678.90.12"))
	fmt.Printf("%+v\n", selfExp.FindStringSubmatchMap("1234.5678.9"))
	fmt.Printf("%+v\n", exp2.FindStringSubmatchMap("1234.abc.901"))
}
