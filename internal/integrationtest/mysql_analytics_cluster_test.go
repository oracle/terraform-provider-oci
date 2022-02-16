// Copyright (c) 2017, 2021, Oracle and/or its affiliates. All rights reserved.
// Licensed under the Mozilla Public License v2.0

package integrationtest

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/oracle/oci-go-sdk/v58/common"
	oci_mysql "github.com/oracle/oci-go-sdk/v58/mysql"

	"github.com/terraform-providers/terraform-provider-oci/httpreplay"
	"github.com/terraform-providers/terraform-provider-oci/internal/acctest"
	tf_client "github.com/terraform-providers/terraform-provider-oci/internal/client"
	"github.com/terraform-providers/terraform-provider-oci/internal/resourcediscovery"
	"github.com/terraform-providers/terraform-provider-oci/internal/tfresource"
	"github.com/terraform-providers/terraform-provider-oci/internal/utils"
)

var (
	AnalyticsClusterRequiredOnlyResource = AnalyticsClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Required, acctest.Create, analyticsClusterRepresentation)

	AnalyticsClusterResourceConfig = AnalyticsClusterResourceDependencies +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Optional, acctest.Update, analyticsClusterRepresentation)

	analyticsClusterSingularDataSourceRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
	}

	analyticsClusterRepresentation = map[string]interface{}{
		"db_system_id": acctest.Representation{RepType: acctest.Required, Create: `${oci_mysql_mysql_db_system.test_mysql_db_system.id}`},
		"cluster_size": acctest.Representation{RepType: acctest.Required, Create: `2`, Update: `3`},
		"shape_name":   acctest.Representation{RepType: acctest.Required, Create: `VM.Standard.E2.2`, Update: `VM.Standard.E2.4`},
		"state":        acctest.Representation{RepType: acctest.Optional, Create: `INACTIVE`, Update: `ACTIVE`},
	}

	AnalyticsClusterResourceDependencies = MysqlConfigurationResourceConfig +
		acctest.GenerateResourceFromRepresentationMap("oci_core_subnet", "test_subnet", acctest.Required, acctest.Create, subnetRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_core_vcn", "test_vcn", acctest.Required, acctest.Create, vcnRepresentation) +
		AvailabilityDomainConfig +
		acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_shapes", "test_shapes", acctest.Required, acctest.Create, mysqlShapeDataSourceRepresentation) +
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_mysql_db_system", "test_mysql_db_system", acctest.Required, acctest.Create, mysqlDbSystemRepresentation)
)

// issue-routing-tag: mysql/default
func TestMysqlAnalyticsClusterResource_basic(t *testing.T) {
	httpreplay.SetScenario("TestMysqlAnalyticsClusterResource_basic")
	defer httpreplay.SaveScenario()

	config := acctest.ProviderTestConfig()

	compartmentId := utils.GetEnvSettingWithBlankDefault("compartment_ocid")
	compartmentIdVariableStr := fmt.Sprintf("variable \"compartment_id\" { default = \"%s\" }\n", compartmentId)

	resourceName := "oci_mysql_analytics_cluster.test_analytics_cluster"
	singularDatasourceName := "data.oci_mysql_analytics_cluster.test_analytics_cluster"

	var resId, resId2 string
	// Save TF content to Create resource with optional properties. This has to be exactly the same as the config part in the "Create with optionals" step in the test.
	acctest.SaveConfigContent(config+compartmentIdVariableStr+AnalyticsClusterResourceDependencies+
		acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Optional, acctest.Create, analyticsClusterRepresentation), "mysql", "analyticsCluster", t)

	acctest.ResourceTest(t, testAccCheckMysqlAnalyticsClusterDestroy, []resource.TestStep{
		// verify Create
		{
			Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Required, acctest.Create, analyticsClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					return err
				},
			),
		},

		// delete before next Create
		{
			Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies,
		},

		// verify Create & stop
		{
			Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Optional, acctest.Create, analyticsClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "2"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "2"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "VM.Standard.E2.2"),
				resource.TestCheckResourceAttr(resourceName, "state", "INACTIVE"),

				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId, err = acctest.FromInstanceState(s, resourceName, "id")
					if isEnableExportCompartment, _ := strconv.ParseBool(utils.GetEnvSettingWithDefault("enable_export_compartment", "true")); isEnableExportCompartment {
						if errExport := resourcediscovery.TestExportCompartmentWithResourceName(&resId, &compartmentId, resourceName); errExport != nil {
							return errExport
						}
					}
					return err
				},
			),
		},

		// verify start & updates to updatable parameters
		{
			Config: config + compartmentIdVariableStr + AnalyticsClusterResourceDependencies +
				acctest.GenerateResourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Optional, acctest.Update, analyticsClusterRepresentation),
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttr(resourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(resourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(resourceName, "shape_name", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttr(resourceName, "state", "ACTIVE"),

				resource.TestCheckResourceAttrSet(resourceName, "db_system_id"),
				resource.TestCheckResourceAttrSet(resourceName, "time_created"),
				resource.TestCheckResourceAttrSet(resourceName, "time_updated"),

				func(s *terraform.State) (err error) {
					resId2, err = acctest.FromInstanceState(s, resourceName, "id")
					if resId != resId2 {
						return fmt.Errorf("Resource recreated when it was supposed to be updated.")
					}
					return err
				},
			),
		},

		// verify singular datasource
		{
			Config: config +
				acctest.GenerateDataSourceFromRepresentationMap("oci_mysql_analytics_cluster", "test_analytics_cluster", acctest.Required, acctest.Create, analyticsClusterSingularDataSourceRepresentation) +
				compartmentIdVariableStr + AnalyticsClusterResourceConfig,
			Check: acctest.ComposeAggregateTestCheckFuncWrapper(
				resource.TestCheckResourceAttrSet(singularDatasourceName, "db_system_id"),

				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_nodes.#", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "cluster_size", "3"),
				resource.TestCheckResourceAttr(singularDatasourceName, "shape_name", "VM.Standard.E2.4"),
				resource.TestCheckResourceAttr(singularDatasourceName, "state", "ACTIVE"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_created"),
				resource.TestCheckResourceAttrSet(singularDatasourceName, "time_updated"),
			),
		},
		// remove singular datasource from previous step so that it doesn't conflict with import tests
		{
			Config: config + compartmentIdVariableStr + AnalyticsClusterResourceConfig,
		},
		// verify resource import
		{
			Config:                  config,
			ImportState:             true,
			ImportStateVerify:       true,
			ImportStateVerifyIgnore: []string{},
			ResourceName:            resourceName,
		},
	})
}

