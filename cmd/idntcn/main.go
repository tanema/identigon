package main

import (
	"bufio"
	"bytes"
	"errors"
	"flag"
	"fmt"
	"image/png"
	"os"
	"strings"

	"github.com/tanema/identigon"
)

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}

func main() {
	size := flag.Int("size", 80, "size of output image")
	pix := flag.Int("blocks", 8, "amount of pixel blocks in the image")
	border := flag.Int("border", 0, "width of border.")
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
			fatal(err)
		}
		data = buf.String()
	} else if len(os.Args) > 1 {
		data = strings.Join(os.Args[1:], " ")
	} else {
		fatal(errors.New("no data provided"))
	}

	if err := png.Encode(os.Stdout, identigon.Generate(strings.TrimSuffix(data, "\n"), *size, *pix, *border)); err != nil {
		fatal(err)
	}
}
