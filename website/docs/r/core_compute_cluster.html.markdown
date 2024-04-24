---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_compute_cluster"
sidebar_current: "docs-oci-resource-core-compute_cluster"
description: |-
  Provides the Compute Cluster resource in Oracle Cloud Infrastructure Core service
---

# oci_core_compute_cluster
This resource provides the Compute Cluster resource in Oracle Cloud Infrastructure Core service.

Creates an empty [compute cluster](https://docs.cloud.oracle.com/iaas/Content/Compute/Tasks/compute-clusters.htm). A compute cluster
is a remote direct memory access (RDMA) network group.

After the compute cluster is created, you can use the compute cluster's OCID with the
[LaunchInstance](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Instance/LaunchInstance) operation to create instances in the compute cluster.
The instances must be created in the same compartment and availability domain as the cluster.

Use compute clusters when you want to manage instances in the cluster individually in the RDMA network group.

If you want predictable capacity for a specific number of identical instances that are managed as a group,
create a cluster network that uses instance pools by using the
[CreateClusterNetwork](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ClusterNetwork/CreateClusterNetwork) operation.


## Example Usage

```hcl
resource "oci_core_compute_cluster" "test_compute_cluster" {
	#Required
	availability_domain = var.compute_cluster_availability_domain
	compartment_id = var.compartment_id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.compute_cluster_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain to place the compute cluster in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the compute cluster is running in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the compute cluster.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compute cluster.
* `state` - The current state of the compute cluster.
* `time_created` - The date and time the compute cluster was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Compute Cluster
	* `update` - (Defaults to 20 minutes), when updating the Compute Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Compute Cluster


## Import

ComputeClusters can be imported using the `id`, e.g.

```
$ terraform import oci_core_compute_cluster.test_compute_cluster "id"
```