func testAccCheckMysqlAnalyticsClusterDestroy(s *terraform.State) error {
	noResourceFound := true
	client := acctest.TestAccProvider.Meta().(*tf_client.OracleClients).DbSystemClient()
	for _, rs := range s.RootModule().Resources {
		if rs.Type == "oci_mysql_analytics_cluster" {
			noResourceFound = false
			request := oci_mysql.GetAnalyticsClusterRequest{}

			if value, ok := rs.Primary.Attributes["db_system_id"]; ok {
				request.DbSystemId = &value
			}

			request.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")

			response, err := client.GetAnalyticsCluster(context.Background(), request)

			if err == nil {
				deletedLifecycleStates := map[string]bool{
					string(oci_mysql.AnalyticsClusterLifecycleStateDeleted): true,
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
	if !acctest.InSweeperExcludeList("MysqlAnalyticsCluster") {
		resource.AddTestSweepers("MysqlAnalyticsCluster", &resource.Sweeper{
			Name:         "MysqlAnalyticsCluster",
			Dependencies: acctest.DependencyGraph["analyticsCluster"],
			F:            sweepMysqlAnalyticsClusterResource,
		})
	}
}

func sweepMysqlAnalyticsClusterResource(compartment string) error {
	dbSystemClient := acctest.GetTestClients(&schema.ResourceData{}).DbSystemClient()
	mysqlDbSystemIds, err := getMysqlDbSystemIds(compartment)
	if err != nil {
		return err
	}
	for _, mysqlDbSystemId := range mysqlDbSystemIds {
		if ok := acctest.SweeperDefaultResourceId[mysqlDbSystemId]; !ok {
			deleteAnalyticsClusterRequest := oci_mysql.DeleteAnalyticsClusterRequest{}
			deleteAnalyticsClusterRequest.DbSystemId = &mysqlDbSystemId

			deleteAnalyticsClusterRequest.RequestMetadata.RetryPolicy = tfresource.GetRetryPolicy(true, "mysql")
			_, error := dbSystemClient.DeleteAnalyticsCluster(context.Background(), deleteAnalyticsClusterRequest)
			if error != nil {
				fmt.Printf("Error deleting AnalyticsCluster of DbSystem %s %s, It is possible that the resource is already deleted. Please verify manually \n", mysqlDbSystemId, error)
				continue
			}
			acctest.WaitTillCondition(acctest.TestAccProvider, &mysqlDbSystemId, analyticsClusterSweepWaitCondition, time.Duration(3*time.Minute),
				analyticsClusterSweepResponseFetchOperation, "mysql", true)
		}
	}
	return nil
}

func analyticsClusterSweepWaitCondition(response common.OCIOperationResponse) bool {
	// Only stop if the resource is available beyond 3 mins. As there could be an issue for the sweeper to delete the resource and manual intervention required.
	if analyticsClusterResponse, ok := response.Response.(oci_mysql.GetAnalyticsClusterResponse); ok {
		return analyticsClusterResponse.LifecycleState != oci_mysql.AnalyticsClusterLifecycleStateDeleted
	}
	return false
}

func analyticsClusterSweepResponseFetchOperation(client *tf_client.OracleClients, resourceId *string, retryPolicy *common.RetryPolicy) error {
	_, err := client.DbSystemClient().GetAnalyticsCluster(context.Background(), oci_mysql.GetAnalyticsClusterRequest{RequestMetadata: common.RequestMetadata{
		RetryPolicy: retryPolicy,
	},
	})
	return err
}
