---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_top_utilized_product_license"
sidebar_current: "docs-oci-datasource-license_manager-top_utilized_product_license"
description: |-
  Provides details about a specific Top Utilized Product License in Oracle Cloud Infrastructure License Manager service
---

# Data Source: oci_license_manager_top_utilized_product_license
This data source provides details about a specific Top Utilized Product License resource in Oracle Cloud Infrastructure License Manager service.

Retrieves the top utilized product licenses for a given compartment.

## Example Usage

```hcl
data "oci_license_manager_top_utilized_product_license" "test_top_utilized_product_license" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	is_compartment_id_in_subtree = var.top_utilized_product_license_is_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration.
* `is_compartment_id_in_subtree` - (Optional) Indicates if the given compartment is the root compartment.


## Attributes Reference

The following attributes are exported:

* `items` - Collection of top utilized product licenses.
    * `is_unlimited` - Specifies if the license unit count is unlimited.
    * `product_license_id` - The product license [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
    * `product_type` - The product type.
    * `status` - The current product license status.
    * `total_license_unit_count` - Total number of license units in the product license provided by the user.
    * `total_units_consumed` - Number of license units consumed.
    * `unit_type` - The product license unit.

