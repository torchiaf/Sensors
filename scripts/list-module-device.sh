#!/bin/sh

REGEX=$1

RET="["

for module in $(ls modules); do
    if [ -z "$REGEX" -o $(echo "$module" | grep "$REGEX") ] ; then
        for device in $(ls modules/$module/devices); do
            RET=$(echo -n "$RET{\"module\":\"$module\",\"device\":\"$device\"},")
        done
    fi
done

RET=$(echo -n "$(echo -n $RET | sed 's/.$//')]")

echo $RET
