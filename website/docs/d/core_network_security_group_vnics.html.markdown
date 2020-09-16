---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_network_security_group_vnics"
sidebar_current: "docs-oci-datasource-core-network_security_group_vnics"
description: |-
  Provides the list of Network Security Group Vnics in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_network_security_group_vnics
This data source provides the list of Network Security Group Vnics in Oracle Cloud Infrastructure Core service.

Lists the VNICs in the specified network security group.


## Example Usage

```hcl
data "oci_core_network_security_group_vnics" "test_network_security_group_vnics" {
	#Required
	network_security_group_id = oci_core_network_security_group.test_network_security_group.id
}
```

## Argument Reference

The following arguments are supported:

* `network_security_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the network security group.


## Attributes Reference

The following attributes are exported:

* `network_security_group_vnics` - The list of network_security_group_vnics.

### NetworkSecurityGroupVnic Reference

The following attributes are exported:

* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent resource that the VNIC is attached to (for example, a Compute instance). 
* `time_associated` - The date and time the VNIC was added to the network security group, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vnic_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VNIC.

