# Automation Scripts
All scripts in this location are referenced for deployment automation as part of Development, Production, and N-Node templates.

## start.sh
This is the first script invoked by Terraform in remote-execution. It simply executes the bastion script in a Linux Screen session as superuser (root).

## bastion.sh
This is the primary script which drives discovery and deployment tasks. It is invoked by start.sh and runs on the Bastion host.

## boot.sh
This script is invoked by cloudinit on each instance creation via Terraform.  It contains steps which perform inital bootstrapping of the instance prior to provisioning.

## cms_install.sh
This is the primary script for installing Cloudera Manager.

## install-postgresql.sh
This installs PostGres on the Utility node for use with Cloudera Manager Metadata.

## node_prep.sh
Top level node bootstrapping script, this is called on each node in parallel and executes the following scripts

### iscsi.sh
Detection and setup for Block Storage via iscsi.

### disk_setup.sh
This script is used for disk formatting to use with HDFS. 

### tune.sh
This script is used for OS performance tuning.

## startup.sh
This script drives the Cloudera EDH install after Cloudera Manager has been installed.  It invokes the following script, and should be customized prior to deployment with User details section and password for Cloudera Manager.

### cmx.py
This Python script drives all cluster deployment automation via the Cloudera Manager Python API.
