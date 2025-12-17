---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_firmware_bundles"
sidebar_current: "docs-oci-datasource-core-firmware_bundles"
description: |-
  Provides the list of Firmware Bundles in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_firmware_bundles
This data source provides the list of Firmware Bundles in Oracle Cloud Infrastructure Core service.

Gets a list of all Firmware Bundles in a compartment for specified platform. Can filter results to include 
only the default (recommended) Firmware Bundle for the given platform.


## Example Usage

```hcl
data "oci_core_firmware_bundles" "test_firmware_bundles" {
	#Required
	platform = var.firmware_bundle_platform

	#Optional
	compartment_id = var.compartment_id
	is_default_bundle = var.firmware_bundle_is_default_bundle
	state = var.firmware_bundle_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `is_default_bundle` - (Optional) If true, return only the default firmware bundle for a given platform. Default is false.
* `platform` - (Required) platform name
* `state` - (Optional) A filter to return only resources that match the given lifecycle state name exactly. 


## Attributes Reference

The following attributes are exported:

* `firmware_bundles_collection` - The list of firmware_bundles_collection.

### FirmwareBundle Reference

The following attributes are exported:

* `allowable_transitions` - A map of firmware bundle upgrades/downgrades validated by OCI.
	* `downgrades` - An array of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of validated firmware bundle downgrades.
	* `upgrades` - An array of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of validated firmware bundle upgrades.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment of this firmware bundle. 
* `description` - A brief description or metadata about this firmware bundle.
* `display_name` - The user-friendly name of this firmware bundle.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of this firmware bundle.
* `platforms` - A map of platforms to pinned firmware versions.
	* `platform` - The name of the platform supported by this bundle.
	* `versions` - An array of pinned components and their respective firmware versions.
		* `component_type` - The type of component.
		* `version` - A list of firmware versions associated with this component type.
* `state` - The current state of the firmware bundle.
* `time_created` - The date and time the firmware bundle was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the firmware bundle was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: `2016-08-25T21:10:29.600Z` 

