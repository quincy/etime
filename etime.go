package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var formats = []string{
	time.ANSIC,
	time.UnixDate,
	time.RubyDate,
	time.RFC822,
	time.RFC822Z,
	time.RFC850,
	time.RFC1123,
	time.RFC1123Z,
	time.RFC3339,
	time.RFC3339Nano,
	time.Kitchen,
	time.Stamp,
	time.StampNano,
	time.StampMilli,
	time.StampMicro,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Print(help())
		os.Exit(1)
	}

	var arg = os.Args[1]

	if arg == "-h" || arg == "--help" {
		fmt.Print(help())
		return
	}

	if parseUnixTime(arg) {
		return
	}

	if parseTimeString(arg) {
		return
	}

	fmt.Printf("Could not parse argument [%s] with any known time format.\n", arg)
	fmt.Print(help())
	os.Exit(1)
}

func help() string {
	return fmt.Sprintf("%s converts between Unix epoch timestamps and human readable forms.\n\n"+
		"usage: %s [-h] {timespec}\n\n"+
		"-h, --help   prints this message\n"+
		"timespec     is a string representation of a time to convert in one of the following formats:\n"+
		"\t\t[Unix epoch time]\n"+
		"\t\t%s\n",
		os.Args[0],
		os.Args[0],
		strings.Join(formats, "\n\t\t"))
}

func parseUnixTime(val string) bool {
	etime, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return false
	}

	fmt.Printf("%s\n", time.Unix(etime, 0).Format(time.RFC1123Z))
	return true
}

func parseTimeString(val string) bool {
	for _, format := range formats {
		if timeValue, err := time.Parse(format, val); err == nil {
			fmt.Printf("%d\n", timeValue.Unix())
			return true
		}
	}

	return false
}
