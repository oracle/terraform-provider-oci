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
* `init_variables` - User-defined service variables set only at DB system initialization. These variables cannot be changed later at runtime.
	* `lower_case_table_names` -  Represents the MySQL server system variable lower_case_table_names (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_lower_case_table_names).

		lowerCaseTableNames controls case-sensitivity of tables and schema names and how they are stored in the DB System.

		Valid values are:
		* CASE_SENSITIVE - (default) Table and schema name comparisons are case-sensitive and stored as specified. (lower_case_table_names=0)
		* CASE_INSENSITIVE_LOWERCASE - Table and schema name comparisons are not case-sensitive and stored in lowercase. (lower_case_table_names=1) 
* `parent_configuration_id` - The OCID of the Configuration from which this Configuration is "derived". This is entirely a metadata relationship. There is no relation between the values in this Configuration and its parent. 
* `shape_name` - The name of the associated Shape.
* `state` - The current state of the Configuration.
* `time_created` - The date and time the Configuration was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `time_updated` - The date and time the Configuration was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `type` - The Configuration type, DEFAULT or CUSTOM.
* `variables` - User-defined service variables.
	* `autocommit` - ("autocommit")
	* `big_tables` - If enabled, the server stores all temporary tables on disk rather than in memory.

		bigTables corresponds to the MySQL server variable [big_tables](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_big_tables). 
	* `binlog_expire_logs_seconds` - Sets the binary log expiration period in seconds. binlogExpireLogsSeconds corresponds to the MySQL binary logging system variable [binlog_expire_logs_seconds](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_expire_logs_seconds). 
	* `binlog_row_metadata` - Configures the amount of table metadata added to the binary log when using row-based logging. binlogRowMetadata corresponds to the MySQL binary logging system variable [binlog_row_metadata](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_metadata). 
	* `binlog_row_value_options` - When set to PARTIAL_JSON, this enables use of a space-efficient binary log format for updates that modify only a small portion of a JSON document. binlogRowValueOptions corresponds to the MySQL binary logging system variable [binlog_row_value_options](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_row_value_options). 
	* `binlog_transaction_compression` - Enables compression for transactions that are written to binary log files on this server. binlogTransactionCompression corresponds to the MySQL binary logging system variable [binlog_transaction_compression](https://dev.mysql.com/doc/refman/8.0/en/replication-options-binary-log.html#sysvar_binlog_transaction_compression). 
	* `completion_type` - ("completion_type")
	* `connect_timeout` - The number of seconds that the mysqld server waits for a connect packet before responding with Bad handshake.

		connectTimeout corresponds to the MySQL system variable [connect_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_connect_timeout)

		Increasing the connect_timeout value might help if clients frequently encounter errors of the form "Lost connection to MySQL server at 'XXX', system error: errno". 
	* `connection_memory_chunk_size` - Set the chunking size for updates to the global memory usage counter Global_connection_memory.

		connectionMemoryChunkSize corresponds to the MySQL system variable [connection_memory_chunk_size](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_chunk_size). 
	* `connection_memory_limit` - Set the maximum amount of memory that can be used by a single user connection.

		connectionMemoryLimit corresponds to the MySQL system variable [connection_memory_limit](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_connection_memory_limit). 
	* `cte_max_recursion_depth` - ("cte_max_recursion_depth")
	* `default_authentication_plugin` - ("default_authentication_plugin")
	* `foreign_key_checks` - ("foreign_key_checks")
	* `generated_random_password_length` - ("generated_random_password_length") DEPRECATED -- variable should not be settable and will be ignored
	* `global_connection_memory_limit` - Set the total amount of memory that can be used by all user connections.

		globalConnectionMemoryLimit corresponds to the MySQL system variable [global_connection_memory_limit](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_limit). 
	* `global_connection_memory_tracking` - Determines whether the MySQL server calculates Global_connection_memory.

		globalConnectionMemoryTracking corresponds to the MySQL system variable [global_connection_memory_tracking](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_global_connection_memory_tracking). 
	* `group_replication_consistency` - 
		* EVENTUAL: Both RO and RW transactions do not wait for preceding transactions to be applied before executing. A RW transaction does not wait for other members to apply a transaction. This means that a transaction could be externalized on one member before the others. This also means that in the event of a primary failover, the new primary can accept new RO and RW transactions before the previous primary transactions are all applied. RO transactions could result in outdated values, RW transactions could result in a rollback due to conflicts.
		* BEFORE_ON_PRIMARY_FAILOVER: New RO or RW transactions with a newly elected primary that is applying backlog from the old primary are held (not applied) until any backlog has been applied. This ensures that when a primary failover happens, intentionally or not, clients always see the latest value on the primary. This guarantees consistency, but means that clients must be able to handle the delay in the event that a backlog is being applied. Usually this delay should be minimal, but does depend on the size of the backlog.
		* BEFORE: A RW transaction waits for all preceding transactions to complete before being applied. A RO transaction waits for all preceding transactions to complete before being executed. This ensures that this transaction reads the latest value by only affecting the latency of the transaction. This reduces the overhead of synchronization on every RW transaction, by ensuring synchronization is used only on RO transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* AFTER: A RW transaction waits until its changes have been applied to all of the other members. This value has no effect on RO transactions. This mode ensures that when a transaction is committed on the local member, any subsequent transaction reads the written value or a more recent value on any group member. Use this mode with a group that is used for predominantly RO operations to ensure that applied RW transactions are applied everywhere once they commit. This could be used by your application to ensure that subsequent reads fetch the latest data which includes the latest writes. This reduces the overhead of synchronization on every RO transaction, by ensuring synchronization is used only on RW transactions. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER.
		* BEFORE_AND_AFTER: A RW transaction waits for 1) all preceding transactions to complete before being applied and 2) until its changes have been applied on other members. A RO transaction waits for all preceding transactions to complete before execution takes place. This consistency level also includes the consistency guarantees provided by BEFORE_ON_PRIMARY_FAILOVER. 
	* `information_schema_stats_expiry` - ("information_schema_stats_expiry")
	* `innodb_buffer_pool_dump_pct` - Specifies the percentage of the most recently used pages for each buffer pool to read out and dump.

		innodbBufferPoolDumpPct corresponds to the MySQL InnoDB system variable [innodb_buffer_pool_dump_pct](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_buffer_pool_dump_pct).

		The range is 1 to 100. The default value is 25.

		For example, if there are 4 buffer pools with 100 pages each, and innodb_buffer_pool_dump_pct is set to 25, the 25 most recently used pages from each buffer pool are dumped. 
	* `innodb_buffer_pool_instances` - ("innodb_buffer_pool_instances")
	* `innodb_buffer_pool_size` - The size (in bytes) of the buffer pool, that is, the memory area where InnoDB caches table and index data.

		innodbBufferPoolSize corresponds to the MySQL server system variable [innodb_buffer_pool_size](https://dev.mysql.com/doc/refman/en/innodb-parameters.html#sysvar_innodb_buffer_pool_size).

		The default and maximum values depend on the amount of RAM provisioned by the shape. See [Default User Variables](/mysql-database/doc/configuring-db-system.html#GUID-B5504C19-F6F4-4DAB-8506-189A4E8F4A6A). 
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
	* `innodb_stats_persistent_sample_pages` - The number of index pages to sample when estimating cardinality and other statistics for an indexed column, such as those calculated by ANALYZE TABLE.

		innodbStatsPersistentSamplePages corresponds to the MySQL InnoDB system variable [innodb_stats_persistent_sample_pages](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_persistent_sample_pages)

		innodb_stats_persistent_sample_pages only applies when innodb_stats_persistent is enabled for a table; when innodb_stats_persistent is disabled, innodb_stats_transient_sample_pages applies instead. 
	* `innodb_stats_transient_sample_pages` - The number of index pages to sample when estimating cardinality and other statistics for an indexed column, such as those calculated by [ANALYZE TABLE](https://dev.mysql.com/doc/refman/8.0/en/analyze-table.html).

		innodbStatsTransientSamplePages corresponds to the MySQL InnoDB system variable [innodb_stats_transient_sample_pages](https://dev.mysql.com/doc/refman/8.0/en/innodb-parameters.html#sysvar_innodb_stats_transient_sample_pages)

		innodb_stats_transient_sample_pages only applies when innodb_stats_persistent is disabled for a table; when innodb_stats_persistent is enabled, innodb_stats_persistent_sample_pages applies instead.

		innodb_stats_persistent is ON by default and cannot be changed. It is possible to override it using the STATS_PERSISTENT clause of the [CREATE TABLE](https://dev.mysql.com/doc/refman/8.0/en/create-table.html) and [ALTER TABLE](https://dev.mysql.com/doc/refman/8.0/en/alter-table.html) statements. 
	* `interactive_timeout` - The number of seconds the server waits for activity on an interactive connection before closing it.

		interactiveTimeout corresponds to the MySQL system variable. [interactive_timeout](https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_interactive_timeout) 
	* `local_infile` - ("local_infile")
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
	* `parser_max_mem_size` - ("parser_max_mem_size")
	* `query_alloc_block_size` - ("query_alloc_block_size") DEPRECATED -- variable should not be settable and will be ignored
	* `query_prealloc_size` - ("query_prealloc_size") DEPRECATED -- variable should not be settable and will be ignored
	* `regexp_time_limit` - regexpTimeLimit corresponds to the MySQL system variable [regexp_time_limit] (https://dev.mysql.com/doc/refman/8.0/en/server-system-variables.html#sysvar_regexp_time_limit) 
	* `sort_buffer_size` - Each session that must perform a sort allocates a buffer of this size.

		sortBufferSize corresponds to the MySQL system variable [sort_buffer_size](https://dev.mysql.com/doc/refman/en/server-system-variables.html#sysvar_sort_buffer_size) 
	* `sql_mode` - ("sql_mode")
	* `sql_require_primary_key` - ("sql_require_primary_key")
	* `sql_warnings` - ("sql_warnings")
	* `thread_pool_dedicated_listeners` - Controls whether the thread pool uses dedicated listener threads. If enabled, a listener thread in each thread group is dedicated to the task of listening for network events from clients, ensuring that the maximum number of query worker threads is no more than the value specified by threadPoolMaxTransactionsLimit. threadPoolDedicatedListeners corresponds to the MySQL Database Service-specific system variable thread_pool_dedicated_listeners. 
	* `thread_pool_max_transactions_limit` - Limits the maximum number of open transactions to the defined value. The default value is 0, which enforces no limit. threadPoolMaxTransactionsLimit corresponds to the MySQL Database Service-specific system variable thread_pool_max_transactions_limit. 
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

