---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cpe_device_shape"
sidebar_current: "docs-oci-datasource-core-cpe_device_shape"
description: |-
  Provides details about a specific Cpe Device Shape in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cpe_device_shape
This data source provides details about a specific Cpe Device Shape resource in Oracle Cloud Infrastructure Core service.

Gets the detailed information about the specified CPE device type. This might include a set of questions
that are specific to the particular CPE device type. The customer must supply answers to those questions
(see [UpdateTunnelCpeDeviceConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelCpeDeviceConfig/UpdateTunnelCpeDeviceConfig)).
The service merges the answers with a template of other information for the CPE device type. The following
operations return the merged content:

  * [GetCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/GetCpeDeviceConfigContent)
  * [GetIpsecCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnection/GetIpsecCpeDeviceConfigContent)
  * [GetTunnelCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfigContent)


## Example Usage

```hcl
data "oci_core_cpe_device_shape" "test_cpe_device_shape" {
	#Required
	cpe_device_shape_id = oci_core_cpe_device_shape.test_cpe_device_shape.id
}
```

## Argument Reference

The following arguments are supported:

* `cpe_device_shape_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE device shape.


## Attributes Reference

The following attributes are exported:

* `cpe_device_info` - Basic information about a particular CPE device type.
	* `platform_software_version` - The platform or software version of the CPE device.
	* `vendor` - The vendor that makes the CPE device.
* `cpe_device_shape_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the CPE device shape. This value uniquely identifies the type of CPE device. 
* `parameters` - For certain CPE devices types, the customer can provide answers to questions that are specific to the device type. This attribute contains a list of those questions. The Networking service merges the answers with other information and renders a set of CPE configuration content. To provide the answers, use [UpdateTunnelCpeDeviceConfig](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelCpeDeviceConfig/UpdateTunnelCpeDeviceConfig). 
	* `display_name` - A descriptive label for the question (for example, to display in a form in a graphical interface). Avoid entering confidential information. 
	* `explanation` - A description or explanation of the question, to help the customer answer accurately. 
	* `key` - A string that identifies the question. 
* `template` - A template of CPE device configuration information that will be merged with the customer's answers to the questions to render the final CPE device configuration content. Also see:
	* [GetCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/Cpe/GetCpeDeviceConfigContent)
	* [GetIpsecCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/IPSecConnection/GetIpsecCpeDeviceConfigContent)
	* [GetTunnelCpeDeviceConfigContent](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/20160918/TunnelCpeDeviceConfig/GetTunnelCpeDeviceConfigContent) 

