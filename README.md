## NOTICE
**The terraform provider has been renamed, see [this wiki](https://github.com/oracle/terraform-provider-oci/wiki/Oracle-Terraform-Provider-Name-Change) for information on migration steps.**

*Legacy provider documentation (for v1.0.18 and earlier) can be found [here](https://github.com/oracle/terraform-provider-oci/tree/v1.0.18/docs)* 
 

    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Terraform provider for Oracle Cloud Infrastructure

[![wercker status](https://app.wercker.com/status/666d2ee10f45dde41189bb03248aadf9/s/master "wercker status")](https://app.wercker.com/project/byKey/666d2ee10f45dde41189bb03248aadf9)

Oracle customers now have access to an enterprise class, developer friendly orchestration tool they can use to manage [Oracle Cloud Infrastructure](https://cloud.oracle.com/cloud-infrastructure).

This Terraform provider is OSS, available to all OCI customers at no charge.

## Compatibility
The provider is compatible with Terraform v0.10.x.

### Coverage
The Terraform provider provides coverage for the entire OCI API, with some minor exceptions.

## Getting started
Be sure to read the FAQ and Writing Terraform configurations for OCI in [/docs](https://github.com/oracle/terraform-provider-oci/tree/master/docs).

## Installation
**NOTE** Terraform v0.10.x introduces a change to plugin management where 
previous v0.9.x configuration no longer applies. See note below.

### On Oracle Linux 7.x
```
$ sudo yum install -y terraform terraform-provider-oci
```

### Other platforms
#### Download Terraform
Download the appropriate **v0.10.x binary** for your platform.  
https://www.terraform.io/downloads.html

#### Install Terraform
https://www.terraform.io/intro/getting-started/install.html

#### Get the Oracle Cloud Infrastructure Terraform provider
https://github.com/oracle/terraform-provider-oci/releases

Unpack the provider. Terraform v0.10.x introduces a change to plugin 
management where v0.9.x configuration no longer applies. To be compatible 
with both terraform v0.9.x and v0.10.x, do the following depending on your
platform.

##### On \*nix other than Oracle Linux 7.x
Copy the provider to the following location:
```
~/.terraform.d/plugins/
```
###### For terraform v0.9.x only
Create the `~/.terraformrc` file that specifies the path to the 
`oci` provider.
```
providers {
  oci = "~/.terraform.d/plugins/terraform-provider-oci"
}
```

##### On Windows
Copy the provider to the following location:
```
%APPDATA%/terraform.d/plugins/
```
Note: `%APPDATA%` is a system path specific to your Windows version.

###### For terraform v0.9.x only
Create `%APPDATA%/terraform.rc` that specifies the path to the 
`oci` provider.
```
providers {
  oci = "%appdata%/terraform.d/plugins/terraform-provider-oci"
}
```

### Export credentials
Required Keys and OCIDs - https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm

If you primarily work in a single compartment consider exporting that compartment's OCID as well. Remember that the tenancy OCID is also the OCID of the root compartment.

#### \*nix
If your TF configurations are limited to a single compartment/user then 
using this `bash_profile` option will work well. For more complex 
environments you may want to maintain multiple sets of environment 
variables. 
See the [compute single instance example](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/compute/instance) for more info.

In your `~/.bash_profile` set these variables
```
export TF_VAR_tenancy_ocid=
export TF_VAR_user_ocid=
export TF_VAR_compartment_ocid=<The tenancy OCID can be used as the compartment OCID of your root compartment>
export TF_VAR_fingerprint=
export TF_VAR_private_key_path=<fully qualified path>
```

Once you've set these values open a new terminal or source your profile changes
```
$ source ~/.bash_profile
```

#### Windows
```
setx TF_VAR_tenancy_ocid <value>
setx TF_VAR_user_ocid <value>
setx TF_VAR_compartment_ocid <value>
setx TF_VAR_fingerprint <value>
setx TF_VAR_private_key_path <value>
```
The variables won't be set for the current session, exit the terminal and reopen.

## Deploy an example configuration
Download the [virtual cloud network example](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/networking/vcn).

You should always plan, then apply a configuration -
```
# From the vcn directory

# Initialize the plugin for this template directory
$ terraform init

# Run the plan command to see what will happen.
$ terraform plan
  
# If the plan looks right, apply it.
$ terraform apply

# If you are done with this infrastructure, take it down
$ terraform destroy
```

## OCI resource and datasource details
https://github.com/oracle/terraform-provider-oci/tree/master/docs

## Getting help
You can file an issue against the project
https://github.com/oracle/terraform-provider-oci/issues

or meet us in the OCI forums
https://community.oracle.com/community/cloud_computing/bare-metal

## Known issues

[Github issues](https://github.com/oracle/terraform-provider-oci/issues)

