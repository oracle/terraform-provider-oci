---
layout: "oci"
page_title: "Provider: Oracle Cloud Infrastructure"
sidebar_current: "docs-oci-index"
description: |-
  The Oracle Cloud Infrastructure provider is used to interact with the many resources supported by the Oracle Cloud Infrastructure services. The provider needs to be configured with credentials for the Oracle Cloud Account.
---

# Oracle Cloud Infrastructure Provider

The Oracle Cloud Infrastructure provider is used to interact with the many resources supported by the [Oracle Cloud Infrastructure](https://cloud.oracle.com/cloud-infrastructure).

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the Oracle Cloud Infrastructure provider
provider "oci" {
  tenancy_ocid = "${var.tenancy_ocid}"
  user_ocid = "${var.user_ocid}"
  fingerprint = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
  region = "${var.region}"
}
```

## Argument Reference

The following arguments are supported:

* `tenancy_ocid` - Every Oracle Cloud Infrastructure resource has an Oracle-assigned unique ID called an Oracle Cloud Identifier (OCID). You need your tenancy's OCID to use the API. You'll also need it when contacting support.

* `user_ocid` - A user's unique ID see [Managing Users](https://docs.cloud.oracle.com/iaas/Content/Identity/Tasks/managingusers.htm)

* `fingerprint` - You specify the key's fingerprint to indicate which key you're using to sign the request.

* `private_key_path` - The path to the private key file

* `region` - A region is composed of several availability domains. Oracle Cloud Infrastructure resources are either region-specific, such as a virtual cloud network, or availability domain-specific, such as a compute instance.

## Export credentials
Required Keys and OCIDs - https://docs.us-phoenix-1.oraclecloud.com/Content/API/Concepts/apisigningkey.htm

If you primarily work in a single compartment consider exporting that compartment's OCID as well. Remember that the tenancy OCID is also the OCID of the root compartment.

### \*nix
If your TF configurations are limited to a single compartment/user then 
using this `bash_profile` option will work well. For more complex 
environments you may want to maintain multiple sets of environment 
variables. 

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

### Windows
```
setx TF_VAR_tenancy_ocid <value>
setx TF_VAR_user_ocid <value>
setx TF_VAR_compartment_ocid <value>
setx TF_VAR_fingerprint <value>
setx TF_VAR_private_key_path <value>
```
The variables won't be set for the current session, exit the terminal and reopen.

## Enabling Instance Principal Authorization
To enable instance principal authorization, you can set 'auth' attribute to "InstancePrincipal"
in the provider definition as follows ('tenancy_ocid', 'user_ocid', 'fingerprint'
and 'private_key_path' are not necessary):
```
variable "region" {}

provider "oci" {
  auth = "InstancePrincipal"
  region = "${var.region}"
}
```

## Testing

Credentials must be provided via the environment variables as shown above in order to run
acceptance tests.
