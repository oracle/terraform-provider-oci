---
subcategory: "Oci Product Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oci_product_catalog_internal_product"
sidebar_current: "docs-oci-resource-oci_product_catalog-internal_product"
description: |-
  Provides the Internal Product resource in Oracle Cloud Infrastructure Oci Product Catalog service
---

# oci_oci_product_catalog_internal_product
This resource provides the Internal Product resource in Oracle Cloud Infrastructure Oci Product Catalog service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/oci_product_catalog

Used by the Oracle Cloud Infrastructure service teams to create a product to control the availability of the relative resources

## Example Usage

```hcl
resource "oci_oci_product_catalog_internal_product" "test_internal_product" {
	#Required
	compartment_id = var.compartment_id
	description = var.internal_product_description
	name = var.internal_product_name
	service_name = oci_announcements_service_service.test_service.name

	#Optional
	backfill_eligibility = var.internal_product_backfill_eligibility
	is_excluded = var.internal_product_is_excluded
	limits {

		#Optional
		public_limit_name = var.internal_product_limits_public_limit_name
		public_service_name = oci_announcements_service_service.test_service.name
	}
	meters {

		#Optional
		name = var.internal_product_meters_name
	}
}
```

## Argument Reference

The following arguments are supported:

* `backfill_eligibility` - (Optional) Optional create-time signal that marks the product as eligible for the backfill flow in existing Alloy regions. Omit this field to preserve the current create behavior. Final runtime backfill and auto-enable decisions remain worker-side.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment (remember that the tenancy is simply the root compartment).
* `description` - (Required) (Updatable) description to the product
* `is_excluded` - (Optional) (Updatable) Optional create-time exclusion flag. When true, the product is created in an excluded state. Omit this field to preserve backward-compatible create behavior.
* `limits` - (Optional) (Updatable) List of limits that this product is associated with.
	* `public_limit_name` - (Optional) (Updatable) public name of the limit
	* `public_service_name` - (Optional) (Updatable) public name of the limit service
* `meters` - (Optional) (Updatable) List of meters associated with this product
	* `name` - (Optional) (Updatable) Name of the meter
* `name` - (Required) (Updatable) Name of the product, defined by service teams. Unique within one service
* `service_name` - (Required) (Updatable) Name of the metering service this product relates to.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - description to the product
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - OCID of the product
* `is_excluded` - Indicates whether the product is currently excluded from SKU processing flows. This field is nullable when the exclusion flag has not been explicitly set.
* `lifecycle_details` - The displayed status of a product. lifecycleState to lifecycleDetails mapping: INACTIVE -> "Not yet launched" ACTIVE -> "Launched" NEEDS_ATTENTION -> "Pricing not set"
* `name` - Name of the product, defined by service teams. Unique within one service
* `service_name` - Name of the metering service this product relates to.
* `skus` - List of skus associated with this product.
	* `bpart_number` - sku bpartNumber
	* `description` - description of the sku
* `state` - The status of a product following the Oracle Cloud Infrastructure standard. ACTIVE: resources of this product can be used by the end customer. INACTIVE: resources of this product can't be used by the end customer. Product needs manual approval. NEEDS_ATTENTION: operator action is needed.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Date and time when the product is created Example: `2022-01-25T21:10:29.600Z`
* `time_launched` - Date and time when the product is launched. Nullable Example: `2022-01-25T21:10:29.600Z`
* `time_ready` - Date and time when the product is ready to be launched. Nullable Example: `2022-01-25T21:10:29.600Z`

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Internal Product
	* `update` - (Defaults to 20 minutes), when updating the Internal Product
	* `delete` - (Defaults to 20 minutes), when destroying the Internal Product


## Import

InternalProducts can be imported using the `id`, e.g.

```
$ terraform import oci_oci_product_catalog_internal_product.test_internal_product "internal/product/{productId}"
```
