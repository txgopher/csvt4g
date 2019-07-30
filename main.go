// Some useful tools for processing csv file
package main

import (
	"bingo/easylog"
	"os"
	"runtime"
	"strconv"
)

func init() {
	easylog.Initialize(os.Stderr)
	runtime.GOMAXPROCS(2)
}

var INFO = easylog.INFO
var WARN = easylog.WARN
var ERROR = easylog.ERROR

func usage() {
	println()
	println("Usage: " + os.Args[0] + " [OPTIONS]")
	println("    -r    remove all quotation marks on fileds for standard csv file")
	println("    -f    filtering out useless lines unmatch the regular expression")
	println("    -s    translate csv file into libsvm format file with label file")
	println("    -i    input file for reading")
	println("    -o    output file for writing")
	println("    -d    delimeter of csv file, should be a valid character, default is \",\"")
	println("    -x    valid regular expression for `-f` option, grep or pcre style")
	println("    -n    column number for `-f` option to match the specified regex")
	println("    -l    label file name for `-s` option to fill the libsvm file\n")
	println("Example:")
	println("    ./csvt4g -r -i in.csv -o out.csv -d \",\"")
	println("    ./csvt4g -f -i in.csv -o out.csv -n 2 -x \"^[a-z0-9]+$\"")
	println("    ./csvt4g -s -i in.csv -o out.libsvm -l label.csv")
	println()
	os.Exit(-1)
}

func main() {
	if len(os.Args) < 2 || os.Args[1] == "-h" || os.Args[1] == "--help" {
		usage()
	}

	if len(os.Args) != 6 && len(os.Args) != 8 && len(os.Args) != 10 && len(os.Args) != 12 {
		println("Error argument, see help for more details")
		os.Exit(-1)
	}

	var schema string
	var input, output, label string
	var delim string
	var regex string
	var ncol int

	for i := 1; i < len(os.Args); i++ {
		s := os.Args[i]
		switch s {
		case "-r":
			fallthrough
		case "-f":
			fallthrough
		case "-s":
			schema = s
		case "-i":
			i++
			input = os.Args[i]
		case "-o":
			i++
			output = os.Args[i]
		case "-l":
			i++
			label = os.Args[i]
		case "-d":
			i++
			delim = os.Args[i]
		case "-x":
			i++
			regex = os.Args[i]
		case "-n":
			i++
			n, err := strconv.Atoi(os.Args[i])
			if err != nil {
				println(err.Error())
				usage()
			}
			ncol = n
		default:
			break
		}
	}

	if delim == "" {
		delim = COMMA
	}

	switch schema {
	case "-r":
		if err := remove_quotations(input, output); err != nil {
			ERROR(err)
		}
	case "-f":
		if err := filter_by_regex(input, output, delim, regex, ncol); err != nil {
			ERROR(err)
		}
	case "-s":
		if err := csv_to_libsvm(input, output, label, delim); err != nil {
			ERROR(err)
		}
	default:
		usage()
	}
}
