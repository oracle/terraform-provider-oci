---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_cross_connect_port_speed_shapes"
sidebar_current: "docs-oci-datasource-core-cross_connect_port_speed_shapes"
description: |-
  Provides the list of Cross Connect Port Speed Shapes in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_cross_connect_port_speed_shapes
This data source provides the list of Cross Connect Port Speed Shapes in Oracle Cloud Infrastructure Core service.

Lists the available port speeds for cross-connects. You need this information
so you can specify your desired port speed (that is, shape) when you create a
cross-connect.


## Example Usage

```hcl
data "oci_core_cross_connect_port_speed_shapes" "test_cross_connect_port_speed_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `cross_connect_port_speed_shapes` - The list of cross_connect_port_speed_shapes.

### CrossConnectPortSpeedShape Reference

The following attributes are exported:

* `name` - The name of the port speed shape.  Example: `10 Gbps` 
* `port_speed_in_gbps` - The port speed in Gbps.  Example: `10` 

