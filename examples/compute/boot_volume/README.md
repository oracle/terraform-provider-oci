## Manage instances with multiple attached volumes
This example creates a new boot volume from an existing instance

### Using this example
* Update env-vars with the required information. Most examples use the same set of environment variables so you only need to do this once.
* Source env-vars
  * `$ . env-vars`
* Update `variables.tf` with your instance options. 

### Files in the configuration

#### `env-vars`
Is used to export the environmental variables used in the configuration. These are usually authentication related, be sure to exclude this file from your version control system. It's typical to keep this file outside of the configuration.

Before you plan, apply, or destroy the configuration source the file -  
`$ . env-vars`

#### `boot_volume.tf`
Defines the boot volumes that are created from the boot volume of the instance

#### `compute.tf`
Defines the compute resource. This demo connects to the running instance 
so you will need to supply public/private keys to create an ssh connection. 
**NOTE**: do not try to use your api keys, see [this doc](https://docs.us-phoenix-1.oraclecloud.com/Content/Compute/Tasks/managingkeypairs.htm)
for more info on configuring keys.

#### `variables.tf`
Defines the variables used in the configuration

#### `outputs.tf`
Defines the outputs of the configuration

#### `provider.tf`
Specifies and passes authentication details to the OCI TF provider
