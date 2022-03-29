package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/tanema/identigon"
)

func fatal(err string) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func printUsage() {
	fmt.Println(`idntcn - Fast idententicon generation.

Usage:	idntcn [options] <string>

Options:`)
	flag.PrintDefaults()
	fmt.Println(`
Examples:
  - idntcn "tim@chips.com" > ident.png
  - idntcn -o ident.png "tim@chips.com"
  - echo "tim@chips.com" | idntcon > ident.png`)
}

func main() {
	flag.Usage = printUsage
	size := flag.Int("s", 80, "square size of the image in pixels")
	pix := flag.Int("p", 8, "amount of pixel blocks in the image")
	out := flag.String("o", "", "path to ouput file (defaults to stdout)")
	flag.Parse()

	var data string
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		var buf bytes.Buffer
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			buf.Write(scanner.Bytes())
		}
		if err := scanner.Err(); err != nil {
			fatal(err.Error())
		}
		data = buf.String()
	} else if len(os.Args) > 1 {
		data = strings.Join(os.Args[1:], " ")
	} else {
		printUsage()
		fatal("\n[Err] no data provided")
	}

	output := os.Stdout
	if *out != "" {
		var err error
		output, err = os.Create(*out)
		if err != nil {
			fatal(err.Error())
		}
	}

	if err := png.Encode(output, identigon.Generate(strings.TrimSuffix(data, "\n"), *size, *pix)); err != nil {
		fatal(err.Error())
	}
}
