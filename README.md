# csvt4g
simple csv tools for golang  

```
Usage: ./csvt4g [OPTIONS]
    -r    Remove all quotation marks on fields in csv file.
    -f    Filtering out useless lines unmatched the regex of -x option.
    -s    Translate csv file into libsvm format with label file.
    -i    Input file for reading.
    -o    Output file for writing.
    -d    Delimeter of csv, should be a valid character, default is ",".
    -x    Valid regular expression for `-f` option, grep or pcre style.
    -n    Column number for `-f` option to match the specified regex.
    -l    Label file name for `-s` option to fill the libsvm file.

Example:
    ./csvt4g -r -i in.csv -o out.csv -d ","
    ./csvt4g -f -i in.csv -o out.csv -n 2 -x "^[a-z0-9]+$"
    ./csvt4g -s -i in.csv -o out.libsvm -l label.csv
```

Test  
```
./test/run.sh
```
