---
subcategory: "Mysql"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_mysql_mysql_configurations"
sidebar_current: "docs-oci-datasource-mysql-mysql_configurations"
description: |-
  Provides the list of Mysql Configurations in Oracle Cloud Infrastructure Mysql service
---

# Data Source: oci_mysql_mysql_configurations
This data source provides the list of Mysql Configurations in Oracle Cloud Infrastructure Mysql service.

Lists the Configurations available when creating a DB System.

This may include DEFAULT configurations per Shape and CUSTOM configurations.

The default sort order is a multi-part sort by:
  - shapeName, ascending
  - DEFAULT-before-CUSTOM
  - displayName ascending


## Example Usage

```hcl
data "oci_mysql_mysql_configurations" "test_mysql_configurations" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	configuration_id = var.mysql_configuration_id
	display_name = var.mysql_configuration_display_name
	shape_name = var.mysql_shape_name
	state = var.mysql_configuration_state
	type = var.mysql_configuration_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `configuration_id` - (Optional) The requested Configuration instance.
* `display_name` - (Optional) A filter to return only the resource matching the given display name exactly.
* `shape_name` - (Optional) The requested Shape name.
* `state` - (Optional) Configuration Lifecycle State
* `type` - (Optional) The requested Configuration types.


## Attributes Reference

The following attributes are exported:

* `configurations` - The list of configurations.

### MysqlConfiguration Reference

The following attributes are exported:

* `compartment_id` - OCID of the Compartment the Configuration exists in.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - User-provided data about the Configuration.
* `display_name` - The display name of the Configuration.
* `freeform_tags` - Simple key-value pair applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the Configuration.
* `parent_configuration_id` - The OCID of the Configuration from which this Configuration is "derived". This is entirely a metadata relationship. There is no relation between the values in this Configuration and its parent. 
* `shape_name` - The name of the associated Shape.
* `state` - The current state of the Configuration.
* `time_created` - The date and time the Configuration was created, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `time_updated` - The date and time the Configuration was last updated, as described by [RFC 3339](https://tools.ietf.org/rfc/rfc3339).
* `type` - The Configuration type, DEFAULT or CUSTOM.
* `variables` - User controllable service variables.
	* `autocommit` - ("autocommit")
	* `binlog_expire_logs_seconds` - ("binlog_expire_logs_seconds")
	* `completion_type` - ("completion_type")
	* `connect_timeout` - ("connect_timeout")
	* `cte_max_recursion_depth` - ("cte_max_recursion_depth")
	* `default_authentication_plugin` - ("default_authentication_plugin")
	* `foreign_key_checks` - ("foreign_key_checks")
	* `generated_random_password_length` - ("generated_random_password_length")
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
	* `mysql_zstd_default_compression_level` - Set the default compression level for the zstd algorithm. ("mysqlx_zstd_default_compression_level")
	* `mysqlx_connect_timeout` - ("mysqlx_connect_timeout")
	* `mysqlx_deflate_default_compression_level` - Set the default compression level for the deflate algorithm. ("mysqlx_deflate_default_compression_level")
	* `mysqlx_deflate_max_client_compression_level` - Limit the upper bound of accepted compression levels for the deflate algorithm. ("mysqlx_deflate_max_client_compression_level")
	* `mysqlx_document_id_unique_prefix` - ("mysqlx_document_id_unique_prefix")
	* `mysqlx_enable_hello_notice` - ("mysqlx_enable_hello_notice")
	* `mysqlx_idle_worker_thread_timeout` - ("mysqlx_idle_worker_thread_timeout")
	* `mysqlx_interactive_timeout` - ("mysqlx_interactive_timeout")
	* `mysqlx_lz4default_compression_level` - Set the default compression level for the lz4 algorithm. ("mysqlx_lz4_default_compression_level")
	* `mysqlx_lz4max_client_compression_level` - Limit the upper bound of accepted compression levels for the lz4 algorithm. ("mysqlx_lz4_max_client_compression_level")
	* `mysqlx_max_allowed_packet` - ("mysqlx_max_allowed_packet")
	* `mysqlx_min_worker_threads` - ("mysqlx_min_worker_threads")
	* `mysqlx_read_timeout` - ("mysqlx_read_timeout")
	* `mysqlx_wait_timeout` - ("mysqlx_wait_timeout")
	* `mysqlx_write_timeout` - ("mysqlx_write_timeout")
	* `mysqlx_zstd_max_client_compression_level` - Limit the upper bound of accepted compression levels for the zstd algorithm. ("mysqlx_zstd_max_client_compression_level")
	* `parser_max_mem_size` - ("parser_max_mem_size")
	* `query_alloc_block_size` - ("query_alloc_block_size")
	* `query_prealloc_size` - ("query_prealloc_size")
	* `sql_mode` - ("sql_mode")
	* `sql_require_primary_key` - ("sql_require_primary_key")
	* `sql_warnings` - ("sql_warnings")
	* `transaction_isolation` - ("transaction_isolation")

