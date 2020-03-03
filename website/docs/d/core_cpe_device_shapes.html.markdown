---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cpe_device_shapes"
sidebar_current: "docs-oci-datasource-core-cpe_device_shapes"
description: |-
  Provides the list of Cpe Device Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cpe_device_shapes
This data source provides the list of Cpe Device Shapes in Oracle Cloud Infrastructure Core service.

Lists the CPE device types that the Networking service provides CPE configuration
content for (example: Cisco ASA). The content helps a network engineer configure
the actual CPE device represented by a [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/) object.

If you want to generate CPE configuration content for one of the returned CPE device types,
ensure that the [Cpe](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/) object's `cpeDeviceShapeId` attribute is set
to the CPE device type's OCID (returned by this operation).

For information about generating CPE configuration content, see these operations:

  * [GetCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/GetCpeDeviceConfigContent)
  * [GetIpsecCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnection/GetIpsecCpeDeviceConfigContent)
  * [GetTunnelCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfigContent)


## Example Usage

```hcl
data "oci_core_cpe_device_shapes" "test_cpe_device_shapes" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `cpe_device_shapes` - The list of cpe_device_shapes.

### CpeDeviceShape Reference

The following attributes are exported:

* `cpe_device_info` - Basic information about this particular CPE device type.
	* `platform_software_version` - The platform or software version of the CPE device.
	* `vendor` - The vendor that makes the CPE device.
* `cpe_device_shape_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE device shape. This value uniquely identifies the type of CPE device. 
* `parameters` - For certain CPE devices types, the customer can provide answers to questions that are specific to the device type. This attribute contains a list of those questions. The Networking service merges the answers with other information and renders a set of CPE configuration content. To provide the answers, use [UpdateTunnelCpeDeviceConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelCpeDeviceConfig/UpdateTunnelCpeDeviceConfig). 
	* `display_name` - A descriptive label for the question (for example, to display in a form in a graphical interface). 
	* `explanation` - A description or explanation of the question, to help the customer answer accurately. 
	* `key` - A string that identifies the question. 
* `template` - A template of CPE device configuration information that will be merged with the customer's answers to the questions to render the final CPE device configuration content. Also see:
	* [GetCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/GetCpeDeviceConfigContent)
	* [GetIpsecCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnection/GetIpsecCpeDeviceConfigContent)
	* [GetTunnelCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfigContent) 

