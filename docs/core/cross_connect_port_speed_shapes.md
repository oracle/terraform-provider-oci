
# oci_core_cross_connect_port_speed_shapes

## CrossConnectPortSpeedShape DataSource

Gets a list of cross_connect_port_speed_shapes.

### List Operation
Lists the available port speeds for cross-connects. You need this information
so you can specify your desired port speed (that is, shape) when you create a
cross-connect.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


The following attributes are exported:

* `cross_connect_port_speed_shapes` - The list of cross_connect_port_speed_shapes.

### Example Usage

```hcl
data "oci_core_cross_connect_port_speed_shapes" "test_cross_connect_port_speed_shapes" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```
### CrossConnectPortSpeedShape Reference

The following attributes are exported:

* `name` - The name of the port speed shape.  Example: `10 Gbps` 
* `port_speed_in_gbps` - The port speed in Gbps.  Example: `10` 
