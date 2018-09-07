#!/bin/bash
## Primary CDH Cluster setup script - invokes python script to setup cluster via CMS API

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
        if [ $memtotal -lt "260" ]; then
                VMSIZE="BM.Standard1.36"
        else
                VMSIZE="BM.DenseIO1.36"
        fi
elif [ $wprocs = "48" ]; then
        VMSIZE="VM.Standard2.24"
elif [ $wprocs = "32" ]; then
        if [ $memtotal -lt "115" ]; then
                VMSIZE="VM.Standard1.16"
        else
                VMSIZE="VM.Standard2.16"
        fi
elif [ $wprocs = "16" ]; then
        if [ $memtotal -lt "60" ]; then
                VMSIZE="VM.Standard1.8"
        else
                VMSIZE="VM.Standard2.8"
        fi
fi

if [ -z $VMSIZE ]; then
        echo -e "VMSIZE NULL - EXITING - check memory and cpu values in /tmp and retry"
        exit
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
pip install --upgrade pip
pip install pyopenssl ndg-httpsclient pyasn1
yum install libffi-devel -y
pip install "cm_api<20"
echo "Starting SCM Server..."
service cloudera-scm-server start 
## Scrape hosts file to gather all IPs - this allows for dynamic number of hosts in cluster
for ip in `cat /home/opc/hosts | sed 1d | gawk '{print $1}'`; do
	if [ -z $cluster_host_ip ]; then
		cluster_host_ip="$ip"
	else
		cluster_host_ip="$cluster_host_ip,$ip"	
	fi
done;
## Setup known_hosts entries for all hosts
for host in `cat /home/opc/hosts | gawk '{print $2}'`; do 
	host_ip=`cat /home/opc/hosts | grep -w $host | gawk '{print $1}'`; 
	host_key=`ssh-keyscan -t rsa -H $host 2>&1 | sed 1d | gawk '{print $3}'`; 
	echo -e $host,$host_ip ecdsa-sha2-nistp256 $host_key >> ~/.ssh/known_hosts; 
done;
## Check that SCM is running - the SCM startup takes some time
echo -n "Waiting for SCM server to be available [*"
scm_chk="1"
while [ "$scm_chk" != "0" ]; do
	scm_lsn=`sudo netstat -tlpn | grep 7180`
	scm_chk=`echo -e $?`
	if [ "$scm_chk" = "0" ]; then
		echo -n "*] [OK]"
		echo -e "\n"
	else
		echo -n "*"
		sleep 1
	fi
done;
## Execute Python cluster setup
mkdir -p /log/cloudera
echo -e "Setup ready to execute... Running Cluster Initialization Script... (output will begin shortly)"
python /home/opc/cmx.py -a -n "$ClusterName" -u "$User" -m "$mip" -w "$cluster_host_ip" -c "$cmUser" -s "$cmPassword" -e -r "$EMAILADDRESS" -b "$BUSINESSPHONE" -f "$FIRSTNAME" -t "$LASTNAME" -o "$JOBROLE" -i "$JOBFUNCTION" -y "$COMPANY" -v "$VMSIZE" -k "$ssh_keypath"
