// Copyright (c) 2017, 2024, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/oracle/oci-go-sdk/v65/common"
	oci_data_safe "github.com/oracle/oci-go-sdk/v65/datasafe"

	"github.com/oracle/terraform-provider-oci/httpreplay"
	"github.com/oracle/terraform-provider-oci/internal/acctest"
	tf_client "github.com/oracle/terraform-provider-oci/internal/client"
	"github.com/oracle/terraform-provider-oci/internal/tfresource"

	"github.com/oracle/terraform-provider-oci/internal/utils"
)

var (
	DataSafeSqlFirewallAllowedSqlSingularDataSourceRepresentation = map[string]interface{}{
		"sql_firewall_allowed_sql_id": acctest.Representation{RepType: acctest.Required, Create: `${var.sql_firewall_allowed_sql_id}`},
	}

	DataSafeSqlFirewallAllowedSqlDataSourceRepresentation = map[string]interface{}{
		"compartment_id":            acctest.Representation{RepType: acctest.Required, Create: `${var.compartment_id}`},
		"access_level":              acctest.Representation{RepType: acctest.Optional, Create: `ACCESSIBLE`},
		"compartment_id_in_subtree": acctest.Representation{RepType: acctest.Optional, Create: `true`},
	}

	DataSafeSqlFirewallAllowedSqlResourceConfig = ""
)

// issue-routing-tag: data_safe/default
func TestDataSafeSqlFirewallAllowedSqlResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestDataSafeSqlFirewallAllowedSqlResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)
	sqlFirewallAllowedSqlId := utils.GetEnvSettingWithBlankDefault("sql_firewall_allowed_sql_ocid")
	sqlFirewallAllowedSqlIdVariableStr := fmt.Sprintf("variable \"sql_firewall_allowed_sql_id\" { default = \"%s\" }\n", sqlFirewallAllowedSqlId)
	datasourceName := "data.oci_data_safe_sql_firewall_allowed_sqls.test_sql_firewall_allowed_sqls"
	singularDatasourceName := "data.oci_data_safe_sql_firewall_allowed_sql.test_sql_firewall_allowed_sql"

	acctest.SaveConfigContent("", "", "", t)

	acctest.ResourceTest(t, nil, []resource.TestStep{
		// verify datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_allowed_sqls", "test_sql_firewall_allowed_sqls", acctest.Optional, acctest.Create, DataSafeSqlFirewallAllowedSqlDataSourceRepresentation) +
				compartmentIdVariableStr + DataSafeSqlFirewallAllowedSqlResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(datasourceName, "access_level", "ACCESSIBLE"),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id", compartmentId),
				resource.TestCheckResourceAttr(datasourceName, "compartment_id_in_subtree", "true"),
				resource.TestCheckResourceAttrSet(datasourceName, "sql_firewall_allowed_sql_collection.#"),
			),
		},
		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_data_safe_sql_firewall_allowed_sql", "test_sql_firewall_allowed_sql", acctest.Required, acctest.Create, DataSafeSqlFirewallAllowedSqlSingularDataSourceRepresentation) +
				sqlFirewallAllowedSqlIdVariableStr + DataSafeSqlFirewallAllowedSqlResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_firewall_allowed_sql_id"),

				resource.TestCheckResourceAttrSet(singularDatasourceName, "compartment_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "current_user"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_user_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "display_name"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_accessed_objects.#"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_firewall_policy_id"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_level"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "sql_text"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "state"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_collected"),
			),
		},
	})
}

func init() {
	if acctest.DependencyGraph == nil {
		acctest.InitDependencyGraph()
	}
	if !acctest.InSweeperExcludeList("DataSafeSqlFirewallAllowedSql") {
		resource.AddTestSweepers("DataSafeSqlFirewallAllowedSql", &resource.Sweeper{
			Name:         "DataSafeSqlFirewallAllowedSql",
			Dependencies: acctest.DependencyGraph["sqlFirewallAllowedSql"],
			F:            sweepDataSafeSqlFirewallAllowedSqlResource,
		})
	}
}

func sweepDataSafeSqlFirewallAllowedSqlResource(compartment string) error {
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()
	sqlFirewallAllowedSqlIds, err := getDataSafeSqlFirewallAllowedSqlIds(compartment)
	if err != nil {
		return err
	}
	for _, sqlFirewallAllowedSqlId := range sqlFirewallAllowedSqlIds {
		if ok := acctest.SweeperDefaultResourceId[sqlFirewallAllowedSqlId]; !ok {
			deleteSqlFirewallAllowedSqlRequest := oci_data_safe.DeleteSqlFirewallAllowedSqlRequest{}

			deleteSqlFirewallAllowedSqlRequest.SqlFirewallAllowedSqlId = &sqlFirewallAllowedSqlId

			deleteSqlFirewallAllowedSqlRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "data_safe")
			_, error := dataSafeClient.DeleteSqlFirewallAllowedSql(context.Background(), deleteSqlFirewallAllowedSqlRequest)
			if error != nil {
				fmt.Printf("Error deleting SqlFirewallAllowedSql %s %s, It is possible that the resource is already deleted. Please verify manually \n", sqlFirewallAllowedSqlId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &sqlFirewallAllowedSqlId, DataSafeSqlFirewallAllowedSqlSweepWaitCondition, time.Duration(3*time.Minute),
				DataSafeSqlFirewallAllowedSqlSweepResponseFetchOperation, "data_safe", true)
		}
	}
	return nil
}

func getDataSafeSqlFirewallAllowedSqlIds(compartment string) ([]string, error) {
	ids := acctest.GetResourceIdsToSweep(compartment, "SqlFirewallAllowedSqlId")
	if ids != nil {
		return ids, nil
	}
	var resourceIds []string
	compartmentId := compartment
	dataSafeClient := acctest.GetTestClients(&schema.ResourceData{}).DataSafeClient()

	listSqlFirewallAllowedSqlsRequest := oci_data_safe.ListSqlFirewallAllowedSqlsRequest{}
	listSqlFirewallAllowedSqlsRequest.CompartmentId = &compartmentId
	listSqlFirewallAllowedSqlsResponse, err := dataSafeClient.ListSqlFirewallAllowedSqls(context.Background(), listSqlFirewallAllowedSqlsRequest)

	if err != nil {
		return resourceIds, fmt.Errorf("Error getting SqlFirewallAllowedSql list for compartment id : %s , %s \n", compartmentId, err)
	}
	for _, sqlFirewallAllowedSql := range listSqlFirewallAllowedSqlsResponse.Items {
		id := *sqlFirewallAllowedSql.Id
		resourceIds = append(resourceIds, id)
		acctest.AddResourceIdToSweeperResourceIdMap(compartmentId, "SqlFirewallAllowedSqlId", id)
	}
	return resourceIds, nil
}

func DataSafeSqlFirewallAllowedSqlSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if sqlFirewallAllowedSqlResponse, ok := response.Response.(oci_data_safe.GetSqlFirewallAllowedSqlResponse); ok {
		return sqlFirewallAllowedSqlResponse.LifecycleState != oci_data_safe.SqlFirewallAllowedSqlLifecycleStateDeleted
	}
	return false
}

func DataSafeSqlFirewallAllowedSqlSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DataSafeClient().GetSqlFirewallAllowedSql(context.Background(), oci_data_safe.GetSqlFirewallAllowedSqlRequest{
		SqlFirewallAllowedSqlId: resourceId,
		RequestMetadata: common.RequestMetadata{
			RetryPolicy: retryPolicy,
		},
	})
	return err
}
