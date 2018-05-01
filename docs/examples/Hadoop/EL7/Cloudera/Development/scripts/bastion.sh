#!/bin/bash
## Bastion script to drive all provisioning tasks 
## by Zachary Smith (Zachary.Smith@oracle.com) 
## Last Update - March 2018
##
### Firewall Configuration
## Set this flag to 1 to enable host firewalls, 0 to disable
firewall_on="0"
### Main execution below this point - all tasks are initiated from Bastion host inside screen session ##
cd /home/opc/

## Set DNS to resolve all subnet domains
sudo rm -f /etc/resolv.conf
sudo echo "search public1.cdhvcn.oraclevcn.com public2.cdhvcn.oraclevcn.com public3.cdhvcn.oraclevcn.com private1.cdhvcn.oraclevcn.com private2.cdhvcn.oraclevcn.com private3.cdhvcn.oraclevcn.com bastion1.cdhvcn.oraclevcn.com bastion2.cdhvcn.oraclevcn.com bastion3.cdhvcn.oraclevcn.com" > /etc/resolv.conf
sudo echo "nameserver 169.254.169.254" >> /etc/resolv.conf

## Continue with Main Setup
if [ -f host_list ]; then 
	rm -f host_list;
	rm -f datanodes;
	rm -f hosts;
fi 
domain="cdhvcn.oraclevcn.com"
utilname=`nslookup cdh-utility1 | grep Name | gawk '{print $2}'`
echo "$utilname" >> host_list;
ct=1;
mcount=0;
while [ $ct -lt 10 ]; do
        nslk=`nslookup cdh-master-${ct}`
        ns_ck=`echo -e $?`
        if [ $ns_ck = 0 ]; then
		hname=`nslookup cdh-master-${ct} | grep Name | gawk '{print $2}'`
                echo "$hname" >> host_list;
		mcount=$((mcount+1))
        else
                break
        fi
        ct=$((ct+1));
done;
ct=1;
while [ $ct -le $mcount ]; do 
	if [ -z $MASTER_LIST ]; then
		MASTER_LIST="cdh-master-$ct"
	else
		MASTER_LIST="${MASTER_LIST}|cdh-master-$ct"
	fi
	ct=$((ct+1))
done;
ct=1; 
while [ $ct -lt 1000 ]; do
        nslk=`nslookup cdh-worker-${ct}`
        ns_ck=`echo -e $?`
        if [ $ns_ck = 0 ]; then
		hname=`nslookup cdh-worker-${ct} | grep Name | gawk '{print $2}'`
		echo "$hname" >> host_list;
		echo "$hname" >> datanodes;
        else
                break
        fi
        ct=$((ct+1));
done;
for host in `cat host_list`; do 
	h_ip=`dig +short $host`
	echo -e "$h_ip\t$host" >> hosts
done;
## REFACTOR THE NETWORK LOOKUP FOR MULTI AD SUPPORT - OR JUST WHITELIST KNOWN SUBNET 10.0.0.0/16
unset local_network
if [ -f hosts ]; then
	local_network="10.0.0.0/16"
	#local_network=`cat hosts | grep -w "cdh-worker-1" | gawk '{print $1}' | cut -d '.' -f 1-3`
