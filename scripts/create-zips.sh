#!/usr/bin/env bash
# This script creates zip files of service examples to add them as release assets.
if [ -d "./examples/zips" ]; then
    rm -r examples/zips
fi
mkdir -p examples/zips
for EXAMPLE_FOLDER in examples/*/;
do
    RESOURCE=$(cut -d "/" -f 2 <<< "$EXAMPLE_FOLDER")
    if [[ "$RESOURCE" != "zips" ]]; then
        FILEZIP=example_${RESOURCE}.zip
        FOLDER=$(echo "$EXAMPLE_FOLDER" | tr -d '"')
        zip -r $FILEZIP $FOLDER || { printf '\n Unable to create zips.\n'; exit 1;  }
        mv $FILEZIP ./examples/zips
    fi
done