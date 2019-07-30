#!/bin/sh

cd ./test
echo "\nSimple testing of csvtools\n"

echo "filter"
../csvt4g -f -i raw.csv -o filter.csv -n 3 -x "^\"[0-9]+\"$"

echo "remove"
../csvt4g -r -i raw.csv -o input.csv

echo "csv2libsvm"
../csvt4g -s -i input.csv -l label.csv -o output.svm 

echo "\nFinish, see test dir for results\n"


