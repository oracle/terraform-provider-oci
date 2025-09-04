---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_firmware_bundle"
sidebar_current: "docs-oci-datasource-core-firmware_bundle"
description: |-
  Provides details about a specific Firmware Bundle in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_firmware_bundle
This data source provides details about a specific Firmware Bundle resource in Oracle Cloud Infrastructure Core service.

Returns the Firmware Bundle matching the provided firmwareBundleId.


## Example Usage

```hcl
data "oci_core_firmware_bundle" "test_firmware_bundle" {
	#Required
	firmware_bundle_id = oci_core_firmware_bundle.test_firmware_bundle.id
}
```

## Argument Reference

The following arguments are supported:

* `firmware_bundle_id` - (Required) Unique identifier for the firmware bundle.


## Attributes Reference

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

