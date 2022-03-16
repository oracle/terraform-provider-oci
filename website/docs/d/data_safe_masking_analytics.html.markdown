---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_analytics"
sidebar_current: "docs-oci-datasource-data_safe-masking_analytics"
description: |-
  Provides the list of Masking Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_masking_analytics
This data source provides the list of Masking Analytics in Oracle Cloud Infrastructure Data Safe service.

Gets consolidated masking analytics data based on the specified query parameters.

## Example Usage

```hcl
data "oci_data_safe_masking_analytics" "test_masking_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.masking_analytic_compartment_id_in_subtree
	group_by = var.masking_analytic_group_by
	masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
	target_id = oci_cloud_guard_target.test_target.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `group_by` - (Optional) Attribute by which the masking analytics data should be grouped.
* `masking_policy_id` - (Optional) A filter to return only the resources that match the specified masking policy OCID.
* `target_id` - (Optional) A filter to return only items related to a specific target OCID.


## Attributes Reference

The following attributes are exported:

* `masking_analytics_collection` - The list of masking_analytics_collection.

### MaskingAnalytic Reference

The following attributes are exported:

* `items` - An array of masking analytics summary objects.
	* `count` - The total count for the aggregation metric.
	* `dimensions` - The scope of analytics data.
		* `policy_id` - The OCID of the masking policy..
		* `target_id` - The OCID of the target database.
	* `metric_name` - The name of the aggregation metric.

