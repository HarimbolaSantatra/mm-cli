#!/bin/bash
# Test if the current structure of the HTML page is the same as the assumed HTML page (scraping-test/result.html)

set -e

dir=$(basename "$(pwd)")
if [ "$dir" = "test" ]; then
    rel_dir=".."
else
    rel_dir="."
fi

# Assumed version
assumed=$(${rel_dir}/mm-parsing ${rel_dir}/scraping-test/result.html)

# Current version
current=$(${rel_dir}/test/test-request.sh)

if [ "$assumed" = "$current" ]; then
    exit 0
else
    exit 3
fi
