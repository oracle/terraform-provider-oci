#!/bin/bash
## Installation script for primary master node in cluster
## by Zachary Smith (Zachary.Smith@oracle.com)
## Last Update - March 2018

sudo yum install -y screen.x86_64
## Initiate and execute scripts in screen parallel screen sessions
sudo screen -dmLS cms 
sleep .001
sudo screen -dmLS tune
sleep .001
sudo screen -XS cms stuff logfile /home/opc/`date +%Y%m%d`.log
sleep .001
sudo screen -XS cms stuff login on
sleep .001
sudo screen -XS cms stuff log on
sleep .001
sudo screen -XS cms stuff logfile flush 1
sleep .001
sudo screen -XS tune stuff logfile /home/opc/`date +%Y%m%d`.1.log
sleep .001
sudo screen -XS tune stuff login on
sleep .001
sudo screen -XS tune stuff log on
sleep .001
sudo screen -XS tune stuff logfile flush 1
sleep .001
## Modification here for CentOS 6.9
#sudo screen -XS tune "/home/opc/tune.sh" 
sudo screen -S tune -X stuff '/home/opc/tune.sh\n'
sleep .001
#sudo screen -XS cms "/home/opc/cms_install.sh"
sudo screen -S cms -X stuff '/home/opc/cms_install.sh\n'
