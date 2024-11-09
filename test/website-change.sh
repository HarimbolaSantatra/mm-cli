#!/bin/bash
temp1="$(mktemp -p test/)"
temp2="$(mktemp -p test/)"

cat scraping-test/result.html | xargs -0 ./scraping > $temp1
./test/test-request.sh | xargs -0 ./scraping > $temp2

cmp -s $temp2 $temp1
