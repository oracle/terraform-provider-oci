---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_data_safe_configuration"
sidebar_current: "docs-oci-resource-data_safe-data_safe_configuration"
description: |-
  Provides the Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_data_safe_configuration
This resource provides the Data Safe Configuration resource in Oracle Cloud Infrastructure Data Safe service.

Enables Data Safe in the tenancy and region.


## Example Usage

```hcl
resource "oci_data_safe_data_safe_configuration" "test_data_safe_configuration" {

	#Optional
	compartment_id = "${var.compartment_id}"
	is_enabled = "${var.data_safe_configuration_is_enabled}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) A filter to return only resources that match the specified compartment OCID.
* `is_enabled` - (Optional) (Updatable) Indicates if Data Safe is enabled.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the tenancy used to enable Data Safe.
* `is_enabled` - Indicates if Data Safe is enabled.
* `state` - The current state of Data Safe configuration.
* `time_enabled` - The specific time when Data Safe configuration was enabled.
* `url` - The URL of the Data Safe service.

## Import

Import is not supported for this resource.

