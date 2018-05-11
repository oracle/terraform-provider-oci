#!/bin/bash
prihost=`hostname | grep 1`
prihostchk=`echo -e $?`
if [ $prihostchk = 1 ]; then 
	sudo pritunl setup-key
else
	sleep 10
fi
TMPDIR=`ls /tmp | grep pritunl`
MONGOPRI=`nslookup mongodb-pri | grep Address | sed 1d | gawk '{print $2}'`
sed -i "s/TMPPATH/\/tmp\/$TMPDIR/g" /home/opc/pritunl.conf
sed -i "s/MONGOPRI/$MONGOPRI/g" /home/opc/pritunl.conf
sudo cp /home/opc/pritunl.conf /etc/pritunl.conf
sudo systemctl restart pritunl
