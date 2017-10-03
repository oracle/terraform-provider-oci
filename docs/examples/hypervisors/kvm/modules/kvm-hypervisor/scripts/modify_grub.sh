#!/bin/bash

GRUBFILE=/etc/default/grub
TMPFILE=`mktemp`

sed -e 's/^\(GRUB_CMDLINE_LINUX=".*\)"/\1 intel_iommu=on"/' $GRUBFILE > $TMPFILE

size=`du -b $GRUBFILE | awk '{print $1}'`
nsize=`du -b $TMPFILE | awk '{print $1}'`

if [[ $nsize -lt $size ]]
then
    echo "Error"
    exit 1
fi

chown --reference=$GRUBFILE $TMPFILE
chmod --reference=$GRUBFILE $TMPFILE

mv $TMPFILE $GRUBFILE
