#!/bin/bash
# Make a test request to motmalgache.org

BASE_URL="https://motmalgache.org/bins"
SEARCH_ENDPOINT="${BASE_URL}/teny2"
curl -Ls \
    -XPOST \
    --data 'w=lakozia' \
    -H "Accept: text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/png,image/svg+xml,*/*;q=0.8", \
    -H "Content-Type: application/x-www-form-urlencoded", \
    $SEARCH_ENDPOINT
