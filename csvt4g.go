// Some useful tools for processing csv file
package main

import (
	"bingo/fio"
	fs "bingo/system"
	"bufio"
	"bytes"
	"errors"
	"os"
	"regexp"
	"strconv"
)

const (
	COMMA = `,`
	QUOTA = `"`
)

func print_sched(nline int) {
	if (nline % 10) == 10000 {
		INFO(nline, "lines parsed ...")
	}
}

func remove_quotations(input, output string) (err error) {
	if !fs.IsExist(input) || !fs.IsFile(input) {
		return errors.New("File not exists!")
	}

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	bq := []byte(QUOTA)

	INFO("input:  ", input)
	INFO("output: ", output)

	nlines := 0
	fio.ReadLineAsBytes(input, false, func(line []byte) {
		nlines++
		print_sched(nlines)
		w.Write(bytes.Replace(line, bq, []byte{}, -1))
		w.WriteByte('\n')
	})

	w.Flush()
	INFO("finished,", nlines, "line parsed")

	return
}

func filter_by_regex(input, output, delim, regex string, col int) (err error) {
	if !fs.IsExist(input) || !fs.IsFile(input) {
		return errors.New("File not exists!")
	}

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()

	x, err := regexp.Compile(regex)
	if err != nil {
		return err
	}

	if col < 1 {
		return errors.New("column index error")
	}
	first, err := fio.FirstLine(input)
	if err != nil || len(first) == 0 {
		return err
	}

	sep := []byte(delim)
	seq := bytes.Split([]byte(first), sep)

	if len(seq) < col {
		return errors.New("column index error")
	}

	INFO("input:  ", input)
	INFO("output: ", output)
	INFO("regex:  ", regex)
	INFO("column: ", col)

	w := bufio.NewWriter(file)
	nlines := 0
	nmatch := 0

	fio.ReadLineAsBytes(input, false, func(line []byte) {
		nlines++
		print_sched(nlines)

		seq = bytes.Split([]byte(line), sep)

		if len(seq) >= col && x.Match(seq[col-1]) {
			w.Write(line)
			w.WriteByte('\n')
			nmatch++
			return
		}
	})

	w.Flush()
	INFO("finished,", nlines, "lines parsed,", nmatch, "lines filtered")

	return
}

func csv_to_libsvm(input, output, labelfile, delim string) (err error) {
	if !fs.IsExist(input) || !fs.IsFile(input) {
		return errors.New("File not exists!")
	}

	if len(labelfile) > 0 && !fs.IsExist(labelfile) {
		return errors.New("File not exists!")
	}

	file, err := os.Create(output)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)

	sep := []byte(delim)

	INFO("input:  ", input)
	INFO("label:  ", labelfile)
	INFO("output: ", output)

	labels := make([][]byte, 0)
	if len(labelfile) > 0 {
		fio.ReadLineAsBytes(labelfile, false, func(line []byte) {
			line = append(line, ' ')
			labels = append(labels, line)
		})
	}

	i := 0

	fio.ReadLineAsBytes(input, false, func(line []byte) {
		i++
		print_sched(i)
		seq := bytes.Split([]byte(line), sep)
		if i <= len(labels) {
			w.Write(labels[i-1])
		} else {
			w.Write([]byte("0 "))
		}

		for k, s := range seq {
			if len(s) > 1 {
				if s[0] == '"' {
					s = s[1:]
				}
				if s[len(s)-1] == '"' {
					s = s[:len(s)-1]
				}
			}

			if v, e := strconv.ParseFloat(string(s), 64); e == nil {
				if v == 0.0 {
					continue
				}
			}

			w.WriteString(strconv.Itoa(k + 1))
			w.WriteByte(':')
			w.Write(s)
			if k < len(seq)-1 {
				w.WriteByte(' ')
			}
		}
		w.WriteByte('\n')
	})

	w.Flush()
	INFO("finished,", i, "lines parsed")

	return
}
