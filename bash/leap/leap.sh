#!/usr/bin/env bash

# check correct number of arguments
if [[ $# -ne 1 ]]; then
    echo "Usage: leap.sh <year>"
    exit 1
fi

# check argument is not float
re='^[0-9]+$'
if ! [[ $1 =~ $re ]] ; then
    echo "Usage: leap.sh <year>"
    exit 1
fi

# check if argument is leap year
LEAP="false"
if [ $(( $1 % 4 )) == 0 ]; then
    LEAP="true"
    if [ $(( $1 % 100 ))  == 0 ]; then
        LEAP="false"
        if [ $(( $1 % 400 )) == 0 ]; then
            LEAP="true"
        fi
    fi
fi

# print result and exit gracefully
echo $LEAP
exit 0
