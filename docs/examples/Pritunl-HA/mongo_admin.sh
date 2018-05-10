#!/bin/bash
sudo yum install screen -y
sleep .001
sudo -u opc /usr/bin/screen -dmLS mongo
sleep .001
sudo -u opc /usr/bin/screen -XS mongo stuff "mongo < /home/opc/mongo.exec\n"
sleep 60 
sudo -u opc /usr/bin/screen -XS mongo stuff "mongo < /home/opc/mongo_admin.exec\n"
sleep 10
sudo -u opc /usr/bin/screen -XS mongo stuff "mongo -u 'admin' -p 'admin' --authenticationDatabase 'admin' < /home/opc/mongo_clusteradmin.exec\n"
sleep 10
sudo -u opc /usr/bin/screen -XS mongo stuff "mongo -u admin -p admin --authenticationDatabase admin < mongo_pritunl.exec\n"
sleep 10
screen -X at 0 hardcopy -h /home/opc/mongo.log
sleep .001
##Debug Step Below
echo -e "--- MONGO REPLICATION SETUP LOG OUTPUT ---\n"
cat /home/opc/mongo.log
