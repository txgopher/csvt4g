# csvt4g
simple csv tools for golang  

```
Usage: ./csvt4g [OPTIONS]
    -r    remove all quotation marks on fields in csv file.
    -f    filtering out useless lines unmatched the regex of -x option.
    -s    translate csv file into libsvm format with label file.
    -i    input file for reading.
    -o    output file for writing.
    -d    delimeter of csv, should be a valid character, default is ",".
    -x    valid regular expression for `-f` option, grep or pcre style.
    -n    column number for `-f` option to match the specified regex.
    -l    label file name for `-s` option to fill the libsvm file.

Example:
    ./csvt4g -r -i in.csv -o out.csv -d ","
    ./csvt4g -f -i in.csv -o out.csv -n 2 -x "^[a-z0-9]+$"
    ./csvt4g -s -i in.csv -o out.libsvm -l label.csv
```

Test  
```
./test/run.sh
```
