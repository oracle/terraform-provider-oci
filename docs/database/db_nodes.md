# oci_database_db_node

## DbNode Singular DataSource

### DbNode Reference

The following attributes are exported:

* `backup_vnic_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the backup VNIC.
* `db_system_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.
* `hostname` - The host name for the database node.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database node.
* `software_storage_size_in_gb` - The size (in GB) of the block storage volume allocation for the DB system. This attribute applies only for virtual machine DB systems. 
* `state` - The current state of the database node.
* `time_created` - The date and time that the database node was created.
* `vnic_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the VNIC.



### Get Operation
Gets information about the specified database node.

The following arguments are supported:

* `db_node_id` - (Required) The database node [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


### Example Usage

```hcl
data "oci_database_db_node" "test_db_node" {
	#Required
	db_node_id = "${var.db_node_db_node_id}"
}
```
# oci_database_db_nodes

## DbNode DataSource

Gets a list of db_nodes.

### List Operation
Gets a list of database nodes in the specified DB system and compartment. A database node is a server running database software.

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Required) The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.


The following attributes are exported:

* `db_nodes` - The list of db_nodes.

### Example Usage

```hcl
data "oci_database_db_nodes" "test_db_nodes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"
}
```
