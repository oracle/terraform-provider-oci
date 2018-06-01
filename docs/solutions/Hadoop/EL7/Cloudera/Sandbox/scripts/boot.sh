#!/bin/bash
## cloud-init bootstrap script
## set speedup="1" to bypass host reboot - should set selinux to permissive mode allowing for faster deployment
speedup="1"
if [ $speedup = "0" ]; then 
	if [ -f /etc/selinux/config ]; then
		selinuxchk=`sudo cat /etc/selinux/config | grep enforcing`
		selinux_chk=`echo -e $?`
		if [ $selinux_chk = "0" ]; then
			sudo sed -i.bak 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config
			sudo reboot
		fi
	fi
elif [ $speedup = "1" ]; then 
        if [ -f /etc/selinux/config ]; then
                selinuxchk=`sudo cat /etc/selinux/config | grep enforcing`
                selinux_chk=`echo -e $?`
                if [ $selinux_chk = "0" ]; then
                        sudo sed -i.bak 's/SELINUX=enforcing/SELINUX=disabled/g' /etc/selinux/config
			sudo setenforce 0
                fi
        fi
fi

## Custom Boot Volume Extension
sudo yum -y install cloud-utils-growpart
sudo yum -y install gdisk
sudo reboot
