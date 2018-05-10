#!/bin/bash
## Automated installer for Terraform & OCI
## Written by Zachary Smith 
## Zachary.Smith@oracle.com
## Last Edit - March 2018 
## Author's Note : This script is intended to be run inside an Oracle Linux VM as the OCI user, in order to deploy Terraform.  
##
############################
#### Function Section ######
############################

SETUP () {
if [ ! -f /home/opc/.tfconfig ]; then
        if [ -f /home/opc/.ssh/id_rsa ]; then
                sleep .001
        else
                echo -e "\e[0;31m/home/opc/.ssh/id_rsa missing\e[0m - generating now."
		ssh-keygen -b 2048 -t rsa -f /home/opc/.ssh/id_rsa -q -N ""
        fi
        if [ -f /home/opc/.ssh/id_rsa.pub ]; then
                sleep .001
        else
                echo -e "/home/opc/.ssh/id_rsa.pub missing.  Creating now...\n"
                ssh-keygen -y -f /home/opc/.ssh/id_rsa > /home/opc/.ssh/id_rsa.pub
		
        fi
        if [ -f /home/opc/.ssh/oci_api_key.pem ]; then
                sleep .001
        else
                echo -e "/home/opc/.ssh/oci_api_key.pem not detected.  Please paste the Object Reference URL where this file can be retrieved:"
                read key_url
        fi
        if [ ! -z $key_url ]; then
                echo -e "Proceeding with initialization.  Downloading and deploying API key file."
                        cd /home/opc/.ssh/
                        wget $key_url
                        chmod 0600 oci_api_key.pem
        fi
        TF_BIN_CHECK
else
        TF_BIN_CHECK
fi
}

TF_BIN_CHECK () {
        if [ -f /opt/pkg/bin/terraform ]; then
                sleep .001
        else
                echo -e "\tTerraform binaries not found in /opt/pkg/bin/.  Installing..."
                cd ~
                if [ -f /home/opc/performance/terraform_0.11.5_freebsd_amd64.zip ]; then
 	               rm -f /home/opc/performance/terraform_0.11.5_freebsd_amd64.zip
                fi
                wget https://releases.hashicorp.com/terraform/0.11.5/terraform_0.11.5_linux_amd64.zip
                unzip terraform_0.11.5_linux_amd64.zip
                sudo mkdir -p /opt/pkg/bin
                sudo chown opc /opt/pkg/bin/
                mv terraform /opt/pkg/bin/
                if [ -f /home/opc/linux.tar.gz ]; then
         	       rm -f /home/opc/linux.tar.gz
                fi
                wget https://github.com/oracle/terraform-provider-oci/releases/download/v2.0.4/linux.tar.gz
                tar -zxvf linux.tar.gz
                echo providers { oci = '"'/home/opc/linux_amd64/terraform-provider-oci_v2.0.4'"'} > ~/.terraformrc
        fi
}

##########
## MAIN ##
##########

SETUP

