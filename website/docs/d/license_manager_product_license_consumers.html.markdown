---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_product_license_consumers"
sidebar_current: "docs-oci-datasource-license_manager-product_license_consumers"
description: |-
  Provides the list of Product License Consumers in Oracle Cloud Infrastructure License Manager service
---

# Data Source: oci_license_manager_product_license_consumers
This data source provides the list of Product License Consumers in Oracle Cloud Infrastructure License Manager service.

Retrieves the product license consumers for a particular product license ID.

## Example Usage

```hcl
data "oci_license_manager_product_license_consumers" "test_product_license_consumers" {
	#Required
	compartment_id = var.compartment_id
	product_license_id = oci_license_manager_product_license.test_product_license.id

	#Optional
	is_compartment_id_in_subtree = var.product_license_consumer_is_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration. 
* `is_compartment_id_in_subtree` - (Optional) Indicates if the given compartment is the root compartment.
* `product_license_id` - (Required) Unique product license identifier.


## Attributes Reference

The following attributes are exported:

* `product_license_consumer_collection` - The list of product_license_consumer_collection.

### ProductLicenseConsumer Reference

The following attributes are exported:

* `items` - Collection of product license consumers.
	* `are_all_options_available` - Specifies if all options are available.
	* `is_base_license_available` - Specifies if the base license is available.
	* `license_unit_type` - The product license unit.
	* `license_units_consumed` - Number of license units consumed by the resource.
	* `missing_products` - Collection of missing product licenses.
		* `category` - Product category base or option.
		* `count` - Units required for the missing product.
		* `name` - Name of the product.
	* `product_name` - The resource product name.
	* `resource_compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the resource.
	* `resource_compartment_name` - The display name of the compartment that contains the resource.
	* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the resource.
	* `resource_name` - The display name of the resource.
	* `resource_unit_count` - Number of units of the resource
	* `resource_unit_type` - The unit type for the resource.

