#!/bin/bash
# Test if the current structure of the HTML page is the same as the assumed HTML page (scraping-test/result.html)

set -e

argnum=$#
verbose=0

print_warning() {
    printf "WARN: %s\n" "$1"
}


# Verbose mode
if [ $argnum -eq 1 ]; then
    if [ "$1" = "-v" ]; then
	verbose=1
	echo "Verbose mode"
    else
	print_warning "Unrecognized argument"
	exit 2
    fi
fi

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
    status=0
else
    status=1
fi

if [ $verbose -eq 1 ]; then
    echo "ASSUMED VALUE:"
    echo "$assumed"
    echo ""
    echo ""
    echo "======"
    echo ""
    echo ""
    echo "CURRENT VALUE:"
    echo "$current"
fi

exit $status