fi
master_ip=`dig +short $utilname`
sed -i "s/MASTERIP/$master_ip/g" startup.sh
for host in `cat host_list | gawk -F '.' '{print $1}'`; do
        echo -e "\tConfiguring $host for deployment."
        echo -ne "\tWaiting for SSH to be available on $host..."
        host_ip=`cat hosts | grep $host | gawk '{print $1}'`
        ## Wait until SSH is available on host, checking every second."
        while [ "$sshchk" != "0" ]; do
                sshtest=`ssh -o BatchMode=yes -o StrictHostKeyChecking=no -q -i /home/opc/.ssh/id_rsa opc@$host 'uptime 2>&1 2>/dev/null' 2>&1 2>/dev/null`
                sshchk=`echo -e $?`
                sleep 1
                echo -n "."
        done;
        unset sshchk
        echo -ne "[OK]\n\tCopying Setup Scripts...\n"
        ## Copy Setup scripts
        scp -o BatchMode=yes -o StrictHostkeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/hosts opc@$host:~/
        scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/iscsi.sh opc@$host:~/
        scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/node_prep.sh opc@$host:~/
        scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/tune.sh opc@$host:~/
        scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/disk_setup.sh opc@$host:~/
        ## Set Execute Flag on scripts
        ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host 'chmod +x *.sh'
        ## Execute Node Prep
        ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host 'sudo ./node_prep.sh &'
        ## Firewall Setup
        if [ $firewall_on = "1" ]; then
                if [ -z "$local_network" ]; then
                        echo -ne "\tSetting up Firewall Ports [ "
                        for dport in `cat /home/opc/firewall.list`; do
                                echo -n "$dport "
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=$dport/tcp"
                        done
                        echo -n "] - DONE"
                        echo -e "\n"
                        if [ $host = "cdh-utility1" ]; then
                                echo -e "\tSetting up Firewall Range 3181-4181 for ZK on $host"
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=3181-4181/tcp"
                                ## Ports for Cloudera Manager UI
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=7180/tcp"
                        fi
                        if [ $host = "cdh-master-2" ]; then
                                echo -e "\tSetting up Firewall Range 3181-4181 for ZK on $host"
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=3181-4181/tcp"
                        fi
                        if [ $host = "cdh-master-3" ]; then
                                echo -e "\tSetting up Firewall Range 3181-4181 for ZK on $host"
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=3181-4181/tcp"
                        fi
                        ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --runtime-to-permanent"
                else
                        if [ $host = "cdh-utility1" ]; then
                                echo -e "\tSetting up Firewall port 7180 for CDH Manager UI access on $host"
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=7180/tcp"
                                echo -e "\tSetting up Firewall port 19888 for Job History Server access on $host"
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=19888/tcp"
                        fi
                        if [[ $host =~ [${MASTER_LIST}] ]]; then
                                echo -e "\tSetting up Firewall port 19888 for Job History Server access on $host"
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=19888/tcp"
                                echo -e "\tSetting up Firewall port 8088 for Resource Manager access on $host"
                                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-port=8088/tcp" 
			fi
                        echo -e "\tAdding whitelist for network ${local_network} to local firewall on $host."
                        ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --zone=public --add-source=${local_network}"
                        ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo firewall-cmd --runtime-to-permanent"
                fi
        else
                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo systemctl stop firewalld"
                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "sudo systemctl disable firewalld"
	fi
        ## Master Setup Files get copied here
        if [ $host = "cdh-utility1" ]; then
                echo -e "\tCopying Master Setup Files..."
                scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/install-postgresql.sh opc@$host:~/
                scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/startup.sh opc@$host:~/
                scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/cms_install.sh opc@$host:~/
                scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/cmx.py opc@$host:~/
                scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/.ssh/id_rsa opc@$host:~/.ssh/
		scp -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa /home/opc/datanodes opc@$host:/tmp/datanodes
                ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@$host "chmod 0600 .ssh/id_rsa"
	fi
        echo -e "\tDone initializing $host.\n\n"
done;
## End Worker Node Setup
echo -e "Checking Resources on Worker Node..."
wprocs=`ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@cdh-worker-1 'cat /proc/cpuinfo | grep processor | wc -l'`
echo -e "$wprocs processors detected.."
ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@cdh-worker-1 "free -hg | grep Mem" > /tmp/meminfo
memtotal=`cat /tmp/meminfo | gawk '{print $2}' | cut -d 'G' -f 1`
echo -e "${memtotal}GB of RAM detected..."
ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@cdh-utility1 "echo $wprocs > /tmp/wprocs"
ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@cdh-utility1 "echo $memtotal > /tmp/memtotal"
## Finish Cluster Setup Below
echo -e "Install Complete..."
echo -n "Waiting for SSH to be ready as root on cdh-utility1..."
count=0
while [ $count -le "30" ]; do 
	sleep 1
	count=$((count+1))
	echo -n "."
done;
        while [ "$sshchk" != "0" ]; do
                sshtest=`ssh -o BatchMode=yes -o StrictHostKeyChecking=no -q -i /home/opc/.ssh/id_rsa opc@cdh-utility1 'uptime 2>&1 2>/dev/null' 2>&1 2>/dev/null`
                sshchk=`echo -e $?`
                sleep 1
                echo -n "."
        done;
        unset sshchk
echo -e "\n"
echo -e "Running CDH Manager Setup..."
## Invoke CMS installer
install_success="1"
while [ $install_success = "1" ]; do
	ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@cdh-utility1 "sudo /home/opc/cms_install.sh"
	install_success=`echo -e $?`
	sleep 10
done
echo -e "CDH Manager Setup Complete... Starting CDH provisioning via SCM..."
## Invoke SCM bootstrapping and initialization 
ssh -o BatchMode=yes -o StrictHostKeyChecking=no -i /home/opc/.ssh/id_rsa opc@cdh-utility1 "sudo /home/opc/startup.sh"
echo -e "---------------------CLUSTER SETUP COMPLETE-------------------------"

