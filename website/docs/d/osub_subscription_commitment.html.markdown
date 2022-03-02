---
subcategory: "Osub Subscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osub_subscription_commitment"
sidebar_current: "docs-oci-datasource-osub_subscription-commitment"
description: |-
  Provides details about a specific Commitment in Oracle Cloud Infrastructure Osub Subscription service
---

# Data Source: oci_osub_subscription_commitment
This data source provides details about a specific Commitment resource in Oracle Cloud Infrastructure Osub Subscription service.

This API returns the commitment details corresponding to the id provided


## Example Usage

```hcl
data "oci_osub_subscription_commitment" "test_commitment" {
	#Required
	commitment_id = oci_osub_subscription_commitment.test_commitment.id

	#Optional
	x_one_gateway_subscription_id = var.commitment_x_one_gateway_subscription_id
	x_one_origin_region = var.commitment_x_one_origin_region
}
```

## Argument Reference

The following arguments are supported:

* `commitment_id` - (Required) The Commitment Id
* `x_one_gateway_subscription_id` - (Optional) This header is meant to be used only for internal purposes and will be ignored on any public request. The purpose of this header is  to help on Gateway to API calls identification.  
* `x_one_origin_region` - (Optional) The Oracle Cloud Infrastructure home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc. 


## Attributes Reference

The following attributes are exported:

* `available_amount` - Commitment available amount 
* `funded_allocation_value` - Funded Allocation line value example: 12000.00 
* `id` - SPM internal Commitment ID 
* `quantity` - Commitment quantity 
* `time_end` - Commitment end date 
* `time_start` - Commitment start date 
* `used_amount` - Commitment used amount 

