#!/bin/bash
# Test the HTML parser script (`mm-parsing`)
# Test if the scraper needs an update.
# That's the case only if the website has been updated by the admin
# and it does not have the same HTML layout anymore.

set -e

dir=$(basename "$(pwd)")
if [ "$dir" = "test" ]; then
    rel_dir="."
else
    rel_dir=".."
fi

HTML_TESTFILE="scraping-test/result.html"

parser="$rel_dir/mm-parsing"
assumed=$($parser $HTML_TESTFILE)

html=$("$rel_dir"/test-request.sh)
current=$($parser $html)

if [ "$assumed" = "$current" ]; then
    exit 0
else
    exit 2
fi
