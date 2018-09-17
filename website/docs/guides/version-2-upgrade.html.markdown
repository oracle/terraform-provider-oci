---
layout: "oci"
page_title: " v2 Provider: Oracle Cloud Infrastructure"
sidebar_current: "docs-oci-guide-version-2-upgrade"
description: |-
  The Oracle Cloud Infrastructure provider. Version 2 Upgrade
---
### Terraform OCI Provider Version 2
All Oracle Bare Metal Cloud (OBMC) technologies and services have been renamed to Oracle Cloud Infrastructure (OCI).

The [V2.0.0 release](https://github.com/terraform-providers/terraform-provider-oci/releases/tag/v2.0.0) will require you to make changes to your configuration files and state file.

You can still use previous versions of the provider without modifying the configuration files and state file.

## Migration tool
As part of the release on September 13th, a migration tool will be supplied that will help you update your terraform plans and state files to work with the new version of the provider. The latest version can be found on the [V2.0.2 release page](https://github.com/terraform-providers/terraform-provider-oci/releases/tag/2.0.2) in the oci-tool.zip file. If you would like to review the source and potentially build your own migration tool, the files can be found [here](https://github.com/terraform-providers/terraform-provider-oci/tree/v2.2.0/tools/oci-tool). A readme on using the tool is also in that directory.

The changes in this release affect all our current Terraform Provider users.
You will be able to run this tool against a plan directory and it will:
- Transform all the "baremetal" references in .tf and .tfstate files to the new "oci" provider name.
- Detect and modify provider blocks that do not specify a region to explicitly use "us-phoenix-1".

Alternatively, you can make the changes manually.

## Provider Name Change
The provider name changes from "baremetal" to "oci". You need to update your Terraform configuration files.

```
Provider "oci" {
	  region = "us-ashburn-1"
	  tenancy_ocid = "${var.tenancy_ocid}"
	  user_ocid = "${var.user_ocid}"
	  fingerprint = "${var.fingerprint}"
	  private_key_path = "${var.private_key_path}"
}
```

The provider binary filename also changes, from "terraform-provider-baremetal" to "terraform-provider-oci".

## Installing the Updated Provider
Use the following guidance to install the updated provider on a Linux or Windows computer.
### On *nix
Copy the unpacked provider into the following directory:
`~/.terraform.d/plugins/`

Your ~/.terraformrc file specifies the path to the baremetal provider (only required for v.9.x).
For example:

```
providers {
  baremetal = "~/.terraform.d/plugins/terraform-provider-baremetal"
}
```

Change the path in your /.terraformrc file to:

```
providers {
  oci = "~/.terraform.d/plugins/terraform-provider-oci_v2.0.0"
}
```

Alternatively you can reference both providers at the same time:

```
providers {
  baremetal = "~/.terraform.d/plugins/terraform-provider-baremetal"
  oci = "~/.terraform.d/plugins/terraform-provider-oci_v2.0.0"
}
```


### On Windows
Copy the unpacked provider into the following directory:

`%APPDATA%/terraform.d/plugins/`

Your %APPDATA%/terraform.rc file specifies the path to the baremetal provider (only required for v.9.x). For example:

```
providers {
  baremetal = "%appdata%/terraform.d/plugins/terraform-provider-baremetal"
}
```

Change the path in your /.terraformrc file to:

```
providers {
  oci = "%appdata%/terraform.d/plugins/terraform-provider-oci"
}
```

## Resource and Datasource Names
Resource and datasource names that use a "baremetal" prefix will now use "oci".
For example, resource "baremetal_core_instance" changes to resource "oci_core_instance".

Code example:

`image = "${lookup(data.baremetal_core_images.OLImageOCID.images[0], "id")}"`

changes to


 `image = "${lookup(data.oci_core_images.OLImageOCID.images[0], "id")}"`.
## Making the Changes
The changes to the provider and the resource names will need to be made to both the configuration files, the state file, and the backup state file.

Once the changes have been made, run terraform plan and verify that there will be no new changes to your infrastructure on the next apply.


## Specifying a Region is Mandatory

In addition to the name change, the region parameter in the provider is a required parameter. In previous releases, the region defaulted to "us-phoenix-1" if no region was specified.
This region parameter is used to determine service endpoints.

```
provider “oci" {
	  region = "us-ashburn-1"
	  tenancy_ocid = "${var.tenancy_ocid}"
	  user_ocid = "${var.user_ocid}"
	  fingerprint = "${var.fingerprint}"
	  private_key_path = "${var.private_key_path}"
}
```

## Building the Code from Source
If you want to build the new code from source you will have to make sure that the root directory of the project is "terraform-provider-oci" instead of "terraform-provider-baremetal"

