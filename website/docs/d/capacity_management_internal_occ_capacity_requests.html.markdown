---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occ_capacity_requests"
sidebar_current: "docs-oci-datasource-capacity_management-internal_occ_capacity_requests"
description: |-
  Provides the list of Internal Occ Capacity Requests in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_internal_occ_capacity_requests
This data source provides the list of Internal Occ Capacity Requests in Oracle Cloud Infrastructure Capacity Management service.

An internal api to list all capacity requests.

## Example Usage

```hcl
data "oci_capacity_management_internal_occ_capacity_requests" "test_internal_occ_capacity_requests" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.internal_occ_capacity_request_display_name
	id = var.internal_occ_capacity_request_id
	namespace = var.internal_occ_capacity_request_namespace
	occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `display_name` - (Optional) A filter to return only the resources that match the entire display name. The match is not case sensitive.
* `id` - (Optional) A filter to return the list of capacity requests based on the OCID of the capacity request. This is done for the users who have INSPECT permission on the resource but do not have READ permission.
* `namespace` - (Optional) The namespace by which we would filter the list.
* `occ_availability_catalog_id` - (Optional) A filter to return the list of capacity requests based on the OCID of the availability catalog against which they were created.
* `occ_customer_group_id` - (Optional) The customer group ocid by which we would filter the list.


## Attributes Reference

The following attributes are exported:

* `occ_capacity_request_collection` - The list of occ_capacity_request_collection.

### InternalOccCapacityRequest Reference

The following attributes are exported:

* `items` - An array of capacity requests.
	* `availability_domain` - The availability domain (AD) for which the capacity request was made.
	* `compartment_id` - The OCID of the tenancy from which the request was made.
	* `date_expected_capacity_handover` - The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
	* `description` - Meaningful text about the capacity request.
	* `display_name` - The display name of the capacity request.
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
	* `id` - The OCID of the capacity request.
	* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
	* `namespace` - The name of the Oracle Cloud Infrastructure service in consideration. For example, Compute, Exadata, and so on. 
	* `occ_availability_catalog_id` - The OCID of the availability catalog against which the capacity request was placed.
	* `occ_customer_group_id` - The OCID of the customer group to which this customer belongs to.
	* `region` - The name of the region for which the capacity request was made.
	* `request_state` - A list of states through which the capacity request goes by.
	* `state` - The current lifecycle state of the customer group.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time when the capacity request was created.
	* `time_updated` - The time when the capacity request was updated.

