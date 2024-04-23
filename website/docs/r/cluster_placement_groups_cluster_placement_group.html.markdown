---
subcategory: "Cluster Placement Groups"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cluster_placement_groups_cluster_placement_group"
sidebar_current: "docs-oci-resource-cluster_placement_groups-cluster_placement_group"
description: |-
  Provides the Cluster Placement Group resource in Oracle Cloud Infrastructure Cluster Placement Groups service
---

# oci_cluster_placement_groups_cluster_placement_group
This resource provides the Cluster Placement Group resource in Oracle Cloud Infrastructure Cluster Placement Groups service.

Creates a new cluster placement group in the specified compartment.


## Example Usage

```hcl
resource "oci_cluster_placement_groups_cluster_placement_group" "test_cluster_placement_group" {
	#Required
	availability_domain = var.cluster_placement_group_availability_domain
	cluster_placement_group_type = var.cluster_placement_group_cluster_placement_group_type
	compartment_id = var.compartment_id
	description = var.cluster_placement_group_description
	name = var.cluster_placement_group_name

	#Optional
	capabilities {
		#Required
		items {
			#Required
			name = var.cluster_placement_group_capabilities_items_name
			service = var.cluster_placement_group_capabilities_items_service
		}
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	opc_dry_run = var.cluster_placement_group_opc_dry_run
	placement_instruction {
		#Required
		type = var.cluster_placement_group_placement_instruction_type
		value = var.cluster_placement_group_placement_instruction_value
	}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain where you want to create the cluster placement group.
* `capabilities` - (Optional) A list of resources that you can create in a cluster placement group. 
	* `items` - (Required) The supported resources.
		* `name` - (Required) The type of resource.
		* `service` - (Required) The service that the resource is part of.
* `cluster_placement_group_type` - (Required) ClusterPlacementGroup Identifier.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want to create the cluster placement group. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Required) (Updatable) A description of the cluster placement group.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) The friendly name of the cluster placement group.
* `opc_dry_run` - (Optional) When set to `true`, the request performs validation on the submitted data without modifying configuration item details. 
* `placement_instruction` - (Optional) Details that inform cluster placement group provisioning.
	* `type` - (Required) The type of placement instruction.
	* `value` - (Required) The value of the token designated for placement of the cluster placement group upon creation.
* `state` - (Optional) (Updatable) The target state for the Cluster Placement Group. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the cluster placement group.
* `capabilities` - A list of resources that you can create in a cluster placement group. 
	* `items` - The supported resources.
		* `name` - The type of resource.
		* `service` - The service that the resource is part of.
* `cluster_placement_group_type` - The type of cluster placement group.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the cluster placement group. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A description of the cluster placement group.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group. 
* `lifecycle_details` - A message describing the current state in more detail. For example, lifecycle details for a resource in a Failed state might include information to act on. 
* `name` - The user-friendly name of the cluster placement group. The display name for a cluster placement must be unique and you cannot change it. Avoid entering confidential information. 
* `placement_instruction` - Details that inform cluster placement group provisioning.
	* `type` - The type of placement instruction.
	* `value` - The value of the token designated for placement of the cluster placement group upon creation.
* `state` - The current state of the ClusterPlacementGroup.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the cluster placement group was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.
* `time_updated` - The time the cluster placement group was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cluster Placement Group
	* `update` - (Defaults to 20 minutes), when updating the Cluster Placement Group
	* `delete` - (Defaults to 20 minutes), when destroying the Cluster Placement Group


## Import

ClusterPlacementGroups can be imported using the `id`, e.g.

```
$ terraform import oci_cluster_placement_groups_cluster_placement_group.test_cluster_placement_group "id"
```

