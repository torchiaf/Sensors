#!/bin/sh

MODULE=$1
DEVICES=$(echo $2 | sed 's/,/ /g')

GITHUB_REGISTRY=$3
GITHUB_REPOSITORY_OWNER=$4
GITHUB_REPO=$5
GITHUB_REF_NAME=$6

OUTPUT=$7

echo "# Module: $MODULE, Devices: $2" > $OUTPUT

for device in $DEVICES; do
    echo "FROM $GITHUB_REGISTRY/$GITHUB_REPOSITORY_OWNER/$GITHUB_REPO/device/$device:$GITHUB_REF_NAME-$MODULE AS $device" >> $OUTPUT
done
echo "" >> $OUTPUT

echo "# Module: $MODULE build" >> $OUTPUT
cat modules/$MODULE/build_template >> $OUTPUT

for device in $DEVICES; do
    echo "COPY --from=$device /dist/$device /rpc_server" >> $OUTPUT
done
echo "" >> $OUTPUT

echo "CMD [\"./init\"]" >> $OUTPUT
