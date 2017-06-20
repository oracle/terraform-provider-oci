    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
## Bare Metal Cloud Instance monitored with Oracle Management Cloud Agent


### Prerequisites

- Subscription to the Oracle Management Cloud Service  -https://cloud.oracle.com/management 
- Download the AgentInstall.zip - https://docs.oracle.com/en/cloud/paas/management-cloud/emcad/deploying-oracle-management-cloud-agents.html
- Generate the registration key for your agent - https://docs.oracle.com/en/cloud/paas/management-cloud/emcad/managing-registration-keys.html
- Instances in your subnet  properly return their FQDN from hostname -f.  This is a requirement for the OMC agent to install properly.
  To ensure this create your subnet with a custom DHCP provider.  The custom DHCP provider should be configured with a search domain that matches 
  "subnet dns label"."vcn dns label".oraclevcn.com
https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingDHCP.htm?Highlight=dhcp%20provider
  
### Optional

- Chef Server - to use example three where Chef is used to configure the instance you need to have access to a Chef server
to upload the provided cookbook for OMC.

### Running this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`

* This example demonstrates three options to create an OMC managed instance.  

    1. Single stage deploy, install and configure the agent.  This approach takes about 12-15 minutes to bring a single server online.
    2. Two stage deploy, where we first stage the agent on an instance and create and image that we can reuse. Staging the agent  allows you to do the time consuming work one time, and create an image to be used to configure the agent on demand.  
    3. Terraform and Chef, where Chef is used to do all the post instance creation configuration of the OMC agent.
 
### 1. Single stage deploy & monitor
* Deploy and configure the agent on an instance in a single step
* `$ cd stage_and_monitor_instance`
* Update example.tfvars with your specific details
    * Region where you want to deploy
    * The compartment name to run your instance
    * The OCID for the subnet where you want your instance to run.
        * It is required that this subnet be public and allow SSH traffic
    * The path to your AgentInstall.zip
    * Your Agent registration key
* Load the resource module 
* `$ terraform get`
* `$ terraform plan -var-file=example.tfvars`
* `$ terraform apply -var-file=example.tfvars`

### 2. Two stage deploy & monitor
* Create your base server with the agent installed
* `$ cd stage_server`
* Update example.tfvars with your specific details
    * Region where you want to deploy
    * The compartment name to run your instance
    * The OCID for the subnet where you want your instance to run.
        * It is required that this subnet be public and allow SSH traffic
    * The path to your AgentInstall.zip
    * Your Agent registration key
* Load the resource module 
* `$ terraform get`
* `$ terraform plan -var-file=example.tfvars`
* `$ terraform apply -var-file=example.tfvars`
* Once your instance is running, goto the console, and create a custom image from this server.  Note the ocid of the image that is created.
* `$ terraform destroy -var-file=example.tfvars `
* With your custom image created change to the create_omc_managed_server directory to create an instance that is managed by OMC.
* `$ cd ../monitor_instance`
* Update example.tfvars with your specific details
    * Region where you want to deploy
    * The compartment name to run your instance
    * The OCID for the subnet where you want your instance to run.
    * It is required that this subnet be public and allow SSH traffic
    * Your Agent registration key
    * OCID of the image you created in the previous step
* `$ terraform get`
* `$ terraform plan -var-file=example.tfvars`
* `$ terraform apply -var-file=example.tfvars`

    
### 3. Chef stage, deploy & monitor
* Use Terraform with Chef to create and configure a server instance 
* `$ cd chef_stage_and_monitor_instance`
* We first need to upload the provided omc cookbook and users databag to your chef server.  Assuming your knife.rb
configuration file is in your current path. Run the following commands
* `$ cd cookbooks/omc`
* Add your AgentInstall.zip file downloaded from OMC to the files/default directory.
* Review the templates/default/omc_entity.json.erb file to see the entity configuration we will load by default.
* `$ berks install`
* `$ berks upload`
* `$ cd ../../data-bag/users`
* Edit the omc.json file providing a ssh public key pair that will be used to ssh to the server as the omc user.
* `$ knife data bag create users`
* `$ knife data bag from file users omc.json`
* Validate on your chef server that the omc cookbook and dependcies have been uploaded
* Validate on your chef server that  the users databag has been created with user omc.
* From the examples root directory
* `$ cd chef_stage_and_monitor_instance`
* Edit the chef_attributes.json file with your OMC Agent Registration key
* There are 3 tf.vars files you can run this example with
    * stage_and_monitor.tfvars 
        * All in one to launch an instance, stage the agent, and monitors a server instance.
    * stage.tfvars
        * Creates a staged server instance, that you will create a custom image from.
    * monitor.tfvars
        * Uses a custom image created in the previous step, to launch a server instance that is monitored by OMC.
