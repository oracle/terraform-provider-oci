---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occm_demand_signal"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occm_demand_signal"
description: |-
  Provides details about a specific Internal Occm Demand Signal in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occm_demand_signal
This data source provides details about a specific Internal Occm Demand Signal resource in Oracle Cloud Infrastructure Capacity Management service.

This is an internal GET API which gets the detailed information about a specific demand signal.


## Example Usage

```hcl
data "oci_capacity_management_internal_occm_demand_signal" "test_internal_occm_demand_signal" {
	#Required
	occm_demand_signal_id = oci_capacity_management_occm_demand_signal.test_occm_demand_signal.id
}
```

## Argument Reference

The following arguments are supported:

* `occm_demand_signal_id` - (Required) The OCID of the demand signal. 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy from which the request to create the demand signal was made. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A short description about the demand signal. 
* `display_name` - The display name of the demand signal. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the demand signal. 
* `lifecycle_details` - The different states associated with a demand signal. 

	CREATED -> A demand signal is by default created in this state.  SUBMITTED -> Once you have reviewed the details of the demand signal, you can transition it to SUBMITTED state so that Oracle Cloud Infrastructure can start working on it. DELETED -> You can delete a demand signal as long as it is in either CREATED or SUBMITTED state. IN_PROGRESS -> Once Oracle Cloud Infrastructure starts working on a given demand signal. They transition it to IN_PROGRESS. REJECTED -> Oracle Cloud Infrastructure can transition the demand signal to this state if all the demand signal items of that demand signal are declined. COMPLETED -> Oracle Cloud Infrastructure will transition the demand signal to COMPLETED state once the quantities which Oracle Cloud Infrastructure committed to deliver to you has been delivered. 
* `occ_customer_group_id` - The OCID of the customer group in which the demand signal is created. 
* `state` - The current lifecycle state of the demand signal. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the demand signal was created. 
* `time_updated` - The time when the demand signal was last updated. 

