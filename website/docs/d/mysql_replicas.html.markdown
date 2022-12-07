---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_replicas"
sidebar_current: "docs-oci-datasource-mysql-replicas"
description: |-
  Provides the list of Replicas in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_replicas
This data source provides the list of Replicas in Oracle Cloud Infrastructure MySQL Database service.

Lists all the read replicas that match the specified filters.

## Example Usage

```hcl
data "oci_mysql_replicas" "test_replicas" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	db_system_id = oci_mysql_mysql_db_system.test_mysql_db_system.id
	display_name = var.replica_display_name
	replica_id = oci_mysql_replica.test_replica.id
	state = var.replica_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Optional) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only the resource matching the given display name exactly.
* `replica_id` - (Optional) The read replica [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `state` - (Optional) The LifecycleState of the read replica.


## Attributes Reference

The following attributes are exported:

* `replicas` - The list of replicas.

### Replica Reference

The following attributes are exported:

* `availability_domain` - The name of the Availability Domain the read replica is located in.
* `compartment_id` - The OCID of the compartment that contains the read replica.
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
* `mysql_version` - The MySQL version used by the read replica.
* `port` - The port the read replica is configured to listen on.
* `port_x` - The TCP network port on which X Plugin listens for connections. This is the X Plugin equivalent of port. 
* `state` - The state of the read replica.
* `time_created` - The date and time the read replica was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The time the read replica was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

