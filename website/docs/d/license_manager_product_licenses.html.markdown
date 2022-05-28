---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_product_licenses"
sidebar_current: "docs-oci-datasource-license_manager-product_licenses"
description: |-
  Provides the list of Product Licenses in Oracle Cloud Infrastructure License Manager service
---

# Data Source: oci_license_manager_product_licenses
This data source provides the list of Product Licenses in Oracle Cloud Infrastructure License Manager service.

Retrieves all the product licenses from a given compartment.

## Example Usage

```hcl
data "oci_license_manager_product_licenses" "test_product_licenses" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	is_compartment_id_in_subtree = var.product_license_is_compartment_id_in_subtree
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) used for the license record, product license, and configuration. 
* `is_compartment_id_in_subtree` - (Optional) Indicates if the given compartment is the root compartment.


## Attributes Reference

The following attributes are exported:

* `product_license_collection` - The list of product_license_collection.

### ProductLicense Reference

The following attributes are exported:

* `active_license_record_count` - The number of active license records associated with the product license.
* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the product license is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - License record name 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The product license [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `images` - The images associated with the product license.
	* `id` - The image ID associated with the product license.
	* `listing_id` - The image listing ID.
	* `listing_name` - The listing name associated with the product license.
	* `package_version` - The image package version.
	* `publisher` - The image publisher.
* `is_over_subscribed` - Specifies whether or not the product license is oversubscribed.
* `is_unlimited` - Specifies if the license unit count is unlimited.
* `is_vendor_oracle` - Specifies whether the vendor is Oracle or a third party.
* `license_unit` - The product license unit.
* `state` - The current product license state.
* `status` - The current product license status.
* `status_description` - Status description for the current product license status. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the product license was created. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.
* `time_updated` - The time the product license was updated. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.
* `total_active_license_unit_count` - The total number of licenses available for the product license, calculated by adding up all the license counts for active license records associated with the product license.
* `total_license_record_count` - The number of license records associated with the product license. 
* `total_license_units_consumed` - The number of license units consumed. Updated after each allocation run. 
* `vendor_name` - The vendor of the ProductLicense 

