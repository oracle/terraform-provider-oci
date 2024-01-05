// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/resourcediscovery"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"
	"github.com/oracle/terraform-provider-oci/internal/utils"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/oracle/oci-go-sdk/v65/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v65/mysql"
)

var (
	MysqlMysqlConfigurationRequiredOnlyResource = MysqlMysqlConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Required, acctest.Create, MysqlMysqlConfigurationRepresentation)

	MysqlMysqlConfigurationResourceConfig = MysqlMysqlConfigurationResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Optional, acctest.Update, MysqlMysqlConfigurationRepresentation)

	MysqlMysqlMysqlConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${var.MysqlConfigurationOCID[var.region]}`},
	}

	MysqlMysqlMysqlConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `MySQL.VM.Standard.E3.1.8GB.Standalone`},
		"shape_name":       acctest.Representation{RepType: acctest.Optional, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`DEFAULT`}},
		"filter":           acctest.RepresentationGroup{RepType: acctest.Required, Group: mysqlConfigurationDataSourceFilterRepresentation}}
	mysqlConfigurationDataSourceFilterRepresentation = map[string]interface{}{
		"name":   acctest.Representation{RepType: acctest.Required, Create: `id`},
		"values": acctest.Representation{RepType: acctest.Required, Create: []string{`${oci_mysql_mysql_configuration.test_mysql_configuration.id}`}},
	}

	MysqlMysqlConfigurationRepresentation = map[string]interface{}{
		"compartment_id":          acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"shape_name":              acctest.Representation{RepType: acctest.Required, Create: `MySQL.VM.Standard.E3.1.8GB`},
		"defined_tags":            acctest.Representation{RepType: acctest.Optional, Create: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "value")}`, Update: `${map("${oci_identity_tag_namespace.tag-namespace1.name}.${oci_identity_tag.tag1.name}", "updatedValue")}`},
		"description":             acctest.Representation{RepType: acctest.Optional, Create: `description`, Update: `description2`},
		"display_name":            acctest.Representation{RepType: acctest.Optional, Create: `displayName`, Update: `displayName2`},
		"freeform_tags":           acctest.Representation{RepType: acctest.Optional, Create: map[string]string{"bar-key": "value"}, Update: map[string]string{"Department": "Accounting"}},
		"init_variables":          acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlConfigurationInitVariablesRepresentation},
		"parent_configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"variables":               acctest.RepresentationGroup{RepType: acctest.Optional, Group: MysqlMysqlConfigurationVariablesRepresentation},
		"lifecycle":               acctest.RepresentationGroup{RepType: acctest.Required, Group: ignoreDefinedTagsChangesForMysqlConfigBasic},
	}
	ignoreDefinedTagsChangesForMysqlConfigBasic = map[string]interface{}{
		"ignore_changes": acctest.Representation{RepType: acctest.Required, Create: []string{"defined_tags"}},
	}
	MysqlMysqlConfigurationInitVariablesRepresentation = map[string]interface{}{
		"lower_case_table_names": acctest.Representation{RepType: acctest.Optional, Create: `CASE_SENSITIVE`},
	}
	MysqlMysqlConfigurationVariablesRepresentation = map[string]interface{}{
		"autocommit":                                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"big_tables":                                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"binlog_expire_logs_seconds":                  acctest.Representation{RepType: acctest.Optional, Create: `3600`},
		"binlog_row_metadata":                         acctest.Representation{RepType: acctest.Optional, Create: `FULL`},
		"binlog_row_value_options":                    acctest.Representation{RepType: acctest.Optional, Create: `PARTIAL_JSON`},
		"binlog_transaction_compression":              acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"completion_type":                             acctest.Representation{RepType: acctest.Optional, Create: `NO_CHAIN`},
		"connect_timeout":                             acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"connection_memory_chunk_size":                acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"connection_memory_limit":                     acctest.Representation{RepType: acctest.Optional, Create: `2097152`},
		"cte_max_recursion_depth":                     acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"default_authentication_plugin":               acctest.Representation{RepType: acctest.Optional, Create: `mysql_native_password`},
		"foreign_key_checks":                          acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"generated_random_password_length":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"global_connection_memory_limit":              acctest.Representation{RepType: acctest.Optional, Create: `2097152`},
		"global_connection_memory_tracking":           acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"group_replication_consistency":               acctest.Representation{RepType: acctest.Optional, Create: `BEFORE_ON_PRIMARY_FAILOVER`},
		"information_schema_stats_expiry":             acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_buffer_pool_dump_pct":                 acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_buffer_pool_instances":                acctest.Representation{RepType: acctest.Optional, Create: `4`},
		"innodb_buffer_pool_size":                     acctest.Representation{RepType: acctest.Optional, Create: `2147483648`},
		"innodb_ddl_buffer_size":                      acctest.Representation{RepType: acctest.Optional, Create: `65536`},
		"innodb_ddl_threads":                          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_ft_enable_stopword":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"innodb_ft_max_token_size":                    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_ft_min_token_size":                    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_ft_num_word_optimize":                 acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_ft_result_cache_limit":                acctest.Representation{RepType: acctest.Optional, Create: `33554432`},
		"innodb_ft_server_stopword_table":             acctest.Representation{RepType: acctest.Optional, Create: `innodbFtServerStopwordTable`},
		"innodb_lock_wait_timeout":                    acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_log_writer_threads":                   acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"innodb_max_purge_lag":                        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_max_purge_lag_delay":                  acctest.Representation{RepType: acctest.Optional, Create: `300000`},
		"innodb_stats_persistent_sample_pages":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"innodb_stats_transient_sample_pages":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"interactive_timeout":                         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"local_infile":                                acctest.Representation{RepType: acctest.Optional, Create: `true`},
		"mandatory_roles":                             acctest.Representation{RepType: acctest.Optional, Create: `public`},
		"max_allowed_packet":                          acctest.Representation{RepType: acctest.Optional, Create: `67108864`},
		"max_binlog_cache_size":                       acctest.Representation{RepType: acctest.Optional, Create: `4096`},
		"max_connect_errors":                          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"max_connections":                             acctest.Representation{RepType: acctest.Optional, Create: `500`},
		"max_execution_time":                          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"max_heap_table_size":                         acctest.Representation{RepType: acctest.Optional, Create: `16384`},
		"max_prepared_stmt_count":                     acctest.Representation{RepType: acctest.Optional, Create: `16382`},
		"mysql_firewall_mode":                         acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"mysql_zstd_default_compression_level":        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_connect_timeout":                      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_deflate_default_compression_level":    acctest.Representation{RepType: acctest.Optional, Create: `9`},
		"mysqlx_deflate_max_client_compression_level": acctest.Representation{RepType: acctest.Optional, Create: `9`},
		"mysqlx_document_id_unique_prefix":            acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_enable_hello_notice":                  acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"mysqlx_idle_worker_thread_timeout":           acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_interactive_timeout":                  acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_lz4default_compression_level":         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_lz4max_client_compression_level":      acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_max_allowed_packet":                   acctest.Representation{RepType: acctest.Optional, Create: `67108864`},
		"mysqlx_min_worker_threads":                   acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_read_timeout":                         acctest.Representation{RepType: acctest.Optional, Create: `30`},
		"mysqlx_wait_timeout":                         acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_write_timeout":                        acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_zstd_default_compression_level":       acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"mysqlx_zstd_max_client_compression_level":    acctest.Representation{RepType: acctest.Optional, Create: `9`},
		"net_read_timeout":                            acctest.Representation{RepType: acctest.Optional, Create: `30`},
		"net_write_timeout":                           acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"parser_max_mem_size":                         acctest.Representation{RepType: acctest.Optional, Create: `10000000`},
		"query_alloc_block_size":                      acctest.Representation{RepType: acctest.Optional, Create: `1024`},
		"query_prealloc_size":                         acctest.Representation{RepType: acctest.Optional, Create: `8192`},
		"regexp_time_limit":                           acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"sort_buffer_size":                            acctest.Representation{RepType: acctest.Optional, Create: `32768`},
		"sql_mode":                                    acctest.Representation{RepType: acctest.Optional, Create: `sqlMode`},
		"sql_require_primary_key":                     acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"sql_warnings":                                acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"thread_pool_dedicated_listeners":             acctest.Representation{RepType: acctest.Optional, Create: `false`},
		"thread_pool_max_transactions_limit":          acctest.Representation{RepType: acctest.Optional, Create: `10`},
		"time_zone":                                   acctest.Representation{RepType: acctest.Optional, Create: `+10:00`},
		"tmp_table_size":                              acctest.Representation{RepType: acctest.Optional, Create: `1024`},
		"transaction_isolation":                       acctest.Representation{RepType: acctest.Optional, Create: `READ-UNCOMMITTED`},
		"wait_timeout":                                acctest.Representation{RepType: acctest.Optional, Create: `10`},
	}

	MysqlMysqlConfigurationResourceDependencies = utils.MysqlConfigurationIdVariable +
		utils.MysqlConfigurationIdVariableE3_2_32_OCID +
		DefinedTagsDependencies
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_mysql_configuration.test_mysql_configuration"
	datasourceName := "data.oci_mysql_mysql_configurations.test_mysql_configurations"
	singularDatasourceName := "data.oci_mysql_mysql_configuration.test_mysql_configuration"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+MysqlMysqlConfigurationResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Optional, acctest.Create, MysqlMysqlConfigurationRepresentation), "mysql", "mysqlConfiguration", t)

	acctest.ResourceTest(t, testAccCheckMysqlMysqlConfigurationDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Required, acctest.Create, MysqlMysqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlConfigurationResourceDependencies,
		},
		// verify Create with optionals
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Optional, acctest.Create, MysqlMysqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "init_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "init_variables.0.lower_case_table_names", "CASE_SENSITIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_configuration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttr(resourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.autocommit", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.big_tables", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_expire_logs_seconds", "3600"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_row_metadata", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_row_value_options", "PARTIAL_JSON"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_transaction_compression", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.completion_type", "NO_CHAIN"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.connect_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.connection_memory_chunk_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.connection_memory_limit", "2097152"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.cte_max_recursion_depth", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.default_authentication_plugin", "mysql_native_password"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.foreign_key_checks", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.generated_random_password_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.global_connection_memory_limit", "2097152"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.global_connection_memory_tracking", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.group_replication_consistency", "BEFORE_ON_PRIMARY_FAILOVER"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.information_schema_stats_expiry", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_buffer_pool_dump_pct", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_buffer_pool_instances", "4"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_buffer_pool_size", "2147483648"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ddl_buffer_size", "65536"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ddl_threads", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_enable_stopword", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_max_token_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_min_token_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_num_word_optimize", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_result_cache_limit", "33554432"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_server_stopword_table", "innodbFtServerStopwordTable"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_lock_wait_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_log_writer_threads", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_max_purge_lag", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_max_purge_lag_delay", "300000"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_stats_persistent_sample_pages", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_stats_transient_sample_pages", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.interactive_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.local_infile", "true"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mandatory_roles", "public"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_allowed_packet", "67108864"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_binlog_cache_size", "4096"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_connect_errors", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_connections", "500"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_execution_time", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_heap_table_size", "16384"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_prepared_stmt_count", "16382"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysql_firewall_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysql_zstd_default_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_connect_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_deflate_default_compression_level", "9"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_deflate_max_client_compression_level", "9"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_document_id_unique_prefix", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_enable_hello_notice", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_idle_worker_thread_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_interactive_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_lz4default_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_lz4max_client_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_max_allowed_packet", "67108864"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_min_worker_threads", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_read_timeout", "30"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_wait_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_write_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_zstd_default_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_zstd_max_client_compression_level", "9"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.net_read_timeout", "30"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.net_write_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.parser_max_mem_size", "10000000"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.query_alloc_block_size", "1024"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.query_prealloc_size", "8192"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.regexp_time_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sort_buffer_size", "32768"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sql_mode", "sqlMode"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sql_require_primary_key", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sql_warnings", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.thread_pool_dedicated_listeners", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.thread_pool_max_transactions_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.time_zone", "+10:00"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.tmp_table_size", "1024"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.transaction_isolation", "READ-UNCOMMITTED"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.wait_timeout", "10"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					// Test case to ensure resource discovery works properly
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + MysqlMysqlConfigurationResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Optional, acctest.Update, MysqlMysqlConfigurationRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(resourceName, "description", "description2"),
				resource.TestCheckResourceAttr(resourceName, "display_name", "displayName2"),
				resource.TestCheckResourceAttr(resourceName, "freeform_tags.%", "1"),
				resource.TestCheckResourceAttrSet(resourceName, "id"),
				resource.TestCheckResourceAttr(resourceName, "init_variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "init_variables.0.lower_case_table_names", "CASE_SENSITIVE"),
				resource.TestCheckResourceAttrSet(resourceName, "parent_configuration_id"),
				resource.TestCheckResourceAttrSet(resourceName, "shape_name"),
				resource.TestCheckResourceAttrSet(resourceName, "state"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),
				resource.TestCheckResourceAttrSet(resourceName, "type"),
				resource.TestCheckResourceAttr(resourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.autocommit", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.big_tables", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_expire_logs_seconds", "3600"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_row_metadata", "FULL"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_row_value_options", "PARTIAL_JSON"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.binlog_transaction_compression", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.completion_type", "NO_CHAIN"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.connect_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.connection_memory_chunk_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.connection_memory_limit", "2097152"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.cte_max_recursion_depth", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.default_authentication_plugin", "mysql_native_password"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.foreign_key_checks", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.generated_random_password_length", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.global_connection_memory_limit", "2097152"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.global_connection_memory_tracking", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.group_replication_consistency", "BEFORE_ON_PRIMARY_FAILOVER"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.information_schema_stats_expiry", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_buffer_pool_dump_pct", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_buffer_pool_instances", "4"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_buffer_pool_size", "2147483648"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ddl_buffer_size", "65536"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ddl_threads", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_enable_stopword", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_max_token_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_min_token_size", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_num_word_optimize", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_result_cache_limit", "33554432"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_ft_server_stopword_table", "innodbFtServerStopwordTable"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_lock_wait_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_log_writer_threads", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_max_purge_lag", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_max_purge_lag_delay", "300000"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_stats_persistent_sample_pages", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.innodb_stats_transient_sample_pages", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.interactive_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.local_infile", "true"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mandatory_roles", "public"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_allowed_packet", "67108864"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_binlog_cache_size", "4096"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_connect_errors", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_connections", "500"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_execution_time", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_heap_table_size", "16384"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.max_prepared_stmt_count", "16382"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysql_firewall_mode", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysql_zstd_default_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_connect_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_deflate_default_compression_level", "9"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_deflate_max_client_compression_level", "9"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_document_id_unique_prefix", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_enable_hello_notice", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_idle_worker_thread_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_interactive_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_lz4default_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_lz4max_client_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_max_allowed_packet", "67108864"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_min_worker_threads", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_read_timeout", "30"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_wait_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_write_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_zstd_default_compression_level", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.mysqlx_zstd_max_client_compression_level", "9"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.net_read_timeout", "30"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.net_write_timeout", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.parser_max_mem_size", "10000000"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.query_alloc_block_size", "1024"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.query_prealloc_size", "8192"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.regexp_time_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sort_buffer_size", "32768"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sql_mode", "sqlMode"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sql_require_primary_key", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.sql_warnings", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.thread_pool_dedicated_listeners", "false"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.thread_pool_max_transactions_limit", "10"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.time_zone", "+10:00"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.tmp_table_size", "1024"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.transaction_isolation", "READ-UNCOMMITTED"),
				resource.TestCheckResourceAttr(resourceName, "variables.0.wait_timeout", "10"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_configurations", "test_mysql_configurations", acctest.Required, acctest.Create, MysqlMysqlMysqlConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlMysqlConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),

				resource.TestCheckResourceAttrSet(datasourceName, "configurations.#"),
				resource.TestCheckResourceAttrSet(datasourceName, "configurations.0.id"),
				resource.TestCheckResourceAttrSet(datasourceName, "configurations.0.shape_name"),
				resource.TestCheckResourceAttrSet(datasourceName, "configurations.0.state"),
				resource.TestCheckResourceAttrSet(datasourceName, "configurations.0.time_created"),
				resource.TestCheckResourceAttrSet(datasourceName, "configurations.0.type"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Required, acctest.Create, MysqlMysqlMysqlConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlMysqlConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "Default Standalone configuration for the MySQL.VM.Standard.E3.1.8GB MySQL Shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "MySQL.VM.Standard.E3.1.8GB.Standalone"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttr(singularDatasourceName, "init_variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "init_variables.0.lower_case_table_names", "CASE_SENSITIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.autocommit", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.big_tables", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_expire_logs_seconds", "3600"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_row_metadata", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_row_value_options", "PARTIAL_JSON"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_transaction_compression", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.completion_type", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.connect_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.connection_memory_chunk_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.connection_memory_limit", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.cte_max_recursion_depth", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.default_authentication_plugin", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.foreign_key_checks", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.generated_random_password_length", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.global_connection_memory_limit", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.global_connection_memory_tracking", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.group_replication_consistency", "BEFORE_ON_PRIMARY_FAILOVER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.information_schema_stats_expiry", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_buffer_pool_dump_pct", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_buffer_pool_instances", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_buffer_pool_size", "2147483648"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ddl_buffer_size", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ddl_threads", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_enable_stopword", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_max_token_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_min_token_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_num_word_optimize", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_result_cache_limit", "33554432"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_server_stopword_table", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_lock_wait_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_log_writer_threads", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_max_purge_lag", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_max_purge_lag_delay", "300000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_stats_persistent_sample_pages", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_stats_transient_sample_pages", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.interactive_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.local_infile", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mandatory_roles", "public"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_allowed_packet", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_binlog_cache_size", "4294967296"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_connect_errors", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_connections", "500"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_execution_time", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_heap_table_size", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_prepared_stmt_count", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysql_firewall_mode", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_connect_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_deflate_default_compression_level", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_deflate_max_client_compression_level", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_document_id_unique_prefix", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_enable_hello_notice", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_idle_worker_thread_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_interactive_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_lz4default_compression_level", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_lz4max_client_compression_level", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_max_allowed_packet", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_min_worker_threads", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_read_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_wait_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_write_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_zstd_default_compression_level", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_zstd_max_client_compression_level", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.net_read_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.net_write_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.parser_max_mem_size", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.query_alloc_block_size", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.query_prealloc_size", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.regexp_time_limit", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.sort_buffer_size", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.sql_mode", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.sql_require_primary_key", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.sql_warnings", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.thread_pool_dedicated_listeners", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.thread_pool_max_transactions_limit", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.time_zone", "UTC"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.tmp_table_size", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.transaction_isolation", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.wait_timeout", "0"),
			),
		},
		// verify resource import
		{
			Config:                  config + MysqlMysqlConfigurationRequiredOnlyResource,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMysqlMysqlConfigurationDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).MysqlaasClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_mysql_configuration" {
			noResourceFound = false
			request := oci_mysql.GetConfigurationRequest{}

			tmp := rs.Primary.ID
			request.ConfigurationId = &tmp

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetConfiguration(context.Background(), request)

			if err == nil {
				configurationTypes := map[string]bool{
					string(oci_mysql.ConfigurationTypeDefault): true,
				}
				if _, ok := configurationTypes[string(response.Type)]; ok {
					//resource is a DEFAULT resource and will always be ACTIVE.
					continue
				}

				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.ConfigurationLifecycleStateDeleted): true,
				}
				if _, ok := deletedLifecycleStates[string(response.LifecycleState)]; !ok {
					//resource lifecycle state is not in expected deleted lifecycle states.
					return fmt.Errorf("resource lifecycle state: %s is not in expected deleted lifecycle states", response.LifecycleState)
				}
				//resource lifecycle state is in expected deleted lifecycle states. continue with next one.
				continue
			}

			//Verify that exception is for '404 not found'.
			if failure, isServiceError := common.IsServiceError(err); !isServiceError || failure.GetHTTPStatusCode() != 404 {
				return err
			}
		}
	}
	if noResourceFound {
		return fmt.Errorf("at least one resource was expected from the state file, but could not be found")
	}

	return nil
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("MysqlMysqlConfiguration") {
		resource.AddTestSweepers("MysqlMysqlConfiguration", &resource.Sweeper{
			Name:         "MysqlMysqlConfiguration",
			Dependencies: acctest.DependencyGraph["mysqlConfiguration"],
			F:            sweepMysqlMysqlConfigurationResource,
		})
	}
}

