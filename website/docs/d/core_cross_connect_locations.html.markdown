---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_locations"
sidebar_current: "docs-oci-datasource-core-cross_connect_locations"
description: |-
  Provides the list of Cross Connect Locations in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect_locations
This data source provides the list of Cross Connect Locations in Oracle Cloud Infrastructure Core service.

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

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `cross_connect_locations` - The list of cross_connect_locations.

### CrossConnectLocation Reference

The following attributes are exported:

* `description` - A description of the location.
* `name` - The name of the location.  Example: `CyrusOne, Chandler, AZ` 