* Each of the .tfvars need to be updated with your specific details
    * Region where you want to deploy
    * The compartment name to run your instance
    * The OCID for the subnet where you want your instance to run.
        * It is required that this subnet be public and allow SSH traffic
        * It is required that this subnet have a custom DHCP provider with a proper seach domain.
    * Chef Server configuration details 
        * chef_server - Server URL
        * chef_user - Chef Client Username
        * chef_key - Chef Client private key
        * chef_node_name - Chef node name
    * The run list and json_attributes do not need to be changed
* Load the resource module 
* `$ terraform get`
* `$ terraform plan -var-file=stage_and_monitor.tfvars`  
* `$ terraform apply -var-file=stage_and_monitor.tfvars` 

* Build the base image
* `$ terraform plan -var-file=stage.tfvars`  
* `$ terraform apply -var-file=stage.tfvars` 
* Once up and running via the console create a custom image from this server, and note the OCID of the image.

* Create a monitored server from the base image
* Update the monitor.tfvars with the ocid of your custom staged image
* `$ terraform plan -var-file=monitor.tfvars`  
* `$ terraform apply -var-file=monitor.tfvars` 


## Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`

### Modules/datasources

#### `datasources.tf`
Reads BMC environment specific values for ADs, and Compartments used in the configuration


### Stage Server 


#### `main.tf`
Defines the compute resource

#### `provider.tf`
Specifies and passes authentication details to the OBMCS TF provider

#### `./userdata/*`
The user-data scripts that get injected into an instance on launch. More information on user-data scripts can be [found at the cloud-init project.](https://cloudinit.readthedocs.io/en/latest/topics/format.html)
This script creates an Oracle user, installs the necessary packages to install the agent, and creates the directory structure of the install. /omc path on the host

#### `variables.tf`
Defines the variables used in the configuration

#### `example.tfvars`
Environment specific values passed into configuration.  

### Monitor Instance

#### `main.tf`
Defines the compute resource that is monitored by OMC

#### `provider.tf`
Specifies and passes authentication details to the OBMCS TF provider

#### `./userdata/*`
The user-data scripts that get injected into an instance on launch. More information on user-data scripts can be [found at the cloud-init project.](https://cloudinit.readthedocs.io/en/latest/topics/format.html)
This script creates an Oracle user, installs the necessary packages to install the agent, and creates the directory structure of the install. /omc path on the host

#### `./omc_config/install_omc.tpl`
Installation script template file for the agent, and the monitoring configuration file.  This file is read by Terraform
and passed the registration key value.

#### `./omc_config/omc_entity.json`
OMC agent entity configuration file.  This file is loaded to the server, and passed to omcli utility to configure 
the agent.

#### `variables.tf`
Defines the variables used in the configuration

#### `example.tfvars`
Environment specific values passed into configuration.  

### Stage and Monitor Instance

#### `main.tf`
Defines the compute resource that is monitored by OMC

#### `provider.tf`
Specifies and passes authentication details to the OBMCS TF provider

#### `./omc_config/install_omc.tpl`
Installation script template file for the agent, and the monitoring configuration file.  This file is read by Terraform
and passed the registration key value.

#### `./omc_config/omc_entity.json`
OMC agent entity configuration file.  This file is loaded to the server, and passed to omcli utility to configure 
the agent.

#### `variables.tf`
Defines the variables used in the configuration

#### `example.tfvars`
Environment specific values passed into configuration.  


### Chef Stage and Monitor Instance

#### `main.tf`
Defines the compute resource that is monitored by OMC

#### `provider.tf`
Specifies and passes authentication details to the OBMCS TF provider

#### `variables.tf`
Defines the variables used in the configuration


#### `stage.tfvars`
Environment specific values passed into configuration to create a staged OMC server.  Once this instance is running, create
a custom image via the console.

#### `monitor.tfvars`
Environment specific values passed into configuration that is based on the staged custom image created by the stage.tfvars run.

#### `stage_and_monitor.tfvars`
Environment specific values passed into configuration that performs a full install and configuration of the OMC agent on a new instance.


#### `./cookbooks/omc`
OMC install and configuration cookbook.  See the cookbook readme for more details.

#### `./cookbooks/data-bags/users/omc.json`
OMC user definition that the cookbook uses for OMC install.  Update this file with the ssh public key to be created 
on the server for this user.





