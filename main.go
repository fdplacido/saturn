package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	rfc3339Format = "rfc3339"
	unixSFormat   = "unix"   // seconds
	unixMsFormat  = "unixms" // milliseconds
	unixUsFormat  = "unixus" // microseconds
	unixNsFormat  = "unixns" // nanoseconds
)

var parseFuncs = map[string]func(string) (time.Time, error){
	rfc3339Format: func(s string) (time.Time, error) {
		return time.Parse(time.RFC3339, s)
	},
	unixSFormat: func(s string) (time.Time, error) { // seconds
		sec, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(sec, 0).UTC(), nil
	},
	unixMsFormat: func(s string) (time.Time, error) { // milliseconds
		ms, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(0, ms*int64(time.Millisecond)).UTC(), nil
	},
	unixUsFormat: func(s string) (time.Time, error) { // microseconds
		us, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(0, us*int64(time.Microsecond)).UTC(), nil
	},
	unixNsFormat: func(s string) (time.Time, error) { // nanoseconds
		ns, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(0, ns).UTC(), nil
	},
}

var formatFuncs = map[string]func(time.Time) string{
	rfc3339Format: func(t time.Time) string {
		return t.UTC().Format(time.RFC3339)
	},
}

func main() {
	input := flag.String("input", "", "Input date/time string")
	inFmt := flag.String("in-format", "date", "Input format: rfc3339, unix, unixms, unixus, unixns")
	outFmt := flag.String("out-format", "rfc3339", "Output format: rfc3339")
	flag.Parse()

	if *input == "" {
		fmt.Fprintln(os.Stderr, "Error: --input is required")
		os.Exit(1)
	}

	parseFunc, ok := parseFuncs[*inFmt]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unsupported in-format: %s\n", *inFmt)
		os.Exit(1)
	}

	formatFunc, ok := formatFuncs[*outFmt]
	if !ok {
		fmt.Fprintf(os.Stderr, "Unsupported out-format: %s\n", *outFmt)
		os.Exit(1)
	}

	t, err := parseFunc(*input)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(formatFunc(t))
}
