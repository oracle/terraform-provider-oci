---
subcategory: "Onesubscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_onesubscription_commitment"
sidebar_current: "docs-oci-datasource-onesubscription-commitment"
description: |-
  Provides details about a specific Commitment in Oracle Cloud Infrastructure Onesubscription service
---

# Data Source: oci_onesubscription_commitment
This data source provides details about a specific Commitment resource in Oracle Cloud Infrastructure Onesubscription service.

This API returns the commitment details corresponding to the id provided


## Example Usage

```hcl
data "oci_onesubscription_commitment" "test_commitment" {
	#Required
	commitment_id = oci_onesubscription_commitment.test_commitment.id
}
```

## Argument Reference

The following arguments are supported:

* `commitment_id` - (Required) The Commitment Id


## Attributes Reference

The following attributes are exported:

* `available_amount` - Commitment available amount 
* `funded_allocation_value` - Funded Allocation line value example: 12000.00 
* `id` - SPM internal Commitment ID 
* `quantity` - Commitment quantity 
* `subscribed_service_id` - SPM internal Subscribed Service ID 
* `time_end` - Commitment end date 
* `time_start` - Commitment start date 
* `used_amount` - Commitment used amount 

