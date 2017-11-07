#!/bin/bash

echo "Install KVM Dependencies: start" >> /home/opc/install-kvm-dependencies.log
yum -y install qemu-kvm qemu-img virt-manager libvirt libvirt-python libvirtclient virt-install virt-viewer bridge-utils wget jq
wget https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
wget https://fedorapeople.org/groups/virt/virtio-win/virtio-win.repo -O /etc/yum.repos.d/virtio-win.repo
yum -y virtio-win
yum -y --nogpgcheck localinstall epel-release-latest-7.noarch.rpm


echo "Install KVM Dependencies: Modify grub" >> /home/opc/install-kvm-dependencies.log
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

systemctl enable tuned
systemctl start tuned
tuned-adm profile virtual-host

echo "Install KVM Dependencies: Finished" >> /home/opc/install-kvm-dependencies.log
