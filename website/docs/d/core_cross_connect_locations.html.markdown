---
layout: "oci"
page_title: "OCI: oci_core_cross_connect_locations"
sidebar_current: "docs-oci-datasource-core-cross_connect_locations"
description: |-
  Provides a list of CrossConnectLocations
---

# Data Source: oci_core_cross_connect_locations
The `oci_core_cross_connect_locations` data source allows access to the list of OCI cross_connect_locations

Lists the available FastConnect locations for cross-connect installation. You need
this information so you can specify your desired location when you create a cross-connect.


## Example Usage

```hcl
data "oci_core_cross_connect_locations" "test_cross_connect_locations" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


## Attributes Reference

The following attributes are exported:

* `cross_connect_locations` - The list of cross_connect_locations.

### CrossConnectLocation Reference

The following attributes are exported:

* `description` - A description of the location.
* `name` - The name of the location.  Example: `CyrusOne, Chandler, AZ` 

