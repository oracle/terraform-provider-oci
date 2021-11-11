---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cpe"
sidebar_current: "docs-oci-resource-core-cpe"
description: |-
  Provides the Cpe resource in Oracle Cloud Infrastructure Core service
---

# oci_core_cpe
This resource provides the Cpe resource in Oracle Cloud Infrastructure Core service.

Creates a new virtual customer-premises equipment (CPE) object in the specified compartment. For
more information, see [Site-to-Site VPN Overview](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/overviewIPsec.htm).

For the purposes of access control, you must provide the [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment where you want
the CPE to reside. Notice that the CPE doesn't have to be in the same compartment as the IPSec
connection or other Networking Service components. If you're not sure which compartment to
use, put the CPE in the same compartment as the DRG. For more information about
compartments and access control, see [Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).
For information about OCIDs, see [Resource Identifiers](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

You must provide the public IP address of your on-premises router. See
[CPE Configuration](https://docs.cloud.oracle.com/iaas/Content/Network/Tasks/configuringCPE.htm).

You may optionally specify a *display name* for the CPE, otherwise a default is provided. It does not have to
be unique, and you can change it. Avoid entering confidential information.


## Example Usage

```hcl
resource "oci_core_cpe" "test_cpe" {
	#Required
	compartment_id = var.compartment_id
	ip_address = var.cpe_ip_address

	#Optional
	cpe_device_shape_id = data.oci_core_cpe_device_shapes.test_cpe_device_shapes.cpe_device_shapes.0.cpe_device_shape_id
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.cpe_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the CPE.
* `cpe_device_shape_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE device type. You can provide a value if you want to later generate CPE device configuration content for IPSec connections that use this CPE. You can also call [UpdateCpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Cpe/UpdateCpe) later to provide a value. For a list of possible values, see [ListCpeDeviceShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CpeDeviceShapeSummary/ListCpeDeviceShapes).

	For more information about generating CPE device configuration content, see:
	* [GetCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Cpe/GetCpeDeviceConfigContent)
	* [GetIpsecCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnection/GetIpsecCpeDeviceConfigContent)
	* [GetTunnelCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfigContent)
	* [GetTunnelCpeDeviceConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfig) 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `ip_address` - (Required) The public IP address of the on-premises router.  Example: `203.0.113.2` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the CPE.
* `cpe_device_shape_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE's device type. The Networking service maintains a general list of CPE device types (for example, Cisco ASA). For each type, Oracle provides CPE configuration content that can help a network engineer configure the CPE. The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) uniquely identifies the type of device. To get the OCIDs for the device types on the list, see [ListCpeDeviceShapes](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/CpeDeviceShapeSummary/ListCpeDeviceShapes).

	For information about how to generate CPE configuration content for a CPE device type, see:
	* [GetCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Cpe/GetCpeDeviceConfigContent)
	* [GetIpsecCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/IPSecConnection/GetIpsecCpeDeviceConfigContent)
	* [GetTunnelCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfigContent)
	* [GetTunnelCpeDeviceConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfig) 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The CPE's Oracle ID (OCID).
* `ip_address` - The public IP address of the on-premises router.
* `time_created` - The date and time the CPE was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cpe
	* `update` - (Defaults to 20 minutes), when updating the Cpe
	* `delete` - (Defaults to 20 minutes), when destroying the Cpe


## Import

Cpes can be imported using the `id`, e.g.

```
$ terraform import oci_core_cpe.test_cpe "id"
```

