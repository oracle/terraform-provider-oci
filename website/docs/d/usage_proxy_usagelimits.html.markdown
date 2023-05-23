---
subcategory: "Usage Proxy"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_usage_proxy_usagelimits"
sidebar_current: "docs-oci-datasource-usage_proxy-usagelimits"
description: |-
  Provides the list of Usagelimits in Oracle Cloud Infrastructure Usage Proxy service
---

# Data Source: oci_usage_proxy_usagelimits
This data source provides the list of Usagelimits in Oracle Cloud Infrastructure Usage Proxy service.

Returns the list of usage limit for the subscription ID and tenant ID.


## Example Usage

```hcl
data "oci_usage_proxy_usagelimits" "test_usagelimits" {
	#Required
	compartment_id = var.compartment_id
	subscription_id = oci_onesubscription_subscription.test_subscription.id

	#Optional
	limit_type = var.usagelimit_limit_type
	resource_type = var.usagelimit_resource_type
	service_type = var.usagelimit_service_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `limit_type` - (Optional) Hard or soft limit. Hard limits lead to breaches, soft to alerts.
* `resource_type` - (Optional) Resource Name.
* `service_type` - (Optional) Service Name.
* `subscription_id` - (Required) The subscription ID for which rewards information is requested for.


## Attributes Reference

The following attributes are exported:

* `usage_limit_collection` - The list of usage_limit_collection.

### Usagelimit Reference

The following attributes are exported:

* `items` - The list of usage limits.
	* `action` - The action when usage limit is hit
	* `alert_level` - The alert level of the usage limit
	* `created_by` - The user who created the limit
	* `entitlement_id` - Entitlement ID of the usage limit
	* `id` - The usage limit ID
	* `limit_type` - The limit type of the usage limit
	* `max_hard_limit` - The maximum hard limit set for the usage limit
	* `modified_by` - The user who modified the limit
	* `resource_name` - The resource for which the limit is defined
	* `service_name` - The service for which the limit is defined
	* `sku_part_id` - The SKU for which the usage limit is set
	* `state` - The usage limit lifecycle state.
	* `time_created` - Time when the usage limit was created
	* `time_modified` - Time when the usage limit was modified
	* `value_type` - The value type of the usage limit

