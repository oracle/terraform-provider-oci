---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_resources"
sidebar_current: "docs-oci-datasource-usage_proxy-resources"
description: |-
  Provides the list of Resources in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_resources
This data source provides the list of Resources in Oracle Cloud Infrastructure Usage Proxy service.

Returns the resource details for a service
> **Important**: Calls to this API will only succeed against the endpoint in the home region.


## Example Usage

```hcl
data "oci_usage_proxy_resources" "test_resources" {
	#Required
	compartment_id = var.compartment_id
	service_name = oci_core_service.test_service.name

	#Optional
	entitlement_id = oci_usage_proxy_entitlement.test_entitlement.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `entitlement_id` - (Optional) Subscription or entitlement Id.
* `service_name` - (Required) Service Name.


## Attributes Reference

The following attributes are exported:

* `resources_collection` - The list of resources_collection.

### Resource Reference

The following attributes are exported:

* `items` - The list of resource details for a service.
	* `child_resources` - The details of any child resources.
	* `daily_unit_display_name` - Units to be used for daily aggregated data.
	* `description` - Description of the resource.
	* `hourly_unit_display_name` - Units to be used for hourly aggregated data.
	* `instance_type` - Instance type for the resource.
	* `is_purchased` - Indicates if the SKU was purchased
	* `name` - Name of the resource.
	* `raw_unit_display_name` - Default units to use when unspecified.
	* `servicename` - Name of the service.
	* `skus` - The details of resource Skus.
		* `cloud_credit_type` - The cloud credit type for the resource.
		* `sku_id` - The Sku Id for the resource.
		* `sku_type` - The Sku type for the resource.
	* `usage_data_type` - Usage data type of the resource.

