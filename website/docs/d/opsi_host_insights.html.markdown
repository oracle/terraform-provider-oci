---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_host_insights"
sidebar_current: "docs-oci-datasource-opsi-host_insights"
description: |-
  Provides the list of Host Insights in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_host_insights
This data source provides the list of Host Insights in Oracle Cloud Infrastructure Opsi service.

Gets a list of host insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.

## Example Usage

```hcl
data "oci_opsi_host_insights" "test_host_insights" {

	#Optional
	compartment_id = var.compartment_id
	host_type = var.host_insight_host_type
	id = var.host_insight_id
	state = var.host_insight_state
	status = var.host_insight_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `host_type` - (Optional) Filter by one or more host types. Possible value is EXTERNAL-HOST. 
* `id` - (Optional) Optional host insight resource [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource. 
* `state` - (Optional) Lifecycle states
* `status` - (Optional) Resource Status


## Attributes Reference

The following attributes are exported:

* `host_insight_summary_collection` - The list of host_insight_summary_collection.

### HostInsight Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `entity_source` - Source of the host entity.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `host_display_name` - The user-friendly name for the host. The name does not have to be unique.
* `host_name` - The host name. The host name is unique amongst the hosts managed by the same management agent.
* `host_type` - Operations Insights internal representation of the host type. Possible value is EXTERNAL-HOST.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the host insight resource.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `processor_count` - Processor count.
* `state` - The current state of the host.
* `status` - Indicates the status of a host insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the the host insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the host insight was updated. An RFC3339 formatted datetime string

