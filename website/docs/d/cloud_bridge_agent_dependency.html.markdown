---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_agent_dependency"
sidebar_current: "docs-oci-datasource-cloud_bridge-agent_dependency"
description: |-
  Provides details about a specific Agent Dependency in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_agent_dependency
This data source provides details about a specific Agent Dependency resource in Oracle Cloud Infrastructure Cloud Bridge service.

Gets an AgentDependency by identifier.

## Example Usage

```hcl
data "oci_cloud_bridge_agent_dependency" "test_agent_dependency" {
	#Required
	agent_dependency_id = oci_cloud_bridge_agent_dependency.test_agent_dependency.id
}
```

## Argument Reference

The following arguments are supported:

* `agent_dependency_id` - (Required) A unique AgentDependency identifier.


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

