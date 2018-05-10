#!/bin/bash
sudo tee -a /etc/yum.repos.d/mongo-org-3.6.repo << EOF
[mongodb-org-3.6]
name=MongoDB Repository
baseurl=https://repo.mongodb.org/yum/redhat/7Server/mongodb-org/3.6/x86_64/
gpgcheck=1
enabled=1
gpgkey=https://www.mongodb.org/static/pgp/server-3.6.asc
EOF
echo -e "---INSTALLATION START---"
sudo yum install mongodb-org -y
sudo cp /home/opc/mongod.conf /etc/mongod.conf
sudo chown mongod:mongod /etc/mongod.conf
hostname=`hostname`
local_ip=`nslookup $hostname | grep Address | gawk '{print $2}' | sed 1d`
sudo sed -i "s/127.0.0.1/127.0.0.1,$local_ip/g" /etc/mongod.conf
sudo firewall-cmd --zone=public --add-port=27017/tcp
sudo firewall-cmd --runtime-to-permanent 
sudo mkdir -p /etc/mongo
sudo chown -R mongod:mongod /etc/mongo
sudo mv /home/opc/mongokey /etc/mongo/
sudo chown mongod:mongod /etc/mongo/mongokey
sudo chmod 0600 /etc/mongo/mongokey
sudo systemctl enable mongod
if [ -f "/home/opc/mongo.exec" ]; then 
	ip1=`nslookup mongodb-pri | grep Name | gawk '{print $2}'`
	ip2=`nslookup mongodb-r1 | grep Name | gawk '{print $2}'`
	ip3=`nslookup mongodb-r2 | grep Name | gawk '{print $2}'`
	sed -i "s/IP1/$ip1/g" /home/opc/mongo.exec
	sed -i "s/IP2/$ip2/g" /home/opc/mongo.exec
	sed -i "s/IP3/$ip3/g" /home/opc/mongo.exec
fi
sudo sed -i.bak "s/SELINUX=enforcing/SELINUX=disabled/g" /etc/selinux/config
echo -e "---REBOOTING SERVER TO DISABLE SELINUX NOW---"
sudo reboot

