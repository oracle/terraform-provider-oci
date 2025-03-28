---
subcategory: "Tenantmanagercontrolplane"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_tenantmanagercontrolplane_assigned_subscription_line_items"
sidebar_current: "docs-oci-datasource-tenantmanagercontrolplane-assigned_subscription_line_items"
description: |-
  Provides the list of Assigned Subscription Line Items in Oracle Cloud Infrastructure Tenantmanagercontrolplane service
---

# Data Source: oci_tenantmanagercontrolplane_assigned_subscription_line_items
This data source provides the list of Assigned Subscription Line Items in Oracle Cloud Infrastructure Tenantmanagercontrolplane service.

List line item summaries that a assigned subscription owns.

## Example Usage

```hcl
data "oci_tenantmanagercontrolplane_assigned_subscription_line_items" "test_assigned_subscription_line_items" {
	#Required
	assigned_subscription_id = oci_tenantmanagercontrolplane_assigned_subscription.test_assigned_subscription.id
}
```

## Argument Reference

The following arguments are supported:

* `assigned_subscription_id` - (Required) OCID of the assigned Oracle Cloud Subscription.


## Attributes Reference

The following attributes are exported:

* `assigned_subscription_line_item_collection` - The list of assigned_subscription_line_item_collection.

### AssignedSubscriptionLineItem Reference

The following attributes are exported:

* `items` - Array containing line item summaries in an assigned subscription.
	* `billing_model` - Billing model supported by the associated line item.
	* `id` - Subscription line item identifier.
	* `product_code` - Product code.
	* `quantity` - Product number.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_ended` - The time the subscription item and associated products should end. An RFC 3339 formatted date and time string.
	* `time_started` - The time the subscription item and associated products should start. An RFC 3339 formatted date and time string.

