---
subcategory: "Osub Subscription"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_osub_subscription_commitments"
sidebar_current: "docs-oci-datasource-osub_subscription-commitments"
description: |-
  Provides the list of Commitments in Oracle Cloud Infrastructure Osub Subscription service
---

# Data Source: oci_osub_subscription_commitments
This data source provides the list of Commitments in Oracle Cloud Infrastructure Osub Subscription service.

This list API returns all commitments for a particular Subscribed Service


## Example Usage

```hcl
data "oci_osub_subscription_commitments" "test_commitments" {
	#Required
	compartment_id = var.compartment_id
	subscribed_service_id = oci_core_service.test_service.id

	#Optional
	x_one_gateway_subscription_id = var.commitment_x_one_gateway_subscription_id
	x_one_origin_region = var.commitment_x_one_origin_region
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `subscribed_service_id` - (Required) This param is used to get the commitments for a particular subscribed service 
* `x_one_gateway_subscription_id` - (Optional) This header is meant to be used only for internal purposes and will be ignored on any public request. The purpose of this header is  to help on Gateway to API calls identification.  
* `x_one_origin_region` - (Optional) The Oracle Cloud Infrastructure home region name in case home region is not us-ashburn-1 (IAD), e.g. ap-mumbai-1, us-phoenix-1 etc. 


## Attributes Reference

The following attributes are exported:

* `commitments` - The list of commitments.

### Commitment Reference

The following attributes are exported:

* `available_amount` - Commitment available amount 
* `funded_allocation_value` - Funded Allocation line value example: 12000.00 
* `id` - SPM internal Commitment ID 
* `quantity` - Commitment quantity 
* `time_end` - Commitment end date 
* `time_start` - Commitment start date 
* `used_amount` - Commitment used amount 

