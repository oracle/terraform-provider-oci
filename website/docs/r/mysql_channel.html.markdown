---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_channel"
sidebar_current: "docs-oci-resource-mysql-channel"
description: |-
  Provides the Channel resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_channel
This resource provides the Channel resource in Oracle Cloud Infrastructure MySQL Database service.

Creates a Channel to establish replication from a source to a target.


## Example Usage

```hcl
resource "oci_mysql_channel" "test_channel" {
	#Required
	source {
		#Required
		hostname = var.channel_source_hostname
		password = var.channel_source_password
		source_type = var.channel_source_source_type
		ssl_mode = var.channel_source_ssl_mode
		username = var.channel_source_username

		#Optional
		anonymous_transactions_handling {
			#Required
			policy = var.channel_source_anonymous_transactions_handling_policy

			#Optional
			last_configured_log_filename = var.channel_source_anonymous_transactions_handling_last_configured_log_filename
			last_configured_log_offset = var.channel_source_anonymous_transactions_handling_last_configured_log_offset
			uuid = var.channel_source_anonymous_transactions_handling_uuid
		}
		port = var.channel_source_port
		ssl_ca_certificate {
			#Required
			certificate_type = var.channel_source_ssl_ca_certificate_certificate_type
			contents = var.channel_source_ssl_ca_certificate_contents
		}
	}
	target {
		#Required
		db_system_id = oci_database_db_system.test_db_system.id
		target_type = var.channel_target_target_type

		#Optional
		applier_username = var.channel_target_applier_username
		channel_name = oci_mysql_channel.test_channel.name
		delay_in_seconds = var.channel_target_delay_in_seconds
		filters {
			#Required
			type = var.channel_target_filters_type
			value = var.channel_target_filters_value
		}
		tables_without_primary_key_handling = var.channel_target_tables_without_primary_key_handling
	}

	#Optional
	compartment_id = var.compartment_id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.channel_description
	display_name = var.channel_display_name
	freeform_tags = {"bar-key"= "value"}
	is_enabled = var.channel_is_enabled
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) User provided information about the Channel.
* `display_name` - (Optional) (Updatable) The user-friendly name for the Channel. It does not have to be unique.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_enabled` - (Optional) (Updatable) Whether the Channel should be enabled upon creation. If set to true, the Channel will be asynchronously started as a result of the create Channel operation. 
* `source` - (Required) (Updatable) Parameters detailing how to provision the source for the given Channel.
	* `anonymous_transactions_handling` - (Optional) (Updatable) Specifies how the replication channel handles replicated transactions without an identifier, enabling replication from a source that does not use transaction-id-based replication to a replica that does. 
		* `last_configured_log_filename` - (Applicable when policy=ASSIGN_MANUAL_UUID | ASSIGN_TARGET_UUID) (Updatable) Specifies one of the coordinates (file) at which the replica should begin reading the source's log. As this value specifies the point where replication starts from, it is only used once, when it starts. It is never used again, unless a new UpdateChannel operation modifies it. 
		* `last_configured_log_offset` - (Applicable when policy=ASSIGN_MANUAL_UUID | ASSIGN_TARGET_UUID) (Updatable) Specifies one of the coordinates (offset) at which the replica should begin reading the source's log. As this value specifies the point where replication starts from, it is only used once, when it starts. It is never used again, unless a new UpdateChannel operation modifies it. 
		* `policy` - (Required) (Updatable) Specifies how the replication channel handles anonymous transactions.
		* `uuid` - (Applicable when policy=ASSIGN_MANUAL_UUID) (Updatable) The UUID that is used as a prefix when generating transaction identifiers for anonymous transactions coming from the source. You can change the UUID later. 
	* `hostname` - (Required) (Updatable) The network address of the MySQL instance.
	* `password` - (Required) (Updatable) The password for the replication user. The password must be between 8 and 32 characters long, and must contain at least 1 numeric character, 1 lowercase character, 1 uppercase character, and 1 special (nonalphanumeric) character. 
	* `port` - (Optional) (Updatable) The port the source MySQL instance listens on.
	* `source_type` - (Required) (Updatable) The specific source identifier.
	* `ssl_ca_certificate` - (Optional) (Updatable) The CA certificate of the server used for VERIFY_IDENTITY and VERIFY_CA ssl modes.
		* `certificate_type` - (Required) (Updatable) The type of CA certificate.
		* `contents` - (Required) (Updatable) The string containing the CA certificate in PEM format.
	* `ssl_mode` - (Required) (Updatable) The SSL mode of the Channel.
	* `username` - (Required) (Updatable) The name of the replication user on the source MySQL instance. The username has a maximum length of 96 characters. For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/change-master-to.html) 
* `target` - (Required) (Updatable) Parameters detailing how to provision the target for the given Channel.
	* `applier_username` - (Optional) (Updatable) The username for the replication applier of the target MySQL DB System.
	* `channel_name` - (Optional) (Updatable) The case-insensitive name that identifies the replication channel. Channel names must follow the rules defined for [MySQL identifiers](https://dev.mysql.com/doc/refman/8.0/en/identifiers.html). The names of non-Deleted Channels must be unique for each DB System. 
	* `db_system_id` - (Required) The OCID of the target DB System.
	* `delay_in_seconds` - (Optional) (Updatable) Specifies the amount of time, in seconds, that the channel waits before  applying a transaction received from the source. 
	* `filters` - (Optional) (Updatable) Replication filter rules to be applied at the DB System Channel target. 
		* `type` - (Required) (Updatable) The type of the filter rule.

			For details on each type, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html) 
		* `value` - (Required) (Updatable) The body of the filter rule. This can represent a database, a table, or a database pair (represented as "db1->db2"). For more information, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html). 
	* `tables_without_primary_key_handling` - (Optional) (Updatable) Specifies how a replication channel handles the creation and alteration of tables  that do not have a primary key. The default value is set to ALLOW. 
	* `target_type` - (Required) (Updatable) The specific target identifier.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User provided description of the Channel.
* `display_name` - The user-friendly name for the Channel. It does not have to be unique.
* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
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
	* `delay_in_seconds` - Specifies the amount of time, in seconds, that the channel waits before  applying a transaction received from the source. 
	* `filters` - Replication filter rules to be applied at the DB System Channel target. 
		* `type` - The type of the filter rule.

			For details on each type, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html) 
		* `value` - The body of the filter rule. This can represent a database, a table, or a database pair (represented as "db1->db2"). For more information, see [Replication Filtering Rules](https://dev.mysql.com/doc/refman/8.0/en/replication-rules.html). 
	* `tables_without_primary_key_handling` - Specifies how a replication channel handles the creation and alteration of tables  that do not have a primary key. 
	* `target_type` - The specific target identifier.
* `time_created` - The date and time the Channel was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 
* `time_updated` - The time the Channel was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 30 minutes), when creating the Channel
	* `update` - (Defaults to 30 minutes), when updating the Channel
	* `delete` - (Defaults to 30 minutes), when destroying the Channel


## Import

Channels can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_channel.test_channel "id"
```

