---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_commitments"
sidebar_current: "docs-oci-datasource-onesubscription-commitments"
description: |-
  Provides the list of Commitments in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_commitments
This data source provides the list of Commitments in Oracle Cloud Infrastructure Onesubscription service.

This list API returns all commitments for a particular Subscribed Service


## Example Usage

```hcl
data "oci_onesubscription_commitments" "test_commitments" {
	#Required
	compartment_id = var.compartment_id
	subscribed_service_id = oci_onesubscription_subscribed_service.test_subscribed_service.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the root compartment.
* `subscribed_service_id` - (Required) This param is used to get the commitments for a particular subscribed service 


## Attributes Reference

The following attributes are exported:

* `commitments` - The list of commitments.

### Commitment Reference

The following attributes are exported:

* `available_amount` - Commitment available amount 
* `funded_allocation_value` - Funded Allocation line value example: 12000.00 
* `id` - SPM internal Commitment ID 
* `quantity` - Commitment quantity 
* `subscribed_service_id` - SPM internal Subscribed Service ID 
* `time_end` - Commitment end date 
* `time_start` - Commitment start date 
* `used_amount` - Commitment used amount 

