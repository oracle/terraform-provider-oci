---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_waas_address_list"
sidebar_current: "docs-oci-datasource-waas-address_list"
description: |-
  Provides details about a specific Address List in Oracle Cloud Infrastructure Waas service
---

# Data Source: oci_waas_address_list
This data source provides details about a specific Address List resource in Oracle Cloud Infrastructure Waas service.

Gets the details of an address list.

## Example Usage

```hcl
data "oci_waas_address_list" "test_address_list" {
	#Required
	address_list_id = "${oci_waas_address_list.test_address_list.id}"
}
```

## Argument Reference

The following arguments are supported:

* `address_list_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the address list. This number is generated when the address list is added to the compartment.


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

