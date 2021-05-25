---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace"
sidebar_current: "docs-oci-datasource-dataintegration-workspace"
description: |-
  Provides details about a specific Workspace in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace
This data source provides details about a specific Workspace resource in Oracle Cloud Infrastructure Data Integration service.

Retrieves a Data Integration workspace using the specified identifier.

## Example Usage

```hcl
data "oci_dataintegration_workspace" "test_workspace" {
	#Required
	workspace_id = oci_dataintegration_workspace.test_workspace.id
}
```

## Argument Reference

The following arguments are supported:

* `workspace_id` - (Required) The workspace ID.


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

