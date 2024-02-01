---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_replica"
sidebar_current: "docs-oci-resource-mysql-replica"
description: |-
  Provides the Replica resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_replica
This resource provides the Replica resource in Oracle Cloud Infrastructure MySQL Database service.

Creates a DB System read replica.

## Example Usage

```hcl
resource "oci_mysql_replica" "test_replica" {
	#Required
	db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.replica_description
	display_name = var.replica_display_name
	freeform_tags = {"bar-key"= "value"}
	is_delete_protected = var.replica_is_delete_protected
	replica_overrides {

		#Optional
		configuration_id = oci_mysql_mysql_configuration.test_mysql_configuration.id
		mysql_version = var.replica_replica_overrides_mysql_version
		shape_name = oci_mysql_shape.test_shape.name
	}
}
```

## Argument Reference

The following arguments are supported:

* `db_system_id` - (Required) The OCID of the DB System the read replica is associated with.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) User provided description of the read replica.
* `display_name` - (Optional) (Updatable) The user-friendly name for the read replica. It does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_delete_protected` - (Optional) (Updatable) Specifies whether the read replica can be deleted. Set to true to prevent deletion, false (default) to allow. Note that if a read replica is delete protected it also prevents the entire DB System from being deleted. If the DB System is delete protected, read replicas can still be deleted individually if they are not delete  protected themselves. 
* `replica_overrides` - (Optional) (Updatable) By default a read replica inherits the MySQL version, shape, and configuration of the source DB system.  If you want to override any of these, provide values in the properties, mysqlVersion, shapeName,  and configurationId. If you set a property value to "", then the value is inherited from its  source DB system. 
	* `configuration_id` - (Optional) (Updatable) The OCID of the Configuration to be used by the read replica.
	* `mysql_version` - (Optional) (Updatable) The MySQL version to be used by the read replica.
	* `shape_name` - (Optional) (Updatable) The shape to be used by the read replica. The shape determines the resources allocated:  CPU cores and memory for VM shapes, CPU cores, memory and storage for non-VM (bare metal) shapes.  To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20190415/ShapeSummary/ListShapes) operation. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the Availability Domain the read replica is located in.
* `compartment_id` - The OCID of the compartment that contains the read replica.
* `configuration_id` - The OCID of the Configuration currently in use by the read replica.
* `db_system_id` - The OCID of the DB System the read replica is associated with.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User provided description of the read replica.
* `display_name` - The user-friendly name for the read replica. It does not have to be unique.
* `fault_domain` - The name of the Fault Domain the read replica is located in.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the read replica.
* `ip_address` - The IP address the read replica is configured to listen on. 
* `is_delete_protected` - Specifies whether the read replica can be deleted. Set to true to prevent deletion, false (default) to allow. Note that if a read replica is delete protected it also prevents the entire DB System from being deleted. If the DB System is delete protected, read replicas can still be deleted individually if they are not delete  protected themselves. 
* `lifecycle_details` - A message describing the state of the read replica.
* `mysql_version` - The MySQL version currently in use by the read replica.
* `port` - The port the read replica is configured to listen on.
* `port_x` - The TCP network port on which X Plugin listens for connections. This is the X Plugin equivalent of port. 
* `replica_overrides` - By default a read replica inherits the MySQL version, shape, and configuration of the source DB system.  If you want to override any of these, provide values in the properties, mysqlVersion, shapeName,  and configurationId. If you set a property value to "", then the value is inherited from its  source DB system. 
	* `configuration_id` - The OCID of the Configuration to be used by the read replica.
	* `mysql_version` - The MySQL version to be used by the read replica.
	* `shape_name` - The shape to be used by the read replica. The shape determines the resources allocated:  CPU cores and memory for VM shapes, CPU cores, memory and storage for non-VM (bare metal) shapes.  To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20190415/ShapeSummary/ListShapes) operation. 
* `secure_connections` - Secure connection configuration details. 
	* `certificate_generation_type` - Select whether to use MySQL Database Service-managed certificate (SYSTEM) or your own certificate (BYOC). 
	* `certificate_id` - The OCID of the certificate to use.
* `shape_name` - The shape currently in use by the read replica. The shape determines the resources allocated:  CPU cores and memory for VM shapes, CPU cores, memory and storage for non-VM (bare metal) shapes.  To get a list of shapes, use the [ListShapes](https://docs.cloud.oracle.com/iaas/api/#/en/mysql/20190415/ShapeSummary/ListShapes) operation. 
* `state` - The state of the read replica.
* `time_created` - The date and time the read replica was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The time the read replica was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Replica
	* `update` - (Defaults to 20 minutes), when updating the Replica
	* `delete` - (Defaults to 20 minutes), when destroying the Replica


## Import

Replicas can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_replica.test_replica "id"
```

