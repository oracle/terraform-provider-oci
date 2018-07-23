---
layout: "oci"
page_title: "OCI: oci_core_internet_gateways"
sidebar_current: "docs-oci-datasource-core-internet_gateways"
description: |-
Provides a list of InternetGateways
---
# Data Source: oci_core_internet_gateways
The InternetGateways data source allows access to the list of OCI internet_gateways

Lists the Internet Gateways in the specified VCN and the specified compartment.


## Example Usage

```hcl
data "oci_core_internet_gateways" "test_internet_gateways" {
	#Required
	compartment_id = "${var.compartment_id}"
	vcn_id = "${oci_core_vcn.test_vcn.id}"

	#Optional
	display_name = "${var.internet_gateway_display_name}"
	state = "${var.internet_gateway_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 
* `vcn_id` - (Required) The OCID of the VCN.


## Attributes Reference

The following attributes are exported:

* `gateways` - The list of internet_gateways.

### InternetGateway Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the Internet Gateway.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `enabled` - Whether the gateway is enabled. When the gateway is disabled, traffic is not routed to/from the Internet, regardless of route rules. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The Internet Gateway's Oracle ID (OCID).
* `state` - The Internet Gateway's current state.
* `time_created` - The date and time the Internet Gateway was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the VCN the Internet Gateway belongs to.

