#!/bin/bash
#yum update -y
#######################################################################################################################################################
### This bootstrap script runs on glusterfs clients and does the following
### 1- install gluster packages on all nodes
### 2- disable local firewall. Feel free to update this script to open only the required ports.
### 3- creates a local mount point in the glusterFS clients that maps to the glusterFS servers
### 4- add a local entry in the clients Fstab with '_netdev' attribute. IMPORTANT to avoid boot issues.
###
######################################################################################################################################################
exec 2>/dev/null

sed -i '/search/d' /etc/resolv.conf
echo "search baremetal.oraclevcn.com publicsubnetad2.baremetal.oraclevcn.com publicsubnetad1.baremetal.oraclevcn.com publicsubnetad3.baremetal.oraclevcn.com localdomain" >> /etc/resolv.conf
chattr -R +i /etc/resolv.conf
#firewall-cmd --zone=public --add-port=111/tcp --add-port=139/tcp --add-port=445/tcp --add-port=965/tcp --add-port=2049/tcp \
#--add-port=38465-38469/tcp --add-port=631/tcp --add-port=111/udp --add-port=963/udp --add-port=49152-49251/tcp  --permanent
#firewall-cmd --reload
systemctl disable firewalld
systemctl stop firewalld
yum install glusterfs glusterfs-fuse attr -y
mkdir /mnt/glustervol
sleep 4m

glusterfshost=$(hostname -s)
case $glusterfshost in
    "glusterfs-client1")
        echo "glusterfs-server1:/glustervol /mnt/glustervol   glusterfs defaults,_netdev  0 0" >> /etc/fstab
        mount -a
        ;;
    "glusterfs-client2")
        echo "glusterfs-server2:/glustervol /mnt/glustervol   glusterfs defaults,_netdev  0 0" >> /etc/fstab
        mount -a
        ;;
    "glusterfs-client3")
        echo "glusterfs-server3:/glustervol /mnt/glustervol   glusterfs defaults,_netdev  0 0" >> /etc/fstab
        mount -a
        ;;
    *)
        echo "$glusterfshost" >>/tmp/notfound.txt
        ;;
esac
