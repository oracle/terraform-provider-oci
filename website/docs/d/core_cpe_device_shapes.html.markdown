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

Lists the customer-premises equipment objects (CPEs)'s hardware information in the specified compartment.


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

* `cpe_device_info` - customer premise equipment hardware information
	* `platform_software_version` - The CPE's vendor/platform version
	* `vendor` - The CPE's hardware information
* `cpe_device_shape_id` - The CPE device type's unique identifier.
* `parameters` - list of questions to ask to cusomter regarding their cpe device in order to generate their cpe device config
	* `display_name` - 
	* `explanation` - 
	* `key` - 
* `template` - the template that will be combined together with customer input to render customer cpe device configuration

