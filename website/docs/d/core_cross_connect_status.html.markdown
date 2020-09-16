---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_status"
sidebar_current: "docs-oci-datasource-core-cross_connect_status"
description: |-
  Provides details about a specific Cross Connect Status in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect_status
This data source provides details about a specific Cross Connect Status resource in Oracle Cloud Infrastructure Core service.

Gets the status of the specified cross-connect.


## Example Usage

```hcl
data "oci_core_cross_connect_status" "test_cross_connect_status" {
	#Required
	cross_connect_id = oci_core_cross_connect.test_cross_connect.id
}
```

## Argument Reference

The following arguments are supported:

* `cross_connect_id` - (Required) The OCID of the cross-connect.


## Attributes Reference

The following attributes are exported:

* `cross_connect_id` - The OCID of the cross-connect.
* `interface_state` - Whether Oracle's side of the interface is up or down.
* `light_level_ind_bm` - The light level of the cross-connect (in dBm).  Example: `14.0` 
* `light_level_indicator` - Status indicator corresponding to the light level.
	* **NO_LIGHT:** No measurable light
	* **LOW_WARN:** There's measurable light but it's too low
	* **HIGH_WARN:** Light level is too high
	* **BAD:** There's measurable light but the signal-to-noise ratio is bad
	* **GOOD:** Good light level 

