######################
### README ###########
######################
## February 2018
## 

-PREREQUISITES-

Installation has a dependency on Terraform being installed and configured for the user tenancy.   As such an "env-vars" file is included with this package which contains all the necessary environment variables.  This file should be updated with the appropriate values prior to installation.  To source this file prior to installation, either reference it in your .rc file for your shell or run the following:

        "source env-vars"

-DATABASE SECURITY-

Files which should be modified prior to installation of Pritunl HA include the following:

	mongo_admin.exec
	mongo_clusteradmin.exec
	mongo_pritunl.exec

These files contain administrative passwords for the Mongo Database.  The default for these is either "admin" or "pritunl" and should be modified to enhance security of the configuration data stored in Mongo DB for this implementation.

Modification is as simple as changing the relevant passwords in the .exec file, and then running the Terraform deployment.

-INSTALLATION-

Installation is a straight forward process using Terraform.  This will deploy a 3-node MongoDB cluster with replication, 2 Pritunl Servers, and 1 Pritunl Link host into the configured tenancy.   If you wish to modify the deployment shapes, this can be done by changing the values in variables.tf prior to installation.   

InstanceShape1 is used for the MongoDB hosts, and InstanceShape is used for Pritunl Hosts.

	variable "InstanceShape1" {
	  default = "VM.Standard1.1"
	}

	variable "InstanceShape" {
	  default = "VM.Standard1.2"
	}

Invoke installation using this syntax (assuming terraform binary is referencable):

	"terraform init"
	"terraform plan"
	"terraform apply"

After completion, output will appear with Private and Public IPs of all hosts setup during the installation.  Copy the Public IP of either Pritunl host (NOT the link host) into your browser using the following format:

	"https://<PUBLIC_IP>/"

Because the initial installation uses self signed certificates, there will be security errors in the browser.  Accept these and you will be presented with the Pritunl login page.  The default login is "pritunl" with a password "pritunl".   You will be prompted at first login to change the password, which is stored in the MongoDB and will affect login to both hosts.

-ENTERPRISE LICENSE-

For Site to Site and Multi-node functionality to work, you will need an Enterprise Licenese for Pritunl.  The license is input through the Pritunl UI, and is required prior to any post configuration steps detailed as part of this package.

-DOCUMENTATION-

Post installation configuration will be detailed and distributed as a separate document or wiki link.  

Pritunl post-install configuration is extensively documented on the vendor website: https://docs.pritunl.com/docs
