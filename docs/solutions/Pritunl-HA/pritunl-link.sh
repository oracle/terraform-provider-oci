#!/bin/bash
sudo tee -a /etc/yum.repos.d/pritunl.repo << EOF
[pritunl]
name=Pritunl Repository
baseurl=https://repo.pritunl.com/stable/yum/centos/7/
gpgcheck=1
enabled=1
EOF
echo -e "---INSTALLATION START---"
gpg --keyserver hkp://keyserver.ubuntu.com --recv-keys 7568D9BB55FF9E5287D586017AE645C0CF8E292A
gpg --armor --export 7568D9BB55FF9E5287D586017AE645C0CF8E292A > key.tmp; sudo rpm --import key.tmp; rm -f key.tmp
sudo yum install pritunl-link -y 
sudo systemctl enable pritunl-link
sudo systemctl start pritunl-link
sudo pritunl-link verify-off
sudo systemctl stop firewalld
sudo systemctl disable firewalld
sudo sed -i.bak "s/SELINUX=enforcing/SELINUX=disabled/g" /etc/selinux/config
echo -e "---REBOOTING SERVER TO DISABLE SELINUX NOW---"
sudo reboot
