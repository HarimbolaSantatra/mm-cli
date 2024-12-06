#!/bin/bash
# Test if the current structure of the HTML page is the same as the assumed HTML page (scraping-test/result.html)

set -e

dir=$(basename $(pwd))
if [ "$dir" = "test" ]; then
    test_dir="."
    rel_dir=".."
else
    test_dir="test"
    rel_dir="."
fi

temp1="$(mktemp -p $test_dir --suffix=ASSUMED)"
temp2="$(mktemp -p $test_dir --suffix=CURRENT)"

# Assumed version
cat ${rel_dir}/scraping-test/result.html | xargs -0 ${rel_dir}/mm-parsing > $temp1

# Current version
${rel_dir}/test/test-request.sh | xargs -0 ${rel_dir}/mm-parsing > $temp2

# beautify temp files
python3 -m json.tool $temp1 $temp1
python3 -m json.tool $temp2 $temp2

cmp -s $temp2 $temp1
