#!/bin/bash
## cloud-init bootstrap script
## Stop SSHD to prevent remote execution during this process
systemctl stop sshd
if [ -f /etc/selinux/config ]; then 
	selinuxchk=`sudo cat /etc/selinux/config | grep enforcing`
	selinux_chk=`echo -e $?`
	if [ $selinux_chk = "0" ]; then
		sudo sed -i.bak 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config
	fi
fi

## NAT SETUP for Private Network 
echo "net.ipv4.ip_forward = 1" >> /etc/sysctl.d/98-ip-forward.conf
firewall-offline-cmd --direct --add-rule ipv4 nat POSTROUTING 0 -o ens3 -j MASQUERADE
firewall-offline-cmd --direct --add-rule ipv4 filter FORWARD 0 -i ens3 -j ACCEPT
/bin/systemctl restart firewalld
sysctl -p /etc/sysctl.d/98-ip-forward.conf

## Custom Boot Volume Extension
sudo yum -y install cloud-utils-growpart screen.x86_64
sudo yum -y install gdisk
sudo reboot
