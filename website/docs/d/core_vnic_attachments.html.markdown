---
layout: "oci"
page_title: "OCI: oci_core_vnic_attachments"
sidebar_current: "docs-oci-datasource-core-vnic_attachments"
description: |-
  Provides a list of VnicAttachments
---

# Data Source: oci_core_vnic_attachments
The VnicAttachments data source allows access to the list of OCI vnic_attachments

Lists the VNIC attachments in the specified compartment. A VNIC attachment
resides in the same compartment as the attached instance. The list can be
filtered by instance, VNIC, or Availability Domain.


## Example Usage

```hcl
data "oci_core_vnic_attachments" "test_vnic_attachments" {
	#Required
	compartment_id = "${var.compartment_id}"

	#Optional
	availability_domain = "${var.vnic_attachment_availability_domain}"
	instance_id = "${oci_core_instance.test_instance.id}"
	vnic_id = "${oci_core_vnic.test_vnic.id}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the Availability Domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `instance_id` - (Optional) The OCID of the instance.
* `vnic_id` - (Optional) The OCID of the VNIC.


## Attributes Reference

The following attributes are exported:

* `vnic_attachments` - The list of vnic_attachments.

### VnicAttachment Reference

The following attributes are exported:

* `availability_domain` - The Availability Domain of the instance.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment the VNIC attachment is in, which is the same compartment the instance is in. 
* `display_name` - A user-friendly name. Does not have to be unique. Avoid entering confidential information. 
* `id` - The OCID of the VNIC attachment.
* `instance_id` - The OCID of the instance.
* `nic_index` - Which physical network interface card (NIC) the VNIC uses. Certain bare metal instance shapes have two active physical NICs (0 and 1). If you add a secondary VNIC to one of these instances, you can specify which NIC the VNIC will use. For more information, see [Virtual Network Interface Cards (VNICs)](https://docs.us-phoenix-1.oraclecloud.com/Content/Network/Tasks/managingVNICs.htm). 
* `state` - The current state of the VNIC attachment.
* `subnet_id` - The OCID of the VNIC's subnet.
* `time_created` - The date and time the VNIC attachment was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 
* `vlan_tag` - The Oracle-assigned VLAN tag of the attached VNIC. Available after the attachment process is complete.  Example: `0` 
* `vnic_id` - The OCID of the VNIC. Available after the attachment process is complete.

