---
subcategory: "Big Data Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_bds_bds_instance_os_patch_action"
sidebar_current: "docs-oci-resource-bds-bds_instance_os_patch_action"
description: |-
  Provides the Bds Instance OS Patch Action resource in Oracle Cloud Infrastructure Big Data Service service
---

# oci_bds_bds_instance_os_patch_action
This resource provides the Bds Instance OS Patch Action resource in Oracle Cloud Infrastructure Big Data Service service.

Install the specified OS patch to this cluster nodes.


## Example Usage

```hcl
resource "oci_bds_bds_instance_os_patch_action" "test_bds_instance_os_patch_action" {
  #Required
  bds_instance_id = oci_bds_bds_instance.test_bds_instance.id
  cluster_admin_password = var.bds_instance_os_patch_action_cluster_admin_password
  os_patch_version = var.bds_instance_os_patch_action_os_patch_version

  #Optional
  patching_configs {
    #Required
    patching_config_strategy = var.bds_instance_os_patch_action_patching_config_strategy

    #Required
    batch_size = var.bds_instance_os_patch_action_batch_size
    wait_time_between_batch_in_seconds = var.bds_instance_os_patch_action_wait_time_between_batch_in_seconds
    tolerance_threshold_per_batch           = var.bds_instance_os_patch_action_tolerance_threshold_per_batch

    wait_time_between_domain_in_seconds = var.bds_instance_os_patch_action_wait_time_between_domain_in_seconds
    tolerance_threshold_per_domain          = var.bds_instance_os_patch_action_tolerance_threshold_per_domain
  }
}
```

## Argument Reference

The following arguments are supported:

* `bds_instance_id` - (Required) The OCID of the cluster.
* `cluster_admin_password` - (Required) Base-64 encoded password for the cluster admin user.
* `patching_configs` - (Optional) Detailed configurations for defining the behavior when installing ODH patches. If not provided, nodes will be patched with down time.
    * `patching_config_strategy` - (Required) Type of strategy used for detailed patching configuration
    * `batch_size` - (Required when patching_config_strategy=BATCHING_BASED) How many nodes to be patched in each iteration.
    * `wait_time_between_batch_in_seconds` - (Required when patching_config_strategy=BATCHING_BASED) The wait time between batches in seconds.
    * `tolerance_threshold_per_batch` - (Required when patching_config_strategy=BATCHING_BASED) Acceptable number of failed-to-be-patched nodes in each batch. The maximum number of failed-to-patch nodes cannot exceed 20% of the number of nodes.
    * `wait_time_between_domain_in_seconds` - (Required when patching_config_strategy=DOMAIN_BASED) The wait time between AD/FD in seconds.
    * `tolerance_threshold_per_domain` - (Required when patching_config_strategy=DOMAIN_BASED) Acceptable number of failed-to-be-patched nodes in each domain. The maximum number of failed-to-patch nodes cannot exceed 20% of the number of nodes.
* `os_patch_version` - (Required) The version of the OS patch to be installed.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 60 minutes), when creating the Bds Instance Patch Action
* `update` - (Defaults to 20 minutes), when updating the Bds Instance Patch Action
* `delete` - (Defaults to 20 minutes), when destroying the Bds Instance Patch Action


## Import

Import is not supported for this resource.
