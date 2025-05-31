package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	rfc3339Format     = "rfc3339"
	rfc3339NanoFormat = "rfc3339nano"
	unixSFormat       = "unix"   // seconds
	unixMsFormat      = "unixms" // milliseconds
	unixUsFormat      = "unixus" // microseconds
	unixNsFormat      = "unixns" // nanoseconds
	autoUnix          = "autounix"
)

func parseAutoUnix(s string) (time.Time, error) {
	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	switch {
	case val < 1e11:
		return time.Unix(val, 0).UTC(), nil // seconds
	case val < 1e14:
		return time.Unix(0, val*int64(time.Millisecond)).UTC(), nil // milliseconds
	case val < 1e17:
		return time.Unix(0, val*int64(time.Microsecond)).UTC(), nil // microseconds
	default:
		return time.Unix(0, val).UTC(), nil // nanoseconds
	}
}

func parseRfc3339(s string) (time.Time, error) {
	return time.Parse(time.RFC3339, s)
}

func parseRfc3339Nano(s string) (time.Time, error) {
	return time.Parse(time.RFC3339Nano, s)
}

func parseUnixSeconds(s string) (time.Time, error) {
	sec, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(sec, 0).UTC(), nil
}

func parseUnixMilliSeconds(s string) (time.Time, error) {
	ms, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, ms*int64(time.Millisecond)).UTC(), nil
}

func parseUnixMicroSeconds(s string) (time.Time, error) {
	us, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, us*int64(time.Microsecond)).UTC(), nil
}

func parseUnixNanoSeconds(s string) (time.Time, error) {
	ns, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(0, ns).UTC(), nil
}

var parseFuncs = map[string]func(string) (time.Time, error){
	rfc3339Format:     parseRfc3339,
	rfc3339NanoFormat: parseRfc3339Nano,
	unixSFormat:       parseUnixSeconds,
	unixMsFormat:      parseUnixMilliSeconds,
	unixUsFormat:      parseUnixMicroSeconds,
	unixNsFormat:      parseUnixNanoSeconds,
	autoUnix:          parseAutoUnix,
}

var formatFuncs = map[string]func(time.Time) string{
	rfc3339Format: func(t time.Time) string {
		return t.UTC().Format(time.RFC3339)
	},
}

func main() {
	input := flag.String("input", "", "Input date/time string")
	inFmt := flag.String("in-format", "date", "Input format: rfc3339, rfc3339nano, unix, unixms, unixus, unixns, autounix")
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
