---
layout: "oci"
page_title: "v3 Provider: Oracle Cloud Infrastructure"
sidebar_current: "docs-oci-guide-version-3-upgrade"
description: |-
  The Oracle Cloud Infrastructure provider. Version 3 Upgrade
---
### Terraform OCI Provider Version 3

#### New Installation

To use the latest OCI Terraform Provider, version 3, run `terraform init` on the directory that contains the configuration referencing the configuration block, 
`provider "oci" {`.

> **Note**: If you have previously configured this environment to run v1 or v2 OCI Provider versions, you will need to 
employ the steps that follow.


#### Upgrading from v2 

The simplest way to begin using the latest provider from the Hashicorp Provider registry is to place an explicit version requirement in your OCI Provider configuration block, as shown here:

```hcl
provider "oci" {
  version          = ">= 3.0.0"
  region           = "${var.region}"
  tenancy_ocid     = "${var.tenancy_ocid}"
  user_ocid        = "${var.user_ocid}"
  fingerprint      = "${var.fingerprint}"
  private_key_path = "${var.private_key_path}"
}
```

For guidance using additional version configuration options, see [Provider Versions](https://www.terraform.io/docs/configuration/providers.html#provider-versions). 

> **Note**: This approach leaves the previous provider configuration unchanged, and is the better approach if you need the environment to remain compatible with previous versions.

To completely update a previously configured environment, remove the old provider from every location where it was added. While this is typically the `plugins` folder of the user's home directory (`~/.terraform.d/plugins`), it can also be found in the plugins folder of any Terraform configuration directory, for example, 
 
```txt
…/my-plan-folder/.terraform/plugins/darwin_amd64/
```


#### Upgrading from v1

Upgrading OCI Provider from v1 to v3 requires the same environment reconfiguration that is shown for the v2 upgrade. However, you must also change all of the plan and statefile references from `bmc` to `oci`. For detailed instructions, see the [Version 2 Upgrade Guide](https://www.terraform.io/docs/providers/oci/guides/version-2-upgrade.html)
