package bytestrings

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"io"
	"os"
	"strings"
)

func SearchString() {
	s := "this is a test"
	fmt.Println(strings.Contains(s, "this"))
	fmt.Println(strings.ContainsAny(s, "abc"))
	fmt.Println(strings.HasPrefix(s, "this"))
	fmt.Println(strings.HasSuffix(s, "test"))
}

func ModifyString() {
	s := "simple string строка"
	fmt.Println(strings.Split(s, " "))
	fmt.Println(strings.Title(s))
	fmt.Println(cases.Title(language.English).String(s))
}

func StringReader() {
	s := "simple string\n"
	r := strings.NewReader(s)
	io.Copy(os.Stdout, r)
}
