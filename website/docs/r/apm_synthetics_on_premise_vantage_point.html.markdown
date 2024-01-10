---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_on_premise_vantage_point"
sidebar_current: "docs-oci-resource-apm_synthetics-on_premise_vantage_point"
description: |-
  Provides the On Premise Vantage Point resource in Oracle Cloud Infrastructure Apm Synthetics service
---

# oci_apm_synthetics_on_premise_vantage_point
This resource provides the On Premise Vantage Point resource in Oracle Cloud Infrastructure Apm Synthetics service.

Registers a new On-premise vantage point.


## Example Usage

```hcl
resource "oci_apm_synthetics_on_premise_vantage_point" "test_on_premise_vantage_point" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	name = var.on_premise_vantage_point_name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.on_premise_vantage_point_description
	freeform_tags = {"bar-key"= "value"}
	type = var.on_premise_vantage_point_type
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) (Updatable) The APM domain ID the request is intended for. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A short description about the On-premise vantage point.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `name` - (Required) Unique On-premise vantage point name that cannot be edited. The name should not contain any confidential information.
* `type` - (Optional) Type of On-premise vantage point.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the On Premise Vantage Point
	* `update` - (Defaults to 20 minutes), when updating the On Premise Vantage Point
	* `delete` - (Defaults to 20 minutes), when destroying the On Premise Vantage Point


## Import

OnPremiseVantagePoints can be imported using the `id`, e.g.

```
$ terraform import oci_apm_synthetics_on_premise_vantage_point.test_on_premise_vantage_point "onPremiseVantagePoints/{onPremiseVantagePointId}/apmDomainId/{apmDomainId}" 
```

