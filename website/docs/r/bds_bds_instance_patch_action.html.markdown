---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_patch_action"
sidebar_current: "docs-oci-resource-bds-bds_instance_patch_action"
description: |-
  Provides the Bds Instance Patch Action resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_patch_action
This resource provides the Bds Instance Patch Action resource in Oracle Cloud Infrastructure Big Data Service service.

Install the specified patch to this cluster.


## Example Usage

```hcl
resource "oci_bds_bds_instance_patch_action" "test_bds_instance_patch_action" {
	#Required
	bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
	cluster_admin_password = var.bds_instance_patch_action_cluster_admin_password
	version = var.bds_instance_patch_action_version
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) Base-64 encoded password for the cluster admin user.
* `version` - (Required) The version of the patch to be installed.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 60 minutes), when creating the Bds Instance Patch Action
	* `update` - (Defaults to 20 minutes), when updating the Bds Instance Patch Action
	* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Patch Action


## Import

Import is not supported for this resource.

