---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_customer_group"
sidebar_current: "docs-oci-resource-capacity_management-occ_customer_group"
description: |-
  Provides the Occ Customer Group resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_occ_customer_group
This resource provides the Occ Customer Group resource in Oracle Cloud Infrastructure Capacity Management service.

Create customer group.

## Example Usage

```hcl
resource "oci_capacity_management_occ_customer_group" "test_occ_customer_group" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.occ_customer_group_display_name

	#Optional
	customers_list {
		#Required
		display_name = var.occ_customer_group_customers_list_display_name
		tenancy_id = oci_identity_tenancy.test_tenancy.id

		#Optional
		description = var.occ_customer_group_customers_list_description
		status = var.occ_customer_group_customers_list_status
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.occ_customer_group_description
	freeform_tags = {"bar-key"= "value"}
	lifecycle_details = var.occ_customer_group_lifecycle_details
	status = var.occ_customer_group_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Since all resources are at tenancy level hence this will be the ocid of the tenancy where operation is to be performed.
* `customers_list` - (Optional) A list containing all the customers that belong to this customer group.
	* `description` - (Optional) The description about the customer group.
	* `display_name` - (Required) The display name for the customer.
	* `status` - (Optional) To determine whether the customer is enabled/disabled.
	* `tenancy_id` - (Required) The OCID of the tenancy belonging to the customer.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A description about the customer group.
* `display_name` - (Required) (Updatable) The name of the customer group.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `lifecycle_details` - (Optional) A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
* `status` - (Optional) (Updatable) To determine whether the customer group is enabled/disabled.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy containing the customer group.
* `customers_list` - A list containing all the customers that belong to this customer group
	* `description` - The description about the customer group.
	* `display_name` - The display name for the customer
	* `occ_customer_group_id` - The OCID of the customer group.
	* `status` - To determine whether the customer is enabled/disabled.`
	* `tenancy_id` - The OCID of the tenancy belonging to the customer.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The description about the customer group.
* `display_name` - The display name of the customer group.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the customer group.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in a Failed State.
* `state` - The current lifecycle state of the resource.
* `status` - To determine whether the customer group is enabled/disabled.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time when the customer group was created.
* `time_updated` - The time when the customer group was last updated.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Occ Customer Group
	* `update` - (Defaults to 20 minutes), when updating the Occ Customer Group
	* `delete` - (Defaults to 20 minutes), when destroying the Occ Customer Group


## Import

OccCustomerGroups can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_occ_customer_group.test_occ_customer_group "id"
```

