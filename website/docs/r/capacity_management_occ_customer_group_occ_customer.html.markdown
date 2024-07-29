---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_customer_group_occ_customer"
sidebar_current: "docs-oci-resource-capacity_management-occ_customer_group_occ_customer"
description: |-
  Provides the Occ Customer Group Occ Customer resource in Oracle Cloud Infrastructure Capacity Management service
---

# oci_capacity_management_occ_customer_group_occ_customer
This resource provides the Occ Customer Group Occ Customer resource in Oracle Cloud Infrastructure Capacity Management service.

Create customer.

## Example Usage

```hcl
resource "oci_capacity_management_occ_customer_group_occ_customer" "test_occ_customer_group_occ_customer" {
	#Required
	display_name = var.occ_customer_group_occ_customer_display_name
	occ_customer_group_id = oci_capacity_management_occ_customer_group.test_occ_customer_group.id
	tenancy_id = oci_identity_tenancy.test_tenancy.id

	#Optional
	description = var.occ_customer_group_occ_customer_description
	status = var.occ_customer_group_occ_customer_status
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Optional) (Updatable) The description about the customer group.
* `display_name` - (Required) (Updatable) The display name for the customer.
* `occ_customer_group_id` - (Required) The OCID of the customer group. 
* `status` - (Optional) (Updatable) To determine whether the customer is enabled/disabled.
* `tenancy_id` - (Required) The OCID of the tenancy belonging to the customer.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `description` - The description about the customer group.
* `display_name` - The display name for the customer
* `occ_customer_group_id` - The OCID of the customer group.
* `status` - To determine whether the customer is enabled/disabled.`
* `tenancy_id` - The OCID of the tenancy belonging to the customer.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Occ Customer Group Occ Customer
	* `update` - (Defaults to 20 minutes), when updating the Occ Customer Group Occ Customer
	* `delete` - (Defaults to 20 minutes), when destroying the Occ Customer Group Occ Customer


## Import

OccCustomerGroupOccCustomers can be imported using the `id`, e.g.

```
$ terraform import oci_capacity_management_occ_customer_group_occ_customer.test_occ_customer_group_occ_customer "id"
```

