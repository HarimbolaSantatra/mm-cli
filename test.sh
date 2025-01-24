#!/bin/bash

executable="./mm-parsing"
html_file="result.html"
json_file="result.json"
tmp_dir="tmp"

mkdir -p $tmp_dir
tmp=$(mktemp -p $tmp_dir)

$executable $html_file > "$tmp"

cmp -s "$tmp" $json_file
