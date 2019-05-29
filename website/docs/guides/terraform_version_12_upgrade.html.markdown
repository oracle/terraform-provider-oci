---
layout: "oci"
page_title: "Provider: Oracle Cloud Infrastructure"
sidebar_current: "docs-oci-guide-terraform_version_12_upgrade"
description: |-
  The Oracle Cloud Infrastructure provider. Upgrading to Terraform v0.12
---

## Upgrading to Terraform v0.12

With the release of Terraform v0.12, Oracle Cloud Infrastructure providers will need to be upgraded and some existing configurations may need to be updated as well.

### Upgrading the Terraform Oracle Cloud Infrastructure Provider
Terraform Oracle Cloud Infrastructure provider versions 3.26.0 and below are not compatible with Terraform v0.12.
Terraform Oracle Cloud Infrastructure provider version 3.27.0 is the earliest version that supports Terraform v0.12.

> **Note**: All Terraform Oracle Cloud Infrastructure provider versions will remain compatible with Terraform v0.11 and v0.10.

The simplest way to begin using the latest v0.12-compatible provider from the Hashicorp Provider registry is to place an explicit version requirement in your provider configuration block, as shown here:

```hcl
provider "oci" {
  version          = ">= 3.27.0"
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}
```

For guidance using additional version configuration options, see [Provider Versions](https://www.terraform.io/docs/configuration/providers.html#provider-versions).

### Upgrading Terraform configurations

Some users may find that no changes are needed to run existing Terraform configurations with v0.12.

For configurations that do need to be upgraded, see [Upgrading Terraform configuration](https://www.terraform.io/upgrade-guides/0-12.html#upgrading-terraform-configuration).

#### Upgrading configurations while preserving v0.11 compatibility

As recommended in [Upgrading Terraform configuration](https://www.terraform.io/upgrade-guides/0-12.html#upgrading-terraform-configuration), 
the simplest way to upgrade your configurations is to use the `terraform 0.12upgrade` command.

However, using the `terraform 0.12upgrade` command to upgrade your configurations will render them incompatible with Terraform v0.11 and earlier.

It may be possible to manually upgrade configurations to work with v0.12 while preserving compatibility with v0.11. 
Not all possible configurations can be made compatible with both v0.11 and v0.12. This method should only be used as a best-effort to preserve compatibility with v0.11. 

Prior to making configuration changes, it is strongly recommended that configuration and state files are backed up using version control or other preferred mechanism. 

The following describes cases where it is possible to convert existing v0.11 usage to be compatible with both v0.11 and v0.12.

#### Attributes vs. Blocks
In v0.11, it was possible to treat attributes and nested blocks interchangeably.
Note how the attribute `metadata` and the nested block `source_details` are both assigned using only braces. 

```hcl
// v0.11 compatible representation
resource "oci_core_instance" "my_instance" {
  metadata {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
    kms_key_id  = "${oci_kms_key.test_key.id}"
  }
}
```

In v0.12, attributes need to be assigned using the `=` operator while blocks need to be assigned using only braces.
Note how the attribute `metadata` is now assigned with `=`. This usage is still compatible with v0.11.

```hcl
// v0.12 and v0.11 compatible representation
resource "oci_core_instance" "my_instance" {
  metadata = {
    ssh_authorized_keys = "${var.ssh_public_key}"
  }

  source_details {
    source_type = "image"
    source_id   = "${var.instance_image_ocid[var.region]}"
    kms_key_id  = "${oci_kms_key.test_key.id}"
  }
```

#### Nested blocks with multiple elements
In v0.11, it was possible to wrap lists of nested blocks inside `[]` like this example.

```hcl
// v0.11 compatible representation
resource "oci_core_virtual_circuit" "virtual_circuit_public" {
  public_prefixes = [
    {
      cidr_block = "${var.virtual_circuit_public_prefixes_cidr_block}"
    },
    {
      cidr_block = "${var.virtual_circuit_public_prefixes_cidr_block2}"
    },
    {
      cidr_block = "${var.virtual_circuit_public_prefixes_cidr_block3}"
    },
  ]
  
  ...
}
```

In v0.12, it is required to specify each nested block individually without wrapping it in `[]`.
This usage is still compatible with v0.11.

```hcl
// v0.11 compatible representation
resource "oci_core_virtual_circuit" "virtual_circuit_public" {
  public_prefixes {
    cidr_block = "${var.virtual_circuit_public_prefixes_cidr_block}"
  }
  public_prefixes {
    cidr_block = "${var.virtual_circuit_public_prefixes_cidr_block2}"
  }
  public_prefixes {
    cidr_block = "${var.virtual_circuit_public_prefixes_cidr_block3}"
  }
  
  ...
}
```

#### Quotes around attribute names
In v0.11, it was possible to put quotation marks `"` around attribute names.
Note how the `min` and `max` attributes have quotation marks around them.

```hcl
// v0.11 compatible representation
resource "oci_core_security_list" "bastion" {
  egress_security_rules {
    destination = "${var.vcn_cidr}"
    protocol    = "${local.tcp_protocol}"

    tcp_options {
      "min" = 22
      "max" = 22
    }
  }
  
  ...
}
```

In v0.12, quotation marks `"` around attribute names are no longer allowed.

```hcl
// v0.11 and v0.12 compatible representation
resource "oci_core_security_list" "bastion" {
  egress_security_rules {
    destination = "${var.vcn_cidr}"
    protocol    = "${local.tcp_protocol}"

    tcp_options {
      min = 22
      max = 22
    }
  }
  
  ...
}
```

#### Variable names starting with non-alphabetical characters
In v0.11, it was possible to specify variable names that begin with non-alphabetical characters.

```hcl
// v0.11 compatible representation
variable "2TB" {
  default = "2048"
}
```

In v0.12, variable names must begin with alphabetical characters.

```hcl
// v0.11 and v0.12 compatible representation
variable "Size2TB" {
  default = "2048"
}
```

#### Computing list index values
In v0.11, division operations often resulted in integer values that could be used as a valid index in a list.

```hcl
// v0.11 compatible representation
instance_id     = "${oci_core_instance.TFInstance.*.id[count.index / var.NumParavirtualizedVolumesPerInstance]}"
```

In v0.12, division operations can result in floating point values that may no longer be valid.
To avoid this situation, use the `floor` interpolation to convert floating point values to an index.

```hcl
// v0.11 and v0.12 compatible representation
instance_id     = "${oci_core_instance.TFInstance.*.id[floor(count.index / var.NumParavirtualizedVolumesPerInstance)]}"
```