// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"

	"github.com/oracle/oci-go-sdk/v56/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v56/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	mysqlConfigurationSingularDataSourceRepresentation = map[string]interface{}{
		"configuration_id": acctest.Representation{RepType: acctest.Required, Create: `${var.MysqlConfigurationOCID[var.region]}`},
	}

	mysqlConfigurationDataSourceRepresentation = map[string]interface{}{
		"compartment_id":   acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"configuration_id": acctest.Representation{RepType: acctest.Optional, Create: `${var.MysqlConfigurationOCID[var.region]}`},
		"display_name":     acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E2.2.Built-in`},
		"shape_name":       acctest.Representation{RepType: acctest.Optional, Create: `VM.Standard.E2.2`},
		"state":            acctest.Representation{RepType: acctest.Optional, Create: `ACTIVE`},
		"type":             acctest.Representation{RepType: acctest.Optional, Create: []string{`DEFAULT`}},
	}

	MysqlConfigurationResourceConfig = utils.MysqlConfigurationIdVariable
)

// issue-routing-tag: mysql/default
func TestMysqlMysqlConfigurationResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlMysqlConfigurationResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	datasourceName := "data.oci_mysql_mysql_configurations.test_mysql_configurations"
	singularDatasourceName := "data.oci_mysql_mysql_configuration.test_mysql_configuration"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_configurations", "test_mysql_configurations", acctest.Required, acctest.Create, mysqlConfigurationDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlConfigurationResourceConfig,
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
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_mysql_configuration", "test_mysql_configuration", acctest.Required, acctest.Create, mysqlConfigurationSingularDataSourceRepresentation) +
				compartmentIdVariableStr + MysqlConfigurationResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "configuration_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "description", "Default Standalone configuration for the VM.Standard.E2.2 MySQL Shape"),
				resource.TestCheckResourceAttr(singularDatasourceName, "display_name", "VM.Standard.E2.2.Standalone"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "type"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.#", "1"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.autocommit", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_expire_logs_seconds", "3600"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_row_metadata", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_row_value_options", "PARTIAL_JSON"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.binlog_transaction_compression", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.completion_type", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.connect_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.cte_max_recursion_depth", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.default_authentication_plugin", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.foreign_key_checks", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.generated_random_password_length", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.group_replication_consistency", "BEFORE_ON_PRIMARY_FAILOVER"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.information_schema_stats_expiry", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_buffer_pool_instances", "4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_buffer_pool_size", "10200547328"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_enable_stopword", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_max_token_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_min_token_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_num_word_optimize", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_result_cache_limit", "33554432"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_ft_server_stopword_table", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_lock_wait_timeout", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_max_purge_lag", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.innodb_max_purge_lag_delay", "300000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.local_infile", "true"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mandatory_roles", "public"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_connections", "1000"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.max_execution_time", "0"),
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
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.mysqlx_zstd_max_client_compression_level", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.parser_max_mem_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.query_alloc_block_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.query_prealloc_size", "0"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.sql_mode", ""),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.sql_require_primary_key", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.sql_warnings", "false"),
				resource.TestCheckResourceAttr(singularDatasourceName, "variables.0.transaction_isolation", ""),
			),
		},
	})
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
	mysqlConfigurationIds, err := getMysqlConfigurationIds(compartment)
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
			acctest.WaitTillCondition(acctest.TestAccProvider, &mysqlConfigurationId, mysqlConfigurationSweepWaitCondition, time.Duration(3*time.Minute),
				mysqlConfigurationSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func getMysqlConfigurationIds(compartment string) ([]string, error) {
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

func mysqlConfigurationSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if mysqlConfigurationResponse, ok := response.Response.(oci_mysql.GetConfigurationResponse); ok {
		return mysqlConfigurationResponse.LifecycleState != oci_mysql.ConfigurationLifecycleStateDeleted
	}
	return false
}

func mysqlConfigurationSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.MysqlaasClient().GetConfiguration(context.Background(), oci_mysql.GetConfigurationRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
