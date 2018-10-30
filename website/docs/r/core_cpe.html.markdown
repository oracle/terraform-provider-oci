---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cpe"
sidebar_current: "docs-oci-resource-core-cpe"
description: |-
  Provides the Cpe resource in Oracle Cloud Infrastructure Core service
---

# oci_core_cpe
This resource provides the Cpe resource in Oracle Cloud Infrastructure Core service.

Creates a new virtual customer-premises equipment (CPE) object in the specified compartment. For
more information, see [IPSec VPNs](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/managingIPsec.htm).

For the purposes of access control, you must provide the OCID of the compartment where you want
the CPE to reside. Notice that the CPE doesn't have to be in the same compartment as the IPSec
connection or other Networking Service components. If you're not sure which compartment to
use, put the CPE in the same compartment as the DRG. For more information about
compartments and access control, see [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You must provide the public IP address of your on-premises router. See
[Configuring Your On-Premises Router for an IPSec VPN](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/configuringCPE.htm).

You may optionally specify a *display name* for the CPE, otherwise a default is provided. It does not have to
be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_cpe" "test_cpe" {
	#Required
	compartment_id = "${var.compartment_id}"
	ip_address = "${var.cpe_ip_address}"

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.cpe_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment to contain the CPE.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ip_address` - (Required) The public IP address of the on-premises router.  Example: `143.19.23.16` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the CPE.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The CPE's Oracle ID (OCID).
* `ip_address` - The public IP address of the on-premises router.
* `time_created` - The date and time the CPE was created, in the format defined by RFC3339.  Example: `2016-08-25T21:10:29.600Z` 

## Import

Cpes can be imported using the `id`, e.g.

```
$ terraform import oci_core_cpe.test_cpe "id"
```

