---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_address_list"
sidebar_current: "docs-oci-resource-waas-address_list"
description: |-
  Provides the Address List resource in Oracle Cloud Infrastructure Waas service
---

# oci_waas_address_list
This resource provides the Address List resource in Oracle Cloud Infrastructure Waas service.

Creates an address list in set compartment and allows it to be used in a WAAS policy.
For more information, see [WAF Settings](https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/wafsettings.htm).

## Example Usage

```hcl
resource "oci_waas_address_list" "test_address_list" {
	#Required
	addresses = "${var.address_list_addresses}"
	compartment_id = "${var.compartment_id}"
	display_name = "${var.address_list_display_name}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `addresses` - (Required) (Updatable) A list of IP addresses or CIDR notations.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the address list.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Required) (Updatable) A unique user-friendly name for the address list.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Import

AddressLists can be imported using the `id`, e.g.

```
$ terraform import oci_waas_address_list.test_address_list "id"
```

