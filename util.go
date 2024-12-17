package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func handleErr(errMsg string) {
	fmt.Fprintf(os.Stderr, "%s %s\n", Bold(Color("error:", RED)), errMsg)
	os.Exit(1)
}

var BLACK string = "0"
var RED string = "1"
var GREEN string = "2"
var YELLOW string = "3"
var BLUE string = "4"
var PURPLE string = "5"
var CYAN string = "6"
var GRAY string = "7"

func Color(text string, color string) string {
	if color[0] == '#' {
		r := color[1:3]
		g := color[3:5]
		b := color[5:7]
		return fmt.Sprintf(
			"\x1b[38;2;%d;%d;%dm%s\x1b[0m",
			UnwrapOr(strconv.ParseUint(r, 16, 8))(0),
			UnwrapOr(strconv.ParseUint(g, 16, 8))(0),
			UnwrapOr(strconv.ParseUint(b, 16, 8))(0),
			text,
		)
	} else {
		return fmt.Sprintf("\x1b[3%sm%s\x1b[0m", color, text)
	}
}

func Italic(text ...string) string {
	return "\x1b[3m" + strings.Join(text, " \x1b[3m") + "\x1b[0m"
}

func Bold(text ...string) string {
	return "\x1b[1m" + strings.Join(text, " \x1b[1m") + "\x1b[0m"
}

func Unwrap[T any](val T, err error) T {
	if err != nil { handleErr(err.Error()) }

	return val
}

func UnwrapOr[T any](val T, err error) func(T) T {
	if err != nil {
		return func(d T) T {
			return d
		}
	} else {
		return func(_ T) T {
			return val
		}
	}
}

func UnwrapOrElse[T any](val T, err error) func(func() T) T {
	if err != nil {
		return func(fn func() T) T {
			return fn()
		}
	} else {
		return func(_ func() T) T {
			return val
		}
	}

}

func Expect(err error) {
	if err != nil { handleErr(err.Error()) }
}
