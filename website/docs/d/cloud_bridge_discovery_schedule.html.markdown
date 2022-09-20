---
subcategory: "Cloud Bridge"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_bridge_discovery_schedule"
sidebar_current: "docs-oci-datasource-cloud_bridge-discovery_schedule"
description: |-
  Provides details about a specific Discovery Schedule in Oracle Cloud Infrastructure Cloud Bridge service
---

# Data Source: oci_cloud_bridge_discovery_schedule
This data source provides details about a specific Discovery Schedule resource in Oracle Cloud Infrastructure Cloud Bridge service.

Reads information about the specified discovery schedule.

## Example Usage

```hcl
data "oci_cloud_bridge_discovery_schedule" "test_discovery_schedule" {
	#Required
	discovery_schedule_id = oci_cloud_bridge_discovery_schedule.test_discovery_schedule.id
}
```

## Argument Reference

The following arguments are supported:

* `discovery_schedule_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the discovery schedule.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the discovery schedule exists.
* `defined_tags` - The defined tags associated with this resource, if any. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name for the discovery schedule. Does not have to be unique, and it's mutable. Avoid entering confidential information. 
* `execution_recurrences` - Recurrence specification for the discovery schedule execution.
* `freeform_tags` - The freeform tags associated with this resource, if any. Each tag is a simple key-value pair with no predefined name, type, or namespace/scope. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the discovery schedule.
* `lifecycle_details` - The detailed state of the discovery schedule.
* `state` - Current state of the discovery schedule.
* `system_tags` - The system tags associated with this resource, if any. The system tags are set by Oracle cloud infrastructure services. Each key is predefined and scoped to namespaces. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{orcl-cloud: {free-tier-retain: true}}` 
* `time_created` - The time when the discovery schedule was created in RFC3339 format.
* `time_updated` - The time when the discovery schedule was last updated in RFC3339 format.

