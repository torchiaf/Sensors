#!/bin/sh

DIRECTORY=$1
OUTPUT=$2

RET=""

for dir in $(ls $DIRECTORY); do
    RET=$(echo -n "$RET\"$dir\",")
done

echo -n "$OUTPUT=["
echo -n $RET | sed 's/.$//'
echo "]"