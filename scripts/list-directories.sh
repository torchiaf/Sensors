#!/bin/sh

ROOT=$1
FORMAT=$2
REGEX=$3

RET=""

for dir in $(ls $ROOT); do
    if [ -z "$REGEX" -o $(echo "$dir" | grep "$REGEX") ] ; then
        if [ "$FORMAT" = "json" ] ; then
            RET=$(echo -n "$RET\"$dir\",")
        elif [ "$FORMAT" = "csv" ] ; then
            RET=$(echo -n "$RET$dir,")
        fi
    fi
done

if [ "$FORMAT" = "json" ] ; then
    echo -n "["
    echo -n $RET | sed 's/.$//'
    echo "]"
elif [ "$FORMAT" = "csv" ] ; then
    echo $RET
fi