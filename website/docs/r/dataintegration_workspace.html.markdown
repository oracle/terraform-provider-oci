---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace"
sidebar_current: "docs-oci-resource-dataintegration-workspace"
description: |-
  Provides the Workspace resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace
This resource provides the Workspace resource in Oracle Cloud Infrastructure Data Integration service.

Creates a new Data Integration workspace ready for performing data integration tasks.


## Example Usage

```hcl
resource "oci_dataintegration_workspace" "test_workspace" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.workspace_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.workspace_description
	dns_server_ip = var.workspace_dns_server_ip
	dns_server_zone = var.workspace_dns_server_zone
	freeform_tags = {"Department"= "Finance"}
	is_private_network_enabled = var.workspace_is_private_network_enabled
	subnet_id = oci_core_subnet.test_subnet.id
	vcn_id = oci_core_vcn.test_vcn.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment containing the workspace.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A user defined description for the workspace.
* `display_name` - (Required) (Updatable) A user-friendly display name for the workspace. Does not have to be unique, and can be modified. Avoid entering confidential information.
* `dns_server_ip` - (Optional) The IP of the custom DNS.
* `dns_server_zone` - (Optional) The DNS zone of the custom DNS to use to resolve names.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `is_private_network_enabled` - (Optional) Specifies whether the private network connection is enabled or disabled.
* `subnet_id` - (Optional) The OCID of the subnet for customer connected databases.
* `vcn_id` - (Optional) The OCID of the VCN the subnet is in.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the workspace.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A user defined description for the workspace.
* `display_name` - A user-friendly display name for the workspace. Does not have to be unique, and can be modified. Avoid entering confidential information.
* `dns_server_ip` - The IP of the custom DNS.
* `dns_server_zone` - The DNS zone of the custom DNS to use to resolve names.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - A system-generated and immutable identifier assigned to the workspace upon creation.
* `is_private_network_enabled` - Specifies whether the private network connection is enabled or disabled.
* `state` - Lifecycle states for workspaces in Data Integration Service CREATING - The resource is being created and may not be usable until the entire metadata is defined UPDATING - The resource is being updated and may not be usable until all changes are commited DELETING - The resource is being deleted and might require deep cleanup of children. ACTIVE   - The resource is valid and available for access INACTIVE - The resource might be incomplete in its definition or might have been made unavailable for administrative reasons DELETED  - The resource has been deleted and isn't available FAILED   - The resource is in a failed state due to validation or other errors STARTING - The resource is being started and may not be usable until becomes ACTIVE again STOPPING - The resource is in the process of Stopping and may not be usable until it Stops or fails STOPPED  - The resource is in Stopped state due to stop operation. 
* `state_message` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in failed state.
* `subnet_id` - The OCID of the subnet for customer connected databases.
* `time_created` - The date and time the workspace was created, in the timestamp format defined by RFC3339. 
* `time_updated` - The date and time the workspace was updated, in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `vcn_id` - The OCID of the VCN the subnet is in.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 1 hours), when creating the Workspace
	* `update` - (Defaults to 1 hours), when updating the Workspace
	* `delete` - (Defaults to 1 hours), when destroying the Workspace


## Import

Workspaces can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace.test_workspace "id"
```

