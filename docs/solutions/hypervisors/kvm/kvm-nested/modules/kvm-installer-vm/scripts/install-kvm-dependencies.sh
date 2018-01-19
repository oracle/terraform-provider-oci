#!/bin/bash

#Configure a KVM Host with UEK4: https://blogs.oracle.com/hlsu/configure-kvm-host-on-oracle-linux-with-uek
yum -y install qemu-kvm qemu-img virt-manager virt-install libvirt libvirt-python libvirt-client
systemctl restart libvirtd

##Modify grub
GRUBFILE=/etc/default/grub
TMPFILE=`mktemp`

sed -e 's/^\(GRUB_CMDLINE_LINUX=".*\)"/\1 intel_iommu=on kvm-intel.nested=1"/' $GRUBFILE > $TMPFILE

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

grub2-mkconfig -o /boot/grub2/grub.cfg

service tuned start
chkconfig tuned on
tuned-adm list
tuned-adm profile virtual-host &
tuned-adm active
