---
subcategory: "MySQL Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_configuration"
sidebar_current: "docs-oci-resource-mysql-mysql_configuration"
description: |-
  Provides the Mysql Configuration resource in Oracle Cloud Infrastructure MySQL Database service
---

# oci_mysql_mysql_configuration
This resource provides the Mysql Configuration resource in Oracle Cloud Infrastructure MySQL Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/mysql/latest/MysqlConfiguration

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/mysql

Creates a new Configuration.

## Example Usage

```hcl
resource "oci_mysql_mysql_configuration" "test_mysql_configuration" {
	#Required
	compartment_id = var.compartment_id
	shape_name = oci_mysql_shape.test_shape.name

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.mysql_configuration_description
	display_name = var.mysql_configuration_display_name
	freeform_tags = {"bar-key"= "value"}
	init_variables {

		#Optional
		lower_case_table_names = var.mysql_configuration_init_variables_lower_case_table_names
	}
	parent_configuration_id = oci_audit_configuration.test_configuration.id
	variables {

		#Optional
		auto_increment_increment = var.mysql_configuration_variables_auto_increment_increment
		auto_increment_offset = var.mysql_configuration_variables_auto_increment_offset
		autocommit = var.mysql_configuration_variables_autocommit
		big_tables = var.mysql_configuration_variables_big_tables
		binlog_expire_logs_seconds = var.mysql_configuration_variables_binlog_expire_logs_seconds
		binlog_group_commit_sync_delay = var.mysql_configuration_variables_binlog_group_commit_sync_delay
		binlog_group_commit_sync_no_delay_count = var.mysql_configuration_variables_binlog_group_commit_sync_no_delay_count
		binlog_row_metadata = var.mysql_configuration_variables_binlog_row_metadata
		binlog_row_value_options = var.mysql_configuration_variables_binlog_row_value_options
		binlog_transaction_compression = var.mysql_configuration_variables_binlog_transaction_compression
		block_encryption_mode = var.mysql_configuration_variables_block_encryption_mode
		character_set_server = var.mysql_configuration_variables_character_set_server
		collation_server = var.mysql_configuration_variables_collation_server
		completion_type = var.mysql_configuration_variables_completion_type
		connect_timeout = var.mysql_configuration_variables_connect_timeout
		connection_memory_chunk_size = var.mysql_configuration_variables_connection_memory_chunk_size
		connection_memory_limit = var.mysql_configuration_variables_connection_memory_limit
		cte_max_recursion_depth = var.mysql_configuration_variables_cte_max_recursion_depth
		default_authentication_plugin = var.mysql_configuration_variables_default_authentication_plugin
		explain_format = var.mysql_configuration_variables_explain_format
		explicit_defaults_for_timestamp = var.mysql_configuration_variables_explicit_defaults_for_timestamp
		foreign_key_checks = var.mysql_configuration_variables_foreign_key_checks
		generated_random_password_length = var.mysql_configuration_variables_generated_random_password_length
		global_connection_memory_limit = var.mysql_configuration_variables_global_connection_memory_limit
		global_connection_memory_tracking = var.mysql_configuration_variables_global_connection_memory_tracking
		group_concat_max_len = var.mysql_configuration_variables_group_concat_max_len
		group_replication_consistency = var.mysql_configuration_variables_group_replication_consistency
		information_schema_stats_expiry = var.mysql_configuration_variables_information_schema_stats_expiry
		innodb_adaptive_hash_index = var.mysql_configuration_variables_innodb_adaptive_hash_index
		innodb_autoinc_lock_mode = var.mysql_configuration_variables_innodb_autoinc_lock_mode
		innodb_buffer_pool_dump_pct = var.mysql_configuration_variables_innodb_buffer_pool_dump_pct
		innodb_buffer_pool_instances = var.mysql_configuration_variables_innodb_buffer_pool_instances
		innodb_buffer_pool_size = var.mysql_configuration_variables_innodb_buffer_pool_size
		innodb_change_buffering = var.mysql_configuration_variables_innodb_change_buffering
		innodb_ddl_buffer_size = var.mysql_configuration_variables_innodb_ddl_buffer_size
		innodb_ddl_threads = var.mysql_configuration_variables_innodb_ddl_threads
		innodb_ft_enable_stopword = var.mysql_configuration_variables_innodb_ft_enable_stopword
		innodb_ft_max_token_size = var.mysql_configuration_variables_innodb_ft_max_token_size
		innodb_ft_min_token_size = var.mysql_configuration_variables_innodb_ft_min_token_size
		innodb_ft_num_word_optimize = var.mysql_configuration_variables_innodb_ft_num_word_optimize
		innodb_ft_result_cache_limit = var.mysql_configuration_variables_innodb_ft_result_cache_limit
		innodb_ft_server_stopword_table = var.mysql_configuration_variables_innodb_ft_server_stopword_table
		innodb_lock_wait_timeout = var.mysql_configuration_variables_innodb_lock_wait_timeout
		innodb_log_writer_threads = var.mysql_configuration_variables_innodb_log_writer_threads
		innodb_max_purge_lag = var.mysql_configuration_variables_innodb_max_purge_lag
		innodb_max_purge_lag_delay = var.mysql_configuration_variables_innodb_max_purge_lag_delay
		innodb_numa_interleave = var.mysql_configuration_variables_innodb_numa_interleave
		innodb_online_alter_log_max_size = var.mysql_configuration_variables_innodb_online_alter_log_max_size
		innodb_redo_log_capacity = var.mysql_configuration_variables_innodb_redo_log_capacity
		innodb_rollback_on_timeout = var.mysql_configuration_variables_innodb_rollback_on_timeout
		innodb_sort_buffer_size = var.mysql_configuration_variables_innodb_sort_buffer_size
		innodb_stats_persistent_sample_pages = var.mysql_configuration_variables_innodb_stats_persistent_sample_pages
		innodb_stats_transient_sample_pages = var.mysql_configuration_variables_innodb_stats_transient_sample_pages
		innodb_strict_mode = var.mysql_configuration_variables_innodb_strict_mode
		innodb_undo_log_truncate = var.mysql_configuration_variables_innodb_undo_log_truncate
		interactive_timeout = var.mysql_configuration_variables_interactive_timeout
		join_buffer_size = var.mysql_configuration_variables_join_buffer_size
		local_infile = var.mysql_configuration_variables_local_infile
		long_query_time = var.mysql_configuration_variables_long_query_time
		mandatory_roles = var.mysql_configuration_variables_mandatory_roles
		max_allowed_packet = var.mysql_configuration_variables_max_allowed_packet
		max_binlog_cache_size = var.mysql_configuration_variables_max_binlog_cache_size
		max_connect_errors = var.mysql_configuration_variables_max_connect_errors
		max_connections = var.mysql_configuration_variables_max_connections
		max_execution_time = var.mysql_configuration_variables_max_execution_time
		max_heap_table_size = var.mysql_configuration_variables_max_heap_table_size
		max_prepared_stmt_count = var.mysql_configuration_variables_max_prepared_stmt_count
		max_seeks_for_key = var.mysql_configuration_variables_max_seeks_for_key
		max_user_connections = var.mysql_configuration_variables_max_user_connections
		mysql_firewall_mode = var.mysql_configuration_variables_mysql_firewall_mode
		mysql_zstd_default_compression_level = var.mysql_configuration_variables_mysql_zstd_default_compression_level
		mysqlx_connect_timeout = var.mysql_configuration_variables_mysqlx_connect_timeout
		mysqlx_deflate_default_compression_level = var.mysql_configuration_variables_mysqlx_deflate_default_compression_level
		mysqlx_deflate_max_client_compression_level = var.mysql_configuration_variables_mysqlx_deflate_max_client_compression_level
		mysqlx_document_id_unique_prefix = var.mysql_configuration_variables_mysqlx_document_id_unique_prefix
		mysqlx_enable_hello_notice = var.mysql_configuration_variables_mysqlx_enable_hello_notice
		mysqlx_idle_worker_thread_timeout = var.mysql_configuration_variables_mysqlx_idle_worker_thread_timeout
		mysqlx_interactive_timeout = var.mysql_configuration_variables_mysqlx_interactive_timeout
		mysqlx_lz4default_compression_level = var.mysql_configuration_variables_mysqlx_lz4default_compression_level
		mysqlx_lz4max_client_compression_level = var.mysql_configuration_variables_mysqlx_lz4max_client_compression_level
		mysqlx_max_allowed_packet = var.mysql_configuration_variables_mysqlx_max_allowed_packet
		mysqlx_min_worker_threads = var.mysql_configuration_variables_mysqlx_min_worker_threads
		mysqlx_read_timeout = var.mysql_configuration_variables_mysqlx_read_timeout
		mysqlx_wait_timeout = var.mysql_configuration_variables_mysqlx_wait_timeout
		mysqlx_write_timeout = var.mysql_configuration_variables_mysqlx_write_timeout
		mysqlx_zstd_default_compression_level = var.mysql_configuration_variables_mysqlx_zstd_default_compression_level
		mysqlx_zstd_max_client_compression_level = var.mysql_configuration_variables_mysqlx_zstd_max_client_compression_level
		net_read_timeout = var.mysql_configuration_variables_net_read_timeout
		net_write_timeout = var.mysql_configuration_variables_net_write_timeout
		optimizer_switch = var.mysql_configuration_variables_optimizer_switch
		parser_max_mem_size = var.mysql_configuration_variables_parser_max_mem_size
		query_alloc_block_size = var.mysql_configuration_variables_query_alloc_block_size
		query_prealloc_size = var.mysql_configuration_variables_query_prealloc_size
		range_optimizer_max_mem_size = var.mysql_configuration_variables_range_optimizer_max_mem_size
		regexp_time_limit = var.mysql_configuration_variables_regexp_time_limit
		relay_log_space_limit = var.mysql_configuration_variables_relay_log_space_limit
		replica_net_timeout = var.mysql_configuration_variables_replica_net_timeout
		replica_parallel_workers = var.mysql_configuration_variables_replica_parallel_workers
		replica_type_conversions = var.mysql_configuration_variables_replica_type_conversions
		require_secure_transport = var.mysql_configuration_variables_require_secure_transport
		skip_name_resolve = var.mysql_configuration_variables_skip_name_resolve
		sort_buffer_size = var.mysql_configuration_variables_sort_buffer_size
		sql_generate_invisible_primary_key = var.mysql_configuration_variables_sql_generate_invisible_primary_key
		sql_mode = var.mysql_configuration_variables_sql_mode
		sql_require_primary_key = var.mysql_configuration_variables_sql_require_primary_key
		sql_warnings = var.mysql_configuration_variables_sql_warnings
		table_definition_cache = var.mysql_configuration_variables_table_definition_cache
		table_open_cache = var.mysql_configuration_variables_table_open_cache
		temptable_max_ram = var.mysql_configuration_variables_temptable_max_ram
		thread_pool_dedicated_listeners = var.mysql_configuration_variables_thread_pool_dedicated_listeners
		thread_pool_max_transactions_limit = var.mysql_configuration_variables_thread_pool_max_transactions_limit
		thread_pool_query_threads_per_group = var.mysql_configuration_variables_thread_pool_query_threads_per_group
		thread_pool_size = var.mysql_configuration_variables_thread_pool_size
		thread_pool_transaction_delay = var.mysql_configuration_variables_thread_pool_transaction_delay
		time_zone = var.mysql_configuration_variables_time_zone
		tmp_table_size = var.mysql_configuration_variables_tmp_table_size
		transaction_isolation = var.mysql_configuration_variables_transaction_isolation
		wait_timeout = var.mysql_configuration_variables_wait_timeout
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) User-provided data about the Configuration.
* `display_name` - (Optional) (Updatable) The display name of the Configuration.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `init_variables` - (Optional) User-defined service variables set only at DB system initialization. These variables cannot be changed later at runtime.
	* `lower_case_table_names` - (Optional)  Represents the MySQL server system variable lower_case_table_names (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_lower_case_table_names).

		lowerCaseTableNames controls case-sensitivity of tables and schema names and how they are stored in the DB System.

		Valid values are:
		* CASE_SENSITIVE - (default) Table and schema name comparisons are case-sensitive and stored as specified. (lower_case_table_names=0)
		* CASE_INSENSITIVE_LOWERCASE - Table and schema name comparisons are not case-sensitive and stored in lowercase. (lower_case_table_names=1) 
* `parent_configuration_id` - (Optional) The OCID of the Configuration from which the new Configuration is derived. The values in CreateConfigurationDetails.variables supersede the variables of the parent Configuration. 
* `shape_name` - (Required) The name of the associated Shape.
* `variables` - (Optional) User-defined service variables.
	* `auto_increment_increment` - (Optional) auto_increment_increment and auto_increment_offset are intended for use with circular (source-to-source) replication, and can be used to control the operation of AUTO_INCREMENT columns. Both variables have global and session values, and each can assume an integer value between 1 and 65,535 inclusive.

		autoIncrementIncrement corresponds to the MySQL Replication Source Options variable [auto_increment_increment] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-source.html#sysvar_auto_increment_increment). 
	* `auto_increment_offset` - (Optional) This variable has a default value of 1. If it is left with its default value, and Group Replication is started on the server in multi-primary mode, it is changed to the server ID.

		autoIncrementOffset corresponds to the MySQL Replication Source Options variable [auto_increment_offset] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-source.html#sysvar_auto_increment_offset). 
	* `autocommit` - (Optional) ("autocommit")
	* `big_tables` - (Optional) If enabled, the server stores all temporary tables on disk rather than in memory.

		bigTables corresponds to the MySQL server variable [big_tables](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_big_tables). 
	* `binlog_expire_logs_seconds` - (Optional) Sets the binary log expiration period in seconds. binlogExpireLogsSeconds corresponds to the MySQL binary logging system variable [binlog_expire_logs_seconds](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_expire_logs_seconds). 
	* `binlog_group_commit_sync_delay` - (Optional) Controls how many microseconds the binary log commit waits before synchronizing the binary log file to disk. There is no delay by default. Setting this variable to a microsecond delay enables more transactions to be synchronized together to disk at once, reducing the overall time to commit a group of transactions because the larger groups required fewer time units per group.

		binlogGroupCommitSyncDelay corresponds to the MySQL Replication system variable [binlog_group_commit_sync_delay](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_group_commit_sync_delay) 
	* `binlog_group_commit_sync_no_delay_count` - (Optional) The maximum number of transactions to wait for before aborting the current delay as specified by binlog_group_commit_sync_delay. If binlog_group_commit_sync_delay is set to 0, then this option has no effect.

		binlogGroupCommitSyncNoDelayCount corresponds to the MySQL Replication system variable [binlog_group_commit_sync_no_delay_count](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_group_commit_sync_no_delay_count) 
	* `binlog_row_metadata` - (Optional) Configures the amount of table metadata added to the binary log when using row-based logging. binlogRowMetadata corresponds to the MySQL binary logging system variable [binlog_row_metadata](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_metadata). 
	* `binlog_row_value_options` - (Optional) When set to PARTIAL_JSON, this enables use of a space-efficient binary log format for updates that modify only a small portion of a JSON document. binlogRowValueOptions corresponds to the MySQL binary logging system variable [binlog_row_value_options](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_value_options). 
	* `binlog_transaction_compression` - (Optional) Enables compression for transactions that are written to binary log files on this server. binlogTransactionCompression corresponds to the MySQL binary logging system variable [binlog_transaction_compression](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_transaction_compression). 
	* `block_encryption_mode` - (Optional) This variable controls the block encryption mode for block-based algorithms such as AES. It affects encryption for AES_ENCRYPT() and AES_DECRYPT(). block_encryption_mode takes a value in aes-keylen-mode format, where keylen is the key length in bits and mode is the encryption mode. The value is not case-sensitive. Permitted keylen values are 128, 192, and 256. Permitted mode values are ECB, CBC, CFB1, CFB8, CFB128, and OFB.

		block_encryption_mode corresponds to the MySQL Server Administration system variable [block_encryption_mode](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_block_encryption_mode) 
	* `character_set_server` - (Optional) The server's default character set. If you set this variable, you should also set collation_server to specify the collation for the character set.

		characterSetServer corresponds to the MySQL server variable [character_set_server](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_character_set_server). 
	* `collation_server` - (Optional) The server's default collation.

		collationServer corresponds to the MySQL server variable [collation_server](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_collation_server). 
	* `completion_type` - (Optional) ("completion_type")
	* `connect_timeout` - (Optional) The number of seconds that the mysqld server waits for a connect packet before responding with Bad handshake.

		connectTimeout corresponds to the MySQL system variable [connect_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_connect_timeout)

		Increasing the connect_timeout value might help if clients frequently encounter errors of the form "Lost connection to MySQL server at 'XXX', system error: errno". 
	* `connection_memory_chunk_size` - (Optional) Set the chunking size for updates to the global memory usage counter Global_connection_memory.

		connectionMemoryChunkSize corresponds to the MySQL system variable [connection_memory_chunk_size](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_chunk_size). 
	* `connection_memory_limit` - (Optional) Set the maximum amount of memory that can be used by a single user connection.

		connectionMemoryLimit corresponds to the MySQL system variable [connection_memory_limit](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_limit). 
	* `cte_max_recursion_depth` - (Optional) ("cte_max_recursion_depth")
	* `default_authentication_plugin` - (Optional) The default authentication plugin. This must be a plugin that uses internal credentials storage, so these values are permitted: mysql_native_password, sha256_password, caching_sha2_password.

		As of MySQL 8.0.27, which introduces multifactor authentication, default_authentication_plugin is still used, but in conjunction with and at a lower precedence than the authentication_policy system variable. For details, see The Default Authentication Plugin. Because of this diminished role, default_authentication_plugin is deprecated as of MySQL 8.0.27 and subject to removal in a future MySQL version.

		defaultAuthenticationPlugin corresponds to the MySQL system variable [default_authentication_plugin](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_default_authentication_plugin). 
	* `explain_format` - (Optional) This variable determines the default output format used by EXPLAIN in the absence of a FORMAT option when displaying a query execution plan.

		explainFormat corresponds to the MySQL system variable [explain_format](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_explain_format). 
	* `explicit_defaults_for_timestamp` - (Optional) This system variable determines whether the server enables certain nonstandard behaviors for default values and NULL-value handling in TIMESTAMP columns. By default, explicit_defaults_for_timestamp is enabled, which disables the nonstandard behaviors. Disabling explicit_defaults_for_timestamp results in a warning.

		explicit_defaults_for_timestamp corresponds to the MySQL Server Administration system variable [explicit_defaults_for_timestamp](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_explicit_defaults_for_timestamp) 
	* `foreign_key_checks` - (Optional) ("foreign_key_checks")
	* `generated_random_password_length` - (Optional) ("generated_random_password_length") DEPRECATED -- variable should not be settable and will be ignored
	* `global_connection_memory_limit` - (Optional) Set the total amount of memory that can be used by all user connections.

		globalConnectionMemoryLimit corresponds to the MySQL system variable [global_connection_memory_limit](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_limit). 
	* `global_connection_memory_tracking` - (Optional) Determines whether the MySQL server calculates Global_connection_memory.

		globalConnectionMemoryTracking corresponds to the MySQL system variable [global_connection_memory_tracking](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_tracking). 
	* `group_concat_max_len` - (Optional) Specifies the maximum permitted result length in bytes for the GROUP_CONCAT() function.

		This is the MySQL variable "group_concat_max_len". For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_group_concat_max_len) 
	* `group_replication_consistency` - (Optional) 
		* EVENTUAL: Both RO and RW transactions do not wait for preceding transactions to be applied before executing. A RW transaction does not wait for other members to apply a transaction. This means that a transaction could be externalized on one member before the others. This also means that in the event of a primary failover, the new primary can accept new RO and RW transactions before the previous primary transactions are all applied. RO transactions could result in outdated values, RW transactions could result in a rollback due to conflicts.
		* BEFORE_ON_PRIMARY_FAILOVER: New RO or RW transactions with a newly elected primary that is applying backlog from the old primary are held (not applied) until any backlog has been applied. This ensures that when a primary failover happens, intentionally or not, clients always see the latest value on the primary. This guarantees consistency, but means that clients must be able to handle the delay in the event that a backlog is being applied. Usually this delay should be minimal, but does depend on the size of the backlog.
		* BEFORE: A RW transaction waits for all preceding transactions to complete before being applied. A RO transaction waits for all preceding transactions to complete before being executed. This ensures that this transaction reads the latest value by only affecting the latency of the transaction. This reduces the overhead of synchronization on every RW transaction, by ensuring synchronization is used only on RO transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* AFTER: A RW transaction waits until its changes have been applied to all of the other members. This value has no effect on RO transactions. This mode ensures that when a transaction is committed on the local member, any subsequent transaction reads the written value or a more recent value on any group member. Use this mode with a group that is used for predominantly RO operations to ensure that applied RW transactions are applied everywhere once they commit. This could be used by your application to ensure that subsequent reads fetch the latest data which includes the latest writes. This reduces the overhead of synchronization on every RO transaction, by ensuring synchronization is used only on RW transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* BEFORE_AND_AFTER: A RW transaction waits for 1) all preceding transactions to complete before being applied and 2) until its changes have been applied on other members. A RO transaction waits for all preceding transactions to complete before execution takes place. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER. 
	* `information_schema_stats_expiry` - (Optional) ("information_schema_stats_expiry")
	* `innodb_adaptive_hash_index` - (Optional) Whether the InnoDB adaptive hash index is enabled or disabled. It may be desirable, depending on your workload, to dynamically enable or disable adaptive hash indexing to improve query performance. Because the adaptive hash index may not be useful for all workloads, conduct benchmarks with it both enabled and disabled, using realistic workloads.

		innodbAdaptiveHashIndex corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_adaptive_hash_index] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_adaptive_hash_index). 
	* `innodb_autoinc_lock_mode` - (Optional) The lock mode to use for generating auto-increment values. Permissible values are 0, 1, or 2, for traditional, consecutive, or interleaved, respectively.

		innodbAutoincLockMode corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_autoinc_lock_mode] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_autoinc_lock_mode). 
	* `innodb_buffer_pool_dump_pct` - (Optional) Specifies the percentage of the most recently used pages for each buffer pool to read out and dump.

		innodbBufferPoolDumpPct corresponds to the MySQL InnoDB system variable [innodb_buffer_pool_dump_pct](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_buffer_pool_dump_pct).

		The range is 1 to 100. The default value is 25.

		For example, if there are 4 buffer pools with 100 pages each, and innodb_buffer_pool_dump_pct is set to 25, the 25 most recently used pages from each buffer pool are dumped. 
	* `innodb_buffer_pool_instances` - (Optional) ("innodb_buffer_pool_instances")
	* `innodb_buffer_pool_size` - (Optional) The size (in bytes) of the buffer pool, that is, the memory area where InnoDB caches table and index data.

		innodbBufferPoolSize corresponds to the MySQL server system variable [innodb_buffer_pool_size](https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_buffer_pool_size).

		The default and maximum values depend on the amount of RAM provisioned by the shape. See [Default User Variables](/mysql-database/doc/configuring-db-system.html#GUID-B5504C19-F6F4-4DAB-8506-189A4E8F4A6A). 
	* `innodb_change_buffering` - (Optional) Whether InnoDB performs change buffering, an optimization that delays write operations to secondary indexes so that the I/O operations can be performed sequentially. Permitted values are described in the following table. Values may also be specified numerically.

		innodbChangeBuffering corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_change_buffering] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_change_buffering). 
	* `innodb_ddl_buffer_size` - (Optional) innodbDdlBufferSize corresponds to the MySQL system variable [innodb_ddl_buffer_size] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_ddl_buffer_size) 
	* `innodb_ddl_threads` - (Optional) innodbDdlThreads corresponds to the MySQL system variable [innodb_ddl_threads] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_ddl_threads) 
	* `innodb_ft_enable_stopword` - (Optional) ("innodb_ft_enable_stopword")
	* `innodb_ft_max_token_size` - (Optional) ("innodb_ft_max_token_size")
	* `innodb_ft_min_token_size` - (Optional) ("innodb_ft_min_token_size")
	* `innodb_ft_num_word_optimize` - (Optional) ("innodb_ft_num_word_optimize")
	* `innodb_ft_result_cache_limit` - (Optional) ("innodb_ft_result_cache_limit")
	* `innodb_ft_server_stopword_table` - (Optional) ("innodb_ft_server_stopword_table")
	* `innodb_lock_wait_timeout` - (Optional) ("innodb_lock_wait_timeout")
	* `innodb_log_writer_threads` - (Optional) Enables dedicated log writer threads for writing redo log records from the log buffer to the system buffers and flushing the system buffers to the redo log files.

		This is the MySQL variable "innodb_log_writer_threads". For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_log_writer_threads) 
	* `innodb_max_purge_lag` - (Optional) The desired maximum purge lag in terms of transactions.

		InnoDB maintains a list of transactions that have index records delete-marked by UPDATE or DELETE operations. The length of the list is the purge lag.

		If this value is exceeded, a delay is imposed on INSERT, UPDATE, and DELETE operations to allow time for purge to catch up.

		The default value is 0, which means there is no maximum purge lag and no delay.

		innodbMaxPurgeLag corresponds to the MySQL server system variable [innodb_max_purge_lag](https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_max_purge_lag). 
	* `innodb_max_purge_lag_delay` - (Optional) The maximum delay in microseconds for the delay imposed when the innodb_max_purge_lag threshold is exceeded.

		The specified innodb_max_purge_lag_delay value is an upper limit on the delay period.

		innodbMaxPurgeLagDelay corresponds to the MySQL server system variable [innodb_max_purge_lag_delay](https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_max_purge_lag_delay). 
	* `innodb_numa_interleave` - (Optional) Enables the NUMA interleave memory policy for allocation of the InnoDB buffer pool. When innodb_numa_interleave is enabled, the NUMA memory policy is set to MPOL_INTERLEAVE for the mysqld process. After the InnoDB buffer pool is allocated, the NUMA memory policy is set back to MPOL_DEFAULT. For the innodb_numa_interleave option to be available, MySQL must be compiled on a NUMA-enabled Linux system.

		innodbNumaInterleave corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_numa_interleave] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_numa_interleave). 
	* `innodb_online_alter_log_max_size` - (Optional) Specifies an upper limit in bytes on the size of the temporary log files used during online DDL operations for InnoDB tables. There is one such log file for each index being created or table being altered. This log file stores data inserted, updated, or deleted in the table during the DDL operation.

		innodbOnlineAlterLogMaxSize corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_online_alter_log_max_size] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_online_alter_log_max_size). 
	* `innodb_redo_log_capacity` - (Optional) Defines the amount of disk space occupied by redo log files. innodb_redo_log_capacity supercedes the innodb_log_files_in_group and innodb_log_file_size variables, which are both ignored if innodb_redo_log_capacity is defined. If innodb_redo_log_capacity is not defined, and if neither innodb_log_file_size or innodb_log_files_in_group are defined, then the default innodb_redo_log_capacity value is used.

		innodbRedoLogCapacity corresponds to the InnoDB Startup Options and System Variables [innodb_redo_log_capacity](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_redo_log_capacity) 
	* `innodb_rollback_on_timeout` - (Optional) InnoDB rolls back only the last statement on a transaction timeout by default. If --innodb-rollback-on-timeout is specified, a transaction timeout causes InnoDB to abort and roll back the entire transaction.

		innodbRollbackOnTimeout corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_rollback_on_timeout] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_rollback_on_timeout). 
	* `innodb_sort_buffer_size` - (Optional) This variable defines:
		* The sort buffer size for online DDL operations that create or rebuild secondary indexes. However, as of MySQL 8.0.27, this responsibility is subsumed by the innodb_ddl_buffer_size variable.
		* The amount by which the temporary log file is extended when recording concurrent DML during an online DDL operation, and the size of the temporary log file read buffer and write buffer.

		innodbSortBufferSize corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_sort_buffer_size] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_sort_buffer_size). 
	* `innodb_stats_persistent_sample_pages` - (Optional) The number of index pages to sample when estimating cardinality and other statistics for an indexed column, such as those calculated by ANALYZE TABLE.

		innodbStatsPersistentSamplePages corresponds to the MySQL InnoDB system variable [innodb_stats_persistent_sample_pages](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_persistent_sample_pages)

		innodb_stats_persistent_sample_pages only applies when innodb_stats_persistent is enabled for a table; when innodb_stats_persistent is disabled, innodb_stats_transient_sample_pages applies instead. 
	* `innodb_stats_transient_sample_pages` - (Optional) The number of index pages to sample when estimating cardinality and other statistics for an indexed column, such as those calculated by [ANALYZE TABLE](https://dev.mysql.com/doc/refman/8.0/en/analyze-table.html).

		innodbStatsTransientSamplePages corresponds to the MySQL InnoDB system variable [innodb_stats_transient_sample_pages](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_transient_sample_pages)

		innodb_stats_transient_sample_pages only applies when innodb_stats_persistent is disabled for a table; when innodb_stats_persistent is enabled, innodb_stats_persistent_sample_pages applies instead.

		innodb_stats_persistent is ON by default and cannot be changed. It is possible to override it using the STATS_PERSISTENT clause of the [CREATE TABLE](https://dev.mysql.com/doc/refman/8.0/en/create-table.html) and [ALTER TABLE](https://dev.mysql.com/doc/refman/8.0/en/alter-table.html) statements. 
	* `innodb_strict_mode` - (Optional) When you enable innodbStrictMode, the InnoDB storage engine returns errors instead of warnings for invalid or incompatible table options.

		innodbStrictMode corresponds to the MySQL InnoDB system variable [innodb_strict_mode](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_strict_mode) 
	* `innodb_undo_log_truncate` - (Optional) When enabled, undo tablespaces that exceed the threshold value defined by innodb_max_undo_log_size are marked for truncation. Only undo tablespaces can be truncated. Truncating undo logs that reside in the system tablespace is not supported. For truncation to occur, there must be at least two undo tablespaces.

		innodbUndoLogTruncate corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_undo_log_truncate] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_undo_log_truncate). 
	* `interactive_timeout` - (Optional) The number of seconds the server waits for activity on an interactive connection before closing it.

		interactiveTimeout corresponds to the MySQL system variable. [interactive_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_interactive_timeout) 
	* `join_buffer_size` - (Optional) The minimum size of the buffer that is used for plain index scans, range index scans, and joins that do not use indexes and thus perform full table scans. In MySQL 8.0.18 and later, this variable also controls the amount of memory used for hash joins. Normally, the best way to get fast joins is to add indexes. Increase the value of join_buffer_size to get a faster full join when adding indexes is not possible. One join buffer is allocated for each full join between two tables. For a complex join between several tables for which indexes are not used, multiple join buffers might be necessary.

		joinBufferSize corresponds to the MySQL Server System variable [join_buffer_size] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_join_buffer_size). 
	* `local_infile` - (Optional) This variable controls server-side LOCAL capability for LOAD DATA statements. Depending on the local_infile setting, the server refuses or permits local data loading by clients that have LOCAL enabled on the client side. 

		local_infile corresponds to the MySQL Server system variable [local_infile](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_local_infile) 
	* `long_query_time` - (Optional) If a query takes longer than this many seconds, the server increments the Slow_queries status variable. If the slow query log is enabled, the query is logged to the slow query log file. This value is measured in real time, not CPU time, so a query that is under the threshold on a lightly loaded system might be above the threshold on a heavily loaded one.

		longQueryTime corresponds to the MySQL Server System variable [long_query_time] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_long_query_time). 
	* `mandatory_roles` - (Optional) ("mandatory_roles")
	* `max_allowed_packet` - (Optional) The maximum size of one packet or any generated/intermediate string.

		This is the mysql variable "max_allowed_packet". 
	* `max_binlog_cache_size` - (Optional) Sets the size of the transaction cache.

		maxBinlogCacheSize corresponds to the MySQL server system variable [max_binlog_cache_size](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_max_binlog_cache_size). 
	* `max_connect_errors` - (Optional) ("max_connect_errors")
	* `max_connections` - (Optional) ("max_connections")
	* `max_execution_time` - (Optional) ("max_execution_time")
	* `max_heap_table_size` - (Optional) This variable sets the maximum size to which user-created MEMORY tables are permitted to grow.

		maxHeapTableSize corresponds to the MySQL system variable [max_heap_table_size](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_heap_table_size) 
	* `max_prepared_stmt_count` - (Optional) ("max_prepared_stmt_count")
	* `max_seeks_for_key` - (Optional) Limit the assumed maximum number of seeks when looking up rows based on a key. The MySQL optimizer assumes that no more than this number of key seeks are required when searching for matching rows in a table by scanning an index, regardless of the actual cardinality of the index (see Section 15.7.7.22, “SHOW INDEX Statement”). By setting this to a low value (say, 100), you can force MySQL to prefer indexes instead of table scans.

		maxSeeksForKey corresponds to the MySQL Server System variable [max_seeks_for_key] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_seeks_for_key). 
	* `max_user_connections` - (Optional) The maximum number of simultaneous connections permitted to any given MySQL user account. A value of 0 (the default) means “no limit.” This variable has a global value that can be set at server startup or runtime. It also has a read-only session value that indicates the effective simultaneous-connection limit that applies to the account associated with the current session.

		maxUserConnections corresponds to the MySQL Server System variable [max_user_connections] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_user_connections). 
	* `mysql_firewall_mode` - (Optional) ("mysql_firewall_mode")
	* `mysql_zstd_default_compression_level` - (Optional) DEPRECATED -- typo of mysqlx_zstd_default_compression_level. variable will be ignored.
	* `mysqlx_connect_timeout` - (Optional) The number of seconds X Plugin waits for the first packet to be received from newly connected clients.

		mysqlxConnectTimeout corresponds to the MySQL X Plugin system variable [mysqlx_connect_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_connect_timeout) 
	* `mysqlx_deflate_default_compression_level` - (Optional) Set the default compression level for the deflate algorithm. ("mysqlx_deflate_default_compression_level")
	* `mysqlx_deflate_max_client_compression_level` - (Optional) Limit the upper bound of accepted compression levels for the deflate algorithm. ("mysqlx_deflate_max_client_compression_level")
	* `mysqlx_document_id_unique_prefix` - (Optional) ("mysqlx_document_id_unique_prefix") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_enable_hello_notice` - (Optional) ("mysqlx_enable_hello_notice") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_idle_worker_thread_timeout` - (Optional) ("mysqlx_idle_worker_thread_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_interactive_timeout` - (Optional) The number of seconds to wait for interactive clients to timeout.

		mysqlxInteractiveTimeout corresponds to the MySQL X Plugin system variable. [mysqlx_interactive_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_interactive_timeout) 
	* `mysqlx_lz4default_compression_level` - (Optional) Set the default compression level for the lz4 algorithm. ("mysqlx_lz4_default_compression_level")
	* `mysqlx_lz4max_client_compression_level` - (Optional) Limit the upper bound of accepted compression levels for the lz4 algorithm. ("mysqlx_lz4_max_client_compression_level")
	* `mysqlx_max_allowed_packet` - (Optional) The maximum size of network packets that can be received by X Plugin.

		This is the mysql variable "mysqlx_max_allowed_packet". 
	* `mysqlx_min_worker_threads` - (Optional) ("mysqlx_min_worker_threads") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_read_timeout` - (Optional) The number of seconds that X Plugin waits for blocking read operations to complete. After this time, if the read operation is not successful, X Plugin closes the connection and returns a warning notice with the error code ER_IO_READ_ERROR to the client application.

		mysqlxReadTimeout corresponds to the MySQL X Plugin system variable [mysqlx_read_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_read_timeout) 
	* `mysqlx_wait_timeout` - (Optional) The number of seconds that X Plugin waits for activity on a connection.

		mysqlxWaitTimeout corresponds to the MySQL X Plugin system variable. [mysqlx_wait_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_wait_timeout) 
	* `mysqlx_write_timeout` - (Optional) The number of seconds that X Plugin waits for blocking write operations to complete. After this time, if the write operation is not successful, X Plugin closes the connection.

		mysqlxReadmysqlxWriteTimeoutTimeout corresponds to the MySQL X Plugin system variable [mysqlx_write_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_write_timeout) 
	* `mysqlx_zstd_default_compression_level` - (Optional) Set the default compression level for the zstd algorithm. ("mysqlx_zstd_default_compression_level")
	* `mysqlx_zstd_max_client_compression_level` - (Optional) Limit the upper bound of accepted compression levels for the zstd algorithm. ("mysqlx_zstd_max_client_compression_level")
	* `net_read_timeout` - (Optional) The number of seconds to wait for more data from a connection before aborting the read.

		netReadTimeout corresponds to the MySQL system variable [net_read_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_net_read_timeout) 
	* `net_write_timeout` - (Optional) The number of seconds to wait for a block to be written to a connection before aborting the write.

		netWriteTimeout corresponds to the MySQL system variable [net_write_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_net_write_timeout) 
	* `optimizer_switch` - (Optional) The optimizer_switch system variable enables control over optimizer behavior. The value of this variable is a set of flags, each of which has a value of on or off to indicate whether the corresponding optimizer behavior is enabled or disabled. This variable has global and session values and can be changed at runtime. The global default can be set at server startup.

		Setting hypergraph_optimizer=on for cloud builds below 9.0.0 will fail.

		optimizerSwitch corresponds to the MySQL Server System variable [optimizer_switch] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_optimizer_switch). 
	* `parser_max_mem_size` - (Optional) ("parser_max_mem_size")
	* `query_alloc_block_size` - (Optional) ("query_alloc_block_size") DEPRECATED -- variable should not be settable and will be ignored
	* `query_prealloc_size` - (Optional) ("query_prealloc_size") DEPRECATED -- variable should not be settable and will be ignored
	* `range_optimizer_max_mem_size` - (Optional) The limit on memory consumption for the range optimizer. A value of 0 means “no limit.” If an execution plan considered by the optimizer uses the range access method but the optimizer estimates that the amount of memory needed for this method would exceed the limit, it abandons the plan and considers other plans. 

		rangeOptimizerMaxMemSize corresponds to the MySQL Server System variable [range_optimizer_max_mem_size] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_range_optimizer_max_mem_size). 
	* `regexp_time_limit` - (Optional) regexpTimeLimit corresponds to the MySQL system variable [regexp_time_limit] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_regexp_time_limit) 
	* `relay_log_space_limit` - (Optional) The maximum amount of space to use for all relay logs.

		relayLogSpaceLimit corresponds to the MySQL Replica Server Options variable [relay_log_space_limit] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_relay_log_space_limit). 
	* `replica_net_timeout` - (Optional) Specifies the number of seconds to wait for more data or a heartbeat signal from the source before the replica considers the connection broken, aborts the read, and tries to reconnect. Setting this variable has no immediate effect. The state of the variable applies on all subsequent START REPLICA commands.

		replicaNetTimeout corresponds to the MySQL Replica server system variable [replica_net_timeout](https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_net_timeout) 
	* `replica_parallel_workers` - (Optional) Beginning with MySQL 8.0.26, slave_parallel_workers is deprecated, and you should use replica_parallel_workers instead. (Prior to MySQL 8.0.26, you must use slave_parallel_workers to set the number of applier threads.)

		replicaParallelWorkers corresponds to the MySQL Replica Server Options variable [replica_parallel_workers] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_parallel_workers). 
	* `replica_type_conversions` - (Optional) From MySQL 8.0.26, use replica_type_conversions in place of slave_type_conversions, which is deprecated from that release. In releases before MySQL 8.0.26, use slave_type_conversions.

		replica_type_conversions controls the type conversion mode in effect on the replica when using row-based replication. Its value is a comma-delimited set of zero or more elements from the list: ALL_LOSSY, ALL_NON_LOSSY, ALL_SIGNED, ALL_UNSIGNED. Set this variable to an empty string to disallow type conversions between the source and the replica. Setting this variable takes effect for all replication channels immediately, including running channels. 

		replica_type_conversions corresponds to the MySQL Replica Server Options variable [replica_type_conversions] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_type_conversions). 
	* `require_secure_transport` - (Optional) Whether client connections to the server are required to use some form of secure transport. When this variable is enabled, the server permits only TCP/IP connections encrypted using TLS/SSL, or connections that use a socket file or shared memory. The server rejects nonsecure connection attempts, which fail with an ER_SECURE_TRANSPORT_REQUIRED error.

		require_secure_transport corresponds to the MySQL Server Administration system variable [require_secure_transport](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_require_secure_transport) 
	* `skip_name_resolve` - (Optional) Whether to resolve host names when checking client connections. If this variable is OFF, mysqld resolves host names when checking client connections. If it is ON, mysqld uses only IP numbers; in this case, all Host column values in the grant tables must be IP addresses. See Section 7.1.12.3, “DNS Lookups and the Host Cache”.

		skipNameResolve corresponds to the MySQL Server System variable [skip_name_resolve] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_skip_name_resolve). 
	* `sort_buffer_size` - (Optional) Each session that must perform a sort allocates a buffer of this size.

		sortBufferSize corresponds to the MySQL system variable [sort_buffer_size](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_sort_buffer_size) 
	* `sql_generate_invisible_primary_key` - (Optional) Whether GIPK mode is in effect, in which case a MySQL replication source server adds a generated invisible primary key to any InnoDB table that is created without one.

		sqlGenerateInvisiblePrimaryKey corresponds to the MySQL system variable [sql_generate_invisible_primary_key] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_sql_generate_invisible_primary_key). 
	* `sql_mode` - (Optional) ("sql_mode")
	* `sql_require_primary_key` - (Optional) ("sql_require_primary_key")
	* `sql_warnings` - (Optional) ("sql_warnings")
	* `table_definition_cache` - (Optional) The number of table definitions that can be stored in the table definition cache. If you use a large number of tables, you can create a large table definition cache to speed up opening of tables. The table definition cache takes less space and does not use file descriptors, unlike the normal table cache.

		table_definition_cache corresponds to the MySQL Server Administration system variable [table_definition_cache](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_table_definition_cache) 
	* `table_open_cache` - (Optional) The number of open tables for all threads. Increasing this value increases the number of file descriptors that mysqld requires.

		table_open_cache corresponds to the MySQL Server Administration system variable [table_open_cache](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_table_open_cache) 
	* `temptable_max_ram` - (Optional) Defines the maximum amount of memory that can be occupied by the TempTable storage engine before it starts storing data on disk. The default value is 1073741824 bytes (1GiB). For more information, see Section 10.4.4, “Internal Temporary Table Use in MySQL”.

		temptableMaxRam corresponds to the MySQL system variable [temptable_max_ram] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_temptable_max_ram). 
	* `thread_pool_dedicated_listeners` - (Optional) Controls whether the thread pool uses dedicated listener threads. If enabled, a listener thread in each thread group is dedicated to the task of listening for network events from clients, ensuring that the maximum number of query worker threads is no more than the value specified by threadPoolMaxTransactionsLimit. threadPoolDedicatedListeners corresponds to the MySQL Database Service-specific system variable thread_pool_dedicated_listeners. 
	* `thread_pool_max_transactions_limit` - (Optional) Limits the maximum number of open transactions to the defined value. The default value is 0, which enforces no limit. threadPoolMaxTransactionsLimit corresponds to the MySQL Database Service-specific system variable thread_pool_max_transactions_limit. 
	* `thread_pool_query_threads_per_group` - (Optional) The maximum number of query threads permitted in a thread group. The maximum value is 4096, but if thread_pool_max_transactions_limit is set, thread_pool_query_threads_per_group must not exceed that value. The default value of 1 means there is one active query thread in each thread group, which works well for many loads. When you are using the high concurrency thread pool algorithm (thread_pool_algorithm = 1), consider increasing the value if you experience slower response times due to long-running transactions. 

		threadPoolQueryThreadsPerGroup corresponds to the MySQL Server system variable [thread_pool_query_threads_per_group](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_query_threads_per_group) 
	* `thread_pool_size` - (Optional) The number of thread groups in the thread pool. This is the most important parameter controlling thread pool performance. It affects how many statements can execute simultaneously. If a value outside the range of permissible values is specified, the thread pool plugin does not load and the server writes a message to the error log.

		threadPoolSize corresponds to the MySQL Server System variable [thread_pool_size] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_size). 
	* `thread_pool_transaction_delay` - (Optional) The delay period before executing a new transaction, in milliseconds. The maximum value is 300000 (5 minutes). A transaction delay can be used in cases where parallel transactions affect the performance of other operations due to resource contention. For example, if parallel transactions affect index creation or an online buffer pool resizing operation, you can configure a transaction delay to reduce resource contention while those operations are running. 

		threadPoolTransactionDelay corresponds to the MySQL Server system variable [thread_pool_transaction_delay](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_transaction_delay) 
	* `time_zone` - (Optional) Initializes the time zone for each client that connects.

		This corresponds to the MySQL System Variable "time_zone".

		The values can be given in one of the following formats, none of which are case-sensitive:
		* As a string indicating an offset from UTC of the form [H]H:MM, prefixed with a + or -, such as '+10:00', '-6:00', or '+05:30'. The permitted range is '-13:59' to '+14:00', inclusive.
		* As a named time zone, as defined by the "IANA Time Zone database", such as 'Europe/Helsinki', 'US/Eastern', 'MET', or 'UTC'. 
	* `tmp_table_size` - (Optional) The maximum size of internal in-memory temporary tables. This variable does not apply to user-created MEMORY tables.

		tmp_table_size corresponds to the MySQL system variable [tmp_table_size](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_tmp_table_size) 
	* `transaction_isolation` - (Optional) ("transaction_isolation")
	* `wait_timeout` - (Optional) The number of seconds the server waits for activity on a noninteractive connection before closing it.

		waitTimeout corresponds to the MySQL system variable. [wait_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_wait_timeout) 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of the Compartment the Configuration exists in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User-provided data about the Configuration.
* `display_name` - The display name of the Configuration.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Configuration.
* `init_variables` - User-defined service variables set only at DB system initialization. These variables cannot be changed later at runtime.
	* `lower_case_table_names` -  Represents the MySQL server system variable lower_case_table_names (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_lower_case_table_names).

		lowerCaseTableNames controls case-sensitivity of tables and schema names and how they are stored in the DB System.

		Valid values are:
		* CASE_SENSITIVE - (default) Table and schema name comparisons are case-sensitive and stored as specified. (lower_case_table_names=0)
		* CASE_INSENSITIVE_LOWERCASE - Table and schema name comparisons are not case-sensitive and stored in lowercase. (lower_case_table_names=1) 
* `parent_configuration_id` - The OCID of the Configuration from which this Configuration is "derived". This is entirely a metadata relationship. There is no relation between the values in this Configuration and its parent. 
* `shape_name` - The name of the associated Shape.
* `state` - The current state of the Configuration.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Configuration was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `time_updated` - The date and time the Configuration was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `type` - The Configuration type, DEFAULT or CUSTOM.
* `variables` - User-defined service variables.
	* `auto_increment_increment` - auto_increment_increment and auto_increment_offset are intended for use with circular (source-to-source) replication, and can be used to control the operation of AUTO_INCREMENT columns. Both variables have global and session values, and each can assume an integer value between 1 and 65,535 inclusive.

		autoIncrementIncrement corresponds to the MySQL Replication Source Options variable [auto_increment_increment] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-source.html#sysvar_auto_increment_increment). 
	* `auto_increment_offset` - This variable has a default value of 1. If it is left with its default value, and Group Replication is started on the server in multi-primary mode, it is changed to the server ID.

		autoIncrementOffset corresponds to the MySQL Replication Source Options variable [auto_increment_offset] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-source.html#sysvar_auto_increment_offset). 
	* `autocommit` - ("autocommit")
	* `big_tables` - If enabled, the server stores all temporary tables on disk rather than in memory.

		bigTables corresponds to the MySQL server variable [big_tables](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_big_tables). 
	* `binlog_expire_logs_seconds` - Sets the binary log expiration period in seconds. binlogExpireLogsSeconds corresponds to the MySQL binary logging system variable [binlog_expire_logs_seconds](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_expire_logs_seconds). 
	* `binlog_group_commit_sync_delay` - Controls how many microseconds the binary log commit waits before synchronizing the binary log file to disk. There is no delay by default. Setting this variable to a microsecond delay enables more transactions to be synchronized together to disk at once, reducing the overall time to commit a group of transactions because the larger groups required fewer time units per group.

		binlogGroupCommitSyncDelay corresponds to the MySQL Replication system variable [binlog_group_commit_sync_delay](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_group_commit_sync_delay) 
	* `binlog_group_commit_sync_no_delay_count` - The maximum number of transactions to wait for before aborting the current delay as specified by binlog_group_commit_sync_delay. If binlog_group_commit_sync_delay is set to 0, then this option has no effect.

		binlogGroupCommitSyncNoDelayCount corresponds to the MySQL Replication system variable [binlog_group_commit_sync_no_delay_count](https://dev.mysql.com/doc/refman/5.7/en/replication-options-binary-log.html#sysvar_binlog_group_commit_sync_no_delay_count) 
	* `binlog_row_metadata` - Configures the amount of table metadata added to the binary log when using row-based logging. binlogRowMetadata corresponds to the MySQL binary logging system variable [binlog_row_metadata](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_metadata). 
	* `binlog_row_value_options` - When set to PARTIAL_JSON, this enables use of a space-efficient binary log format for updates that modify only a small portion of a JSON document. binlogRowValueOptions corresponds to the MySQL binary logging system variable [binlog_row_value_options](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_value_options). 
	* `binlog_transaction_compression` - Enables compression for transactions that are written to binary log files on this server. binlogTransactionCompression corresponds to the MySQL binary logging system variable [binlog_transaction_compression](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_transaction_compression). 
	* `block_encryption_mode` - This variable controls the block encryption mode for block-based algorithms such as AES. It affects encryption for AES_ENCRYPT() and AES_DECRYPT(). block_encryption_mode takes a value in aes-keylen-mode format, where keylen is the key length in bits and mode is the encryption mode. The value is not case-sensitive. Permitted keylen values are 128, 192, and 256. Permitted mode values are ECB, CBC, CFB1, CFB8, CFB128, and OFB.

		block_encryption_mode corresponds to the MySQL Server Administration system variable [block_encryption_mode](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_block_encryption_mode) 
	* `character_set_server` - The server's default character set. If you set this variable, you should also set collation_server to specify the collation for the character set.

		characterSetServer corresponds to the MySQL server variable [character_set_server](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_character_set_server). 
	* `collation_server` - The server's default collation.

		collationServer corresponds to the MySQL server variable [collation_server](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_collation_server). 
	* `completion_type` - ("completion_type")
	* `connect_timeout` - The number of seconds that the mysqld server waits for a connect packet before responding with Bad handshake.

		connectTimeout corresponds to the MySQL system variable [connect_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_connect_timeout)

		Increasing the connect_timeout value might help if clients frequently encounter errors of the form "Lost connection to MySQL server at 'XXX', system error: errno". 
	* `connection_memory_chunk_size` - Set the chunking size for updates to the global memory usage counter Global_connection_memory.

		connectionMemoryChunkSize corresponds to the MySQL system variable [connection_memory_chunk_size](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_chunk_size). 
	* `connection_memory_limit` - Set the maximum amount of memory that can be used by a single user connection.

		connectionMemoryLimit corresponds to the MySQL system variable [connection_memory_limit](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_limit). 
	* `cte_max_recursion_depth` - ("cte_max_recursion_depth")
	* `default_authentication_plugin` - The default authentication plugin. This must be a plugin that uses internal credentials storage, so these values are permitted: mysql_native_password, sha256_password, caching_sha2_password.

		As of MySQL 8.0.27, which introduces multifactor authentication, default_authentication_plugin is still used, but in conjunction with and at a lower precedence than the authentication_policy system variable. For details, see The Default Authentication Plugin. Because of this diminished role, default_authentication_plugin is deprecated as of MySQL 8.0.27 and subject to removal in a future MySQL version.

		defaultAuthenticationPlugin corresponds to the MySQL system variable [default_authentication_plugin](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_default_authentication_plugin). 
	* `explain_format` - This variable determines the default output format used by EXPLAIN in the absence of a FORMAT option when displaying a query execution plan.

		explainFormat corresponds to the MySQL system variable [explain_format](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_explain_format). 
	* `explicit_defaults_for_timestamp` - This system variable determines whether the server enables certain nonstandard behaviors for default values and NULL-value handling in TIMESTAMP columns. By default, explicit_defaults_for_timestamp is enabled, which disables the nonstandard behaviors. Disabling explicit_defaults_for_timestamp results in a warning.

		explicit_defaults_for_timestamp corresponds to the MySQL Server Administration system variable [explicit_defaults_for_timestamp](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_explicit_defaults_for_timestamp) 
	* `foreign_key_checks` - ("foreign_key_checks")
	* `generated_random_password_length` - ("generated_random_password_length") DEPRECATED -- variable should not be settable and will be ignored
	* `global_connection_memory_limit` - Set the total amount of memory that can be used by all user connections.

		globalConnectionMemoryLimit corresponds to the MySQL system variable [global_connection_memory_limit](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_limit). 
	* `global_connection_memory_tracking` - Determines whether the MySQL server calculates Global_connection_memory.

		globalConnectionMemoryTracking corresponds to the MySQL system variable [global_connection_memory_tracking](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_tracking). 
	* `group_concat_max_len` - Specifies the maximum permitted result length in bytes for the GROUP_CONCAT() function.

		This is the MySQL variable "group_concat_max_len". For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_group_concat_max_len) 
	* `group_replication_consistency` - 
		* EVENTUAL: Both RO and RW transactions do not wait for preceding transactions to be applied before executing. A RW transaction does not wait for other members to apply a transaction. This means that a transaction could be externalized on one member before the others. This also means that in the event of a primary failover, the new primary can accept new RO and RW transactions before the previous primary transactions are all applied. RO transactions could result in outdated values, RW transactions could result in a rollback due to conflicts.
		* BEFORE_ON_PRIMARY_FAILOVER: New RO or RW transactions with a newly elected primary that is applying backlog from the old primary are held (not applied) until any backlog has been applied. This ensures that when a primary failover happens, intentionally or not, clients always see the latest value on the primary. This guarantees consistency, but means that clients must be able to handle the delay in the event that a backlog is being applied. Usually this delay should be minimal, but does depend on the size of the backlog.
		* BEFORE: A RW transaction waits for all preceding transactions to complete before being applied. A RO transaction waits for all preceding transactions to complete before being executed. This ensures that this transaction reads the latest value by only affecting the latency of the transaction. This reduces the overhead of synchronization on every RW transaction, by ensuring synchronization is used only on RO transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* AFTER: A RW transaction waits until its changes have been applied to all of the other members. This value has no effect on RO transactions. This mode ensures that when a transaction is committed on the local member, any subsequent transaction reads the written value or a more recent value on any group member. Use this mode with a group that is used for predominantly RO operations to ensure that applied RW transactions are applied everywhere once they commit. This could be used by your application to ensure that subsequent reads fetch the latest data which includes the latest writes. This reduces the overhead of synchronization on every RO transaction, by ensuring synchronization is used only on RW transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* BEFORE_AND_AFTER: A RW transaction waits for 1) all preceding transactions to complete before being applied and 2) until its changes have been applied on other members. A RO transaction waits for all preceding transactions to complete before execution takes place. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER. 
	* `information_schema_stats_expiry` - ("information_schema_stats_expiry")
	* `innodb_adaptive_hash_index` - Whether the InnoDB adaptive hash index is enabled or disabled. It may be desirable, depending on your workload, to dynamically enable or disable adaptive hash indexing to improve query performance. Because the adaptive hash index may not be useful for all workloads, conduct benchmarks with it both enabled and disabled, using realistic workloads.

		innodbAdaptiveHashIndex corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_adaptive_hash_index] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_adaptive_hash_index). 
	* `innodb_autoinc_lock_mode` - The lock mode to use for generating auto-increment values. Permissible values are 0, 1, or 2, for traditional, consecutive, or interleaved, respectively.

		innodbAutoincLockMode corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_autoinc_lock_mode] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_autoinc_lock_mode). 
	* `innodb_buffer_pool_dump_pct` - Specifies the percentage of the most recently used pages for each buffer pool to read out and dump.

		innodbBufferPoolDumpPct corresponds to the MySQL InnoDB system variable [innodb_buffer_pool_dump_pct](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_buffer_pool_dump_pct).

		The range is 1 to 100. The default value is 25.

		For example, if there are 4 buffer pools with 100 pages each, and innodb_buffer_pool_dump_pct is set to 25, the 25 most recently used pages from each buffer pool are dumped. 
	* `innodb_buffer_pool_instances` - ("innodb_buffer_pool_instances")
	* `innodb_buffer_pool_size` - The size (in bytes) of the buffer pool, that is, the memory area where InnoDB caches table and index data.

		innodbBufferPoolSize corresponds to the MySQL server system variable [innodb_buffer_pool_size](https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_buffer_pool_size).

		The default and maximum values depend on the amount of RAM provisioned by the shape. See [Default User Variables](/mysql-database/doc/configuring-db-system.html#GUID-B5504C19-F6F4-4DAB-8506-189A4E8F4A6A). 
	* `innodb_change_buffering` - Whether InnoDB performs change buffering, an optimization that delays write operations to secondary indexes so that the I/O operations can be performed sequentially. Permitted values are described in the following table. Values may also be specified numerically.

		innodbChangeBuffering corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_change_buffering] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_change_buffering). 
	* `innodb_ddl_buffer_size` - innodbDdlBufferSize corresponds to the MySQL system variable [innodb_ddl_buffer_size] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_ddl_buffer_size) 
	* `innodb_ddl_threads` - innodbDdlThreads corresponds to the MySQL system variable [innodb_ddl_threads] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_ddl_threads) 
	* `innodb_ft_enable_stopword` - ("innodb_ft_enable_stopword")
	* `innodb_ft_max_token_size` - ("innodb_ft_max_token_size")
	* `innodb_ft_min_token_size` - ("innodb_ft_min_token_size")
	* `innodb_ft_num_word_optimize` - ("innodb_ft_num_word_optimize")
	* `innodb_ft_result_cache_limit` - ("innodb_ft_result_cache_limit")
	* `innodb_ft_server_stopword_table` - ("innodb_ft_server_stopword_table")
	* `innodb_lock_wait_timeout` - ("innodb_lock_wait_timeout")
	* `innodb_log_writer_threads` - Enables dedicated log writer threads for writing redo log records from the log buffer to the system buffers and flushing the system buffers to the redo log files.

		This is the MySQL variable "innodb_log_writer_threads". For more information, please see the [MySQL documentation](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_log_writer_threads) 
	* `innodb_max_purge_lag` - The desired maximum purge lag in terms of transactions.

		InnoDB maintains a list of transactions that have index records delete-marked by UPDATE or DELETE operations. The length of the list is the purge lag.

		If this value is exceeded, a delay is imposed on INSERT, UPDATE, and DELETE operations to allow time for purge to catch up.

		The default value is 0, which means there is no maximum purge lag and no delay.

		innodbMaxPurgeLag corresponds to the MySQL server system variable [innodb_max_purge_lag](https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_max_purge_lag). 
	* `innodb_max_purge_lag_delay` - The maximum delay in microseconds for the delay imposed when the innodb_max_purge_lag threshold is exceeded.

		The specified innodb_max_purge_lag_delay value is an upper limit on the delay period.

		innodbMaxPurgeLagDelay corresponds to the MySQL server system variable [innodb_max_purge_lag_delay](https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_max_purge_lag_delay). 
	* `innodb_numa_interleave` - Enables the NUMA interleave memory policy for allocation of the InnoDB buffer pool. When innodb_numa_interleave is enabled, the NUMA memory policy is set to MPOL_INTERLEAVE for the mysqld process. After the InnoDB buffer pool is allocated, the NUMA memory policy is set back to MPOL_DEFAULT. For the innodb_numa_interleave option to be available, MySQL must be compiled on a NUMA-enabled Linux system.

		innodbNumaInterleave corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_numa_interleave] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_numa_interleave). 
	* `innodb_online_alter_log_max_size` - Specifies an upper limit in bytes on the size of the temporary log files used during online DDL operations for InnoDB tables. There is one such log file for each index being created or table being altered. This log file stores data inserted, updated, or deleted in the table during the DDL operation.

		innodbOnlineAlterLogMaxSize corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_online_alter_log_max_size] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_online_alter_log_max_size). 
	* `innodb_redo_log_capacity` - Defines the amount of disk space occupied by redo log files. innodb_redo_log_capacity supercedes the innodb_log_files_in_group and innodb_log_file_size variables, which are both ignored if innodb_redo_log_capacity is defined. If innodb_redo_log_capacity is not defined, and if neither innodb_log_file_size or innodb_log_files_in_group are defined, then the default innodb_redo_log_capacity value is used.

		innodbRedoLogCapacity corresponds to the InnoDB Startup Options and System Variables [innodb_redo_log_capacity](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_redo_log_capacity) 
	* `innodb_rollback_on_timeout` - InnoDB rolls back only the last statement on a transaction timeout by default. If --innodb-rollback-on-timeout is specified, a transaction timeout causes InnoDB to abort and roll back the entire transaction.

		innodbRollbackOnTimeout corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_rollback_on_timeout] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_rollback_on_timeout). 
	* `innodb_sort_buffer_size` - This variable defines:
		* The sort buffer size for online DDL operations that create or rebuild secondary indexes. However, as of MySQL 8.0.27, this responsibility is subsumed by the innodb_ddl_buffer_size variable.
		* The amount by which the temporary log file is extended when recording concurrent DML during an online DDL operation, and the size of the temporary log file read buffer and write buffer.

		innodbSortBufferSize corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_sort_buffer_size] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_sort_buffer_size). 
	* `innodb_stats_persistent_sample_pages` - The number of index pages to sample when estimating cardinality and other statistics for an indexed column, such as those calculated by ANALYZE TABLE.

		innodbStatsPersistentSamplePages corresponds to the MySQL InnoDB system variable [innodb_stats_persistent_sample_pages](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_persistent_sample_pages)

		innodb_stats_persistent_sample_pages only applies when innodb_stats_persistent is enabled for a table; when innodb_stats_persistent is disabled, innodb_stats_transient_sample_pages applies instead. 
	* `innodb_stats_transient_sample_pages` - The number of index pages to sample when estimating cardinality and other statistics for an indexed column, such as those calculated by [ANALYZE TABLE](https://dev.mysql.com/doc/refman/8.0/en/analyze-table.html).

		innodbStatsTransientSamplePages corresponds to the MySQL InnoDB system variable [innodb_stats_transient_sample_pages](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_transient_sample_pages)

		innodb_stats_transient_sample_pages only applies when innodb_stats_persistent is disabled for a table; when innodb_stats_persistent is enabled, innodb_stats_persistent_sample_pages applies instead.

		innodb_stats_persistent is ON by default and cannot be changed. It is possible to override it using the STATS_PERSISTENT clause of the [CREATE TABLE](https://dev.mysql.com/doc/refman/8.0/en/create-table.html) and [ALTER TABLE](https://dev.mysql.com/doc/refman/8.0/en/alter-table.html) statements. 
	* `innodb_strict_mode` - When you enable innodbStrictMode, the InnoDB storage engine returns errors instead of warnings for invalid or incompatible table options.

		innodbStrictMode corresponds to the MySQL InnoDB system variable [innodb_strict_mode](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_strict_mode) 
	* `innodb_undo_log_truncate` - When enabled, undo tablespaces that exceed the threshold value defined by innodb_max_undo_log_size are marked for truncation. Only undo tablespaces can be truncated. Truncating undo logs that reside in the system tablespace is not supported. For truncation to occur, there must be at least two undo tablespaces.

		innodbUndoLogTruncate corresponds to the MySQL InnoDB Startup Options and System Variables [innodb_undo_log_truncate] (https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_undo_log_truncate). 
	* `interactive_timeout` - The number of seconds the server waits for activity on an interactive connection before closing it.

		interactiveTimeout corresponds to the MySQL system variable. [interactive_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_interactive_timeout) 
	* `join_buffer_size` - The minimum size of the buffer that is used for plain index scans, range index scans, and joins that do not use indexes and thus perform full table scans. In MySQL 8.0.18 and later, this variable also controls the amount of memory used for hash joins. Normally, the best way to get fast joins is to add indexes. Increase the value of join_buffer_size to get a faster full join when adding indexes is not possible. One join buffer is allocated for each full join between two tables. For a complex join between several tables for which indexes are not used, multiple join buffers might be necessary.

		joinBufferSize corresponds to the MySQL Server System variable [join_buffer_size] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_join_buffer_size). 
	* `local_infile` - This variable controls server-side LOCAL capability for LOAD DATA statements. Depending on the local_infile setting, the server refuses or permits local data loading by clients that have LOCAL enabled on the client side. 

		local_infile corresponds to the MySQL Server system variable [local_infile](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_local_infile) 
	* `long_query_time` - If a query takes longer than this many seconds, the server increments the Slow_queries status variable. If the slow query log is enabled, the query is logged to the slow query log file. This value is measured in real time, not CPU time, so a query that is under the threshold on a lightly loaded system might be above the threshold on a heavily loaded one.

		longQueryTime corresponds to the MySQL Server System variable [long_query_time] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_long_query_time). 
	* `mandatory_roles` - ("mandatory_roles")
	* `max_allowed_packet` - The maximum size of one packet or any generated/intermediate string.

		This is the mysql variable "max_allowed_packet". 
	* `max_binlog_cache_size` - Sets the size of the transaction cache.

		maxBinlogCacheSize corresponds to the MySQL server system variable [max_binlog_cache_size](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_max_binlog_cache_size). 
	* `max_connect_errors` - ("max_connect_errors")
	* `max_connections` - ("max_connections")
	* `max_execution_time` - ("max_execution_time")
	* `max_heap_table_size` - This variable sets the maximum size to which user-created MEMORY tables are permitted to grow.

		maxHeapTableSize corresponds to the MySQL system variable [max_heap_table_size](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_heap_table_size) 
	* `max_prepared_stmt_count` - ("max_prepared_stmt_count")
	* `max_seeks_for_key` - Limit the assumed maximum number of seeks when looking up rows based on a key. The MySQL optimizer assumes that no more than this number of key seeks are required when searching for matching rows in a table by scanning an index, regardless of the actual cardinality of the index (see Section 15.7.7.22, “SHOW INDEX Statement”). By setting this to a low value (say, 100), you can force MySQL to prefer indexes instead of table scans.

		maxSeeksForKey corresponds to the MySQL Server System variable [max_seeks_for_key] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_seeks_for_key). 
	* `max_user_connections` - The maximum number of simultaneous connections permitted to any given MySQL user account. A value of 0 (the default) means “no limit.” This variable has a global value that can be set at server startup or runtime. It also has a read-only session value that indicates the effective simultaneous-connection limit that applies to the account associated with the current session.

		maxUserConnections corresponds to the MySQL Server System variable [max_user_connections] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_max_user_connections). 
	* `mysql_firewall_mode` - ("mysql_firewall_mode")
	* `mysql_zstd_default_compression_level` - DEPRECATED -- typo of mysqlx_zstd_default_compression_level. variable will be ignored.
	* `mysqlx_connect_timeout` - The number of seconds X Plugin waits for the first packet to be received from newly connected clients.

		mysqlxConnectTimeout corresponds to the MySQL X Plugin system variable [mysqlx_connect_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_connect_timeout) 
	* `mysqlx_deflate_default_compression_level` - Set the default compression level for the deflate algorithm. ("mysqlx_deflate_default_compression_level")
	* `mysqlx_deflate_max_client_compression_level` - Limit the upper bound of accepted compression levels for the deflate algorithm. ("mysqlx_deflate_max_client_compression_level")
	* `mysqlx_document_id_unique_prefix` - ("mysqlx_document_id_unique_prefix") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_enable_hello_notice` - ("mysqlx_enable_hello_notice") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_idle_worker_thread_timeout` - ("mysqlx_idle_worker_thread_timeout") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_interactive_timeout` - The number of seconds to wait for interactive clients to timeout.

		mysqlxInteractiveTimeout corresponds to the MySQL X Plugin system variable. [mysqlx_interactive_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_interactive_timeout) 
	* `mysqlx_lz4default_compression_level` - Set the default compression level for the lz4 algorithm. ("mysqlx_lz4_default_compression_level")
	* `mysqlx_lz4max_client_compression_level` - Limit the upper bound of accepted compression levels for the lz4 algorithm. ("mysqlx_lz4_max_client_compression_level")
	* `mysqlx_max_allowed_packet` - The maximum size of network packets that can be received by X Plugin.

		This is the mysql variable "mysqlx_max_allowed_packet". 
	* `mysqlx_min_worker_threads` - ("mysqlx_min_worker_threads") DEPRECATED -- variable should not be settable and will be ignored
	* `mysqlx_read_timeout` - The number of seconds that X Plugin waits for blocking read operations to complete. After this time, if the read operation is not successful, X Plugin closes the connection and returns a warning notice with the error code ER_IO_READ_ERROR to the client application.

		mysqlxReadTimeout corresponds to the MySQL X Plugin system variable [mysqlx_read_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_read_timeout) 
	* `mysqlx_wait_timeout` - The number of seconds that X Plugin waits for activity on a connection.

		mysqlxWaitTimeout corresponds to the MySQL X Plugin system variable. [mysqlx_wait_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_wait_timeout) 
	* `mysqlx_write_timeout` - The number of seconds that X Plugin waits for blocking write operations to complete. After this time, if the write operation is not successful, X Plugin closes the connection.

		mysqlxReadmysqlxWriteTimeoutTimeout corresponds to the MySQL X Plugin system variable [mysqlx_write_timeout](https://dev.mysql.com/doc/refman/8.0/en/x-plugin-options-system-variables.html#sysvar_mysqlx_write_timeout) 
	* `mysqlx_zstd_default_compression_level` - Set the default compression level for the zstd algorithm. ("mysqlx_zstd_default_compression_level")
	* `mysqlx_zstd_max_client_compression_level` - Limit the upper bound of accepted compression levels for the zstd algorithm. ("mysqlx_zstd_max_client_compression_level")
	* `net_read_timeout` - The number of seconds to wait for more data from a connection before aborting the read.

		netReadTimeout corresponds to the MySQL system variable [net_read_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_net_read_timeout) 
	* `net_write_timeout` - The number of seconds to wait for a block to be written to a connection before aborting the write.

		netWriteTimeout corresponds to the MySQL system variable [net_write_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_net_write_timeout) 
	* `optimizer_switch` - The optimizer_switch system variable enables control over optimizer behavior. The value of this variable is a set of flags, each of which has a value of on or off to indicate whether the corresponding optimizer behavior is enabled or disabled. This variable has global and session values and can be changed at runtime. The global default can be set at server startup.

		Setting hypergraph_optimizer=on for cloud builds below 9.0.0 will fail.

		optimizerSwitch corresponds to the MySQL Server System variable [optimizer_switch] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_optimizer_switch). 
	* `parser_max_mem_size` - ("parser_max_mem_size")
	* `query_alloc_block_size` - ("query_alloc_block_size") DEPRECATED -- variable should not be settable and will be ignored
	* `query_prealloc_size` - ("query_prealloc_size") DEPRECATED -- variable should not be settable and will be ignored
	* `range_optimizer_max_mem_size` - The limit on memory consumption for the range optimizer. A value of 0 means “no limit.” If an execution plan considered by the optimizer uses the range access method but the optimizer estimates that the amount of memory needed for this method would exceed the limit, it abandons the plan and considers other plans. 

		rangeOptimizerMaxMemSize corresponds to the MySQL Server System variable [range_optimizer_max_mem_size] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_range_optimizer_max_mem_size). 
	* `regexp_time_limit` - regexpTimeLimit corresponds to the MySQL system variable [regexp_time_limit] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_regexp_time_limit) 
	* `relay_log_space_limit` - The maximum amount of space to use for all relay logs.

		relayLogSpaceLimit corresponds to the MySQL Replica Server Options variable [relay_log_space_limit] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_relay_log_space_limit). 
	* `replica_net_timeout` - Specifies the number of seconds to wait for more data or a heartbeat signal from the source before the replica considers the connection broken, aborts the read, and tries to reconnect. Setting this variable has no immediate effect. The state of the variable applies on all subsequent START REPLICA commands.

		replicaNetTimeout corresponds to the MySQL Replica server system variable [replica_net_timeout](https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_net_timeout) 
	* `replica_parallel_workers` - Beginning with MySQL 8.0.26, slave_parallel_workers is deprecated, and you should use replica_parallel_workers instead. (Prior to MySQL 8.0.26, you must use slave_parallel_workers to set the number of applier threads.)

		replicaParallelWorkers corresponds to the MySQL Replica Server Options variable [replica_parallel_workers] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_parallel_workers). 
	* `replica_type_conversions` - From MySQL 8.0.26, use replica_type_conversions in place of slave_type_conversions, which is deprecated from that release. In releases before MySQL 8.0.26, use slave_type_conversions.

		replica_type_conversions controls the type conversion mode in effect on the replica when using row-based replication. Its value is a comma-delimited set of zero or more elements from the list: ALL_LOSSY, ALL_NON_LOSSY, ALL_SIGNED, ALL_UNSIGNED. Set this variable to an empty string to disallow type conversions between the source and the replica. Setting this variable takes effect for all replication channels immediately, including running channels. 

		replica_type_conversions corresponds to the MySQL Replica Server Options variable [replica_type_conversions] (https://dev.mysql.com/doc/refman/8.0/en/replication-options-replica.html#sysvar_replica_type_conversions). 
	* `require_secure_transport` - Whether client connections to the server are required to use some form of secure transport. When this variable is enabled, the server permits only TCP/IP connections encrypted using TLS/SSL, or connections that use a socket file or shared memory. The server rejects nonsecure connection attempts, which fail with an ER_SECURE_TRANSPORT_REQUIRED error.

		require_secure_transport corresponds to the MySQL Server Administration system variable [require_secure_transport](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_require_secure_transport) 
	* `skip_name_resolve` - Whether to resolve host names when checking client connections. If this variable is OFF, mysqld resolves host names when checking client connections. If it is ON, mysqld uses only IP numbers; in this case, all Host column values in the grant tables must be IP addresses. See Section 7.1.12.3, “DNS Lookups and the Host Cache”.

		skipNameResolve corresponds to the MySQL Server System variable [skip_name_resolve] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_skip_name_resolve). 
	* `sort_buffer_size` - Each session that must perform a sort allocates a buffer of this size.

		sortBufferSize corresponds to the MySQL system variable [sort_buffer_size](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_sort_buffer_size) 
	* `sql_generate_invisible_primary_key` - Whether GIPK mode is in effect, in which case a MySQL replication source server adds a generated invisible primary key to any InnoDB table that is created without one.

		sqlGenerateInvisiblePrimaryKey corresponds to the MySQL system variable [sql_generate_invisible_primary_key] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_sql_generate_invisible_primary_key). 
	* `sql_mode` - ("sql_mode")
	* `sql_require_primary_key` - ("sql_require_primary_key")
	* `sql_warnings` - ("sql_warnings")
	* `table_definition_cache` - The number of table definitions that can be stored in the table definition cache. If you use a large number of tables, you can create a large table definition cache to speed up opening of tables. The table definition cache takes less space and does not use file descriptors, unlike the normal table cache.

		table_definition_cache corresponds to the MySQL Server Administration system variable [table_definition_cache](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_table_definition_cache) 
	* `table_open_cache` - The number of open tables for all threads. Increasing this value increases the number of file descriptors that mysqld requires.

		table_open_cache corresponds to the MySQL Server Administration system variable [table_open_cache](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_table_open_cache) 
	* `temptable_max_ram` - Defines the maximum amount of memory that can be occupied by the TempTable storage engine before it starts storing data on disk. The default value is 1073741824 bytes (1GiB). For more information, see Section 10.4.4, “Internal Temporary Table Use in MySQL”.

		temptableMaxRam corresponds to the MySQL system variable [temptable_max_ram] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_temptable_max_ram). 
	* `thread_pool_dedicated_listeners` - Controls whether the thread pool uses dedicated listener threads. If enabled, a listener thread in each thread group is dedicated to the task of listening for network events from clients, ensuring that the maximum number of query worker threads is no more than the value specified by threadPoolMaxTransactionsLimit. threadPoolDedicatedListeners corresponds to the MySQL Database Service-specific system variable thread_pool_dedicated_listeners. 
	* `thread_pool_max_transactions_limit` - Limits the maximum number of open transactions to the defined value. The default value is 0, which enforces no limit. threadPoolMaxTransactionsLimit corresponds to the MySQL Database Service-specific system variable thread_pool_max_transactions_limit. 
	* `thread_pool_query_threads_per_group` - The maximum number of query threads permitted in a thread group. The maximum value is 4096, but if thread_pool_max_transactions_limit is set, thread_pool_query_threads_per_group must not exceed that value. The default value of 1 means there is one active query thread in each thread group, which works well for many loads. When you are using the high concurrency thread pool algorithm (thread_pool_algorithm = 1), consider increasing the value if you experience slower response times due to long-running transactions. 

		threadPoolQueryThreadsPerGroup corresponds to the MySQL Server system variable [thread_pool_query_threads_per_group](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_query_threads_per_group) 
	* `thread_pool_size` - The number of thread groups in the thread pool. This is the most important parameter controlling thread pool performance. It affects how many statements can execute simultaneously. If a value outside the range of permissible values is specified, the thread pool plugin does not load and the server writes a message to the error log.

		threadPoolSize corresponds to the MySQL Server System variable [thread_pool_size] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_size). 
	* `thread_pool_transaction_delay` - The delay period before executing a new transaction, in milliseconds. The maximum value is 300000 (5 minutes). A transaction delay can be used in cases where parallel transactions affect the performance of other operations due to resource contention. For example, if parallel transactions affect index creation or an online buffer pool resizing operation, you can configure a transaction delay to reduce resource contention while those operations are running. 

		threadPoolTransactionDelay corresponds to the MySQL Server system variable [thread_pool_transaction_delay](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_thread_pool_transaction_delay) 
	* `time_zone` - Initializes the time zone for each client that connects.

		This corresponds to the MySQL System Variable "time_zone".

		The values can be given in one of the following formats, none of which are case-sensitive:
		* As a string indicating an offset from UTC of the form [H]H:MM, prefixed with a + or -, such as '+10:00', '-6:00', or '+05:30'. The permitted range is '-13:59' to '+14:00', inclusive.
		* As a named time zone, as defined by the "IANA Time Zone database", such as 'Europe/Helsinki', 'US/Eastern', 'MET', or 'UTC'. 
	* `tmp_table_size` - The maximum size of internal in-memory temporary tables. This variable does not apply to user-created MEMORY tables.

		tmp_table_size corresponds to the MySQL system variable [tmp_table_size](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_tmp_table_size) 
	* `transaction_isolation` - ("transaction_isolation")
	* `wait_timeout` - The number of seconds the server waits for activity on a noninteractive connection before closing it.

		waitTimeout corresponds to the MySQL system variable. [wait_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_wait_timeout) 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Mysql Configuration
	* `update` - (Defaults to 20 minutes), when updating the Mysql Configuration
	* `delete` - (Defaults to 20 minutes), when destroying the Mysql Configuration


## Import

MysqlConfigurations can be imported using the `id`, e.g.

```
$ terraform import oci_mysql_mysql_configuration.test_mysql_configuration "configurations/{configurationId}" 
```

