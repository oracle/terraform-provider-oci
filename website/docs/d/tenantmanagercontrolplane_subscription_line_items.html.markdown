---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_subscription_line_items"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-subscription_line_items"
description: |-
  Provides the list of Subscription Line Items in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_subscription_line_items
This data source provides the list of Subscription Line Items in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

Lists the line items in a subscription.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_subscription_line_items" "test_subscription_line_items" {
	#Required
	subscription_id = oci_tenantmanagercontrolplane_subscription.test_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `subscription_id` - (Required) OCID of the subscription.


## Attributes Reference

The following attributes are exported:

* `subscription_line_item_collection` - The list of subscription_line_item_collection.

### SubscriptionLineItem Reference

The following attributes are exported:

* `items` - Array containing line item summaries in a subscription.
	* `billing_model` - Billing model supported by the associated line item.
	* `id` - Subscription line item identifier.
	* `product_code` - Product code.
	* `quantity` - Product number.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_ended` - The time the subscription item and associated products should end. An RFC 3339 formatted date and time string.
	* `time_started` - The time the subscription item and associated products should start. An RFC 3339 formatted date and time string.

