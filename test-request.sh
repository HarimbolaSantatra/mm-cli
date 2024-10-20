#!/bin/bash

BASE_URL="https://motmalgache.org/bins"
SEARCH_ENDPOINT="${BASE_URL}/teny2"
curl -Lv \
    -XPOST \
    --data '{"w": "lakozia"}' \
    $SEARCH_ENDPOINT
