    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
## GlusterFS Infrastructure
This configuration creates a glusterfs volume that is replicated across 3 GlusterFS servers.

It creates a VCN with a route table, internet gateway, and security list.
The VCN spans 3 ADs with each AD containing a subnet and 2 instances: a GlusterFS server and a GlusterFS client. 

### Using this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`
* Update `variables.tf` with your instance options.
* Update `./userdata/bootstrap` by replacing instances of `baremetal.oraclevcn.com` with `[Your VCN's DnsLabel].oraclevcn.com`.

### Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`

#### `compute.tf`
Defines the compute resources

#### `networking.tf`
Defines the virtual cloud network resources used in the configuration

#### `variables.tf`
Defines the variables used in the configuration

#### `datasources.tf`
Defines the datasources used in the configuration

#### `provider.tf`
Specifies and passes authentication details to the OCI TF provider

#### `./userdata/bootstrap`
The script gets injected into an instance on launch.
The script configures the glusterfs volumes on each server and sets up the glusterfs clients.