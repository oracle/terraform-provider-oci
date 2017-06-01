    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
## Bare Metal Cloud Instance managed with Oracle Management Cloud Agent
These two configurations 

It creates a VN

### Using this example
* This example is predicated on your usage of Oracle Management Cloud, and having downloaded the AgentInstall.zip and generating an Agent Registration Key
Details can be found here - <some link>
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`

* Two paths are available to create an OMC managed instance. 1. The reccomended approach is a two stage deploy, where we first load the agent, and download necessary agent files and create and image that we can reuse. This approach allows you to do the time consuming work one time, and create an image to be used to configure the agent on demand.  2. Install and configure the agent in one single step.  This approach takes about 20 minutes to bring a server online.

### Two  Stage Deploy
* Create your base server with the agent installed
* `$ cd create_omc_base_server`
* Update examples.tfvars with your specific details
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
* `$ cd ../create_omc_managed_server`
* Update examples.tfvars with your specific details
    * Region where you want to deploy
    * The compartment name to run your instance
    * The OCID for the subnet where you want your instance to run.
    * It is required that this subnet be public and allow SSH traffic
    * Your Agent registration key
    * OCID of the image you created in the previous step
* `$ terraform get`
* `$ terraform plan -var-file=example.tfvars`
* `$ terraform apply -var-file=example.tfvars`

### Single Stage Deploy
* Deploy and configure the agent on an instance in a single step
* `$ cd create_single_omc_server`
* Update examples.tfvars with your specific details
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
    



## Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`

### Create OMC Base Server 


#### `omc_base_server_image.tf`
Defines the compute resource

#### `provider.tf`
Specifies and passes authentication details to the OBMCS TF provider

#### `./userdata/*`
The user-data scripts that get injected into an instance on launch. More information on user-data scripts can be [found at the cloud-init project.](https://cloudinit.readthedocs.io/en/latest/topics/format.html)
This script creates an Oracle user, installs the necessary packages to install the agent, and creates the directory structure of the install. /omc path on the host
#### `variables.tf`
Defines the variables used in the configuration


### Create OMC Managed Server 

#### `omc_base_managed_instance.tf`
Defines the compute resource

#### `provider.tf`
Specifies and passes authentication details to the OBMCS TF provider

#### `./omc/*`
Installation script for the agent, and the monitoring configuration file.

#### `variables.tf`
Defines the variables used in the configuration


### Modules/datasources

#### `datasources.tf`
Defines the datasources used in the configuration



