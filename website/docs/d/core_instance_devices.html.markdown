---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_devices"
sidebar_current: "docs-oci-datasource-core-instance_devices"
description: |-
  Provides the list of Instance Devices in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_devices
This data source provides the list of Instance Devices in Oracle Cloud Infrastructure Core service.

Gets a list of all the devices for given instance. You can optionally filter results by device availability.

## Example Usage

```hcl
data "oci_core_instance_devices" "test_instance_devices" {
	#Required
	instance_id = oci_core_instance.test_instance.id

	#Optional
	is_available = var.instance_device_is_available
	name = var.instance_device_name
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
* `is_available` - (Optional) A filter to return only available devices or only used devices. 
* `name` - (Optional) A filter to return only devices that match the given name exactly. 


## Attributes Reference

The following attributes are exported:

* `devices` - The list of devices.

### InstanceDevice Reference

The following attributes are exported:

* `is_available` - The flag denoting whether device is available.
* `name` - The device name.

