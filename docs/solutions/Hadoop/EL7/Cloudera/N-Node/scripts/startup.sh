#!/bin/bash
## CDH Cluster Startup Script by Zachary Smith
## by Zachary Smith (Zachary.Smith@oracle.com)
## Last Update - March 2018

##
## DO NOT MODIFY THESE VARIABLES
##

User="root"
mip="MASTERIP"
LOG_FILE="/home/opc/cdh_setup.log"
ssh_keypath="/home/opc/.ssh/id_rsa"

wprocs=`cat /tmp/wprocs`
memtotal=`cat /tmp/memtotal`
if [ $wprocs = "104" ]; then
	VMSIZE="BM.DenseIO2.52"
elif [ $wprocs = "72" ]; then 
	if [ $memtotal = "251" ]; then 
		VMSIZE="BM.Standard1.36"
	elif [ $memtotal = "503" ]; then 
		VMSIZE="BM.DenseIO1.36"
	fi
elif [ $wprocs = "48" ]; then 
	VMSIZE="VM.Standard2.24"
elif [ $wprocs = "32" ]; then 
	if [ $memtotal = "109" ]; then 
		VMSIZE="VM.Standard1.16"
	elif [ $memtotal = "240" ]; then 
		VMSIZE="VM.Standard2.16"
	fi
elif [ $wprocs = "16" ]; then 
	if [ $memtotal = "54" ]; then 
		VMSIZE="VM.Standard1.8"
	elif [ $memtotal = "117" ]; then 
		VMSIZE="VM.Standard2.8"
	fi
fi

if [ -z $VMSIZE ]; then
	echo -e "VMSIZE NULL - EXITING"
fi

##
## MAIN CLUSTER CONFIGURATION - MODIFY THESE VARIABLES PRIOR TO INSTALLATION
##

ClusterName="TestCluster"
cmUser="cdhadmin"
cmPassword="somepassword"
EMAILADDRESS="someguy@oracle.com"
BUSINESSPHONE="8675309"
FIRSTNAME="Big"
LASTNAME="Data"
JOBROLE="root"
JOBFUNCTION="root"
COMPANY="Oracle"

## 
## END CONFIGURATION
##

## MAIN
echo "Installing Postgres, Python, Paramiko..."
yum install postgresql-server python-pip python-paramiko.noarch -y
echo "Configuring Postgres Database..."
bash /home/opc/install-postgresql.sh >> /var/log/postgresql_cdh_setup.log
echo "Installing CM API via PIP plus dependencies..."
pip install pyopenssl ndg-httpsclient pyasn1
yum install libffi-devel -y
pip install cm_api
echo "Starting SCM Server..."
service cloudera-scm-server start 
## Scrape hosts file to gather all IPs - this allows for dynamic number of hosts in cluster
for ip in `cat /home/opc/hosts | sed 1d | gawk '{print $1}'`; do
	if [ -z $worker_ip ]; then
		worker_ip="$ip"
	else
		worker_ip="$worker_ip,$ip"	
	fi
done;
## Setup known_hosts entries for all hosts
for host in `cat /home/opc/hosts | gawk '{print $2}'`; do host_ip=`cat /home/opc/hosts | grep -w $host | gawk '{print $1}'`; host_key=`ssh-keyscan -t rsa -H $host 2>&1 | sed 1d | gawk '{print $3}'`; echo -e $host,$host_ip ecdsa-sha2-nistp256 $host_key >> ~/.ssh/known_hosts; done;
## Check that SCM is running - the SCM startup takes some time
echo -n "Waiting for SCM server to be available..."
scm_chk="1"
while [ "$scm_chk" != "0" ]; do
	scm_lsn=`sudo netstat -tlpn | grep 7180`
	scm_chk=`echo -e $?`
	if [ "$scm_chk" = "0" ]; then
		echo -n " [OK]"
		echo -e "\n"
	else
		echo -n "."
		sleep 1
	fi
done;
## Execute Python cluster setup
mkdir -p /log/cloudera
echo -e "Setup ready to execute... Running Cluster Initialization Script... (output will begin shortly)"
python /home/opc/cmx.py -n "$ClusterName" -u "$User" -m "$mip" -w "$worker_ip" -a -c "$cmUser" -s "$cmPassword" -e -r "$EMAILADDRESS" -b "$BUSINESSPHONE" -f "$FIRSTNAME" -t "$LASTNAME" -o "$JOBROLE" -i "$JOBFUNCTION" -y "$COMPANY" -v "$VMSIZE" -k "$ssh_keypath"
