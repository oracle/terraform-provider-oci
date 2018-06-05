# Usage Guide

## PREREQUISITES

Installation has a dependency on Terraform being installed and configured for the user tenancy.   As such an "env-vars" file is included with this package which contains all the necessary environment variables.  This file should be updated with the appropriate values prior to installation.  To source this file prior to installation, either reference it in your .rc file for your shell or run the following:

        "source env-vars"

## Scaling 

Modify the env-vars file prior to deployment and modify the number of workers to scale your cluster dynamically.

## Password & User Details

Please modify scripts/startup.sh and look for the MAIN CLUSTER CONFIGURATION section - this is where you can input your contact information, and set up the Cloudera Manager credentials prior to deployment.

## Deployment

Deploy using standard terraform commands

        "terraform init && terraform plan && terraform apply

## Post Deployment

All post deployment is automated via a scriped process using Bash and CM API via Python.  Clusters are pre-configured with tunings based around instance type (in the cmx.py script).  Login to the Bastion host after terrafor
m completes, and run the following commands to watch installation progress.  The public IP will output as a result of the Terraform completion:

        "ssh -i ~/.ssh/id_rsa opc@<public_ip_of_bastion>"
        "sudo su -"
        "screen -r"

Cluster provisioning can take up to a half hour.  After SCM setup is complete, progress can also be monitored directly via Cloudera Manager - the URL for this is also output as part of the Terraform completion.

## Security and Post-Deployment Auditing

Note that as part of this deployment, ssh keys are used for root level access to provisioned hosts in order to setup software.  The key used is the same as the OPC user which has super-user access to the hosts by default.   If enhanced security is desired, then the following steps should be taken after the Cluster is up and running:

	1. Remove ssh private keys from Bastion & Utility hosts "rm -f /home/opc/.ssh/id_rsa"
	2. Replace the authorized_keys file in /root/.ssh/ on all hosts with the backup copy "mv /root/.ssh/authorized_keys.bak /root/.ssh/authorized_keys"


