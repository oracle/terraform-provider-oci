
# oci_core_cross_connect_locations

## CrossConnectLocation DataSource

Gets a list of cross_connect_locations.

### List Operation
Lists the available FastConnect locations for cross-connect installation. You need
this information so you can specify your desired location when you create a cross-connect.

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.


The following attributes are exported:

* `cross_connect_locations` - The list of cross_connect_locations.

### Example Usage

```hcl
data "oci_core_cross_connect_locations" "test_cross_connect_locations" {
	#Required
	compartment_id = "${var.compartment_id}"
}
```
### CrossConnectLocation Reference

The following attributes are exported:

* `description` - A description of the location.
* `name` - The name of the location.  Example: `CyrusOne, Chandler, AZ` 
