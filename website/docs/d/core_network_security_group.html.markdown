---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_network_security_group"
sidebar_current: "docs-oci-datasource-core-network_security_group"
description: |-
  Provides details about a specific Network Security Group in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_network_security_group
This data source provides details about a specific Network Security Group resource in Oracle Cloud Infrastructure Core service.

Gets the specified network security group's information.

To list the VNICs in an NSG, see
[ListNetworkSecurityGroupVnics](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/NetworkSecurityGroupVnic/ListNetworkSecurityGroupVnics).

To list the security rules in an NSG, see
[ListNetworkSecurityGroupSecurityRules](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/SecurityRule/ListNetworkSecurityGroupSecurityRules).


## Example Usage

```hcl
data "oci_core_network_security_group" "test_network_security_group" {
	#Required
	network_security_group_id = oci_core_network_security_group.test_network_security_group.id
}
```

## Argument Reference

The following arguments are supported:

* `network_security_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment the network security group is in. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.
* `state` - The network security group's current state.
* `time_created` - The date and time the network security group was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group's VCN.

