---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_channels"
sidebar_current: "docs-oci-datasource-mysql-channels"
description: |-
  Provides the list of Channels in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_channels
This data source provides the list of Channels in Oracle Cloud Infrastructure MySQL Database service.

Lists all the Channels that match the specified filters.

## Example Usage

```hcl
data "oci_mysql_channels" "test_channels" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	channel_id = oci_mysql_channel.test_channel.id
	db_system_id = oci_database_db_system.test_db_system.id
	display_name = var.channel_display_name
	is_enabled = var.channel_is_enabled
	state = var.channel_state
}
```

## Argument Reference

The following arguments are supported:

* `channel_id` - (Optional) The OCID of the Channel.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Optional) The DB System [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only the resource matching the given display name exactly.
* `is_enabled` - (Optional) If true, returns only Channels that are enabled. If false, returns only Channels that are disabled. 
* `state` - (Optional) The LifecycleState of the Channel.


## Attributes Reference

The following attributes are exported:

* `channels` - The list of channels.

### Channel Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User provided description of the Channel.
* `display_name` - The user-friendly name for the Channel. It does not have to be unique.
* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Channel.
* `is_enabled` - Whether the Channel has been enabled by the user.
* `lifecycle_details` - A message describing the state of the Channel.
* `source` - Parameters detailing how to provision the source for the given Channel.
	* `anonymous_transactions_handling` - Specifies how the replication channel handles replicated transactions without an identifier, enabling replication from a source that does not use transaction-id-based replication to a replica that does. 
		* `last_configured_log_filename` - Specifies one of the coordinates (file) at which the replica should begin reading the source's log. As this value specifies the point where replication starts from, it is only used once, when it starts. It is never used again, unless a new UpdateChannel operation modifies it. 
		* `last_configured_log_offset` - Specifies one of the coordinates (offset) at which the replica should begin reading the source's log. As this value specifies the point where replication starts from, it is only used once, when it starts. It is never used again, unless a new UpdateChannel operation modifies it. 
		* `policy` - Specifies how the replication channel handles anonymous transactions.
		* `uuid` - The UUID that is used as a prefix when generating transaction identifiers for anonymous transactions coming from the source. You can change the UUID later. 
	* `hostname` - The network address of the MySQL instance.
	* `port` - The port the source MySQL instance listens on.
	* `source_type` - The specific source identifier.
	* `ssl_ca_certificate` - The CA certificate of the server used for VERIFY_IDENTITY and VERIFY_CA ssl modes.
		* `certificate_type` - The type of CA certificate.
		* `contents` - The string containing the CA certificate in PEM format.
	* `ssl_mode` - The SSL mode of the Channel.
	* `username` - The name of the replication user on the source MySQL instance. The username has a maximum length of 96 characters. For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html) 
* `state` - The state of the Channel.
* `target` - Details about the Channel target.
	* `applier_username` - The username for the replication applier of the target MySQL DB System.
	* `channel_name` - The case-insensitive name that identifies the replication channel. Channel names must follow the rules defined for [MySQL identifiers](https://dev.mysql.com/doc/refman/8.0/en/identifiers.html). The names of non-Deleted Channels must be unique for each DB System. 
	* `db_system_id` - The OCID of the source DB System.
	* `filters` - Replication filter rules to be applied at the DB System Channel target. 
		* `type` - The type of the filter rule.

			For details on each type, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html) 
		* `value` - The body of the filter rule. This can represent a database, a table, or a database pair (represented as "db1->db2"). For more information, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html). 
	* `target_type` - The specific target identifier.
* `time_created` - The date and time the Channel was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The time the Channel was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

