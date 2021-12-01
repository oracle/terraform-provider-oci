---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_vnic_attachments"
sidebar_current: "docs-oci-datasource-core-vnic_attachments"
description: |-
  Provides the list of Vnic Attachments in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_vnic_attachments
This data source provides the list of Vnic Attachments in Oracle Cloud Infrastructure Core service.

Lists the VNIC attachments in the specified compartment. A VNIC attachment
resides in the same compartment as the attached instance. The list can be
filtered by instance, VNIC, or availability domain.


## Example Usage

```hcl
data "oci_core_vnic_attachments" "test_vnic_attachments" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.vnic_attachment_availability_domain
	instance_id = oci_core_instance.test_instance.id
	vnic_id = oci_core_vnic.test_vnic.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `instance_id` - (Optional) The OCID of the instance.
* `vnic_id` - (Optional) The OCID of the VNIC.


## Attributes Reference

The following attributes are exported:

* `vnic_attachments` - The list of vnic_attachments.

### VnicAttachment Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment the VNIC attachment is in, which is the same compartment the instance is in. 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `id` - The OCID of the VNIC attachment.
* `instance_id` - The OCID of the instance.
* `nic_index` - Which physical network interface card (NIC) the VNIC uses. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingVNICs.htm). 
* `state` - The current state of the VNIC attachment.
* `subnet_id` - The OCID of the subnet to create the VNIC in.
* `time_created` - The date and time the VNIC attachment was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `vlan_id` - The OCID of the VLAN to create the VNIC in. Creating the VNIC in a VLAN (instead of a subnet) is possible only if you are an Oracle Cloud VMware Solution customer. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

	An error is returned if the instance already has a VNIC attached to it from this VLAN. 
* `vlan_tag` - The Oracle-assigned VLAN tag of the attached VNIC. Available after the attachment process is complete.

	However, if the VNIC belongs to a VLAN as part of the Oracle Cloud VMware Solution, the `vlanTag` value is instead the value of the `vlanTag` attribute for the VLAN. See [Vlan](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Vlan).

	Example: `0` 
* `vnic_id` - The OCID of the VNIC. Available after the attachment process is complete. 

