---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occm_demand_signals"
sidebar_current: "docs-oci-datasource-capacity_management-occm_demand_signals"
description: |-
  Provides the list of Occm Demand Signals in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occm_demand_signals
This data source provides the list of Occm Demand Signals in Oracle Cloud Infrastructure Capacity Management service.

This GET call is used to list all demand signals within the compartment passed as a query parameter.


## Example Usage

```hcl
data "oci_capacity_management_occm_demand_signals" "test_occm_demand_signals" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.occm_demand_signal_display_name
	id = var.occm_demand_signal_id
	lifecycle_details = var.occm_demand_signal_lifecycle_details
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `display_name` - (Optional) A filter to return only the resources that match the entire display name. The match is not case sensitive.
* `id` - (Optional) A query parameter to filter the list of demand signals based on it's OCID. 
* `lifecycle_details` - (Optional) A query parameter to filter the list of demand signals based on its state. 


## Attributes Reference

The following attributes are exported:

* `occm_demand_signal_collection` - The list of occm_demand_signal_collection.

### OccmDemandSignal Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy from which the request to create the demand signal was made. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Meaningful text about the demand signal. 
* `display_name` - The display name of the demand signal. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the demand signal. 
* `lifecycle_details` - The different states associated with a demand signal. 

	CREATED -> A demand signal is by default created in this state.  SUBMITTED -> Once you have reviewed the details of the demand signal, you can transition it to SUBMITTED state so that Oracle Cloud Infrastructure can start working on it. DELETED -> You can delete a demand signal as long as it is in either CREATED or SUBMITTED state. IN_PROGRESS -> Once Oracle Cloud Infrastructure starts working on a given demand signal. They transition it to IN_PROGRESS. CANCELLED -> Oracle Cloud Infrastructure can transition the demand signal to this state. COMPLETED -> Oracle Cloud Infrastructure will transition the demand signal to COMPLETED state once the quantities which Oracle Cloud Infrastructure committed to deliver to you has been delivered. 
* `state` - The current lifecycle state of the resource. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the demand signal was created. 
* `time_updated` - The time when the demand signal was last updated. 

