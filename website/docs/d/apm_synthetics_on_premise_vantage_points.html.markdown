---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_on_premise_vantage_points"
sidebar_current: "docs-oci-datasource-apm_synthetics-on_premise_vantage_points"
description: |-
  Provides the list of On Premise Vantage Points in Oracle Cloud Infrastructure APM Availability Monitoring service (aka APM Synthetics Service)
---

# Data Source: oci_apm_synthetics_on_premise_vantage_points
This data source provides the list of On Premise Vantage Points in Oracle Cloud Infrastructure APM Availability Monitoring service (aka APM Synthetics Service).

Returns a list of On-premise vantage points.


## Example Usage

```hcl
data "oci_apm_synthetics_on_premise_vantage_points" "test_on_premise_vantage_points" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id

	#Optional
	display_name = var.on_premise_vantage_point_display_name
	name = var.on_premise_vantage_point_name
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) The APM domain ID the request is intended for. 
* `display_name` - (Optional) A filter to return only the resources that match the entire display name.
* `name` - (Optional) A filter to return only the resources that match the entire name.


## Attributes Reference

The following attributes are exported:

* `on_premise_vantage_point_collection` - The list of on_premise_vantage_point_collection.

### OnPremiseVantagePoint Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description about the On-premise vantage point.
* `display_name` - Unique permanent name of the On-premise vantage point.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the On-premise vantage point.
* `name` - Unique On-premise vantage point name that cannot be edited. The name should not contain any confidential information.
* `time_created` - The time the resource was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-12T22:47:12.613Z` 
* `time_updated` - The time the resource was updated, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2020-02-13T22:47:12.613Z` 
* `type` - Type of On-premise vantage point.
* `workers_summary` - Details of the workers in a specific On-premise vantage point. 
	* `available` - Number of available workers in a specific On-premise vantage point.
	* `available_capabilities` - List of available capabilities in a specific On-premise vantage point.
		* `capability` - Capability of an On-premise vantage point worker.
		* `on_premise_vantage_point_count` - Count of available capability in a specific On-premise vantage point.
	* `disabled` - Number of disabled workers in a specific On-premise vantage point.
	* `min_version` - Minimum version among the workers in a specific On-premise vantage point.
	* `total` - Total number of workers in a specific On-premise vantage point.
	* `used` - Number of occupied workers in a specific On-premise vantage point.

