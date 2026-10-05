---
subcategory: "Oci Product Catalog"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_oci_product_catalog_internal_product"
sidebar_current: "docs-oci-datasource-oci_product_catalog-internal_product"
description: |-
  Provides details about a specific Internal Product in Oracle Cloud Infrastructure Oci Product Catalog service
---

# Data Source: oci_oci_product_catalog_internal_product
This data source provides details about a specific Internal Product resource in Oracle Cloud Infrastructure Oci Product Catalog service.

Get the product by ID

## Example Usage

```hcl
data "oci_oci_product_catalog_internal_product" "test_internal_product" {
	#Required
	product_id = oci_oci_product_catalog_product.test_product.id
}
```

## Argument Reference

The following arguments are supported:

* `product_id` - (Required) OCID of the product


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
