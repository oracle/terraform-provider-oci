    #     ___  ____     _    ____ _     _____
    #    / _ \|  _ \   / \  / ___| |   | ____|
    #   | | | | |_) | / _ \| |   | |   |  _|
    #   | |_| |  _ < / ___ | |___| |___| |___
    #    \___/|_| \_/_/   \_\____|_____|_____|
***
# Terraform provider for Oracle Bare Metal Cloud Services

[![wercker status](https://app.wercker.com/status/666d2ee10f45dde41189bb03248aadf9/s/master "wercker status")](https://app.wercker.com/project/byKey/666d2ee10f45dde41189bb03248aadf9)

Oracle customers now have access to an enterprise class, developer friendly orchestration tool they can use to manage [Oracle Bare Metal Cloud Service](https://cloud.oracle.com/en_US/bare-metal) resources as well as the [Oracle Compute Cloud](https://github.com/oracle/terraform-provider-compute).

This Terraform provider is OSS, available to all OBMCS customers at no charge.

## Compatibility
The provider is compatible with Terraform .9.\*.

### Coverage
The Terraform provider provides coverage for the entire BMC API, with some minor exceptions.

## Getting started
Be sure to read the FAQ and Writing Terraform configurations for OBMCS in [/docs](https://github.com/oracle/terraform-provider-baremetal/tree/master/docs).

### Download Terraform
Download the appropriate **.9.x binary** for your platform.  
https://www.terraform.io/downloads.html

### Install Terraform
https://www.terraform.io/intro/getting-started/install.html

### Get the Oracle Bare Metal Cloud Terraform provider
https://github.com/oracle/terraform-provider-baremetal/releases

Unpack the provider to an appropriate location then -
#### On \*nix
Create `~/.terraformrc` that specifies the path to the `baremetal` provider.
```
providers {
  baremetal = "<path_to_provider_binary>/terraform-provider-baremetal"
  }
```

#### On Windows
Create `%APPDATA%/terraform.rc` that specifies the path to the `baremetal` provider.
```
providers {
  baremetal = "<path_to_provider_binary>/terraform-provider-baremetal"
  }
```
### Export credentials
Required Keys and OCIDs - https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm

If you primarily work in a single compartment consider exporting that compartment's OCID as well. Remember that the tenancy OCID is also the OCID of the root compartment.

#### \*nix
If your TF configurations are limited to a single compartment/user then using this `bash_profile` option will work well. For more complex environments you may want to maintain multiple sets of environment variables. [See the single-compute example for an example.](https://github.com/oracle/terraform-provider-baremetal/tree/master/docs/examples/compute/single_instance)

In your ~/.bash_profile set these variables
```
export TF_VAR_tenancy_ocid=
export TF_VAR_user_ocid=
export TF_VAR_fingerprint=
export TF_VAR_private_key_path=<fully qualified path>`
```

Once you've set these values open a new terminal or source your profile changes
```
$ source ~/.bash_profile
```

#### Windows
```
setx TF_VAR_tenancy_ocid <value>
setx TF_VAR_user_ocid <value>
setx TF_VAR_fingerprint <value>
setx TF_VAR_private_key_path <value>
```
The variables won't be set for the current session, exit the terminal and reopen.

## Deploy an example configuration
Download the [Single instance example.](https://github.com/oracle/terraform-provider-baremetal/tree/master/docs/examples/compute/single-instance)

Edit it to include the OCID of the compartment you want to create the VCN. Remember that the tenancy OCID is the compartment OCID of your root compartment.

You should always plan, then apply a configuration -
```
$ terraform plan ./single_instance
# Make sure the plan looks right.
$ terraform apply ./single_instance
```
## OBMC resource and datasource details
https://github.com/oracle/terraform-provider-baremetal/tree/master/docs

## Getting help
You can file an issue against the project
https://github.com/oracle/terraform-provider-baremetal/issues

or meet us in the OBMCS forums
https://community.oracle.com/community/cloud_computing/bare-metal

## Known issues

[Github issues](https://github.com/oracle/terraform-provider-baremetal/issues)

## About the provider
This provider was written on behalf of Oracle by [MustWin.](http://mustwin.com/)
