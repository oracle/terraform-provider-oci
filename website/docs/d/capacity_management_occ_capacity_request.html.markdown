---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_capacity_request"
sidebar_current: "docs-oci-datasource-capacity_management-occ_capacity_request"
description: |-
  Provides details about a specific Occ Capacity Request in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occ_capacity_request
This data source provides details about a specific Occ Capacity Request resource in Oracle Cloud Infrastructure Capacity Management service.

Get details about the capacity request.

## Example Usage

```hcl
data "oci_capacity_management_occ_capacity_request" "test_occ_capacity_request" {
	#Required
	occ_capacity_request_id = oci_capacity_management_occ_capacity_request.test_occ_capacity_request.id
}
```

## Argument Reference

The following arguments are supported:

* `occ_capacity_request_id` - (Required) The OCID of the capacity request.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain (AD) for which the capacity request was made.
* `compartment_id` - The OCID of the tenancy from which the request was made.
* `date_expected_capacity_handover` - The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Meaningful text about the capacity request.
* `details` - A list of resources requested as part of this request
	* `actual_handover_quantity` - The actual handed over quantity of resources at the time of request resolution.
	* `date_actual_handover` - The date on which the actual handover quantity of resources is delivered.
	* `date_expected_handover` - The date on which the latest increment to supplied quantity of resources was delivered.
	* `demand_quantity` - The number of compute server's with name <resourceName> required by the user.
	* `expected_handover_quantity` - The incremental quantity of resources supplied as the provisioning is underway.
	* `resource_name` - The name of the COMPUTE server shape for which the request is made. Do not use CAPACITY_CONSTRAINT as the resource name.
	* `resource_type` - The type of the resource against which the user wants to place a capacity request.
	* `workload_type` - The type of the workload (Generic/ROW).
* `display_name` - The display name of the capacity request.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the capacity request.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
* `namespace` - The name of the Oracle Cloud Infrastructure service in consideration. For example, Compute, Exadata, and so on. 
* `occ_availability_catalog_id` - The OCID of the availability catalog against which the capacity request was placed.
* `occ_customer_group_id` - The OCID of the customer group to which this customer belongs to.
* `region` - The name of the region for which the capacity request was made.
* `request_state` - The different states the capacity request goes through.
* `state` - The current lifecycle state of the resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the capacity request was created.
* `time_updated` - The time when the capacity request was updated.

