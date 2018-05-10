#!/bin/bash
sudo firewall-cmd --zone=public --add-port=14643/udp
sudo firewall-cmd --zone=public --add-service=https 
sudo firewall-cmd --runtime-to-permanent
sudo tee -a /etc/yum.repos.d/pritunl.repo << EOF
[pritunl]
name=Pritunl Repository
baseurl=https://repo.pritunl.com/stable/yum/centos/7/
gpgcheck=1
enabled=1
EOF
sudo rpm -Uvh https://dl.fedoraproject.org/pub/epel/epel-release-latest-7.noarch.rpm
gpg --keyserver hkp://keyserver.ubuntu.com --recv-keys 7568D9BB55FF9E5287D586017AE645C0CF8E292A
gpg --armor --export 7568D9BB55FF9E5287D586017AE645C0CF8E292A > key.tmp; sudo rpm --import key.tmp; rm -f key.tmp
echo -e "---INSTALLATION START---"
sudo yum install pritunl -y
sudo systemctl start pritunl
sudo systemctl enable pritunl
#private_key=`sudo pritunl setup-key`
#echo -e "---Pritunl Private Key for UI Access: $private_key"
sudo sh -c 'echo "* hard nofile 64000" >> /etc/security/limits.conf'
sudo sh -c 'echo "* soft nofile 64000" >> /etc/security/limits.conf'
sudo sh -c 'echo "root hard nofile 64000" >> /etc/security/limits.conf'
sudo sh -c 'echo "root soft nofile 64000" >> /etc/security/limits.conf'
sudo sed -i.bak "s/SELINUX=enforcing/SELINUX=disabled/g" /etc/selinux/config
echo -e "---REBOOTING SERVER TO DISABLE SELINUX NOW---"
sudo reboot
