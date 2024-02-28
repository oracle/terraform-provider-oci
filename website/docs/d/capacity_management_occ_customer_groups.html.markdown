---
subcategory: "Capacity Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_capacity_management_occ_customer_groups"
sidebar_current: "docs-oci-datasource-capacity_management-occ_customer_groups"
description: |-
  Provides the list of Occ Customer Groups in Oracle Cloud Infrastructure Capacity Management service
---

# Data Source: oci_capacity_management_occ_customer_groups
This data source provides the list of Occ Customer Groups in Oracle Cloud Infrastructure Capacity Management service.

Lists all the customer groups.


## Example Usage

```hcl
data "oci_capacity_management_occ_customer_groups" "test_occ_customer_groups" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.occ_customer_group_display_name
	id = var.occ_customer_group_id
	status = var.occ_customer_group_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ocid of the compartment or tenancy in which resources are to be listed. This will also be used for authorization purposes.
* `display_name` - (Optional) A filter to return only the resources that match the entire display name. The match is not case sensitive.
* `id` - (Optional) A query filter to return the list result based on the customer group OCID. This is done for users who have INSPECT permission but do not have READ permission.
* `status` - (Optional) A query filter to return the list result based on status.


## Attributes Reference

The following attributes are exported:

* `occ_customer_group_collection` - The list of occ_customer_group_collection.

### OccCustomerGroup Reference

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

