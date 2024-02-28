---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_internal_occ_capacity_request"
sidebar_current: "docs-oci-resource-capacity_management-internal_occ_capacity_request"
description: |-
  Provides the Internal Occ Capacity Request resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_internal_occ_capacity_request
This resource provides the Internal Occ Capacity Request resource in Oracle Cloud Infrastructure Capacity Management service.

The internal api to update the capacity request. This api will be used by operators for updating the capacity request to either completed, resubmitted or rejected.

## Example Usage

```hcl
resource "oci_capacity_management_internal_occ_capacity_request" "test_internal_occ_capacity_request" {
	#Required
	occ_capacity_request_id = oci_capacity_management_occ_capacity_request.test_occ_capacity_request.id

	#Optional
	lifecycle_details = var.internal_occ_capacity_request_lifecycle_details
	request_state = var.internal_occ_capacity_request_request_state
}
```

## Argument Reference

The following arguments are supported:

* `lifecycle_details` - (Optional) (Updatable) A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
* `occ_capacity_request_id` - (Required) The OCID of the capacity request.
* `request_state` - (Optional) (Updatable) The subset of request states available internally for updating the capacity request.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Internal Occ Capacity Request
	* `update` - (Defaults to 20 minutes), when updating the Internal Occ Capacity Request
	* `delete` - (Defaults to 20 minutes), when destroying the Internal Occ Capacity Request


## Import

InternalOccCapacityRequests can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_internal_occ_capacity_request.test_internal_occ_capacity_request "internal/occCapacityRequests/{occCapacityRequestId}/compartmentId/{compartmentId}" 
```

