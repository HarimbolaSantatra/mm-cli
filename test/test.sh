#!/bin/bash
# Test the HTML parser script (`mm-parsing`)
# Test if the scraper needs an update.
# That's the case only if the website has been updated by the admin
# and it does not have the same HTML layout anymore.

HTML_TESTFILE="scraping-test/result.html"
JSON_RESULT="scraping-test/parsing-result.json"

parser="../mm-parsing"
assumed=$($parser $HTML_TESTFILE)

html=$(./test-request.sh)
current=$($parser $html)

if [ "$assumed" = "$current" ]; then
    exit 0
else
    exit 2
fi
