---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_billing_schedules"
sidebar_current: "docs-oci-datasource-onesubscription-billing_schedules"
description: |-
  Provides the list of Billing Schedules in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_billing_schedules
This data source provides the list of Billing Schedules in Oracle Cloud Infrastructure Onesubscription service.

This list API returns all billing schedules for given subscription id and
for a particular Subscribed Service if provided


## Example Usage

```hcl
data "oci_onesubscription_billing_schedules" "test_billing_schedules" {
	#Required
	compartment_id = var.compartment_id
	subscription_id = oci_onesubscription_subscription.test_subscription.id

	#Optional
	subscribed_service_id = oci_onesubscription_subscribed_service.test_subscribed_service.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `subscribed_service_id` - (Optional) This param is used to get only the billing schedules for a particular Subscribed Service 
* `subscription_id` - (Required) This param is used to get only the billing schedules for a particular Subscription Id 


## Attributes Reference

The following attributes are exported:

* `billing_schedules` - The list of billing_schedules.

### BillingSchedule Reference

The following attributes are exported:

* `amount` - Billing schedule line net amount 
* `ar_customer_transaction_id` - Indicates the associated AR Customer transaction id a unique identifier existing on AR. 
* `ar_invoice_number` - Indicates the associated AR Invoice Number 
* `billing_frequency` - Billing frequency 
* `invoice_status` - Billing schedule invoice status 
* `net_unit_price` - Billing schedule net unit price 
* `order_number` - Order number associated with the Subscribed Service 
* `product` - Product description
	* `name` - Product name 
	* `part_number` - Indicates the associated AR Invoice Number 
* `quantity` - Billing schedule quantity 
* `subscribed_service_id` - SPM internal Subscribed Service ID 
* `time_end` - Billing schedule end date 
* `time_invoicing` - Billing schedule invoicing date 
* `time_start` - Billing schedule start date 

