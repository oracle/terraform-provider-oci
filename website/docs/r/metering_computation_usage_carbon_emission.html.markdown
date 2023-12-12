---
subcategory: "Metering Computation"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_metering_computation_usage_carbon_emission"
sidebar_current: "docs-oci-resource-metering_computation-usage_carbon_emission"
description: |-
  Provides the Usage Carbon Emission resource in Oracle Cloud Infrastructure Metering Computation service
---

# oci_metering_computation_usage_carbon_emission
This resource provides the Usage Carbon Emission resource in Oracle Cloud Infrastructure Metering Computation service.

Returns usage carbon emission for the given account.


## Example Usage

```hcl
resource "oci_metering_computation_usage_carbon_emission" "test_usage_carbon_emission" {
	#Required
	tenant_id = oci_metering_computation_tenant.test_tenant.id
	time_usage_ended = var.usage_carbon_emission_time_usage_ended
	time_usage_started = var.usage_carbon_emission_time_usage_started

	#Optional
	compartment_depth = var.usage_carbon_emission_compartment_depth
	group_by = var.usage_carbon_emission_group_by
	group_by_tag {

		#Optional
		key = var.usage_carbon_emission_group_by_tag_key
		namespace = var.usage_carbon_emission_group_by_tag_namespace
		value = var.usage_carbon_emission_group_by_tag_value
	}
	is_aggregate_by_time = var.usage_carbon_emission_is_aggregate_by_time
	usage_carbon_emission_filter = var.usage_carbon_emission_usage_carbon_emission_filter
}
```

## Argument Reference

The following arguments are supported:

* `compartment_depth` - (Optional) The compartment depth level.
* `group_by` - (Optional) Aggregate the result by. For example: `["tagNamespace", "tagKey", "tagValue", "service", "skuName", "skuPartNumber", "unit", "compartmentName", "compartmentPath", "compartmentId", "platform", "region", "logicalAd", "resourceId", "resourceName", "tenantId", "tenantName", "subscriptionId"]` 
* `group_by_tag` - (Optional) GroupBy a specific tagKey. Provide the tagNamespace and tagKey in the tag object. Only supports one tag in the list. For example: `[{"namespace":"oracle", "key":"createdBy"]` 
	* `key` - (Optional) The tag key.
	* `namespace` - (Optional) The tag namespace.
	* `value` - (Optional) The tag value.
* `is_aggregate_by_time` - (Optional) Specifies whether aggregated by time. If isAggregateByTime is true, all usage carbon emissions over the query time period will be added up.
* `tenant_id` - (Required) Tenant ID.
* `time_usage_ended` - (Required) The usage end time.
* `time_usage_started` - (Required) The usage start time.
* `usage_carbon_emission_filter` - (Optional) The filter object for query usage.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `group_by` - Aggregate the result by.
* `items` - A list of usage carbon emission items.
	* `ad` - The availability domain of the usage.
	* `compartment_id` - The compartment OCID.
	* `compartment_name` - The compartment name.
	* `compartment_path` - The compartment path, starting from root.
	* `computed_carbon_emission` - The carbon emission in MTCO2 unit.
	* `emission_calculation_method` - The method used to calculate carbon emission.
	* `platform` - Platform for the cost.
	* `region` - The region of the usage.
	* `resource_id` - The resource OCID that is incurring the cost.
	* `resource_name` - The resource name that is incurring the cost.
	* `service` - The service name that is incurring the cost.
	* `sku_name` - The SKU friendly name.
	* `sku_part_number` - The SKU part number.
	* `subscription_id` - The subscription ID.
	* `tags` - For grouping, a tag definition. For filtering, a definition and key.
		* `key` - The tag key.
		* `namespace` - The tag namespace.
		* `value` - The tag value.
	* `tenant_id` - The tenancy OCID.
	* `tenant_name` - The tenancy name.
	* `time_usage_ended` - The usage end time.
	* `time_usage_started` - The usage start time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Usage Carbon Emission
	* `update` - (Defaults to 20 minutes), when updating the Usage Carbon Emission
	* `delete` - (Defaults to 20 minutes), when destroying the Usage Carbon Emission


## Import

UsageCarbonEmissions can be imported using the `id`, e.g.

```
$ terraform import oci_metering_computation_usage_carbon_emission.test_usage_carbon_emission "id"
```

