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

get single cpeDeviceShape object and list of the questions need to asked for that single cpeDeviceShape.


## Example Usage

```hcl
data "oci_core_cpe_device_shape" "test_cpe_device_shape" {
	#Required
	cpe_device_shape_id = "${oci_core_cpe_device_shape.test_cpe_device_shape.id}"
}
```

## Argument Reference

The following arguments are supported:

* `cpe_device_shape_id` - (Required) The OCID of the CPE device shape.


## Attributes Reference

The following attributes are exported:

* `cpe_device_info` - customer premise equipment hardware information
	* `platform_software_version` - The CPE's vendor/platform version
	* `vendor` - The CPE's hardware information
* `cpe_device_shape_id` - The CPE device type's unique identifier.
* `parameters` - list of questions to ask to cusomter regarding their cpe device in order to generate their cpe device config
	* `display_name` - 
	* `explanation` - 
	* `key` - 
* `template` - the template that will be combined together with customer input to render customer cpe device configuration

