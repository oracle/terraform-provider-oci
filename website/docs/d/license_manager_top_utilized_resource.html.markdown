---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_top_utilized_resource"
sidebar_current: "docs-oci-datasource-license_manager-top_utilized_resource"
description: |-
Provides details about a specific Top Utilized Resource in Oracle Cloud Infrastructure License Manager service
---

# Data Source: oci_license_manager_top_utilized_resource
This data source provides details about a specific Top Utilized Resource resource in Oracle Cloud Infrastructure License Manager service.

Retrieves the top utilized resources for a given compartment.

## Example Usage

```hcl
data "oci_license_manager_top_utilized_resource" "test_top_utilized_resource" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	is_compartment_id_in_subtree = var.top_utilized_resource_is_compartment_id_in_subtree
	resource_unit_type = var.top_utilized_resource_resource_unit_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration.
* `is_compartment_id_in_subtree` - (Optional) Indicates if the given compartment is the root compartment.
* `resource_unit_type` - (Optional) A filter to return only resources whose unit matches the given resource unit.


## Attributes Reference

The following attributes are exported:

* `items` - The top utilized resource summary collection.
    * `resource_compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) that contains the resource.
    * `resource_compartment_name` - The display name of the compartment that contains the resource.
    * `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
    * `resource_name` - Resource canonical name.
    * `total_units` - Number of license units consumed by the resource.
    * `unit_type` - The resource unit.

