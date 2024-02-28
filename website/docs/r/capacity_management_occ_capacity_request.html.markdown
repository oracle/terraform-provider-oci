---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_capacity_request"
sidebar_current: "docs-oci-resource-capacity_management-occ_capacity_request"
description: |-
  Provides the Occ Capacity Request resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_occ_capacity_request
This resource provides the Occ Capacity Request resource in Oracle Cloud Infrastructure Capacity Management service.

Create Capacity Request.
  Updates the OccCapacityRequest by evaluating a sequence of instructions.

## Example Usage

```hcl
resource "oci_capacity_management_occ_capacity_request" "test_occ_capacity_request" {
	#Required
	availability_domain = var.occ_capacity_request_availability_domain
	compartment_id = var.compartment_id
	date_expected_capacity_handover = var.occ_capacity_request_date_expected_capacity_handover
	details {
		#Required
		demand_quantity = var.occ_capacity_request_details_demand_quantity
		resource_name = oci_usage_proxy_resource.test_resource.name
		resource_type = var.occ_capacity_request_details_resource_type
		workload_type = var.occ_capacity_request_details_workload_type

		#Optional
		actual_handover_quantity = var.occ_capacity_request_details_actual_handover_quantity
		date_actual_handover = var.occ_capacity_request_details_date_actual_handover
		date_expected_handover = var.occ_capacity_request_details_date_expected_handover
		expected_handover_quantity = var.occ_capacity_request_details_expected_handover_quantity
	}
	display_name = var.occ_capacity_request_display_name
	namespace = var.occ_capacity_request_namespace
	occ_availability_catalog_id = oci_capacity_management_occ_availability_catalog.test_occ_availability_catalog.id
	occ_capacity_request_id = var.occ_capacity_request_occ_capacity_request_id
	region = var.occ_capacity_request_region

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.occ_capacity_request_description
	freeform_tags = {"bar-key"= "value"}
	lifecycle_details = var.occ_capacity_request_lifecycle_details
	patch_operations {
		#Required
		operation = var.occ_capacity_request_patch_operations_operation
		selection = var.occ_capacity_request_patch_operations_selection

		#Optional
		from = var.occ_capacity_request_patch_operations_from
		position = var.occ_capacity_request_patch_operations_position
		selected_item = var.occ_capacity_request_patch_operations_selected_item
		value = var.occ_capacity_request_patch_operations_value
		values = var.occ_capacity_request_patch_operations_values
	}
	request_state = var.occ_capacity_request_request_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain (AD) for which the capacity request is made. If this is specified then the capacity will be validated and fulfilled within the scope of this AD.
* `compartment_id` - (Required) Since all resources are at tenancy level hence this will be the ocid of the tenancy where operation is to be performed.
* `date_expected_capacity_handover` - (Required) The date by which the capacity requested by customers before dateFinalCustomerOrder needs to be fulfilled.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) Meaningful text about the capacity request.
* `details` - (Required) A list of different resources requested by the user.
	* `actual_handover_quantity` - (Optional) The actual handed over quantity of resources at the time of request resolution.
	* `date_actual_handover` - (Optional) The date on which the actual handover quantity of resources is delivered.
	* `date_expected_handover` - (Optional) The date on which the latest increment to supplied quantity of resources was delivered.
	* `demand_quantity` - (Required) The number of compute server's with name <resourceName> required by the user.
	* `expected_handover_quantity` - (Optional) The incremental quantity of resources supplied as the provisioning is underway.
	* `resource_name` - (Required) The name of the COMPUTE server shape for which the request is made. Do not use CAPACITY_CONSTRAINT as the resource name.
	* `resource_type` - (Required) The type of the resource against which the user wants to place a capacity request.
	* `workload_type` - (Required) The type of the workload (Generic/ROW).
* `display_name` - (Required) (Updatable) An user-friendly name for the capacity request. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `lifecycle_details` - (Optional) A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
* `namespace` - (Required) The name of the Oracle Cloud Infrastructure service in consideration. For example, Compute, Exadata, and so on.
* `occ_availability_catalog_id` - (Required) The OCID of the availability catalog against which capacity request is made.
* `occ_capacity_request_id` - (Required) 
* `patch_operations` - (Optional) (Updatable) 
	* `from` - (Required when operation=MOVE) (Updatable) 
	* `operation` - (Required) (Updatable) The operation can be one of these values: `INSERT`, `INSERT_MULTIPLE`, `MERGE`, `MOVE`, `PROHIBIT`, `REMOVE`, `REPLACE`, `REQUIRE`
	* `position` - (Applicable when operation=INSERT | INSERT_MULTIPLE | MOVE) (Updatable) 
	* `selected_item` - (Applicable when operation=INSERT | INSERT_MULTIPLE) (Updatable) 
	* `selection` - (Required) (Updatable) 
	* `value` - (Required when operation=INSERT | MERGE | PROHIBIT | REPLACE | REQUIRE) (Updatable) 
	* `values` - (Required when operation=INSERT_MULTIPLE) (Updatable) 
* `region` - (Required) The name of the region for which the capacity request is made.
* `request_state` - (Optional) (Updatable) The subset of request states available for creating the capacity request.


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
	* `create` - (Defaults to 20 minutes), when creating the Occ Capacity Request
	* `update` - (Defaults to 20 minutes), when updating the Occ Capacity Request
	* `delete` - (Defaults to 20 minutes), when destroying the Occ Capacity Request


## Import

OccCapacityRequests can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_occ_capacity_request.test_occ_capacity_request "id"
```

