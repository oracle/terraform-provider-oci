---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_dedicated_vantage_points"
sidebar_current: "docs-oci-datasource-apm_synthetics-dedicated_vantage_points"
description: |-
  Provides the list of Dedicated Vantage Points in Oracle Cloud Infrastructure Apm Synthetics service
---

# Data Source: oci_apm_synthetics_dedicated_vantage_points
This data source provides the list of Dedicated Vantage Points in Oracle Cloud Infrastructure Apm Synthetics service.

Returns a list of dedicated vantage points.


## Example Usage

```hcl
data "oci_apm_synthetics_dedicated_vantage_points" "test_dedicated_vantage_points" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

	#Optional
	display_name = var.dedicated_vantage_point_display_name
	name = var.dedicated_vantage_point_name
	status = var.dedicated_vantage_point_status
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM domain ID the request is intended for. 
* `display_name` - (Optional) A filter to return only the resources that match the entire display name.
* `name` - (Optional) A filter to return only the resources that match the entire name.
* `status` - (Optional) A filter to return only the dedicated vantage points that match a given status.


## Attributes Reference

The following attributes are exported:

* `dedicated_vantage_point_collection` - The list of dedicated_vantage_point_collection.

### DedicatedVantagePoint Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Unique dedicated vantage point name that cannot be edited. The name should not contain any confidential information.
* `dvp_stack_details` - Details of a Dedicated Vantage Point (DVP) stack in Resource Manager.
	* `dvp_stack_id` - Stack [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Resource Manager stack for dedicated vantage point.
	* `dvp_stack_type` - Type of stack.
	* `dvp_stream_id` - Stream [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Resource Manager stack for dedicated vantage point.
	* `dvp_version` - Version of the dedicated vantage point.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated vantage point.
* `monitor_status_count_map` - Details of the monitor count per state. Example: `{ "total" : 5, "enabled" : 3 , "disabled" : 2, "invalid" : 0 }` 
	* `disabled` - Number of disabled monitors using the script.
	* `enabled` - Number of enabled monitors using the script.
	* `invalid` - Number of invalid monitors using the script.
	* `total` - Total number of monitors using the script.
* `name` - Unique permanent name of the dedicated vantage point. This is the same as the displayName.
* `region` - Name of the region.
* `status` - Status of the dedicated vantage point.
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 

