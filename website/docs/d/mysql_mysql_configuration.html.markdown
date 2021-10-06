---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_configuration"
sidebar_current: "docs-oci-datasource-mysql-mysql_configuration"
description: |-
  Provides details about a specific Mysql Configuration in Oracle Cloud Infrastructure MySQL Database service
---

# Data Source: oci_mysql_mysql_configuration
This data source provides details about a specific Mysql Configuration resource in Oracle Cloud Infrastructure MySQL Database service.

Get the full details of the specified Configuration, including the list of MySQL Variables and their values.


## Example Usage

```hcl
data "oci_mysql_mysql_configuration" "test_mysql_configuration" {
	#Required
	configuration_id = var.mysql_configuration_id
}
```

## Argument Reference

The following arguments are supported:

* `configuration_id` - (Required) The OCID of the Configuration.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of the Compartment the Configuration exists in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User-provided data about the Configuration.
* `display_name` - The display name of the Configuration.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Configuration.
* `parent_configuration_id` - The OCID of the Configuration from which this Configuration is "derived". This is entirely a metadata relationship. There is no relation between the values in this Configuration and its parent. 
* `shape_name` - The name of the associated Shape.
* `state` - The current state of the Configuration.
* `time_created` - The date and time the Configuration was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `time_updated` - The date and time the Configuration was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `type` - The Configuration type, DEFAULT or CUSTOM.
* `variables` - User controllable service variables.
	* `autocommit` - ("autocommit")
	* `binlog_expire_logs_seconds` - Sets the binary log expiration period in seconds. binlogExpireLogsSeconds corresponds to the MySQL binary logging system variable [binlog_expire_logs_seconds](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_expire_logs_seconds). 
	* `binlog_row_metadata` - Configures the amount of table metadata added to the binary log when using row-based logging. binlogRowMetadata corresponds to the MySQL binary logging system variable [binlog_row_metadata](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_metadata). 
	* `binlog_row_value_options` - When set to PARTIAL_JSON, this enables use of a space-efficient binary log format for updates that modify only a small portion of a JSON document. binlogRowValueOptions corresponds to the MySQL binary logging system variable [binlog_row_value_options](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_value_options). 
	* `binlog_transaction_compression` - Enables compression for transactions that are written to binary log files on this server. binlogTransactionCompression corresponds to the MySQL binary logging system variable [binlog_transaction_compression](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_transaction_compression). 
	* `completion_type` - ("completion_type")
	* `connect_timeout` - ("connect_timeout")
	* `cte_max_recursion_depth` - ("cte_max_recursion_depth")
	* `default_authentication_plugin` - ("default_authentication_plugin")
	* `foreign_key_checks` - ("foreign_key_checks")
	* `generated_random_password_length` - ("generated_random_password_length") DEPRECATED -- variable should not be settable and will be ignored
	* `group_replication_consistency` - 
		* EVENTUAL: Both RO and RW transactions do not wait for preceding transactions to be applied before executing. A RW transaction does not wait for other members to apply a transaction. This means that a transaction could be externalized on one member before the others. This also means that in the event of a primary failover, the new primary can accept new RO and RW transactions before the previous primary transactions are all applied. RO transactions could result in outdated values, RW transactions could result in a rollback due to conflicts.
		* BEFORE_ON_PRIMARY_FAILOVER: New RO or RW transactions with a newly elected primary that is applying backlog from the old primary are held (not applied) until any backlog has been applied. This ensures that when a primary failover happens, intentionally or not, clients always see the latest value on the primary. This guarantees consistency, but means that clients must be able to handle the delay in the event that a backlog is being applied. Usually this delay should be minimal, but does depend on the size of the backlog.
		* BEFORE: A RW transaction waits for all preceding transactions to complete before being applied. A RO transaction waits for all preceding transactions to complete before being executed. This ensures that this transaction reads the latest value by only affecting the latency of the transaction. This reduces the overhead of synchronization on every RW transaction, by ensuring synchronization is used only on RO transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* AFTER: A RW transaction waits until its changes have been applied to all of the other members. This value has no effect on RO transactions. This mode ensures that when a transaction is committed on the local member, any subsequent transaction reads the written value or a more recent value on any group member. Use this mode with a group that is used for predominantly RO operations to ensure that applied RW transactions are applied everywhere once they commit. This could be used by your application to ensure that subsequent reads fetch the latest data which includes the latest writes. This reduces the overhead of synchronization on every RO transaction, by ensuring synchronization is used only on RW transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* BEFORE_AND_AFTER: A RW transaction waits for 1) all preceding transactions to complete before being applied and 2) until its changes have been applied on other members. A RO transaction waits for all preceding transactions to complete before execution takes place. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER. 
	* `information_schema_stats_expiry` - ("information_schema_stats_expiry")
	* `innodb_buffer_pool_instances` - ("innodb_buffer_pool_instances")
	* `innodb_buffer_pool_size` - ("innodb_buffer_pool_size")
	* `innodb_ft_enable_stopword` - ("innodb_ft_enable_stopword")
	* `innodb_ft_max_token_size` - ("innodb_ft_max_token_size")
	* `innodb_ft_min_token_size` - ("innodb_ft_min_token_size")
	* `innodb_ft_num_word_optimize` - ("innodb_ft_num_word_optimize")
	* `innodb_ft_result_cache_limit` - ("innodb_ft_result_cache_limit")
	* `innodb_ft_server_stopword_table` - ("innodb_ft_server_stopword_table")
	* `innodb_lock_wait_timeout` - ("innodb_lock_wait_timeout")
	* `innodb_max_purge_lag` - ("innodb_max_purge_lag")
	* `innodb_max_purge_lag_delay` - ("innodb_max_purge_lag_delay")
	* `local_infile` - ("local_infile")
	* `mandatory_roles` - ("mandatory_roles")
	* `max_connections` - ("max_connections")
	* `max_execution_time` - ("max_execution_time")
	* `max_prepared_stmt_count` - ("max_prepared_stmt_count")
	* `mysql_firewall_mode` - ("mysql_firewall_mode")
	* `mysql_zstd_default_compression_level` - DEPRECATED -- typo of mysqlx_zstd_default_compression_level. variable will be ignored.
	* `mysqlx_connect_timeout` - ("mysqlx_connect_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_deflate_default_compression_level` - Set the default compression level for the deflate algorithm. ("mysqlx_deflate_default_compression_level")
	* `mysqlx_deflate_max_client_compression_level` - Limit the upper bound of accepted compression levels for the deflate algorithm. ("mysqlx_deflate_max_client_compression_level")
	* `mysqlx_document_id_unique_prefix` - ("mysqlx_document_id_unique_prefix") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_enable_hello_notice` - ("mysqlx_enable_hello_notice") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_idle_worker_thread_timeout` - ("mysqlx_idle_worker_thread_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_interactive_timeout` - ("mysqlx_interactive_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_lz4default_compression_level` - Set the default compression level for the lz4 algorithm. ("mysqlx_lz4_default_compression_level")
	* `mysqlx_lz4max_client_compression_level` - Limit the upper bound of accepted compression levels for the lz4 algorithm. ("mysqlx_lz4_max_client_compression_level")
	* `mysqlx_max_allowed_packet` - ("mysqlx_max_allowed_packet") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_min_worker_threads` - ("mysqlx_min_worker_threads") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_read_timeout` - ("mysqlx_read_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_wait_timeout` - ("mysqlx_wait_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_write_timeout` - ("mysqlx_write_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_zstd_default_compression_level` - Set the default compression level for the zstd algorithm. ("mysqlx_zstd_default_compression_level")
	* `mysqlx_zstd_max_client_compression_level` - Limit the upper bound of accepted compression levels for the zstd algorithm. ("mysqlx_zstd_max_client_compression_level")
	* `parser_max_mem_size` - ("parser_max_mem_size")
	* `query_alloc_block_size` - ("query_alloc_block_size") DEPRECATED -- variable should not be settable and will be ignored
	* `query_prealloc_size` - ("query_prealloc_size") DEPRECATED -- variable should not be settable and will be ignored
	* `sql_mode` - ("sql_mode")
	* `sql_require_primary_key` - ("sql_require_primary_key")
	* `sql_warnings` - ("sql_warnings")
	* `transaction_isolation` - ("transaction_isolation")

