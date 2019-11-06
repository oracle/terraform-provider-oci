---
subcategory: "Waas"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_address_lists"
sidebar_current: "docs-oci-datasource-waas-address_lists"
description: |-
  Provides the list of Address Lists in Oracle Cloud Infrastructure Waas service
---

# Data Source: oci_waas_address_lists
This data source provides the list of Address Lists in Oracle Cloud Infrastructure Waas service.

Gets a list of address lists that can be used in a WAAS policy.

## Example Usage

```hcl
data "oci_waas_address_lists" "test_address_lists" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	ids = "${var.address_list_ids}"
	names = "${var.address_list_names}"
	states = "${var.address_list_states}"
	time_created_greater_than_or_equal_to = "${var.address_list_time_created_greater_than_or_equal_to}"
	time_created_less_than = "${var.address_list_time_created_less_than}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment. This number is generated when the compartment is created.
* `ids` - (Optional) Filter address lists using a list of address lists OCIDs.
* `names` - (Optional) Filter address lists using a list of names.
* `states` - (Optional) Filter address lists using a list of lifecycle states.
* `time_created_greater_than_or_equal_to` - (Optional) A filter that matches address lists created on or after the specified date-time.
* `time_created_less_than` - (Optional) A filter that matches address lists created before the specified date-time.


## Attributes Reference

The following attributes are exported:

* `address_lists` - The list of address_lists.

### AddressList Reference

The following attributes are exported:

* `address_count` - The total number of unique IP addresses in the address list.
* `addresses` - The list of IP addresses or CIDR notations.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the address list's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The user-friendly name of the address list.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the address list.
* `state` - The current lifecycle state of the address list.
* `time_created` - The date and time the address list was created, expressed in RFC 3339 timestamp format.

