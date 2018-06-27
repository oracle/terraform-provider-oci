
# oci_database_db_system_shapes

## DbSystemShape DataSource

Gets a list of db_system_shapes.

### List Operation
Gets a list of the shapes that can be used to launch a new DB System. The shape determines resources to allocate to the DB system - CPU cores and memory for VM shapes; CPU cores, memory and storage for non-VM (or bare metal) shapes.
The following arguments are supported:

* `availability_domain` - (Required) The name of the Availability Domain.
* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


The following attributes are exported:

* `db_system_shapes` - The list of db_system_shapes.

### Example Usage

```hcl
data "oci_database_db_system_shapes" "test_db_system_shapes" {
	#Required
	availability_domain = "${var.db_system_shape_availability_domain}"
	compartment_id = "${var.compartment_id}"
}
```
### DbSystemShape Reference

The following attributes are exported:

* `available_core_count` - The maximum number of CPU cores that can be enabled on the DB System for this shape.
* `core_count_increment` - The discrete number by which the CPU core count for this shape can be increased or decreased.
* `maximum_node_count` - The maximum number of database nodes available for this shape.
* `minimum_core_count` - The minimum number of CPU cores that can be enabled on the DB System for this shape.
* `minimum_node_count` - The minimum number of database nodes available for this shape.
* `name` - The name of the shape used for the DB System.
* `shape` - Deprecated. Use `name` instead of `shape`.
