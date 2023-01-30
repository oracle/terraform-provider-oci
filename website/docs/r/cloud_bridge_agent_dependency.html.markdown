---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_agent_dependency"
sidebar_current: "docs-oci-resource-cloud_bridge-agent_dependency"
description: |-
  Provides the Agent Dependency resource in Oracle Cloud Infrastructure Cloud Bridge service
---

# oci_cloud_bridge_agent_dependency
This resource provides the Agent Dependency resource in Oracle Cloud Infrastructure Cloud Bridge service.

Creates an AgentDependency.


## Example Usage

```hcl
resource "oci_cloud_bridge_agent_dependency" "test_agent_dependency" {
	#Required
	bucket = var.agent_dependency_bucket
	compartment_id = var.compartment_id
	dependency_name = var.agent_dependency_dependency_name
	display_name = var.agent_dependency_display_name
	namespace = var.agent_dependency_namespace
	object = var.agent_dependency_object

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	dependency_version = var.agent_dependency_dependency_version
	description = var.agent_dependency_description
	freeform_tags = {"Department"= "Finance"}
	system_tags = var.agent_dependency_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `bucket` - (Required) (Updatable) Object storage bucket where the dependency is uploaded.
* `compartment_id` - (Required) (Updatable) Compartment identifier.
* `defined_tags` - (Optional) (Updatable) The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `dependency_name` - (Required) (Updatable) Name of the dependency type. This should match the whitelisted enum of dependency names.
* `dependency_version` - (Optional) (Updatable) Version of the Agent dependency.
* `description` - (Optional) (Updatable) Description about the Agent dependency.
* `display_name` - (Required) (Updatable) Display name of the Agent dependency.
* `freeform_tags` - (Optional) (Updatable) The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `namespace` - (Required) (Updatable) Object storage namespace associated with the customer's tenancy.
* `object` - (Required) (Updatable) Name of the dependency object uploaded by the customer.
* `system_tags` - (Optional) (Updatable) The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `bucket` - Object storage bucket where the Agent dependency is uploaded.
* `checksum` - The checksum associated with the dependency object returned by Object Storage.
* `compartment_id` - Compartment identifier
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `dependency_name` - Name of the dependency type. This should match the whitelisted enum of dependency names.
* `dependency_version` - Version of the Agent dependency.
* `description` - Description about the Agent dependency.
* `display_name` - Display name of the Agent dependency.
* `e_tag` - The eTag associated with the dependency object returned by Object Storage.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `namespace` - Object storage namespace associated with the customer's tenancy.
* `object` - Name of the dependency object uploaded by the customer.
* `state` - The current state of AgentDependency.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the AgentDependency was created. An RFC3339 formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Agent Dependency
	* `update` - (Defaults to 20 minutes), when updating the Agent Dependency
	* `delete` - (Defaults to 20 minutes), when destroying the Agent Dependency


## Import

AgentDependencies can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_bridge_agent_dependency.test_agent_dependency "id"
```

