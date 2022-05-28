---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_product_license"
sidebar_current: "docs-oci-resource-license_manager-product_license"
description: |-
  Provides the Product License resource in Oracle Cloud Infrastructure License Manager service
---

# oci_license_manager_product_license
This resource provides the Product License resource in Oracle Cloud Infrastructure License Manager service.

Creates a new product license.

## Example Usage

```hcl
resource "oci_license_manager_product_license" "test_product_license" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.product_license_display_name
	is_vendor_oracle = var.product_license_is_vendor_oracle
	license_unit = var.product_license_license_unit

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	images {
		#Required
		listing_id = oci_marketplace_listing.test_listing.id
		package_version = var.product_license_images_package_version
	}
	vendor_name = var.product_license_vendor_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where product licenses are created.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) Name of the product license. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `images` - (Optional) (Updatable) The image details associated with the product license.
	* `listing_id` - (Required) (Updatable) Marketplace image listing ID.
	* `package_version` - (Required) (Updatable) Image package version.
* `is_vendor_oracle` - (Required) Specifies if the product license vendor is Oracle or a third party.
* `license_unit` - (Required) The product license unit.
* `vendor_name` - (Optional) The product license vendor name, for example: Microsoft, RHEL, and so on. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Product License
	* `update` - (Defaults to 20 minutes), when updating the Product License
	* `delete` - (Defaults to 20 minutes), when destroying the Product License


## Import

ProductLicenses can be imported using the `id`, e.g.

```
$ terraform import oci_license_manager_product_license.test_product_license "id"
```

