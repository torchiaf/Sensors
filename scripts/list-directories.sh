#!/bin/sh

ROOT=$1
RET=""

for dir in $(ls $ROOT); do
    RET=$(echo -n "$RET$dir,")
done

echo $RET