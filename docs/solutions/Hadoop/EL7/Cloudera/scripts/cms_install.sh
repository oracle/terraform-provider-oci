#!/bin/bash
## Set cdh_version to "5" to use latest version, otherwise custom version can be specified
#cdh_version="5.12.2"
cdh_version="5"
rpm --import http://archive.cloudera.com/cm5/redhat/7/x86_64/cm/RPM-GPG-KEY-cloudera
wget http://archive.cloudera.com/cm5/redhat/7/x86_64/cm/cloudera-manager.repo -O /etc/yum.repos.d/cloudera-manager.repo
if [ $cdh_version = "5" ]; then 
	sleep .001
else
	sed -i "s/cm\/5/cm\/${cdh_version}/g" /etc/yum.repos.d/cloudera-manager.repo
fi
yum install -y oracle-j2sdk* cloudera-manager-daemons cloudera-manager-server
