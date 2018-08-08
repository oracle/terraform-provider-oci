## NOTICE
**OCI Terraform Provider v2.2.0 and above is not compatible with Terraform binaries below v0.10.1. To use the latest OCI 
Provider, upgrade your version of Terraform to v0.10.1 or higher.** 

    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Terraform Provider for Oracle Cloud Infrastructure

[![wercker status](https://app.wercker.com/status/666d2ee10f45dde41189bb03248aadf9/s/master "wercker status")](https://app.wercker.com/project/byKey/666d2ee10f45dde41189bb03248aadf9)

OCI Terraform Provider gives Oracle customers access to an enterprise class, developer friendly orchestration tool they 
can use to manage [Oracle Cloud Infrastructure](https://cloud.oracle.com/cloud-infrastructure). This Terraform provider 
is open-source software, available to all OCI customers at no charge.

## Compatibility
The OCI Terraform Provider is compatible with Terraform v0.10.1 or greater.

### Coverage
The OCI Terraform Provider supports the entire OCI API, with a few minor exceptions.

## Getting started

To see supported OCI resources and view documentation go to the OCI resource and datasource documentation 
[Table of Contents](https://github.com/oracle/terraform-provider-oci/tree/master/docs/Table%20of%20Contents.md).

Be sure to read the [FAQ](https://github.com/oracle/terraform-provider-oci/tree/master/docs/FAQ.md) 
and [Writing Terraform configurations for OCI](https://github.com/oracle/terraform-provider-oci/tree/master/docs/Writing%20Terraform%20configurations%20for%20OCI.md) document
in the [docs](https://github.com/oracle/terraform-provider-oci/tree/master/docs) section.

## Installation

### On Oracle Linux 7.x
```
$ sudo yum install -y terraform terraform-provider-oci
```

### Other platforms
#### Download Terraform
Download the appropriate **v0.11.x binary** for your platform.  
https://www.terraform.io/downloads.html

#### Install Terraform
https://www.terraform.io/intro/getting-started/install.html

#### Get the Oracle Cloud Infrastructure Terraform provider
https://github.com/oracle/terraform-provider-oci/releases

Unpack the provider.

##### On Mac or other Linux flavors
Copy the provider to the following location:
```
~/.terraform.d/plugins/
```

##### On Windows
Copy the provider to the following location:
```
%APPDATA%/terraform.d/plugins/
```
Note: `%APPDATA%` is a system path specific to your Windows version.


### Setup credentials for using OCI
Every call to OCI infrastructure requires a minimum of four credentials. 
These are `tenancy_ocid`, `user_ocid`, `fingerprint` and `private_key_path`. It is common to export these values as 
environment variables, or source them in different bash profiles when executing Terraform commands. See the next 
section for OS specific instructions on configuring these environment values.

Here is breakdown of required as well as commonly set configuration values:  
- `tenancy_ocid` - The global identifier for your account, always shown on the bottom of the web console. 
- `user_ocid` - The identifier of the user account you will be using Terraform with.
- `private_key_path` - The path to the private key stored on your computer. The public key portion must be added to the 
user account above in the _API Keys_ section of the web console. 
- `fingerprint` - The fingerprint of the public key added in the above user's _API Keys_ section of the web console. 

For details on how to create and configure keys see [Required Keys and OCIDs](https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm).

If you primarily work in a single compartment consider defining a `compartment_ocid` as well. The tenancy OCID is also 
the OCID of the root compartment, so that can be used where resources expect a `compartment_id` or `compartment_ocid`.

#### \*nix
If your Terraform configurations are limited to a single compartment or user then using this `bash_profile` option 
will work well. For more complex environments you may want to maintain multiple sets of environment variables. 
See the [compute single instance example](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/compute/instance) for more info.

In your `~/.bash_profile` set these variables
```
export TF_VAR_tenancy_ocid=<value>
export TF_VAR_compartment_ocid=<value>
export TF_VAR_user_ocid=<value>
export TF_VAR_fingerprint=<value>
export TF_VAR_private_key_path=<value>
```

Once you've set these values open a new terminal or source your profile changes
```
$ source ~/.bash_profile
```

#### Windows
```
setx TF_VAR_tenancy_ocid <value>
setx TF_VAR_compartment_ocid <value>
setx TF_VAR_user_ocid <value>
setx TF_VAR_fingerprint <value>
setx TF_VAR_private_key_path <value>
```
The variables won't be set for the current session, exit the terminal and reopen.

## Deploy an example configuration
Download the [virtual cloud network example](https://github.com/oracle/terraform-provider-oci/tree/master/docs/examples/networking/vcn).

```
# Change to the directory of an example like: 
$ cd doc/examples/networking/vcn

# Initialize the plugin for this template directory
$ terraform init

# Run the plan command to see what will happen
$ terraform plan
  
# If the plan looks right, apply it
$ terraform apply

# If you are done with this infrastructure, take it down
$ terraform destroy
```

## Getting help
If you are having trouble getting the OCI Provider working, check the 
[troubleshooting doc](https://github.com/oracle/terraform-provider-oci/tree/master/docs/Troubleshooting.md)

To see known issues or report unexpected behavior go to the 
[Github issues](https://github.com/oracle/terraform-provider-oci/issues) page.

For questions or information visit the 
[OCI forums](https://cloudcustomerconnect.oracle.com/resources/9c8fa8f96f/summary). 
