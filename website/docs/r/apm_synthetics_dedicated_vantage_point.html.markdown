---
subcategory: "Apm Synthetics"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_apm_synthetics_dedicated_vantage_point"
sidebar_current: "docs-oci-resource-apm_synthetics-dedicated_vantage_point"
description: |-
  Provides the Dedicated Vantage Point resource in Oracle Cloud Infrastructure Apm Synthetics service
---

# oci_apm_synthetics_dedicated_vantage_point
This resource provides the Dedicated Vantage Point resource in Oracle Cloud Infrastructure Apm Synthetics service.

Registers a new dedicated vantage point.


## Example Usage

```hcl
resource "oci_apm_synthetics_dedicated_vantage_point" "test_dedicated_vantage_point" {
	#Required
	apm_domain_id = oci_apm_apm_domain.test_apm_domain.id
	display_name = var.dedicated_vantage_point_display_name
	dvp_stack_details {
		#Required
		dvp_stack_id = oci_resourcemanager_stack.test_stack.id
		dvp_stack_type = var.dedicated_vantage_point_dvp_stack_details_dvp_stack_type
		dvp_stream_id = oci_streaming_stream.test_stream.id
		dvp_version = var.dedicated_vantage_point_dvp_stack_details_dvp_version
	}
	region = var.dedicated_vantage_point_region

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	status = var.dedicated_vantage_point_status
}
```

## Argument Reference

The following arguments are supported:

* `apm_domain_id` - (Required) (Updatable) The APM domain ID the request is intended for. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) Unique dedicated vantage point name that cannot be edited. The name should not contain any confidential information.
* `dvp_stack_details` - (Required) (Updatable) Details of a Dedicated Vantage Point (DVP) stack in Resource Manager.
	* `dvp_stack_id` - (Required) (Updatable) Stack [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Resource Manager stack for dedicated vantage point.
	* `dvp_stack_type` - (Required) (Updatable) Type of stack.
	* `dvp_stream_id` - (Required) (Updatable) Stream [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Resource Manager stack for dedicated vantage point.
	* `dvp_version` - (Required) (Updatable) Version of the dedicated vantage point.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `region` - (Required) (Updatable) Name of the region.
* `status` - (Optional) (Updatable) Status of the dedicated vantage point.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dedicated Vantage Point
	* `update` - (Defaults to 20 minutes), when updating the Dedicated Vantage Point
	* `delete` - (Defaults to 20 minutes), when destroying the Dedicated Vantage Point


## Import

DedicatedVantagePoints can be imported using the `id`, e.g.

```
$ terraform import oci_apm_synthetics_dedicated_vantage_point.test_dedicated_vantage_point "dedicatedVantagePoints/{dedicatedVantagePointId}/apmDomainId/{apmDomainId}" 
```