func sweepMysqlMysqlConfigurationResource(compartment string) error {
	mysqlaasClient := acctest.GetTestClients(&schema.ResourceData{}).MysqlaasClient()
	mysqlConfigurationIds, err := getMysqlMysqlConfigurationIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlConfigurationId := range mysqlConfigurationIds {
		if ok := acctest.SweeperDefaultResourceId[mysqlConfigurationId]; !ok {
			deleteConfigurationRequest := oci_mysql.DeleteConfigurationRequest{}
			deleteConfigurationRequest.ConfigurationId = &mysqlConfigurationId

			deleteConfigurationRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := mysqlaasClient.DeleteConfiguration(context.Background(), deleteConfigurationRequest)
			if error != nil {
				fmt.Printf("Error deleting MysqlConfiguration %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlConfigurationId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mysqlConfigurationId, MysqlMysqlConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				MysqlMysqlConfigurationSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlMysqlConfigurationIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "MysqlConfigurationId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	mysqlaasClient := acctest.GetTestClients(&schema.ResourceData{}).MysqlaasClient()

	listConfigurationsRequest := oci_mysql.ListConfigurationsRequest{}
	listConfigurationsRequest.CompartmentId = &compartmentId
	listConfigurationsRequest.LifecycleState = oci_mysql.ConfigurationLifecycleStateActive
	listConfigurationsResponse, err := mysqlaasClient.ListConfigurations(context.Background(), listConfigurationsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting MysqlConfiguration list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, mysqlConfiguration := range listConfigurationsResponse.Items {
		id := *mysqlConfiguration.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "MysqlConfigurationId", id)
	}
	return resourceIds, nil
}

func MysqlMysqlConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mysqlConfigurationResponse, ok := response.Response.(oci_mysql.GetConfigurationResponse); ok {
		return mysqlConfigurationResponse.LifecycleState != oci_mysql.ConfigurationLifecycleStateDeleted
	}
	return false
}

func MysqlMysqlConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MysqlaasClient().GetConfiguration(context.Background(), oci_mysql.GetConfigurationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
