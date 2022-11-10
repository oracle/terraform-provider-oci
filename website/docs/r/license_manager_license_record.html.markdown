---
subcategory: "License Manager"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_license_manager_license_record"
sidebar_current: "docs-oci-resource-license_manager-license_record"
description: |-
  Provides the License Record resource in Oracle Cloud Infrastructure License Manager service
---

# oci_license_manager_license_record
This resource provides the License Record resource in Oracle Cloud Infrastructure License Manager service.

Creates a new license record for the given product license ID.

## Example Usage

```hcl
resource "oci_license_manager_license_record" "test_license_record" {
	#Required
	display_name = var.license_record_display_name
	is_perpetual = var.license_record_is_perpetual
	is_unlimited = var.license_record_is_unlimited
	product_license_id = oci_license_manager_product_license.test_product_license.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	expiration_date = var.license_record_expiration_date
	freeform_tags = {"bar-key"= "value"}
	license_count = var.license_record_license_count
	product_id = oci_license_manager_product.test_product.id
	support_end_date = var.license_record_support_end_date
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) License record name. 
* `expiration_date` - (Optional) (Updatable) The license record end date in [RFC 3339](https://tools.ietf.org/html/rfc3339) date format. Example: `2018-09-12` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_perpetual` - (Required) (Updatable) Specifies if the license record term is perpertual.
* `is_unlimited` - (Required) (Updatable) Specifies if the license count is unlimited.
* `license_count` - (Optional) (Updatable) The number of license units added by a user in a license record. Default 1 
* `product_id` - (Optional) (Updatable) The license record product ID.
* `product_license_id` - (Required) Unique product license identifier.
* `support_end_date` - (Optional) (Updatable) The license record support end date in [RFC 3339](https://tools.ietf.org/html/rfc3339) date format. Example: `2018-09-12` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) where the license record is created.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - The license record display name. Avoid entering confidential information. 
* `expiration_date` - The license record end date in [RFC 3339](https://tools.ietf.org/html/rfc3339) date format. Example: `2018-09-12` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type, or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The license record [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `is_perpetual` - Specifies if the license record term is perpertual.
* `is_unlimited` - Specifies if the license count is unlimited.
* `license_count` - The number of license units added by the user for the given license record. Default 1 
* `license_unit` - The product license unit.
* `product_id` - The license record product ID.
* `product_license` - The product license name with which the license record is associated.
* `product_license_id` - The product license [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) with which the license record is associated.
* `state` - The current license record state.
* `support_end_date` - The license record support end date in [RFC 3339](https://tools.ietf.org/html/rfc3339) date format. Example: `2018-09-12` 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the license record was created. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.
* `time_updated` - The time the license record was updated. An [RFC 3339](https://tools.ietf.org/html/rfc3339)-formatted datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the License Record
	* `update` - (Defaults to 20 minutes), when updating the License Record
	* `delete` - (Defaults to 20 minutes), when destroying the License Record


## Import

LicenseRecords can be imported using the `id`, e.g.

```
$ terraform import oci_license_manager_license_record.test_license_record "id"
```

