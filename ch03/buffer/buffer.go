package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	method := flag.String("method", "buffer", "use \"concat\", \"join\" or \"buffer\" to create result")
	output := flag.String("output", "length", "show \"length\" or \"result\"")
	flag.Parse()

	if len(os.Args) < 2 || (*method != "concat" && *method != "join" && *method != "buffer") ||
		(*output != "length" && *output != "result") {
		fmt.Fprintf(os.Stderr, "usage: %s [flags] number-of-strings\n", filepath.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(1)
	}

	count, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: invalid number %s",
			filepath.Base(os.Args[0]), flag.Arg(0))
		os.Exit(1)
	}

	var result string

	if *method == "buffer" {
		var buf bytes.Buffer
		for i := 0; i < count; i++ {
			buf.WriteString(fmt.Sprint(i))
			if i < count-1 {
				buf.WriteString(", ")
			}
		}
		result = buf.String()
	} else if *method == "join" {
		var buf []string
		for i := 0; i < count; i++ {
			buf = append(buf, fmt.Sprint(i))
		}
		result = strings.Join(buf, ", ")
	} else if *method == "concat" {
		var buf string
		for i := 0; i < count; i++ {
			buf += fmt.Sprint(i)
			if i < count-1 {
				buf += ", "
			}
		}
		result = buf
	}

	if *output == "result" {
		fmt.Println(result)
	} else {
		fmt.Println(len(result))
	}
}
