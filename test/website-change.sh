#!/bin/bash
# Test if the current structure of the HTML page is the same as the assumed HTML page (scraping-test/result.html)

set -e

argnum=$#
verbose=0
compare_mode=0

print_warning() {
    printf "WARN: %s\n" "$1"
}

print_usage() {
    echo "$0 [-d | -v]"
}

compare_strings() {
    # Compare if two string are the same and print the difference
    echo "Compare mode"
    echo "Old file is first; new file is last"
    echo ""
    old=$(mktemp)
    new=$(mktemp)
    echo "$1" > $old
    echo "$2" > $new
    diff $old $new
}

# Add verbose mode and compare mode
if [ $argnum -eq 1 ]; then
    if [ "$1" = "-v" ]; then
	verbose=1
	echo "Verbose mode"
    elif [ "$1" = "-d" ]; then
	compare_mode=1
    else
	print_warning "Unrecognized argument"
	print_usage
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
assumed=$(cat ${rel_dir}/result.html)

# Current version
current=$(${rel_dir}/test/test-request.sh)

if [ "$assumed" = "$current" ]; then
    status=0
else
    status=6
fi

if [ $compare_mode -eq 1 ]; then
    compare_strings "$assumed" "$current"
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
