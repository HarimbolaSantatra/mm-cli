#!/bin/bash

set -e
executable="./mm-parsing"
html_file="result.html"
json_file="result.json"
tmp_dir="tmp"

mkdir -p $tmp_dir
tmp1=$(mktemp -p $tmp_dir)
tmp2=$(mktemp -p $tmp_dir)

# Test if the json output is expected
$executable -f $html_file > "$tmp1"
cmp -s "$tmp1" $json_file

# the two method of ./mm-parsing should return the same value
cat $html_file | xargs -0 $executable -c > "$tmp2"
cmp -s "$tmp1" "$tmp2"

rm "$tmp1" "$tmp2"
