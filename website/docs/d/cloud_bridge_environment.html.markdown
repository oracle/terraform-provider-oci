---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_environment"
sidebar_current: "docs-oci-datasource-cloud_bridge-environment"
description: |-
  Provides details about a specific Environment in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_environment
This data source provides details about a specific Environment resource in Oracle Cloud Infrastructure Cloud Bridge service.

Gets a source environment by identifier.

## Example Usage

```hcl
data "oci_cloud_bridge_environment" "test_environment" {
	#Required
	environment_id = oci_cloud_bridge_environment.test_environment.id
}
```

## Argument Reference

The following arguments are supported:

* `environment_id` - (Required) Unique environment identifier.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - Environment identifier, which can be renamed.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in Failed state.
* `state` - The current state of the source environment.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the source environment was created. An RFC3339 formatted datetime string.
* `time_updated` - The time when the source environment was updated. An RFC3339 formatted datetime string.

